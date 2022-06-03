package main

import (
	"github.com/go-redis/redis"
)

func SetSession(rHandle *redis.Client, email string) (string, error) {
	token, err := GenerateRandomString(64)

	if err != nil {
		return "", err
	}

	err = rHandle.Set(token, email, SESSION_DURATION).Err()

	if err != nil {
		return "", err
	}

	return token, err
}

func GetSession() {}
