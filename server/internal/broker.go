package internal

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kinjalrk2k/Kaho-aur-Suno/lib/go/proto/broker"
)

type BrokerService struct {
	broker.UnimplementedBrokerServer
	Topics map[string][]chan *broker.Baat
}

func (b *BrokerService) RemoveChannel(topic string, channelToRemove chan *broker.Baat) {
	// Find the index of the channel to remove
	indexToRemove := -1
	for i, ch := range b.Topics[topic] {
		if ch == channelToRemove {
			indexToRemove = i
			break
		}
	}

	// Check if the channel was found and remove it if so
	if indexToRemove != -1 {
		b.Topics[topic] = append(b.Topics[topic][:indexToRemove], b.Topics[topic][indexToRemove+1:]...)
	}
}

func (b *BrokerService) Kaho(ctx context.Context, in *broker.KahoRequest) (*broker.Baat, error) {
	id := uuid.New().String()

	baatBolo := broker.Baat{
		Topic:     in.Topic,
		Id:        id,
		Message:   in.Message,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	go func() {
		for _, ch := range b.Topics[baatBolo.Topic] {
			ch <- &baatBolo
		}
	}()

	return &baatBolo, nil
}

func (b *BrokerService) Suno(req *broker.SunoRequest, stream broker.Broker_SunoServer) error {
	var ch = make(chan *broker.Baat)
	b.Topics[req.Topic] = append(b.Topics[req.Topic], ch)

	go func() {
		<-stream.Context().Done()
		ch <- nil
	}()

	for {
		baatSuno := <-ch
		if baatSuno != nil {
			stream.Send(baatSuno)
		} else {
			b.RemoveChannel(req.Topic, ch)
			break
		}
	}

	return nil
}
