package cache

import (
	"context"
	"sync"
	"time"
)

/*Ï
На задачу отводится 40 минут.

Необходимо написать in-memory кэш.

Начало 15:15

Условия:

1. У кэша должен быть TTL
- должна быть возможность задать кастомный TTL для каждого элемента
2. Написать функции для работы с кешэм:
- Получение пользователя по его ID (при чтении не обновляем TTL)
- Добавление пользователя в кеш
- Удаление пользователя из кеша по его ID
1. Написать тестовые сценарии для проверки работы кеша
*/

const defaultCleanTick = time.Second

const deafultСacheCap = 1000

type User struct {
ID int64
Name string
// more fields...
}


type CacheItem struct {
data User
deadline time.Time
}


type TTLCache struct {
userMap map[int64]CacheItem
mu sync.RWMutex
}

func NewTTLCache(ctx context.Context) *TTLCache {
cache := &TTLCache{
userMap: make(map[int64]CacheItem, deafultСacheCap)
}
go cache.cleanup(ctx)
return cache
}

func (t *TTLCache) cleanup(ctx context.Context) error {
tick := time.NewTicker(defaultCleanTick)
for {
select {
case <-ctx.Done():
return ctx.Err()
case <- tick.C:
t.mu.RLock()
users := map.Values(t.userMap)
t.mu.RUnLock()
for _, u := range users {
if !time.Now().Before(u.dedline) {
actual, err := t.Get()
if err != nil {
t.Delete(u.ID)
}
}
}
}
}
}


func (t *TTLCache) Set(user User, ttl time.Duration) error {
t.mu.Lock()
defer t.mu.UnLock()
deadline := time.Now().Add(ttl)
t.userMap[user.ID] = CacheItem{
data: user,
deadline: deadline,
}
return nil

}

func (t *TTLCache) Get(userID int64) (User, error) {
t.mu.RLock()
defer t.mu.RUnLock()
if data, ok := t.userMap[userID]; ok && time.Now().Before(data.dedline) {
return data, nil
}
return User{}, errors.New("not found")
}

func (t *TTLCache) Delete(userID int64)  {
t.mu.Lock()
defer t.mu.UnLock()
del(t.userMap, userID)
}



func main() {
cache := NewTTLCache(context.Background())

user := User{ID: 1}
// first happy case
if err := cache.Set(user, time.Second*2); err != nil {
panic(err.Error())
}

actualUser,  err := cache.Get(user.ID)
if err != nil {
panic(err.Error())
}
if actualUser != user {
panic("err")
}

cache.Delete(user.ID)

actualUser,  err := cache.Get(user.ID)
if err == nil {
panic("not correct")
}
if actualUser != User{} {
panic("err")
}


// second happy case
if err := cache.Set(user, time.Second); err != nil {
panic(err.Error())
}
time.Sleep(time.Milisecond*1001)

actualUser,  err := cache.Get(user.ID)
if err == nil {
panic("not correct")
}
if actualUser != User{} {
panic("err")
}

}



