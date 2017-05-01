package usersess

import ()

type UserStore interface {
	Find(string) (*User, error)
	FindByEmail(string) (*User, error)
	FindByUsername(string) (*User, error)
	Save(User) error
}


var globalUserStore UserStore

func ImplementUserStore(store UserStore) {
	globalUserStore = store
}