package cache

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetKeyValue(t *testing.T) {
	err := InitRedis(Config{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	assert.Nil(t, err, "redis init fail")
	key, val := "hello", "nihaoa"
	err = SetKeyValue(context.Background(), key, val)
	assert.Nil(t, err, "set data fail")
	res, err := GetKey(context.Background(), key)
	assert.Nil(t, err, "get data fail")
	assert.Equal(t, val, res, "got data not equal original data")
}
