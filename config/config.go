package config

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"

	sharedGRPC "github.com/wwi21seb-projekt/alpha-shared/grpc"
)

// AlphaConfig holds all possible configuration values
type AlphaConfig struct {
	DatabaseConfig   DatabaseConfig
	ServiceEndpoints ServiceEndpoints
	JaegerConfig     JaegerConfig
	Port             string
	Environment      string
	GRPCClients      GRPCClients
}

type DatabaseConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	SchemaName       string
}

type ServiceEndpoints struct {
	UserServiceURL         string
	PostServiceURL         string
	MailServiceURL         string
	NotificationServiceURL string
	ChatServiceURL         string
}

type JaegerConfig struct {
	TracesEndpoint string
	ServiceName    string
}

type GRPCClients struct {
	UserService         grpc.ClientConnInterface
	PostService         grpc.ClientConnInterface
	MailService         grpc.ClientConnInterface
	NotificationService grpc.ClientConnInterface
	ChatService         grpc.ClientConnInterface
}

// LoadConfig loads environment variables into an AlphaConfig struct
func LoadConfig() (*AlphaConfig, error) {
	config := &AlphaConfig{
		DatabaseConfig: DatabaseConfig{
			PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
			PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
			PostgresDB:       getEnv("POSTGRES_DB", "mydatabase"),
			PostgresUser:     getEnv("POSTGRES_USER", "myuser"),
			PostgresPassword: getEnv("POSTGRES_PASSWORD", "mypassword"),
			SchemaName:       getEnv("SCHEMA_NAME", "public"),
		},
		ServiceEndpoints: ServiceEndpoints{
			UserServiceURL:         getEnv("USER_SERVICE_URL", "http://user-service:50051"),
			PostServiceURL:         getEnv("POST_SERVICE_URL", "http://post-service:50052"),
			MailServiceURL:         getEnv("MAIL_SERVICE_URL", "http://mail-service:50053"),
			NotificationServiceURL: getEnv("NOTIFICATION_SERVICE_URL", "http://notification-service:50054"),
			ChatServiceURL:         getEnv("CHAT_SERVICE_URL", "http://chat-service:50055"),
		},
		JaegerConfig: JaegerConfig{
			TracesEndpoint: getEnv("TRACES_ENDPOINT", "jaeger-collector.observability:4317"),
			ServiceName:    getEnv("SERVICE_NAME", "alpha"),
		},
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}

	// Add any necessary validation here
	if config.Port == "" {
		return nil, fmt.Errorf("missing required environment variable: PORT")
	}

	return config, nil
}

// getEnv reads an environment variable and returns its value or a default value if not set
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}

// InitializeClients initializes the gRPC clients based on the configuration.
func (cfg *AlphaConfig) InitializeClients(logger *zap.Logger) error {
	dialOpts := sharedGRPC.NewClientOptions(logger)

	var err error

	// Initialize each client connection based on the service URL in the configuration
	if cfg.GRPCClients.UserService, err = initializeClient(cfg.ServiceEndpoints.UserServiceURL, dialOpts); err != nil {
		return err
	}
	if cfg.GRPCClients.PostService, err = initializeClient(cfg.ServiceEndpoints.PostServiceURL, dialOpts); err != nil {
		return err
	}
	if cfg.GRPCClients.MailService, err = initializeClient(cfg.ServiceEndpoints.MailServiceURL, dialOpts); err != nil {
		return err
	}
	if cfg.GRPCClients.NotificationService, err = initializeClient(cfg.ServiceEndpoints.NotificationServiceURL, dialOpts); err != nil {
		return err
	}
	if cfg.GRPCClients.ChatService, err = initializeClient(cfg.ServiceEndpoints.ChatServiceURL, dialOpts); err != nil {
		return err
	}

	return nil
}

// initializeClient creates a gRPC client connection if the service URL is provided.
func initializeClient(serviceURL string, dialOpts []grpc.DialOption) (grpc.ClientConnInterface, error) {
	if serviceURL == "" {
		return nil, nil
	}
	conn, err := grpc.NewClient(serviceURL, dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to service at %s: %w", serviceURL, err)
	}
	return conn, nil
}
