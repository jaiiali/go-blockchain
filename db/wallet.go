package db

type Account string

func NewAccount(v string) Account {
	return Account(v)
}
