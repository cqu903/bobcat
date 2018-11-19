package bobcat

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cqu903/bobcat/config"
	"github.com/cqu903/bobcat/utils"
)

type HandleFunction func(url string, params map[string]string, request *http.Request, responseWriter http.ResponseWriter)

//控制器信息，用于封装控制器
type ControllerInfo struct {
	Url            string // example   /index/person/:id/:name/
	regexp         *regexp.Regexp
	paramNames     []string //给定的url解析出来的参数名称列表[id,name]
	GetHandleFunc  HandleFunction
	PostHandleFunc HandleFunction
}

var HTTP_GET string = "GET"
var HTTP_POST string = "POST"

//判断对应的控制器是否可以匹配给定的url和http method
func (c *ControllerInfo) isMatch(url string, method string) bool {
	isMatch := false
	switch method {
	case HTTP_GET:
		if c.GetHandleFunc != nil && c.regexp.MatchString(url) {
			isMatch = true
		}
	case HTTP_POST:
		if c.PostHandleFunc != nil && c.regexp.MatchString(url) {
			isMatch = true
		}
	}
	return isMatch
}
func NewControllerInfo(url string, get HandleFunction, post HandleFunction) (*ControllerInfo, error) {
	if url == "" {
		return nil, errors.New("invalid controllerInfo,the url field is nil!")
	}
	if get == nil && post == nil {
		return nil, errors.New("invalid controllerInfo,the GetHandler And the PostHandler must at least one is not empty")
	}
	regexp, params, err := complieRegexp(url)
	if err != nil {
		return nil, err
	}
	return &ControllerInfo{url, regexp, params, get, post}, nil
}

/*根据url解析正则表达式
支持的url格式如下：
/index
/pserson/:id
/person/:id/:name
将url进行正则表达式替换后，并将可能存在的参数以切片的形式返回，供后续参数封装使用
*/
func complieRegexp(url string) (*regexp.Regexp, []string, error) {
	regex, _ := regexp.Compile(":[^/]*")
	paramNames := regex.FindAllString(url, -1)
	for index, value := range paramNames {
		paramNames[index] = value[1:]
	}
	url = regex.ReplaceAllString(url, "([^/]*)") //启用分组捕获
	newReg, err := regexp.Compile(url)
	if err != nil {
		return nil, nil, err
	}
	return newReg, paramNames, nil
}

//总体控制器集合，管理所有的控制器
type Router struct {
	controllerInfoList []*ControllerInfo
}

var BobCat = new(Router)

func (router *Router) AddControllerInfo(c *ControllerInfo) {
	if c == nil || c.Url == "" || c.regexp == nil || (c.GetHandleFunc == nil && c.PostHandleFunc == nil) {
		panic("invalid controllerInfo,the required field of the controller info is nil,please check!")
	}
	router.controllerInfoList = append(router.controllerInfoList, c)
}
func (router Router) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	//handleFilter() 暂未实现
	//session() 暂未实现
	url := request.URL
	method := request.Method
	if method != HTTP_GET && method != HTTP_POST {
		http.Error(responseWriter, "Unsupported Http Method", 500)
	}
	//处理静态文件请求
	if strings.HasPrefix(url.Path, config.Conf.StaticDir) {
		processStaticContent(responseWriter, request)
		return
	}
	requestPath := url.Path + "?" + url.RawQuery + url.Fragment
	requestParams := generateParams(request)
	isFindController := false
	for _, controllerInfo := range router.controllerInfoList {
		if controllerInfo.isMatch((url.Path), method) {
			addPathVariables(url.Path, requestParams, controllerInfo)
			switch method {
			case HTTP_GET:
				isFindController = true
				controllerInfo.GetHandleFunc(requestPath, requestParams, request, responseWriter)
			case HTTP_POST:
				isFindController = true
				controllerInfo.PostHandleFunc(requestPath, requestParams, request, responseWriter)
			}
			break
		}
	}
	//if don't find the controllerInfo to process the request,use unfindControllerInfo
	if !isFindController {
		applyDefaultUnfound(responseWriter, request)
	}
}

//处理静态文件内容
func processStaticContent(responseWriter http.ResponseWriter, request *http.Request) {
	requestUrlPath := request.URL.Path
	arrays := strings.Split(requestUrlPath, "/")
	if len(arrays) <= 1 {
		http.Error(responseWriter, "invalid static content reqeust path:"+requestUrlPath, 200)
		return
	}
	path := filepath.Join(utils.GetBaseDir(), filepath.Join(arrays...))
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		http.Error(responseWriter, "read static file error!", 404)
		return
	}
	responseWriter.Header().Set("Content-Type", utils.DecideRightContentType(requestUrlPath[strings.LastIndex(requestUrlPath, ".")+1:]))
	responseWriter.Write(bytes)
}

//对于无法找到对应的controller，返回404
func applyDefaultUnfound(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "text/html")
	responseWriter.WriteHeader(404)
	responseWriter.Write([]byte("<html><body>oops!the request url is missing from out website!</body></html>"))
}
func addPathVariables(path string, paramsMap map[string]string, c *ControllerInfo) {
	values := c.regexp.FindAllStringSubmatch(path, -1)
	//see regepx分组捕获，第0个元素是原始字符串
	fmt.Println(values[0])
	for index, paramName := range c.paramNames {
		paramsMap[paramName] = values[0][index+1]
	}
}
func generateParams(request *http.Request) map[string]string {
	paramsMap := make(map[string]string)
	request.ParseForm()
	form := request.Form
	for key, value := range form {
		//由于同名参数的问题，例如在post请求中的body部分和query string中存在同名参数，在本实现中只取第一个参数作为有效参数，请不要使用同名参数
		paramsMap[key] = value[0]
	}

	return paramsMap
}
func init() {
	config.Config()
}

//启动http服务
func StartHttpServe() error {
	server := &http.Server{
		Addr:         ":" + strconv.Itoa(config.Conf.Port),
		Handler:      BobCat,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	if config.Conf.EnableHTTPS {
		return server.ListenAndServeTLS(config.Conf.Server.CertFile, config.Conf.Server.KeyFile)
	} else {
		return server.ListenAndServe()
	}
}
