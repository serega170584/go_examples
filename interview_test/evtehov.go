/*
In-Memory TTL Cache
Цель:
    Реализовать in-memory кеш с индивидуальным TTL для каждого элемента
Требования:
    Поддержка базовых операций: Get/Set
    Удаление просроченных записей
*/

type item struct {
value     any
ttl       time.Time
}

type TtlCache struct {
storage map[string]item
}

func NewTtlCache(size int) TtlCache {
res := TtlCache{
storage: make(map[string]item, size)
}

go func(){
t := time.NewTicker(2000)
for _ <- t.C {
for k, v := range res.storage {
if time.Since(v.ttl) >= 0 {
delete(res, k)
}
}
}
}()

return res
}

func (t TtlCache)Get(key string) (item, bool) {
if i, ok := t.storage[key]; !ok {
return nil, false
}

if time.Since(i.ttl) >= 0 {
return nil, false
}

return i, true
}

func (t TtlCache)Set(key string, val any, ttl time.Duration) {
t.storage[key] = item{
value:     val,
ttl:       time.Now().Add(ttl),
}
}