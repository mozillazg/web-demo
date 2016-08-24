package views

import (
	"net/http"
	"log"
)

type router struct {
	url string
	methods []string
	handler http.HandlerFunc
}


func (r *router) match(req *http.Request) bool {
	if req.URL.Path == r.url {
		for _, method := range(r.methods) {
			if req.Method == method {
				return true
			}
		}
	}
	return false
}

var routers = []router{}

func addRouter(r router) {
	routers = append(routers, r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request url '%s' with '%s'", r.URL.Path, r.Method)
	log.Printf("routers: %v", routers)
	for _, router := range(routers) {
		if router.match(r) {
			router.handler(w, r)
			return
		}
	}
	http.NotFound(w, r)
}
