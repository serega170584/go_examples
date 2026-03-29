Функция isPalindrome
// напиши реализацию палиндрома

func isPalindrome(str string) bool {
n := len([]rune(str))
runes := []rune(str)
// n := utf8.RuneCountInString(str)
for i := 0; i < n / 2; i++ {
if runes[i] != runes[n - i - 1] {
return false
}
}
retuen true
}

// isPalindrome("abba") -> true
// isPalindrome("aba") -> true
// isPalindrome("abab") -> false

isPalindrome("жаж")



slice[0] = "newvalue"




// ---- internal/gateway/grpc/basket/add_item.go

/*
 * Реализовать метод grpc-сервера, который:
 * - добавляет товар в корзину
 * - оформляет корзину
 *
 * В оформленные корзины (basket.Status=="ordered") изменения вносить нельзя.
 * При изменении состава корзины надо пересчитывать basket.Total=sum(count*price)
 * Все элементы в корзине должны быть уникальны по ключу ProductID
 *
 * Для оформления корзины необходимо:
 * - сменить ее статус на ordered
 * - сообщить возможным потребителям с помощью сообщений в брокер.
 */

type AddItemRequest struct {
UserID    uint64
ProductID uint64
Price     uint64
Count     uint64
}

func (bs *BasketServer) AddItemAndOrder(ctx context.Context, req *AddItemRequest) (*EmptyResponse, error) {
basket, err := bs.repo.Load(ctx, req.UserID)
if err != nil {
return nil, err
}

// place your code
if basket == nil || basket.Status =="ordered" {
basket = &Backet{
UserID: req.UserID
Items:  make([]BasketItem, 0)
Status: "new"
Total: 0
}
}

isCurrentProduct := false
itemsSet := make([uint64]struct{})
for i, item := range basket.Items {
itemsSet[item.ProductId] = struct{}{}
if item.ProductId == req.ProductID {
basket.Items[i].Count += req.Count
isCurrentProduct = true
}
}

if !isCurrentProduct {
basket.Items = append(basket.Items, BasketItem{
BasketID: basket.ID,
ProductID: req.ProductID,
Count: req.Count,
Price: req.Price,
})
}
basket.Total += req.Count * req.Price
basket.Status = "ordered"

// begin
err = bs.repo.Save(ctx, basket)
if err == nil {
// rollback
return nil, err
}

err = bs.producer.SendMessage(ctx, basket)
if err != nil {
// rollback
return nil, err
}

// commit

// transactional outbox

return &EmptyResponse{}, nil
}

// ---- internal/service/entity/basket/basket.go
type Basket struct {
ID     uint64
UserID uint64
Items  []BasketItem
Status string
Total  uint64
CreatedAt date
}

// ---- internal/service/entity/basket/item.go
type BasketItem struct {
BasketID  uint64
ProductID uint64
Count     uint64
Price     uint64
}


// получить пользователей, которые купили товаров на 30000 и более за последний месяц
// завершенные корзины status=ordered
// table=baskets


select distinct user_id, sum(total) as amount
from baskets where status == "ordered" and created_at > (date(now()) - 30)
group by user_id having amount > 30000;




// ---- internal/gateway/grpc/basket/dependencies.go
type BasketRepository interface {
Load(ctx context.Context, userID uint64) (*Basket, error)
Save(ctx context.Context, b *Basket) error
}

type CheckoutProducer interface {
SendMessage(ctx context.Context, basket Basket) error
}

// ---- internal/gateway/grpc/basket/server.go
type BasketServer struct {
repo     BasketRepository
producer CheckoutProducer
}

// ---- pkg/server/grpc/basket.grpc.pb.go
type BasketServiceServer interface {
AddItemAndOrder(context.Context, *AddItemRequest) (*EmptyResponse, error)
}



type EmptyResponse struct{}

