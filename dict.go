package main

import (
  "io/ioutil"
  "os"
  "path/filepath"
  "encoding/json"
  "gopkg.in/yaml.v2"
  "github.com/golang/glog"
  "github.com/google/uuid"
  "github.com/Lunkov/lib-env"
  "github.com/Lunkov/lib-tr"
)

type DictInfo struct {
  CODE           string        `db:"code"         json:"code"           yaml:"code"`
  Name           string        `db:"name"         json:"name"           yaml:"name"`
}

var mapDicts = make(map[string]DictInfo)
var mapDictsItems = make(map[string][]Dictionary)
var mapDictsJSON = make(map[string][]byte)
var cnt_records = 0

func dictInit(configPath string) {
  env.LoadFromYMLFiles(configPath + "/dicts", loadFormYAML)
  for _, item := range mapDicts {
    filepath.Walk(configPath + "/data/" + item.CODE, func(filename string, f os.FileInfo, err error) error {
      if f != nil && f.IsDir() == false {
        if glog.V(2) {
          glog.Infof("FILE: %s\n", filename)
        }
        var err error
        yamlFile, err := ioutil.ReadFile(filename)
        if err != nil {
          glog.Errorf("ERR: ReadFile.yamlFile(%s)  #%v ", filename, err)
        } else {
          loadFormYAMLDict(filename, item.CODE, yamlFile)
        }
      }
      return nil
    })

  }
}

func loadFormYAML(filename string, yamlFile []byte) int {
  var mapTmp = make(map[string]DictInfo)

  err := yaml.Unmarshal(yamlFile, &mapTmp)
  if err != nil {
    glog.Errorf("ERR: yamlFile(%s): YAML: %v", filename, err)
    return 0
  }
  if(len(mapTmp) > 0) {
    for _, item := range mapTmp {
      mapDicts[item.CODE] = item
      index_def := item.CODE + "#" + tr.LangDefault()
      mapDictsItems[index_def] = make([]Dictionary, 0)
    }
  }
  return 1
}

func loadFormYAMLDict(filename string, dict_id string, yamlFile []byte) int {
  var mapTmp = make(map[string]Dictionary)

  err := yaml.Unmarshal(yamlFile, &mapTmp)
  if err != nil {
    glog.Errorf("ERR: yamlFile(%s): YAML: %v", filename, err)
    return 0
  }
  if(len(mapTmp) > 0) {
    update := false
    for i, item := range mapTmp {
      tr.SetDef(item.Name)
      tr.SetDef(item.Description)
      cnt_records ++
      index_def := dict_id + "#" + tr.LangDefault()
      if item.ID == uuid.Nil {
        item.ID, _ = uuid.NewUUID()
        mapTmp[i] = item
        update = true
      }
      mapDictsItems[index_def] = append(mapDictsItems[index_def], item)
    }
    if update {
      Save2YAML(filename, &mapTmp)
    }
  }
  return 1
}

func Save2YAML(filename string, dict_items *map[string]Dictionary) int {
  d, err := yaml.Marshal(dict_items)
  if err != nil {
    glog.Errorf("ERR: Save2YAML: %v", err)
    return 0
  }
  err = ioutil.WriteFile(filename, d, 0644)
  if err != nil {
    glog.Errorf("ERR: SAVE DICTIONARY: (%s) err = %v \n", filename, err)
    return 0
  }
  return len((*dict_items))
}

func jsonDict(dict_id string, lang string) []byte {
  index := dict_id + "#" + lang
  js, ok := mapDictsJSON[index]
  if ok {
    return js
  }
  dict, okd := mapDictsItems[index]
  if !okd {
    index_def := dict_id + "#" + tr.LangDefault()
    dict, okd = mapDictsItems[index_def]
    if !okd || len(dict) > 1 {
      new_tr_dict := make([]Dictionary, len(dict))
      for key, item := range dict {
        item.Name = tr.Tr(lang, item.Name)
        item.Description = tr.Tr(lang, item.Description)
        new_tr_dict[key] = item
      }
      dict = new_tr_dict
    }
  }
  jso, err := json.Marshal(dict)
  if err != nil {
    glog.Errorf("ERR: Dict(%s): JSON: %v", index, err)
    return []byte("")
  }
  mapDictsJSON[index] = jso
  return jso
}

