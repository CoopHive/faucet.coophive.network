package config

import (
	"github.com/CoopHive/hive/utils"
)

// this is the address of the user
var X_COOPHIVE_USER_HEADER = utils.Base64DecodeFast("WC1MaWx5cGFkLVVzZXI=")

// this is the signature of the message
var X_COOPHIVE_SIGNATURE_HEADER = utils.Base64DecodeFast("WC1MaWx5cGFkLVNpZ25hdHVyZQ==")

// the context name we keep the address
const CONTEXT_ADDRESS = "address"

// the sub path any API's are served over
const API_SUB_PATH = "/api/v1"

// the sub path the websocket server is mounted on
const WEBSOCKET_SUB_PATH = "/ws"
