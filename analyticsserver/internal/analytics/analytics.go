package analytics

import (
	"context"
	"encoding/json"

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
// 		Addr:     "redis:6379",
// 		Password: "", // No password set
// 		DB:       0,  // Use default DB
// 		Protocol: 2,  // Connection protocol
// 	})
// 	client.FlushAll(context.Background())
// 	client.Close()
// }

func NewAnalyticsQueue() *AnalyticsQueue {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-analytics:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	return &AnalyticsQueue{
		client: client,
	}
}

func (w *AnalyticsQueue) PopRequestEventQueue() (RequestInfo, error) {
	ctx := context.Background()
	val, err := w.client.LPop(ctx, "request_event").Result()

	if err != nil {
		return RequestInfo{}, err
	}

	var event RequestInfo

	err = json.Unmarshal([]byte(val), &event)

	if err != nil {
		return RequestInfo{}, err
	}

	return event, nil
}
