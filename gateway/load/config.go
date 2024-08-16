package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTPPort string

	MongoHost     string
	MongoUser     string
	MongoPassword string
	MongoDatabase string

	KafkaBrokers     []string

	DefaultOffset string
	DefaultLimit  string

	// JWT
	JWTSecretKey string
	JWTExpiry    int

	LOG_PATH        string
	ProductAddr string
	
}

// Load ...
func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found", err)
	}

	config := Config{}

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8070"))

	config.MongoHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "mongo-db"))
	config.MongoUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "mubina"))
	config.MongoPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1234"))
	config.MongoDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "carwash"))


	config.KafkaBrokers = cast.ToStringSlice(getOrReturnDefaultValue("KAFKA_BROKERS", []string{"kafka:9092"}))


	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))


	// JWT Configuration
	config.JWTSecretKey = cast.ToString(getOrReturnDefaultValue("JWT_SECRET_KEY", "your_secret_key"))
	config.JWTExpiry = cast.ToInt(getOrReturnDefaultValue("JWT_EXPIRY", 60))

	config.ProductAddr = cast.ToString(getOrReturnDefaultValue("PRODUCT_port", "product:8085"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}