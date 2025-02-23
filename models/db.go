package models

type DbConfig struct {
	AppName  string
	Scheme   string
	Username string
	Password string
	Host     string
	Port     int64
}