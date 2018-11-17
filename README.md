# bobcat
a light weight  web application framework by golang

# install
just use the shell to install it!
    
    go get github.com/cqu903/bobcat 

# Usage
    package main

    import (
      "encoding/json"
      "fmt"
      "net/http"

      "github.com/cqu903/bobcat"
      "github.com/cqu903/bobcat/config"
    )

    func main() {
      /*
        define a controller,a controller has url pattern and request handler,the pattern likes 
        /home  /personInfo/:id etc,so it's support the path variable binding.the handler 
        is bobcat.HandleFunction type,in this sample is the ShowConifg function.so
        you can put it to process both get and post http request.
      */
      c, _ := bobcat.NewControllerInfo("/config", ShowConfig, ShowConfig)
      //add the controler to the router
      bobcat.BobCat.AddControllerInfo(c)
      //start the server
      err := bobcat.StartHttpServe()
      if(err!=nil){
	      log.Fatalf("try to start http server failed ,errors:%v",err)
      }
    }

    func ShowConfig(url string, params map[string]string, request *http.Request, responseWriter http.ResponseWriter) {
      responseWriter.Header().Set("Content-Type", "application/json")
      jsonValue, _ := json.Marshal(config.Conf)
      fmt.Println(jsonValue)
      responseWriter.Write(jsonValue)
    }
# config file
you can create *.yaml in your src path,bobcat will find and put it in it's config.so you can use your own config to replace the default config.the next is a full config for reference.
    
    server:
      port: 80
      enableHTTPS: false
      staticDir: /static/
    app:
      name: simple web app
      version: 1.0
      author: roy
