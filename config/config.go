package config

import (
    "os"
)

type Config struct {
    Env      string
    Address  string
    DBFile    string
    DBMaxCon int
}

func LoadFromEnv() *Config {
		
    env := getEnv("APP_ENV", "development")
    addr := getEnv("APP_ADDRESS", ":8080")
    dbfile := getEnv("SQLITE_FILE", "meetings.sqlite")

    return &Config{
        Env:     env,
        Address: addr,
        DBFile:   dbfile,
    }
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
