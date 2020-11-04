package main

import (
  "flag"
  "testing"
  "github.com/stretchr/testify/assert"
  "net/http"
  "net/http/httptest"
  "github.com/gorilla/mux"
  "github.com/Lunkov/lib-tr"
)

/////////////////////////
// TESTS
/////////////////////////

/////
func TestWWWDict(t *testing.T) {
  flag.Set("alsologtostderr", "true")
  flag.Set("log_dir", ".")
  flag.Set("v", "9")
  flag.Parse()

  tr.LoadLangs("./etc.test")
  tr.LoadTrs("./etc.test/tr")
  dictInit("./etc.test")


  req, err := http.NewRequest("GET", "/gis_layer/ru", nil)
  if err != nil {
    t.Fatal(err)
  }
  req = mux.SetURLVars(req, map[string]string{
			"dict_id": "gis_layer",
      "lang": "ru_RU",
		})
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(showDict)
  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
      t.Errorf("handler returned wrong status code: got %v want %v",
          status, http.StatusOK)
  }
  
  body_need := "[{\"id\":\"a6dd28d7-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Buildings\",\"parent_code\":\"\",\"name\":\"Здания и сооружения\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"a6dfae5d-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Electro.Charging.Stations\",\"parent_code\":\"\",\"name\":\"Электро заправочные станции\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"a6e234ae-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Street.Lighting\",\"parent_code\":\"\",\"name\":\"Уличное освещение\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d397ec60-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Parkings\",\"parent_code\":\"\",\"name\":\"Паркоматы\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d39a43c7-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Waste.Sites\",\"parent_code\":\"\",\"name\":\"Пункт сбора отходов\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d39cafdd-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"water.leaks\",\"parent_code\":\"\",\"name\":\"Протечки воды\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"a6ec7f00-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"wc\",\"parent_code\":\"\",\"name\":\"Санузел\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d39f6c6b-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Weather.Stations\",\"parent_code\":\"\",\"name\":\"Метео-станции\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"}]"
  
  assert.Equal(t, body_need, string(rr.Body.Bytes()))
}

func TestWWWDictStat(t *testing.T) {
  req, err := http.NewRequest("GET", "/stat", nil)
  if err != nil {
    t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(Stats)
  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
      t.Errorf("handler returned wrong status code: got %v want %v",
          status, http.StatusOK)
  }

}
