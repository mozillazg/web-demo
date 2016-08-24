package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mozillazg/go-slugify"
)

type slugifyRequest struct {
	String string `json:"string"`
}

type slugifyResponse struct {
	Slug string `json:"slug"`
}

var slugifyRouter = router{
	url:     "/api/v1/go-slugify",
	methods: []string{"POST"},
	handler: func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		req := slugifyRequest{}
		err = json.Unmarshal(b, &req)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		s := slugify.Slugify(req.String)
		resp := slugifyResponse{s}
		b, _ = json.Marshal(resp)
		fmt.Fprintf(w, string(b))
	},
}

func init() {
	addRouter(slugifyRouter)
}
