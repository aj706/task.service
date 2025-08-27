package config

import "os"

type Config struct {
    MongoURI string
    DBName   string
    Port     string
}

func Load() Config {
    return Config{
        MongoURI: getenv("MONGO_URI", "mongodb://localhost:27017"),
        DBName:   getenv("DB_NAME", "users"),
        Port:     getenv("PORT", "8081"),
    }
}

func getenv(k, d string) string {
    if v := os.Getenv(k); v != "" {
        return v
    }
    return d
}
