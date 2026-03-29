package main


— //Задача на код ревью

//Дан куcок кода, который выполняет финансовую транзакцию в приложении и выполняет нотификацию внешний системы через брокер
// Необходимо провести code review и найти проблемы, если таковые присутствуют
// Примечание: для упрощения логики некоторые части системы представлены в виде интерфейсов


type Transaction struct {
	FromUser int32 // uint64, либо uuid
	ToUser int32 // uint64, либо uuid
	Amount float64 // uint64, в копейках
}

type Account struct {
	Id int64 // uint64, id в верхнем регистре на любителя
	Balance float64 // uint64, в копейках

type BrokerMessage struct{
	FromUser int32 // uint64, либо uuid
	ToUser int32 // uint64, либо uuid
	Amount float64 // uint64, в копейках
}

type AccountService interface {
	GetAccount(sqlQuery string) Account // для запросов БД отдельный слой, не возвращаем ошибку 2м параметром, лучше указатель - 1ый параметр
	SaveAccount(acc Account) // cсылка Account, ошибка возвращается
}

type BrokerService interface {
	SendMessage(msg BrokerMessage) // ошибка должна возвращаться
}

type TransactionService struct {
	brokerService BrokerService
	accountService AccountService
}

func(t *TransactionService) SendMoney(fromAcc, toAcc int64, amount float64) { // передавал бы Account и копейки, не возвращается ошибка
	fromAcc := t.accountService.GetAccount("SELECT * FROM account WHERE id =" + fromAcc) // нет слоя БД
	toAcc := t.accountService.GetAccount("SELECT * FROM account WHERE id =" + toAcc) // нет слоя БД


	if fromAcc.Balance >= amount {
		fromAcc.Balance -= amount // обернуть в транзакцию
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

