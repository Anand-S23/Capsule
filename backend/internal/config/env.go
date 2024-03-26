package config

import (
	"fmt"
	"log"
	"os"
)

type EnvVars struct {
    PRODUCTION       bool
    PORT             string
    DB_URI           string
    FE_URI           string
    JWT_SECRET       []byte
    COOKIE_HASH_KEY  []byte
    COOKIE_BLOCK_KEY []byte
}

func LoadEnv() (*EnvVars, error) {
    envMode     := GetEnv("MODE", "development")
    port        := GetEnv("PORT", "8080")
    dbUri       := GetEnvOrPanic("DB_URI")
    frontendURI := GetEnvOrPanic("FE_URI")

    secret   := GetEnvOrPanic("JWT_SECRET")
    hashKey  := GetEnvOrPanic("COOKIE_HASH_KEY")
    blockKey := GetEnvOrPanic("COOKIE_BLOCK_KEY")

    return &EnvVars {
        PRODUCTION: (envMode == "production"),
        PORT: port,
        DB_URI: dbUri,
        FE_URI: frontendURI,
        JWT_SECRET: []byte(secret),
        COOKIE_HASH_KEY: []byte(hashKey),
        COOKIE_BLOCK_KEY: []byte(blockKey),
    }, nil
}

func GetEnv(env, defaultValue string) string {
	variable := os.Getenv(env)
	if variable == "" {
		return defaultValue
	}

	return variable
}

func GetEnvOrPanic(env string) string {
	variable := os.Getenv(env)
	if variable == "" {
        message := fmt.Sprintf("Must provide %s variable in .env file", env)
        log.Fatal(message)
	}

	return variable
} 

