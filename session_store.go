package usersess

type SessionStore interface {
	Find(string) (*Session, error)
	Save(*Session) error
	Delete(*Session) error
}

var globalSessionStore SessionStore

func ImplementSessionStore(store SessionStore) {
	globalSessionStore = store
}