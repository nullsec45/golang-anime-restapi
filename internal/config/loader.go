package config

import(
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func Get() * Config{
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error when loading file configuration", err.Error())
	}

	expInt, _ := strconv.Atoi(os.Getenv("JWT_EXP"))

	return &Config{
		App:App{
			AppEnv:os.Getenv("APP_ENV"),
		},
		Server:Server{
			Host:os.Getenv("SERVER_HOST"),
			Port:os.Getenv("SERVER_PORT"),
			Asset:os.Getenv("SERVER_ASSET_URL"),
			AssetPrivate:os.Getenv("SERVER_ASSET_URL_PRIVATE"),
		},
		Database: Database{
			Host:os.Getenv("DB_HOST"),
			Port:os.Getenv("DB_PORT"),
			Name:os.Getenv("DB_NAME"),
			User:os.Getenv("DB_USER"),
			Pass:os.Getenv("DB_PASS"),
			Tz:os.Getenv("DB_TZ"),
		},
		Jwt: Jwt {
			Key:os.Getenv("JWT_KEY"),
			Exp:expInt,
		},
		Storage: Storage{	
			BasePath:os.Getenv("STORAGE_PATH"),
		},
		Redis: Redis{
			Host:os.Getenv("REDIS_HOST"),
			Port:os.Getenv("REDIS_PORT"),
			Password:os.Getenv("REDIS_PASSWORD"),
		},
	}
}