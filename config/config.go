package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/AleksK1NG/nats-streaming/pkg/constants"
	"github.com/spf13/viper"
)

// Config microservice config
type Config struct {
	AppVersion  string
	HTTP        HTTP
	GRPC        GRPC
	Logger      Logger
	Metrics     Metrics
	Jaeger      Jaeger
	Nats        Nats
	Redis       Redis
	MailService MailService
	PostgreSQL  PostgreSQL
}

// HTTP server config
type HTTP struct {
	Port              string
	Development       bool
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

// Logger config
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

// Redis config
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

// Nats config
type Nats struct {
	URL       string
	ClusterID string
	ClientID  string
}

// MailService config
type MailService struct {
	URL            string
	From           string
	Host           string
	Port           int
	Username       string
	Password       string
	KeepAlive      bool
	ConnectTimeout time.Duration
	SendTimeout    time.Duration
}

// PostgreSQL config
type PostgreSQL struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDBName   string
	PostgresqlSSLMode  string
	PgDriver           string
}

// GRPC gRPC service config
type GRPC struct {
	Port              string
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
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

	metricsPort := os.Getenv(constants.METRICS_PORT)
	if metricsPort != "" {
		c.Metrics.Port = metricsPort
	}

	natsUrl := os.Getenv(constants.NATS_URL)
	if natsUrl != "" {
		c.Nats.URL = natsUrl
	}
	natsClientID := os.Getenv(constants.NATS_CLIENT_ID)
	if natsClientID != "" {
		c.Nats.ClusterID = natsClientID
	}
	natsClusterID := os.Getenv(constants.CLUSTER_ID)
	if natsClusterID != "" {
		c.Nats.ClusterID = natsClusterID
	}

	redisURL := os.Getenv(constants.REDIS_URL)
	if redisURL != "" {
		c.Redis.RedisAddr = redisURL
	}
	redisPassword := os.Getenv(constants.REDIS_PASSWORD)
	if redisPassword != "" {
		c.Redis.RedisPassword = redisPassword
	}

	mailServiceURL := os.Getenv(constants.MAIL_SERVICE)
	if mailServiceURL != "" {
		c.MailService.URL = mailServiceURL
	}

	postgresPORT := os.Getenv(constants.POSTGRES_HOST)
	if postgresPORT != "" {
		c.PostgreSQL.PostgresqlHost = postgresPORT
	}
	postgresHost := os.Getenv(constants.POSTGRES_HOST)
	if postgresHost != "" {
		c.PostgreSQL.PostgresqlHost = postgresHost
	}
	postgresqlPort := os.Getenv(constants.POSTGRES_PORT)
	if postgresqlPort != "" {
		c.PostgreSQL.PostgresqlPort = postgresqlPort
	}
	postgresUser := os.Getenv(constants.POSTGRES_USER)
	if postgresUser != "" {
		c.PostgreSQL.PostgresqlUser = postgresUser
	}
	postgresPassword := os.Getenv(constants.POSTGRES_PASSWORD)
	if postgresPassword != "" {
		c.PostgreSQL.PostgresqlPassword = postgresPassword
	}
	postgresDB := os.Getenv(constants.POSTGRES_DB)
	if postgresDB != "" {
		c.PostgreSQL.PostgresqlDBName = postgresDB
	}
	postgresSSL := os.Getenv(constants.POSTGRES_SSL)
	if postgresSSL != "" {
		c.PostgreSQL.PostgresqlSSLMode = postgresSSL
	}

	jaegerHost := os.Getenv(constants.JAEGER_HOST)
	if jaegerHost != "" {
		c.Jaeger.Host = jaegerHost
	}

	gRPCPort := os.Getenv(constants.GRPC_PORT)
	if gRPCPort != "" {
		c.GRPC.Port = gRPCPort
	}

	mailHost := os.Getenv(constants.MAIL_HOST)
	if mailHost != "" {
		c.MailService.Host = mailHost
	}
	mailPort := os.Getenv(constants.MAIL_PORT)
	if mailPort != "" {
		mailPortEnv, err := strconv.Atoi(mailPort)
		if err != nil {
			return nil, err
		}
		c.MailService.Port = mailPortEnv
	}
	mailUsername := os.Getenv(constants.MAIL_USERNAME)
	if mailUsername != "" {
		c.MailService.Username = mailUsername
	}
	mailPassword := os.Getenv(constants.MAIL_PASSWORD)
	if mailUsername != "" {
		c.MailService.Password = mailPassword
	}

	return &c, nil
}
