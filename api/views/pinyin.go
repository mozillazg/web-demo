package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mozillazg/go-pinyin"
)

type pinyinRequest struct {
	Han string `json:"han"`
}

type pinyinResponse struct {
	Pinyin string `json:"pinyin"`
}

var pinyinRouter = router{
	url:     "/api/v1/pinyin",
	methods: []string{"POST"},
	handler: func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		req := pinyinRequest{}
		err = json.Unmarshal(b, &req)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		pinyins := pinyin.LazyPinyin(req.Han, pinyin.NewArgs())
		resp := stpinyinResponse{strings.Join(pinyins, " ")}
		b, _ = json.Marshal(resp)
		fmt.Fprintf(w, string(b))
	},
}

func init() {
	addRouter(pinyinRouter)
}
