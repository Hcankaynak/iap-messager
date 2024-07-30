package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

const InProgressMessages = "inProgressMessages"

/*
RedisConnection struct
Redis Connection data.
*/
type RedisConnection struct {
	Host string
	Port string
}

/*
RedisManager struct
*/
type RedisManager struct {
	redis          *redis.Client
	ConnectionData RedisConnection
}

type MessagesInProgress struct {
	Items []string
}

/*
NewRedisManager function
Creating a new RedisManager.
*/
func NewRedisManager(connectionData RedisConnection) RedisManager {
	return RedisManager{
		ConnectionData: connectionData,
	}
}

/*
ConnectRedis function
Connecting to the Redis database.
*/
func (rm *RedisManager) ConnectRedis() {
	log.Println("Connecting to Redis...")
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", rm.ConnectionData.Host, rm.ConnectionData.Port),
	})

	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	rm.redis = rdb
	log.Println("Successfully connected to Redis!")
}

/*
SetItem function
Setting the key and value to the Redis database.
*/
func (rm *RedisManager) SetItem(key string, value interface{}) {
	ctx := context.Background()
	err := rm.redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatalf("could not set key: %v", err)
	}
}

/*
GetItem function
Getting the value of the key from the Redis database.
*/
func (rm *RedisManager) GetItem(key string) string {
	ctx := context.Background()
	val, err := rm.redis.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("could not get key: %v", err)
	}
	return val

}

/*
GetInProgressMessages function
Getting the inProgressMessages list from Redis.
*/
func (rm *RedisManager) GetInProgressMessages() MessagesInProgress {
	var messagesInProgress MessagesInProgress
	messagesInProgressValue := rm.GetItem(InProgressMessages)

	err := json.Unmarshal([]byte(messagesInProgressValue), &messagesInProgress)
	if err != nil {
		log.Fatalf("could not unmarshal data that retrieved from redis: %v", err)
	}
	return messagesInProgress
}

func (rm *RedisManager) setInProgressMessages(messagesInProgress MessagesInProgress) {
	messagesInProgressBytes, err := json.Marshal(messagesInProgress)
	if err != nil {
		log.Fatalf("could not marshal data for redis.: %v", err)
	}
	rm.SetItem(InProgressMessages, messagesInProgressBytes)
}

/*
AddInProgressMessage function
Adding a message to the inProgressMessages list. This list is used to keep track of the messages that are being processed.
First get the inProgressMessages list from Redis, then append the new message to the list, and finally set the updated list to Redis.
*/
func (rm *RedisManager) AddInProgressMessage(message string) {
	messagesInProgress := rm.GetInProgressMessages()
	messagesInProgress.Items = append(messagesInProgress.Items, message)
	rm.setInProgressMessages(messagesInProgress)
}

func (rm *RedisManager) RemoveFromInProgressMessages(message string) {
	messagesInProgress := rm.GetInProgressMessages()
	messagesInProgress.Items = remove(messagesInProgress.Items, message)
	rm.setInProgressMessages(messagesInProgress)
}

func (rm *RedisManager) CreateEmptyListForInProgress() {
	messages := []string{}
	rm.setInProgressMessages(MessagesInProgress{Items: messages})
}

func remove(slice []string, value string) []string {
	result := []string{}
	for _, item := range slice {
		if item != value {
			result = append(result, item)
		}
	}
	return result
}
