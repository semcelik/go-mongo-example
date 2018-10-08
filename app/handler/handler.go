package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/semcelik/go-mongo-example/app/model"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetUsersHandler")
	dummyDatas := []model.Dummy{
		model.Dummy{DummyId: 1, DummyName: "dummy1"},
		model.Dummy{DummyId: 2, DummyName: "dummy2"},
	}
	json.NewEncoder(w).Encode(dummyDatas)
	return
}
