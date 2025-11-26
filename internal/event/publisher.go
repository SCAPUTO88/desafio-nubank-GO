package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

type EventPublisher interface {
	Publish(topicID string, message interface{}) error
	Close()  
}

type GCPPublisher struct {
	client *pubsub.Client
}

func NewGCPPublisher(projectID string) (*GCPPublisher, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar cliente pubsub: %v", err)
	}

	return &GCPPublisher{client: client}, nil
}

func (p *GCPPublisher) Publish(topicID string, message interface{}) error {
	ctx := context.Background()

	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("falha ao serializar mensagem: %v", err)
	}

	topic := p.client.Topic(topicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return fmt.Errorf("falha ao verificar existencia do topico: %v", err)
	}

	if !exists {
		log.Printf("Criando tópicos: %s", topicID)
		topic, err = p.client.CreateTopic(ctx, topicID)
		if err != nil {
			return fmt.Errorf("falha ao criar topico: %v", err)
		}
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("falha ao publicar mensagem: %v", err)
	}

	log.Printf("Mensagem publicada no tópico %s com ID: %s", topicID, id)
	return nil
}

func (p *GCPPublisher) Close() {
	p.client.Close()
}
