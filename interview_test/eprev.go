type Order struct {
ID        int64
UserID    int64
Status    string
Amount    int64
CreatedAt time.Time
}

type Repository interface {
GetOrder(ctx context.Context, id int64) (*Order, error)
SaveOrder(ctx context.Context, order *Order) error
UpdateUserBalance(ctx context.Context, userID int64, delta int64) error
}

type EventProducer interface {
Send(ctx context.Context, topic string, key string, value []byte) error
}

type Service struct {
repo     Repository
producer EventProducer
}

func (s *Service) PayOrder(ctx context.Context, orderID int64) error {
order, err := s.repo.GetOrder(ctx, orderID)
if err != nil {
return nil
}

// enum
if order.Status == "paid" {
return nil
}

// orderNew := -order.Amount

// begin;
// update order set amount = amount - xx where user_id = yy
// end;

// select ... for update

if err := s.repo.UpdateUserBalance(ctx, order.UserID, -order.Amount); err != nil {
return err
}

// enum
order.Status = "paid"
if err := s.repo.SaveOrder(ctx, order); err != nil {
return err
}

event := map[string]any{
"order_id": order.ID,
"status":   order.Status,
"ts":       time.Now().Unix(),
}

// err
data, _ := json.Marshal(event)

eGrp, eCtx := errgroup.WithContext(ctx)

go func() {
// err
// s.outbox.Send(eCtx )
err := s.producer.Send(context.Background(), "orders.paid", strconv.FormatInt(order.ID, 10), data)

}()

return nil
}