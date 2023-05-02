package database

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "",
		DBName:   "miniproject_wardah",
	}
}
