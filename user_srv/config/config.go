package config

type PgsqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type SeverConfig struct {
	Name        string `mapstructure:"name" json:"name"`
	PgsqlConfig `mapstructure:"pgsql" json:"pgsql"`
}
