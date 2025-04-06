package analytics

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

type RequestInfo struct {
	ServiceName string `json:"service_name"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Sender      string `json:"sender"`
	Timestamp   string `json:"timestamp"`
}

type AnalyticsQueue struct {
	client *redis.Client
}

// func init() {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:9000",
// 		Password: "", // No password set
// 		DB:       0,  // Use default DB
// 		Protocol: 2,  // Connection protocol
// 	})
// 	client.FlushAll(context.Background())
// 	client.Close()
// }

func NewAnalyticsQueue() *AnalyticsQueue {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:9000",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	return &AnalyticsQueue{
		client: client,
	}
}

func (w *AnalyticsQueue) PushRequestEventQueue(event RequestInfo) {
	ctx := context.Background()

	b, err := json.Marshal(event)

	if err != nil {
		log.Printf("Error marshalling request event: %s", err)
		return
	}

	_, err = w.client.RPush(ctx, "request_event", string(b)).Result()

	if err != nil {
		log.Printf("Error pushing request event to queue: %s", err)
	}
}
