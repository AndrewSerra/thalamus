package lookup

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type LookupWorker struct {
	client *redis.Client
}

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	client.FlushAll(context.Background())
}

func NewLookupWorker() *LookupWorker {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	return &LookupWorker{
		client: client,
	}
}

func (w *LookupWorker) GetAddresses(serviceName string) []string {
	ctx := context.Background()
	val := w.client.LRange(ctx, serviceName, 0, -1).Val()
	return val
}

func (w *LookupWorker) SetAddress(serviceName string, address string) int64 {
	ctx := context.Background()
	size := w.client.RPush(ctx, serviceName, address)
	log.Printf("Setting new address for service %s: %s, number of available workers: %d", serviceName, address, size.Val())
	return size.Val()
}

func (w *LookupWorker) DeleteAddress(serviceName string, address string) {
	ctx := context.Background()
	w.client.LRem(ctx, serviceName, 1, address).Val()
}
