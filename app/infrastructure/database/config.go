package database

type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewMySQLConfig() *MySQLConfig {
	return &MySQLConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "123",
		Database: "miniproject",
	}
}
