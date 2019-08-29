package main

import (
	"fmt"
	"encoding/json"
)

type CommonResp struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type Api struct {
	NodeAddress string
}

var api *Api

func InitApi(nodeAddress string) {
	api = &Api{
		NodeAddress: nodeAddress,
	}
}

func GetApi() *Api {
	return api
}

func (a *Api) Post(path string, data interface{}) ([]byte, error) {
	return Post(fmt.Sprintf("%s%s", a.NodeAddress, path), data)
}

func (a *Api) Get(path string) ([]byte, error) {
	return UrlGetContent(fmt.Sprintf("%s%s", a.NodeAddress, path))
}


func GetAppInfo(path string, app *application) (resp AppResponse, err error) {
	data, err := GetApi().Get(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}
