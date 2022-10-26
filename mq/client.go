package mq

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-stomp/stomp/v3"
	"github.com/rizefinance/rize-go-sdk/internal"
	"golang.org/x/exp/slices"
)

// Service type to store the client reference
type service struct {
	client *Client
}

// Config stores RMQ configuration values
type Config struct {
	// Message queue username
	Username string
	// Message queue password
	Password string
	// Message queue topic
	ClientID string
	// Rize infrastructure target environment. Defaults to `sandbox``
	Environment string
	// Enable debug logging
	Debug bool
}

// Client is the top-level interface for the RMQ
type Client struct {
	// All configuration values
	cfg *Config
	// Message queue AMQP endpoint
	Endpoint string
	// Connection to a STOMP server
	Connection *stomp.Conn
	// Message queue service
	MessageQueue *messageQueueService
}

// NewClient initializes the RMQ Client
func NewClient(cfg *Config) (*Client, error) {
	// Enable debug logging
	internal.EnableLogging(cfg.Debug)

	log.Println("Creating MQ client...")

	// Validate client config
	if err := cfg.validateConfig(); err != nil {
		return nil, err
	}

	rc := &Client{}
	rc.cfg = cfg
	rc.Endpoint = fmt.Sprintf("mq-%s.rizefs.com:61614", cfg.Environment)
	rc.MessageQueue = &messageQueueService{client: rc}

	return rc, nil
}

// Make sure that we have the proper configuration variables
func (cfg *Config) validateConfig() error {
	if cfg.Username == "" {
		return fmt.Errorf("Config error: Username is required")
	}

	if cfg.Password == "" {
		return fmt.Errorf("Config error: Password is required")
	}

	if cfg.ClientID == "" {
		return fmt.Errorf("Config error: ClientID is required")
	}

	if ok := slices.Contains(internal.Environments, strings.ToLower(cfg.Environment)); !ok {
		log.Printf("Environment %s not recognized. Defaulting to sandbox...\n", cfg.Environment)
		cfg.Environment = "sandbox"
	}

	return nil
}
