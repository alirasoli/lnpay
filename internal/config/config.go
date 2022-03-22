package config

// viper uses mapstructure library to decode configs
type (
	Config struct {
		Server  Server  `mapstructure:"server"`
		Payment Payment `mapstructure:"payment"`
	}

	Server struct {
		Http Http `mapstructure:"http"`
	}

	Http struct {
		Port string `mapstructure:"port"`
	}

	Payment struct {
		LNbits LNbits `mapstructure:"lnbits"`
	}

	LNbits struct {
		URL        string `mapstructure:"url"`
		InvoiceKey string `mapstructure:"invoice_key"`
	}
)
