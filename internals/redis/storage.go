package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr, user, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: user,
		Password: password,
		DB:       db,
	})
	return &RedisClient{client: client}
}

func (r *RedisClient) AllowToken(token string, limit, ttl int) (bool, error) {
	ctx := context.Background()
	ttlSeconds := time.Second * time.Duration(ttl)

	val, err := r.client.Get(ctx, token).Result()
	if errors.Is(err, redis.Nil) {
		err = r.client.Set(ctx, token, 1, ttlSeconds).Err()
		return true, err
	}

	count, _ := strconv.Atoi(val)
	if count >= limit {
		return false, nil
	}

	err = r.client.Incr(ctx, token).Err()
	return true, err
}

func (r *RedisClient) AllowIp(ip string, limit, ttl int) (bool, error) {
	ctx := context.Background()
	//limit := 10
	ttlSeconds := time.Second * time.Duration(ttl)

	val, err := r.client.Get(ctx, ip).Result()
	if errors.Is(err, redis.Nil) {
		err = r.client.Set(ctx, ip, 1, ttlSeconds).Err()
		return true, err
	}

	count, _ := strconv.Atoi(val)
	if count >= limit {
		return false, nil
	}

	err = r.client.Incr(ctx, ip).Err()
	return true, err
}
