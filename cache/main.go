package main

import (
	"fmt"
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

type User struct {
	ID   int64
	Name string
	// more fields...
}

type CacheItem struct {
	data     *User
	deadline time.Time
}

type TTLCache struct {
	userMap map[int64]CacheItem
	mu      sync.RWMutex
}

func NewTTLCache() *TTLCache {
	cache := &TTLCache{
		userMap: make(map[int64]CacheItem),
	}
	return cache
}

func (t *TTLCache) Set(user *User, ttl time.Duration) {
	t.mu.Lock()
	defer t.mu.Unlock()
	deadline := time.Now().Add(ttl)
	t.userMap[user.ID] = CacheItem{
		data:     user,
		deadline: deadline,
	}
}

func (t *TTLCache) Get(userID int64) *User {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if _, ok := t.userMap[userID]; !ok {
		return nil
	}

	cacheItem := t.userMap[userID]

	if time.Now().Before(cacheItem.deadline) {
		return cacheItem.data
	}

	return nil
}

func (t *TTLCache) Delete(userID int64) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.userMap, userID)
}

func main() {
	cache := NewTTLCache()

	user := &User{ID: 1, Name: "TEST"}
	// first happy case
	cache.Set(user, 2*time.Second)

	time.Sleep(time.Second)

	cur := cache.Get(1)
	if cur == user {
		fmt.Println("Success")
	} else {
		fmt.Println("Fail")
	}

	time.Sleep(2 * time.Second)
	cur = cache.Get(1)
	if cur == nil {
		fmt.Println("Success")
	} else {
		fmt.Println("Fail")
	}

	cache.Set(user, 2*time.Second)
	cache.Delete(1)

	if cache.Get(1) == nil {
		fmt.Println("Success")
	} else {
		fmt.Println("Fail")
	}

	cache.Set(user, time.Second)
	user.Name = "Mike"
	cache.Set(user, time.Second)

	if cache.Get(1).Name == "Mike" {
		fmt.Println("Success")
	} else {
		fmt.Println("Fail")
	}

	cache.Set(user, time.Second)
	user.Name = "Nick"
	cache.Set(user, 2*time.Second)

	time.Sleep(time.Second)
	cur = cache.Get(1)
	if cur == nil {
		fmt.Println("Fail")
	} else if cur.Name == "Nick" {
		fmt.Println("Success")
	}
}
