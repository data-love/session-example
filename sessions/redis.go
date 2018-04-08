package sessions

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type redisStore struct {
	client *redis.Client
}

func NewRedisStore() Store {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to ping Redis: %v", err)
	}
	return &redisStore{
		client: client,
	}
}

func (r redisStore) Get(id string) (Session, error) {
	var session Session
	bs, err := r.client.Get(id).Bytes()
	if err != nil {
		return session, fmt.Errorf("failed to get session from redis, %s", err)
	}
	if err := json.Unmarshal(bs, &session); err != nil {
		return session, fmt.Errorf("failed to unmarshall session data, %s", err)
	}
	return session, nil
}

func (r redisStore) Set(id string, session Session) error {
	bs, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to save session to redis,  %s", err)
	}
	if err := r.client.Set(id, bs, 0).Err(); err != nil {
		return fmt.Errorf("failed to save session to redis,  %s", err)
	}
	return nil
}
