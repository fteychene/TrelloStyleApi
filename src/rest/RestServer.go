package rest

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)

type RestServer struct {
	matcher *RouteMatcher 
}

func (server *RestServer) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	var data interface{}

	request.ParseForm()
	method := request.Method
	values := request.Form

	route, computedDatas := server.matcher.getHandlerAndUrlDatas(method, request.URL.Path)

	if route == nil {
		log.Printf("warn: %s ascked, no route found", request.URL.Path)
		http.NotFound(rw, request)
		return
	}
	
	for k, v := range computedDatas {
	    values[k] = append(values[k], v...)
	}
	
	code, data := route.handler(values)
	content, err := json.Marshal(data)
	if err != nil {
		panic(err)
		}
	rw.WriteHeader(code)
	rw.Write(content)
}

func recoverHandler(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    defer func() {
      if err := recover(); err != nil {
        log.Printf("panic: %+v", err)
        http.Error(w, http.StatusText(500), 500)
      }
    }()

    next.ServeHTTP(w, r)
  }

  return http.HandlerFunc(fn)
}

func (server *RestServer) Start(port int, routeMatcher *RouteMatcher) {
	log.Printf("Starting HTTP Server on %d\n", port)
	
	portString := fmt.Sprintf(":%d", port)
	server.matcher = routeMatcher
	http.Handle("/", recoverHandler(server))
	http.ListenAndServe(portString, nil)
}
