package session
import (
	"testing"
	"github.com/cqu903/bobcat/session/memorySession"
	"github.com/stretchr/testify/assert"
)
func TestSession(t *testing.T){
	sessionContext := memorySession.NewSessionManager()
	session:=sessionContext.GetSession("roy",true)
	session.AddParam("address","shenzhen")
	assert.Equal(t,false,session.IsExpire,"session init is not expire")
	session.ExpireImmediately()
	assert.True(t,session.IsExpire(),"session expire faild")
}