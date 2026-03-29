/*
In-Memory TTL Cache
Цель:
    Реализовать in-memory кеш с индивидуальным TTL для каждого элемента
Требования:
    Поддержка базовых операций: Get/Set
    Удаление просроченных записей
*/

type Item struct {
value any
expireDate time.Time
}

type Cache struct {
mu sync.RWMutex
cache map[string]Item
}

func NewCache(maxCa) *Cache {
cache := &Cache{cache: make(map[string]Item)}
go c.worker()
return cache
}

func (c *Cache) Get(key string) (any, bool) {
c.mu.RLock()
item, ok := c.cache[key]
c.mu.RUnlock()

if ok {
return item.value, true
}

// if ttl> retun nil

return nil, false
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
c.mu.Lock()
defer c.mu.Unlock()

c.cache[key] = Item{
value: value,
expireDate: time.Now().Add(ttl)
}

}

func (c *Cache) worker(interval time.Duration) {
ticker := time.NewTicker(interval)
go func() {
defer ticker.Stop()
for range ticker.C {
c.cleanup()
}
}
}

func (c *Cache) cleanup() {
c.mu.Lock()
defer c.mu.Unlock()
for k, v := c.cache {
now := time.Now()
if now.After(v.expireDate) {
delete(c.cache, l)
}
}

}