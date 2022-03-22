package config

type (
	Config struct {
		Server Server `yaml:"server"`
		LNbits LNbits `yaml:"lnbits"`
	}

	Server struct {
		Http Http `yaml:"http"`
	}

	Http struct {
		Port string `yaml:"port"`
	}

	LNbits struct {
		URL        string `yaml:"url"`
		InvoiceKey string `yaml:"invoiceKey"`
	}
)
