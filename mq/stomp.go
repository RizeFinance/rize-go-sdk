package mq

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/go-stomp/stomp/v3"
	"github.com/rizefinance/rize-go-sdk/internal"
)

// Handles all Message Queue related functionality
type messageQueueService service

// Connect will create a new STOMP connection
func (m *messageQueueService) Connect() error {
	// Create TLS connection
	tlsConn, err := tls.Dial("tcp", m.client.Endpoint, &tls.Config{})
	if err != nil {
		return err
	}

	// Connect to MQ
	conn, err := stomp.Connect(
		tlsConn,
		stomp.ConnOpt.Login(m.client.cfg.Username, m.client.cfg.Password),
		stomp.ConnOpt.HeartBeat(internal.MQSendTimeout, internal.MQReceiveTimeout),
		stomp.ConnOpt.Header("client-id", m.client.cfg.ClientID),
	)
	if err != nil {
		return err
	}

	log.Println("Connection Successful")

	m.client.Connection = conn

	return nil
}

// Subscribe will create a new STOMP topic subscription
func (m *messageQueueService) Subscribe(topic string) (*stomp.Subscription, error) {
	sub, err := m.client.Connection.Subscribe(fmt.Sprintf("/topic/%s", topic), stomp.AckAuto)
	if err != nil {
		return nil, err
	}

	log.Println("Subscribed to topic", topic)

	return sub, err
}

// Unsubscribe from the subscription and close the channel
func (m *messageQueueService) Unsubscribe(sub *stomp.Subscription) error {
	if err := sub.Unsubscribe(); err != nil {
		return err
	}

	log.Println("Unsubscribed successfully")

	return nil
}
