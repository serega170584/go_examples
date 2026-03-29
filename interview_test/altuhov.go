
— Задача на код ревью

//Дан куcок кода, который выполняет финансовую транзакцию в приложении и выполняет нотификацию внешний системы через брокер
// Необходимо провести code review и найти проблемы, если таковые присутствуют
// Примечание: для упрощения логики некоторые части системы представлены в виде интерфейсов


type Transaction struct {
FromUser int32
ToUser int32
Amount float64
}

type Account struct {
Id int64
Balance float64
}

type BrokerMessage struct{
FromUser int32
ToUser int32
Amount float64
}

type AccountService interface {
GetAccount(sqlQuery string) Account
SaveAccount(acc Account)
}

type BrokerService interface {
SendMessage(msg BrokerMessage)
}

type TransactionService struct {
brokerService BrokerService
accountService AccountService
}

func(t *TransactionService) SendMoney(fromAcc, toAcc int64, amount float64) {
fromAcc := t.accountService.GetAccount("SELECT * FROM account WHERE id =" + fromAcc)
toAcc := t.accountService.GetAccount("SELECT * FROM account WHERE id =" + toAcc)


if fromAcc.Balance >= amount {
fromAcc.Balance -= amount
toAcc.Balance += amount
t.accountService.SaveAccount(fromAcc)
t.accountService.SaveAccount(toAcc)
t.brokerService.SendMessage(BrokerMessage{
FromUser: fromAcc
ToUser: toAcc
Amount: amount
})

}
}
