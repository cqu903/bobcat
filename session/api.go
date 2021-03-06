package session

type Session interface {
	//add param into a session
	AddParam(key string, val interface{})
	//remove param from a session,and return the param
	RemoveParam(key string) interface{}
	//expore the session
	ExpireImmediately()
	//return is the session expired
	IsExpire() bool
	//return the session key,a session key is the unique identifier
	GetSessionKey() string
}

type SessionManager interface {
	GetSession(token string,isCreateOnNil bool) Session
}
