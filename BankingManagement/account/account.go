package account

import (
	"fmt"
	history "hello/Assignment/github.com/BankingManagement/TransactionHistory"
	"time"

	uuid "github.com/google/uuid"
)

type accInterface interface {
	Create(id string, name string) bool
	Withdraw(amount int) bool
	Deposit(amount int) bool
	History()
	Delete()
}
type Account struct {
	Name, Id    string
	money       int
	Trandetails []history.Transaction
}

func (acc *Account) Create(id string, name string) bool {
	acc.Name = name
	acc.Id = id
	acc.money = 0
	return true
}
func (acc *Account) Withdraw(amount int) bool {

	if amount > acc.money {
		return false
	} else {
		acc.money = acc.money - amount

		var new history.Transaction
		new.Id = uuid.NewString()
		new.Amount = amount
		new.TimeStamp = time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
		new.Type = "credit"

		acc.Trandetails = append(acc.Trandetails, new)
		return true
	}
}
func (acc *Account) Deposit(amount int) bool {
	acc.money = acc.money + amount
	var new history.Transaction
	new.Id = uuid.NewString()
	new.Amount = amount
	new.TimeStamp = time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
	new.Type = "dedit"
	acc.Trandetails = append(acc.Trandetails, new)
	//fmt.Println(acc.Trandetails)

	return true
}
func (acc *Account) History() {
	fmt.Printf("|%4v|%37v|%7v|%5v|%10v|\n", "sno", "transactionId", "type", "amount", "Time")
	for i, trans := range acc.Trandetails {
		fmt.Printf("|%4v|%37v|%7v|%5v|%10v|\n", i+1, trans.Id, trans.Type, trans.Amount, trans.TimeStamp)
	}
}

func (acc *Account) Delete() {
	acc = nil
}
