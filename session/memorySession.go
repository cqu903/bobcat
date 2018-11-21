package session

import (
	"time"
)

type memorySession struct {
	token      string
	paramMap   map[string]interface{}
	createTime time.Time
	expireTime time.Time
	expire     bool
}

var MemorySession memorySession

func (s *memorySession) AddParam(key string, val interface{}) {

}
func (s *memorySession) RemoveParam(key string) interface{} {
	return s.paramMap[key]
}
func (s *memorySession) ExpireImmediately() {
	s.expire = true
}
func (s *memorySession) IsExpire() bool {
	return s.expire
}

type memorySessionManager struct {
	sessionMap map[string]*memorySession
}

func (m *memorySessionManager) GetSession(isCreateOnNil bool, sessionToken string) Session {
	s := m.sessionMap[sessionToken]
	if isCreateOnNil {
		if s == nil {
			m.sessionMap[sessionToken] = &memorySession{}
		}
		return s
	} else {
		if s == nil {
			return nil
		} else {
			return s
		}
	}
}
