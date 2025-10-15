package config

type Config struct {
	Server Server
	Database Database
	Jwt Jwt
	Storage Storage
	Redis Redis
	App App
}

type App struct{
	AppEnv string
}

type Server struct {
	Host string
	Port string
	Asset string
}

type Jwt struct {
	Key string
	Exp int
}

type Database struct {
	Host string
	Port string
	Name string
	User string
	Pass string
	Tz string
}

type Storage struct {
	BasePath string
}

type Redis struct {
	Host string
	Port string
	Password string
}