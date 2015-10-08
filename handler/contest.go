package handler

import (
	"github.com/zenazn/goji/web"
	"ojapi/model"

	"encoding/json"
	"net/http"
)

//@URL: /api/contests @method: GET
func ContestList(c web.C, w http.ResponseWriter, r *http.Request) {

	CModel := model.ContestModel{}
	conetestList, err := CModel.List(nil)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(conetestList)
}