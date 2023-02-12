package config

type Server struct {
	RunMode  string `yaml:"RunMode"`
	HttpPort string `yaml:"HttpPort"`
}

type DataBase struct {
	DBType   string `yaml:"DBType"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	DBName   string `yaml:"DBName"`
}
