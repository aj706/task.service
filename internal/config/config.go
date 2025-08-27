package config

import "os"

type Config struct {
    MongoURI string
    DBName   string
    Port     string
}

func Load() Config {
    return Config{
        MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
        DBName:   getEnv("DB_NAME", "tasks"),
        Port:     getEnv("PORT", "8080"),
    }
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
