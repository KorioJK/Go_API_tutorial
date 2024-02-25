package initializers

import "github.com/joho/godotenv"

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		panic("init env error")
	}
}
