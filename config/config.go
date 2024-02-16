package config

type configMap[T ~string] map[T]*argvMeta

type argvMeta struct {
	desc       string
	defaultVal string
}
