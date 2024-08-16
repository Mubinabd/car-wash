package config

import (
  "fmt"
  "os"
  "strings"

  "github.com/joho/godotenv"
  "github.com/spf13/cast"
)

type Config struct {
  HTTPPort string

  MONGOHOST     string
  MONGOPORT     int
  MONGOUSER     string
  MONGOPASSWORD string

  KafkaBrokers []string

  DefaultOffset string
  DefaultLimit  string

  TokenKey string
}

func Load() Config {
  if err := godotenv.Load(); err != nil {
    fmt.Println("No .env file found")
  }

  config := Config{}

  config.HTTPPort = cast.ToString(GetOrReturnDefaultValue("HTTP_PORT", "8050"))

  config.MONGOHOST = cast.ToString(GetOrReturnDefaultValue("MONGOHOST", "mongo-db"))
  config.MONGOPORT = cast.ToInt(GetOrReturnDefaultValue("MONGOPORT", 5432))
  config.MONGOUSER = cast.ToString(GetOrReturnDefaultValue("MONGOUSER", "mubina"))
  config.MONGOPASSWORD = cast.ToString(GetOrReturnDefaultValue("MONGOPASSWORD", "1234"))

  config.KafkaBrokers = parseKafkaBrokers(GetOrReturnDefaultValue("KAFKA_BROKERS", "kafka:9092"))

  config.DefaultOffset = cast.ToString(GetOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
  config.DefaultLimit = cast.ToString(GetOrReturnDefaultValue("DEFAULT_LIMIT", "10"))
  config.TokenKey = cast.ToString(GetOrReturnDefaultValue("TokenKey", "my_secret_key"))
  return config
}

func GetOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
  val, exists := os.LookupEnv(key)

  if exists {
    return val
  }

  return defaultValue
}

func parseKafkaBrokers(brokers interface{}) []string {
  switch v := brokers.(type) {
  case string:
    return strings.Split(v, ",")
  case []string:
    return v
  default:
    return []string{"kafka:9092"}
  }
}