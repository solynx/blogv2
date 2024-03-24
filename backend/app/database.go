package app

type DatabaseConfig struct {
	UserName string
	Password string
	IPAddress string
	Port string
	DatabaseName string
}

func (config *DatabaseConfig) Setup() {
	config.GetConfigFile()
}

func (config *DatabaseConfig) GetConfigFile() {
	config.UserName = "alidoan"
	config.Password = "alidoan"
	config.IPAddress = "127.0.0.1"
	config.Port = "3306"
	config.DatabaseName = "v2blog"
}

func (config *DatabaseConfig) Connect() {

}