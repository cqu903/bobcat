package bobcat

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewControllerInfo(t *testing.T) {
	c, _ := NewControllerInfo("/home/:id", handleFunc, handleFunc)
	assert.NotNil(t, c.regexp, "controllerInfo.regexp can not be null")
	assert.Equal(t, "id", c.paramNames[0], "the first param must equals 'id'")
}
func handleFunc(url string, params map[string]string, request *http.Request, responseWriter http.ResponseWriter) {
}
