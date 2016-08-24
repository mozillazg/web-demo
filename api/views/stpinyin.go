package views

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"strings"

	"github.com/mozillazg/stpinyin"
)

type stpinyinRequest struct {
	Pinyin string `json:"pinyin"`
}

type stpinyinResponse struct {
	Pinyin string `json:"pinyin"`
}

var stpinyinRouter = router{
	url: "/api/v1/stpinyin",
	methods: []string{"POST"},
	handler: func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		req := stpinyinRequest{}
		err = json.Unmarshal(b, &req)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		pinyins := strings.Split(req.Pinyin, " ")
		for n, v := range(pinyins) {
			v = stpinyin.Convert(v)
			pinyins[n] = v
		}
		resp := stpinyinResponse{strings.Join(pinyins, " ")}
		b, _ = json.Marshal(resp)
		fmt.Fprintf(w, string(b))
	},
}

func init() {
	addRouter(stpinyinRouter)
}
