package main

import (
  "flag"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/Lunkov/lib-tr"
)

/////////////////////////
// TESTS
/////////////////////////

/////
func TestDict(t *testing.T) {
  flag.Set("alsologtostderr", "true")
  flag.Set("log_dir", ".")
  flag.Set("v", "9")
  flag.Parse()

  tr.LoadLangs("./etc.test/langs.yaml")
  tr.LoadTrs("./etc.test/tr")
  dictInit("./etc.test")
  
  cn_dicts := len(mapDicts)
  assert.Equal(t, 1, cn_dicts)
  assert.Equal(t, 8, cnt_records)
  
  d1_need := "[]"
  d1 := string(jsonDict("4543werfgdfgdgd", "ru_RU"))
  assert.Equal(t, d1_need, d1)

  d1_need = "[{\"id\":\"a6dd28d7-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Buildings\",\"parent_code\":\"\",\"name\":\"Здания и сооружения\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"a6dfae5d-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Electro.Charging.Stations\",\"parent_code\":\"\",\"name\":\"Электро заправочные станции\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"a6e234ae-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Street.Lighting\",\"parent_code\":\"\",\"name\":\"Уличное освещение\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d397ec60-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Parkings\",\"parent_code\":\"\",\"name\":\"Паркоматы\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d39a43c7-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Waste.Sites\",\"parent_code\":\"\",\"name\":\"Пункт сбора отходов\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d39cafdd-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"water.leaks\",\"parent_code\":\"\",\"name\":\"Протечки воды\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"a6ec7f00-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"wc\",\"parent_code\":\"\",\"name\":\"Санузел\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d39f6c6b-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Weather.Stations\",\"parent_code\":\"\",\"name\":\"Метео-станции\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"}]"
  d1 = string(jsonDict("gis_layer", "ru_RU"))
  assert.Equal(t, d1_need, d1)

  d1_need = "[{\"id\":\"a6dd28d7-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Buildings\",\"parent_code\":\"\",\"name\":\"Buildings\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"a6dfae5d-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Electro.Charging.Stations\",\"parent_code\":\"\",\"name\":\"Electric gas stations\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"a6e234ae-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Street.Lighting\",\"parent_code\":\"\",\"name\":\"Street lighting\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d397ec60-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Parkings\",\"parent_code\":\"\",\"name\":\"Parking meters\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d39a43c7-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Waste.Sites\",\"parent_code\":\"\",\"name\":\"Waste collection point\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d39cafdd-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"water.leaks\",\"parent_code\":\"\",\"name\":\"Water leaks\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"a6ec7f00-f6f4-11e9-a4cc-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"wc\",\"parent_code\":\"\",\"name\":\"Bathroom\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"},{\"id\":\"d39f6c6b-f6f2-11e9-a488-bcaec5b972a6\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"parent_id\":\"00000000-0000-0000-0000-000000000000\",\"code_id\":0,\"code\":\"Weather.Stations\",\"parent_code\":\"\",\"name\":\"Weather stations\",\"description\":\"\",\"symbol\":\"\",\"image\":\"\"}]"
  d1 = string(jsonDict("gis_layer", "en_US"))
  
  assert.Equal(t, d1_need, d1)
  
  tr.SaveNew("./etc.test/tr")
}
