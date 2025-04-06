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
		Addr:     "redis",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	client.FlushAll(context.Background())
	client.Close()
}

func NewLookupWorker() *LookupWorker {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis",
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
	val, err := w.client.LRange(ctx, serviceName, 0, -1).Result()

	if err != nil {
		log.Printf("Error getting addresses for service %s: %s", serviceName, err)
		return []string{}
	}
	return val
}

func (w *LookupWorker) SetAddress(serviceName string, address string) int64 {
	ctx := context.Background()
	res, err := w.client.RPush(ctx, serviceName, address).Result()

	if err != nil {
		log.Printf("Error setting address for service %s: %s", serviceName, err)
		return -1
	}
	log.Printf("Setting new address for service %s: %s, number of available workers: %d", serviceName, address, res)
	return res
}

func (w *LookupWorker) DeleteAddress(serviceName string, address string) {
	ctx := context.Background()
	res, err := w.client.LRem(ctx, serviceName, 1, address).Result()

	if err != nil {
		log.Printf("Error deleting address for service %s: %s", serviceName, err)
		return
	}
	log.Printf("Deleting address for service %s: %s, number of available workers: %d", serviceName, address, res)
}
