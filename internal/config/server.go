package config

type APIServer struct {
	DatabaseConfig
	ServerConfig
	GeneralConfig
	AuthConfig
	OrderConfig
	ReaderConfig
	EbookConfig
	SsoConfig
}
