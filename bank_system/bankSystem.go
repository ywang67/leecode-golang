package banksystem

import (
	"errors"
	"sync"
)

type Account struct {
	ID        string
	CreatedAt int
	amount    int
}

type Banksystem struct {
	accounts []*Account
	mu       *sync.RWMutex
}

type BankSystemInterface interface {
	CreateAccount(timestamp int, accountId string) bool
	Deposit(timestamp, amount int, accountId string) (int error)
	Transfer(timestamp, amount int, FromAccountId, toAccountId string) error
	Withdraw(timestamp, amount int, acountId string) (int, error)
}

func NewBankSystem() *Banksystem {
	return &Banksystem{
		accounts: []*Account{},
		mu:       new(sync.RWMutex),
	}
}

func (bs *Banksystem) CreateAccount(timestamp int, accountId string) bool {
	if accountId == "" {
		return false
	}
	for _, acc := range bs.accounts {
		if acc.ID == accountId {
			return false
		}
	}

	bs.mu.Lock()
	defer bs.mu.Unlock()

	bs.accounts = append(bs.accounts, &Account{ID: accountId, CreatedAt: timestamp})
	return true
}

func (bs *Banksystem) Deposit(timestamp, amount int, accountId string) (int, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	for _, acc := range bs.accounts {
		if acc.ID == accountId {
			acc.amount += amount
			return acc.amount, nil
		}
	}

	return 0, errors.New("account not found")
}

func (bs *Banksystem) Transfer(timestamp, amount int, fromAccountId, toAccountId string) error {
	var fromAccount *Account
	var toAccount *Account

	bs.mu.RLock()
	for _, acc := range bs.accounts {
		if acc.ID == fromAccountId {
			fromAccount = acc
		}
	}
	for _, acc := range bs.accounts {
		if acc.ID == toAccountId {
			toAccount = acc
		}
	}
	bs.mu.RUnlock()

	if fromAccount == nil || toAccount == nil {
		return errors.New("account not found")
	}

	bs.mu.Lock()
	defer bs.mu.Unlock()
	fromAccount.amount -= amount
	toAccount.amount += amount

	return nil
}
