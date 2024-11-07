package account

import (
	"errors"
)

type Account struct {
	Id    uint64 `json:"id"`
	Email string `json:"email"`
}

type AccountDAO interface {
	CreateAccount(email string) (*Account, error)
	GetAccountByID(id uint64) (*Account, error)
}

type InMemoryAccountDAO struct {
	db     map[uint64]*Account
	nextId uint64
}

func NewInMemory() *InMemoryAccountDAO {
	return &InMemoryAccountDAO{nextId: 1, db: make(map[uint64]*Account)}
}

func (dao *InMemoryAccountDAO) CreateAccount(email string) (*Account, error) {
	id := dao.nextId
	dao.nextId++
	newAccount := &Account{Id: id, Email: email}

	dao.db[id] = newAccount

	return newAccount, nil
}

func (dao *InMemoryAccountDAO) GetAccountByID(id uint64) (*Account, error) {
	val, ok := dao.db[id]

	if !ok {
		return nil, errors.New("account does not exist")
	}

	return val, nil
}
