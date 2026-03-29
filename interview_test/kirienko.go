func longFunction() int {
// что-то долгое
}

// реализовать, чтобы работала отмена по дедлайну
func longFunctionWrapper(ctx context.Context) int {
ch := make(chan int, 1)

go func() {
ch <- longFunction()
close(ch)
}()

var res int
select {
case <-ctx.Done():
case res <- ch:
}

return res
}

// ---


// Дан кусок кода, который выполняет финансовую транзакцию в приложении и выполняет нотификацию внешний системы через брокер
// Необходимо провести code review и найти проблемы, если таковые присутствуют
// Примечание: для упрощения логики некоторые части системы представлены в виде интерфейсов

type Transaction struct {
FromUser int32
ToUser int32
Amount float64 // Поменять на decimal/int64 в копейках
}

type Account struct {
Id int64
Balance float64  // Поменять на decimal/int64 в копейках
}

type BrokerMessage struct{
FromUser int32
ToUser int32
Amount float64  // Поменять на decimal/int64 в копейках
}

type AccountService interface {
GetAccount(sqlQuery string) Account
SaveAccount(acc Account)
}

type KafkaService interface {
SendMessage(msg BrokerMessage)
}

type TransactionService struct {
kafka KafkaService
accountService AccountService
}

// Прокинь context.Context во все функции
func(t *TransactionService) SendMoney(fromAcc, toAcc string, amount float64) {
// Перепиши на repository (инкапсулируй запрос на инфра слое)
// Запрос нужно переписать: на SELECT конкретных полей
fromAccInfo := t.accountService.GetAccount("SELECT * FROM account WHERE id =" + fromAcc) // <- инъекция
// Нет обработки ошибок
toAccInfo := t.accountService.GetAccount("SELECT * FROM account WHERE id =" + toAcc) // <- инъекция

if fromAcc.Balance >= amount {
fromAcc.Balance -= amount

toAcc.Balance += amount

// Нужна транзакция на уровне БД
t.accountService.SaveAccount(fromAcc)

t.accountService.SaveAccount(toAcc)

// Добавь outbox!!!
t.kafka.SendMessage(BrokerMessage{
FromUser: fromAcc
ToUser: toAcc
Amount: amount
})
}
}

// ---

// Имеется сторонний сервис погоды (его имитация — это функция WeatherForecast).
// Сторонний сервис работает за секунду, что для нас долго.
// На наш сервис идёт большая нагрузка. Как доработать текущую реализацию?

// 1. Предложить и реализовать решение.
// 2. Дополнительное задание: сторонний сервис может давать данные не только по одному городу.
// Функция WeatherForecast(city).
// Доработать реализацию из первого пункта с учётом этого факта. Городов - много!

import (
"math/rand"
"time"
"fmt"
"net/http"
)

func WeatherForecast() int {
time.Sleep(1 * time.Second)
return rand.Intn(70) - 30
}

func main() {
var cachedRes atomic.Int64
cachedRes.Store(WeatherForecast())

go func() {
ticker:=time.NewTicker(1*time.Second)

for {
select {
case <-ticker.C:
cachedRes.Store(WeatherForecast())
}
}
}()

http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, `{"temperature":%d}`+"\n", cachedRes.Load())
})

err := http.ListenAndServe(":3333", nil)
if err != nil {
panic(err)
}
}
