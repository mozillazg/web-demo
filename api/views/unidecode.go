package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mozillazg/go-unidecode"
)

type unidecodeRequest struct {
	String string `json:"string"`
}

type unidecodeResponse struct {
	String string `json:"string"`
}

var unidecodeRouter = router{
	url:     "/api/v1/go-unidecode",
	methods: []string{"POST"},
	handler: func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		req := unidecodeRequest{}
		err = json.Unmarshal(b, &req)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		s := unidecode.Unidecode(req.String)
		resp := unidecodeResponse{s}
		b, _ = json.Marshal(resp)
		fmt.Fprintf(w, string(b))
	},
}

func init() {
	addRouter(unidecodeRouter)
}
