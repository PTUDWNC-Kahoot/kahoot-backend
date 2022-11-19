package config

type Specification struct {
	DBConnection string `envconfig:"identity_dbconnection"`
	Port         string `envconfig:"identity_port"`
	SecretKey    string `envconfig:"identity_secretkey"`
}
