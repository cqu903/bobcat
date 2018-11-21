package memorySession

import (
	"github.com/cqu903/bobcat/session"
	"time" 
)

type memorySession struct {
	token      string
	paramMap   map[string]interface{}
	createTime time.Time
	expireTime time.Time
	expire     bool

}

func (s *memorySession) AddParam(key string, val interface{}) {
	s.paramMap[key]=val
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
func (s *memorySession) GetSessionKey() string{
	return s.token
}

type memorySessionManager struct {
	sessionMap map[string]*memorySession
}

func (m *memorySessionManager) GetSession(sessionToken string,isCreateOnNil bool) session.Session {
	s := m.sessionMap[sessionToken]
	if isCreateOnNil {
		if s == nil {
			s = &memorySession{token:sessionToken,createTime:time.Now(),expire:false,expireTime:time.Now().Add(30*time.Minute)}
			m.sessionMap[sessionToken] = s
		}
		return s
	} 
	if s == nil { 
		return nil
	} 
	return s
}
func  NewSessionManager() session.SessionManager{
	return &memorySessionManager{sessionMap:make(map[string]*memorySession)}
}
