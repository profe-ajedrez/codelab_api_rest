package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restexample/api/v1/mocks"
	"restexample/config"
	"restexample/db"
	"restexample/db/model"
	"testing"
)

func TestActorGetByID(t *testing.T) {
	config.SetMockingOn()

	expectedActor, err := (model.MockActorCrud{}).Get(db.DB(), 1)

	if err != nil {
		t.Logf("couldnt get mock entity actor %v", err)
		t.FailNow()
	}

	client := newActorClient()

	for _, tesCase := range actorRequestCases {
		req, err := http.NewRequest("GET", tesCase.url, nil)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		resp, err := client.Do(req)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		defer resp.Body.Close()

		if resp.StatusCode != tesCase.expectedCode {
			t.Logf("Expected status code %d, got %d", tesCase.expectedCode, resp.StatusCode)
			t.FailNow()
		}

		var dataMap map[string]interface{}

		err = json.NewDecoder(resp.Body).Decode(&dataMap)

		if err != nil {
			t.Logf("expected error to be nil while reading body reponse. got %v", err)
			t.FailNow()
		}

		actor, ok := dataMap["data"].(map[string]interface{})

		if !ok {
			t.Log("couldnt extract actor info from response")
			t.FailNow()
		}

		if actor["firstName"] != expectedActor.FirstName {
			t.Logf("Expected actor name to be %s, got, got %s", expectedActor.FirstName, actor["firstName"])
			t.FailNow()
		}

		if actor["lastName"] != expectedActor.LastName {
			t.Logf("Expected actor name to be %v, got %s", expectedActor.LastName, actor["lastName"])
			t.FailNow()
		}

		if actor["lastUpdate"] != expectedActor.LastUpdate.Format("2006-01-02T15:04:05Z07:00") {
			t.Logf("Expected actor name to be %v, got %s", expectedActor.LastUpdate, actor["lastUpdate"])
			t.FailNow()
		}

	}
}

func BenchmarkActorGetByID(b *testing.B) {
	config.SetMockingOn()

	client := newActorClient()

	testCase := actorRequestCases[0]
	req, _ := http.NewRequest("GET", testCase.url, nil)

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		_, _ = client.Do(req)
	}

}

func newActorClient() mocks.HTTPClient {

	client := &mocks.MockClient{
		DoFunc: func(r *http.Request) (*http.Response, error) {
			w := httptest.NewRecorder()

			actorHandler := ActorHandler{}
			actorHandler.GetActorByID(w, r)
			resp := w.Result()

			return resp, nil
		},
	}

	return client
}

var actorRequestCases = []struct {
	url          string
	expectedCode int
	expectedBody string
}{
	{
		url:          "/actor/10",
		expectedCode: 200,
		expectedBody: `{
	"code": 200,
	"data": {
		"firstName": "Pepito",
		"id": 10,
		"lastName": "Grillo",
		"lastUpdate": "2014-02-04T18:05:00Z"
	},
	"status": "ok"
}`,
	},
}
