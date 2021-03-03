package config

import (
	"log"
	"os"
	"time"

	"github.com/AleksK1NG/nats-streaming/pkg/constants"
	"github.com/spf13/viper"
)

// Config
type Config struct {
	AppVersion  string
	HTTP        HTTP
	Logger      Logger
	Metrics     Metrics
	Jaeger      Jaeger
	Nats        Nats
	Redis       Redis
	MailService MailService
	PostgreSQL  PostgreSQL
}

// HTTP
type HTTP struct {
	Port              string
	Development       bool
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

// Logger
type Logger struct {
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Metrics config
type Metrics struct {
	Port        string
	URL         string
	ServiceName string
}

// Jaeger config
type Jaeger struct {
	Host        string
	ServiceName string
	LogSpans    bool
}

type Redis struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultDB string
	MinIdleConn    int
	PoolSize       int
	PoolTimeout    int
	DB             int
}

type Nats struct {
	URL       string
	ClusterID string
	ClientID  string
}

type MailService struct {
	URL  string
	From string
}

// Postgresql config
type PostgreSQL struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDBName   string
	PostgresqlSSLMode  string
	PgDriver           string
}

func exportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// ParseConfig Parse config file
func ParseConfig() (*Config, error) {
	if err := exportConfig(); err != nil {
		return nil, err
	}

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	httpPort := os.Getenv(constants.SERVER_PORT)
	if httpPort != "" {
		c.HTTP.Port = httpPort
	}

	natsUrl := os.Getenv(constants.NATS_URL)
	if httpPort != "" {
		c.Nats.URL = natsUrl
	}
	natsClientID := os.Getenv(constants.NATS_CLIENT_ID)
	if httpPort != "" {
		c.Nats.ClusterID = natsClientID
	}
	natsClusterID := os.Getenv(constants.CLUSTER_ID)
	if httpPort != "" {
		c.Nats.ClusterID = natsClusterID
	}

	redisURL := os.Getenv(constants.REDIS_URL)
	if redisURL != "" {
		c.Redis.RedisAddr = redisURL
	}
	redisPassword := os.Getenv(constants.REDIS_PASSWORD)
	if redisURL != "" {
		c.Redis.RedisPassword = redisPassword
	}

	mailServiceURL := os.Getenv(constants.MAIL_SERVICE)
	if redisURL != "" {
		c.MailService.URL = mailServiceURL
	}

	postgresPORT := os.Getenv(constants.POSTGRES_HOST)
	if redisURL != "" {
		c.PostgreSQL.PostgresqlHost = postgresPORT
	}
	postgresHost := os.Getenv(constants.POSTGRES_HOST)
	if redisURL != "" {
		c.PostgreSQL.PostgresqlHost = postgresHost
	}
	postgresqlPort := os.Getenv(constants.POSTGRES_PORT)
	if redisURL != "" {
		c.PostgreSQL.PostgresqlPort = postgresqlPort
	}
	postgresUser := os.Getenv(constants.POSTGRES_USER)
	if redisURL != "" {
		c.PostgreSQL.PostgresqlUser = postgresUser
	}
	postgresPassword := os.Getenv(constants.POSTGRES_PASSWORD)
	if redisURL != "" {
		c.PostgreSQL.PostgresqlPassword = postgresPassword
	}
	postgresDB := os.Getenv(constants.POSTGRES_DB)
	if redisURL != "" {
		c.PostgreSQL.PostgresqlDBName = postgresDB
	}
	postgresSSL := os.Getenv(constants.POSTGRES_SSL)
	if redisURL != "" {
		c.PostgreSQL.PostgresqlSSLMode = postgresSSL
	}

	jaegerHost := os.Getenv(constants.JAEGER_HOST)
	if redisURL != "" {
		c.Jaeger.Host = jaegerHost
	}

	return &c, nil
}
