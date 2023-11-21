package ient

type Config struct {
	DBName string `yaml:"DBName" env:"DBNAME"`

	Username string `yaml:"Username" env:"USERNAME" env-default:""`
	Password string `yaml:"Password" env:"PASSWORD" env-default:""`
	Hostname string `yaml:"Hostname" env:"HOSTNAME"                env-required:"true"`
	Port     int    `yaml:"Port"     env:"PORT"                    env-required:"true"`
}
