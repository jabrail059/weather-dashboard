package models

type Config struct {
	AppAddr       string
	SqlitePath    string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}
