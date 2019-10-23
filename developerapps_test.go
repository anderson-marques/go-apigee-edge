package apigee

import (
  "encoding/json"
  "testing"
)

const (
	appJson1 = `{
  "attributes": [ {
    "name" : "tag1",
    "value" : "created by golang" }],
  "name" : "will-be-replaced",
  "keyExpiresIn" : "3600000"
}`
)

func randomAppFromTemplate() (DeveloperApp, error) {
	got := DeveloperApp{}
	e := json.Unmarshal([]byte(appJson1), &got)
	
	if e != nil {
		return got, e
	}
	// assign values
	tag := pretag + randomString(7)
	got.Name = tag + "app"
	return got, e
}


func TestDeveloperAppCreateDelete(t *testing.T) {
  opts := &EdgeClientOptions{Org: orgName, Auth: nil, Debug: false }
  client, e := NewEdgeClient(opts)
  if e != nil {
		t.Errorf("while initializing Edge client, error:\n%#v\n", e)
    return
  }
  //wait(1)

	dev, e := randomDeveloperFromTemplate()
  createdDeveloper, resp, e := client.Developers.Create(dev)
  if e != nil {
		t.Errorf("while creating Edge developer, error:\n%#v\n", e)
    return
  }
	t.Logf("Create: got=%+v", createdDeveloper)
	t.Logf("resp: got=%+v", resp)
	
  wait(1)

	devapps := client.Developers.Apps(createdDeveloper.Email)
	devapp, e := randomAppFromTemplate()
	
  createdApp, resp, e := devapps.Create(devapp)	
  if e != nil {
		t.Errorf("while creating developer app, error:\n%#v\n", e)
    return
  }
	t.Logf("CreateApp: got=%v", createdApp)
	
  wait(1)
	
  deletedApp, resp, e := devapps.Delete(createdApp.Name)	
  if e != nil {
		t.Errorf("while creating developer app, error:\n%#v\n", e)
    return
  }
	t.Logf("DeleteApp: got=%v", deletedApp)

  wait(1)
	
  deletedDeveloper, resp, e := client.Developers.Delete(createdDeveloper.Email)
  if e != nil {
		t.Errorf("while deleting Edge developer, error:\n%#v\n", e)
    return
  }
	t.Logf("Delete: got=%v", deletedDeveloper)
	
}
