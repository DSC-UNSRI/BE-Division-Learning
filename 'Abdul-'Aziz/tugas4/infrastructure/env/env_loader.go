package env

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvLoader struct{}

func NewEnvLoader() *EnvLoader {
    return &EnvLoader{}
}

func (l *EnvLoader) Load() error {
    if err := godotenv.Load(); err != nil {
        return err
    }
    return nil
}

func (l *EnvLoader) Get(key string) string {
    return os.Getenv(key)
}