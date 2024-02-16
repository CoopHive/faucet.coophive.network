package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"github.com/CoopHive/faucet.coophive.network/internal/chain"
	"github.com/CoopHive/faucet.coophive.network/web"
)

type Server struct {
	chain.TxBuilder
	cfg *Config
}

func NewServer(builder chain.TxBuilder, cfg *Config) *Server {
	return &Server{
		TxBuilder: builder,
		cfg:       cfg,
	}
}

func (s *Server) setupRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(web.Dist()))
	limiter := NewLimiter(s.cfg.proxyCount, time.Duration(s.cfg.interval)*time.Minute)
	hcaptcha := NewCaptcha(s.cfg.hcaptchaSiteKey, s.cfg.hcaptchaSecret)
	router.Handle("/api/claim", negroni.New(limiter, hcaptcha, negroni.Wrap(s.handleClaim())))
	router.Handle("/api/info", s.handleInfo())

	return router
}

func (s *Server) Run() {
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.UseHandler(s.setupRouter())
	log.Infof("Starting http server %d", s.cfg.httpPort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.cfg.httpPort), n))
}

func (s *Server) handleClaim() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.NotFound(w, r)
			return
		}

		// The error always be nil since it has already been handled in limiter
		address, _ := readAddress(r)
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()
		txHash, err := s.Transfer(ctx, address, chain.EtherToWei(int64(s.cfg.Payout)))
		if err != nil {
			log.WithError(err).Error("Failed to send transaction")
			renderJSON(w, claimResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		tokenTxHash, err := s.TransferTokens(ctx, address, chain.EtherToWei(int64(s.cfg.tokenPayout)))
		if err != nil {
			log.WithError(err).Error("Failed to send transaction")
			renderJSON(w, claimResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		log.WithFields(log.Fields{
			"txHash":      txHash,
			"tokenTxHash": tokenTxHash,
			"address":     address,
		}).Info("Transaction sent successfully")
		resp := claimResponse{Message: fmt.Sprintf("Txhash: %s, TokenTxhash: %s", txHash, tokenTxHash)}
		renderJSON(w, resp, http.StatusOK)
	}
}

func (s *Server) handleInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}
		renderJSON(w, infoResponse{
			Account:         s.Sender().String(),
			Network:         s.cfg.Network,
			Symbol:          s.cfg.Symbol,
			Payout:          strconv.Itoa(s.cfg.Payout),
			HcaptchaSiteKey: s.cfg.HcaptchaSiteKey,
		}, http.StatusOK)
	}
}
