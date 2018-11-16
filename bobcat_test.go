package bobcat

import (
	"net/http"
	"testing"
)

func TestNewControllerInfo(t *testing.T) {
	c, _ := NewControllerInfo("/home/:id", handleFunc, handleFunc)
	t.Log(c)
	if c.regexp == nil {
		t.Fail()
	}
	if c.paramNames[0] != "id" {
		t.Fail()
	}
}
func handleFunc(url string, params map[string]string, request *http.Request, responseWriter http.ResponseWriter) {

}
