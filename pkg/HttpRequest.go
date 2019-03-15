package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpRequest struct {
	HttpClient *http.Client
	Server     *http.Server
}

type jsRender struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}
func (hq *HttpRequest) RegisterServers() error {
	app := GetCfg().App
	http.HandleFunc("/", DealRequest)
	err := http.ListenAndServe(app["addr"]+":"+app["port"], nil)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	return err
}

func DealRequest(w http.ResponseWriter, r *http.Request) {
	Info("xxxxx")
	//writer.Write([]byte(request.URL.Query()))
	//info := fmt.Sprintln(r.Header.Get("Content-Type"))
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	ResultJson(200, "ok", string(body) , w)
}

func ResultJson(code int, message string, data interface{}, w http.ResponseWriter) {
	jsData := jsRender{code, message, data}
	js, err := json.Marshal(jsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}