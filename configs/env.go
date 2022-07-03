package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVar string

const (
	MONGOURI   EnvVar = "MONGOURI"
	JWT_SECRET EnvVar = "JWT_SECRET"
)

var envs = [...]string{
	string(MONGOURI),
}

func GetEnv(env EnvVar) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env - ", string(env), err)
	}
	envValues := map[string]string{}
	for _, env := range envs {
		val := os.Getenv(env)
		if val == "" {
			fmt.Printf("Env not found - %s", string(env))
			log.Fatal("Cannot find a particular ENV")
		}
		envValues[env] = val
	}
	return envValues[string(env)]
}
