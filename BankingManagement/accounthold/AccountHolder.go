package holder

import (
	"fmt"
	account "hello/Assignment/github.com/BankingManagement/account"

	"github.com/google/uuid"
)

var holders = make(map[string]*account.Account)

func CreateAccount(s string) {
	var id string = uuid.NewString()
	var acc account.Account
	success := acc.Create(id, s)
	if success {
		fmt.Println("Your account successfully created")
		fmt.Printf("Your id is %v\n", id)
		fmt.Println()
		fmt.Println()

		holders[id] = &acc
	} else {
		fmt.Println("sorry account cannot be created")
	}
}

func WithdrawMoney(id string, amount int) {
	var acc *account.Account
	acc, ok := holders[id]
	if ok {
		success := acc.Withdraw(amount)
		if success {
			fmt.Println("Transaction Successful")
		} else {
			fmt.Println("Insufficeint balance")
		}
	} else {
		fmt.Println("wrong id re enter your id")
	}
	fmt.Println()
	fmt.Println()
}

func DepositMoney(id string, amount int) {
	var acc *account.Account
	acc, ok := holders[id]
	if ok {
		success := acc.Deposit(amount)
		if success {
			fmt.Println("Transaction Successful")
		} else {
			fmt.Println("Transaction Unsuccessful")
		}
	} else {
		fmt.Println("wrong id re enter your id")
	}
	fmt.Println()
	fmt.Println()
}

func History(id string) {
	var acc *account.Account
	acc, ok := holders[id]
	if ok {
		acc.History()
	} else {
		fmt.Println("wrong id re enter your id")
	}
	fmt.Println()
	fmt.Println()
}

func DeleteAccount(id string) {
	var acc *account.Account
	acc, ok := holders[id]
	if ok {
		acc.Delete()
		delete(holders, id)
		fmt.Println("account successfully deleted")
	} else {
		fmt.Println("wrong id re enter your id")
	}
	fmt.Println()
	fmt.Println()
}
