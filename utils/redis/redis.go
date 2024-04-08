package redis

import (
    "fmt"
    "time"
    "encoding/json"
    "context"
    "sync"

    "github.com/redis/go-redis/v9"

    "gintemplate/utils/config"
)

var lock *sync.RWMutex
var cache *redis.Client
var ctx context.Context
var cursor uint64

func init() {
    lock = new(sync.RWMutex)
    lock.Lock()
    defer lock.Unlock()
    ctx = context.Background()
    cache = redis.NewClient(&redis.Options{
        Addr: config.RedisURL,
        Password: config.RedisPasswd,
        DB: 0,
    })
    _, err := cache.Ping(ctx).Result()
    if err != nil {
        panic(err)
    }
}

func Set(key string, value any, expiration time.Duration) (err error) {
    lock.Lock()
    defer lock.Unlock()
    var data []byte
    data, err = json.Marshal(value)
    if err != nil {
        return
    }
    status := cache.Set(ctx, key, string(data), expiration)
    err = status.Err()
    return
}

func Get(key string, value any) (err error) {
    lock.RLock()
    defer lock.RUnlock()
    var data string
    data, err = cache.Get(ctx, key).Result()
    if err != nil {
        return
    }
    err = json.Unmarshal([]byte(data), value)
    return
}

func Scan(match string) (keys []string, err error) {
    lock.RLock()
    defer lock.RUnlock()
    keys, cursor, err = cache.Scan(ctx, cursor, match, 0).Result()
    return
}

func Clear() {
    lock.Lock()
    defer lock.Unlock()
    cache.FlushDB(ctx)
}

func Close() {
    lock.Lock()
    defer lock.Unlock()
    cache.Close()
}

