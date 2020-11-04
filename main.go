package main

import (
  "flag"
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/golang/glog"
  "github.com/Lunkov/lib-tr"
)

var service_port = ":3000"

func showDict(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Methods", "GET")
  params := mux.Vars(r)
  dict_id, ok := params["dict_id"]
  if !ok {
    glog.Errorf("ERR: URL '%s': DoN`t Set Dictionary `%v`\n", r.URL.Path, params)
    glog.Errorf("ERR: URL '%s': DoN`t Set Dictionary `%s`\n", r.URL.Path, dict_id)
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  lang, ok := params["lang"]
  if !ok {
    lang = "default"
  }
  w.WriteHeader(http.StatusOK)
  w.Write(jsonDict(dict_id, lang))
}

type StatInfo struct {
  CntDicts     int ``
  CntRecords   int ``
}

func Stats(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Methods", "GET")
  w.WriteHeader(http.StatusOK)
  var tui StatInfo
  tui.CntRecords = cnt_records
  tui.CntDicts = len(mapDicts)
  resJSON, _ := json.Marshal(tui)
  w.Write(resJSON)
}

func Health(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET")
  w.Header().Set("Access-Control-Allow-Credentials", "true")
  w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
  status := "OK"
  if true {
    status = "ERROR"
  }
  fmt.Fprintf(w, "{\"status\": \"%s\"}", status)
}

func main() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", ".")
  configPath := flag.String("config_path", "./etc/", "Config path")

  flag.Parse()

  tr.LoadLangs((*configPath) + "/langs.yaml")
  tr.LoadTrs((*configPath) + "/tr")
  dictInit((*configPath))
  tr.SaveNew((*configPath) + "/tr")

  router := mux.NewRouter()
  router.HandleFunc("/health",     Health)
  router.HandleFunc("/stat",       Stats)
  router.HandleFunc("/{dict_id}/{lang}", showDict)

  glog.Infof("LOG: Count dictionaries: %d\n", len(mapDicts))
  glog.Infof("LOG: Starting HTTP Dictionary server on %s\n", service_port)
  http.ListenAndServe(service_port, router)
}
