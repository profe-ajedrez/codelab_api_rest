package model

import (
	"os"
	"restexample/config"
	"restexample/db"
	"testing"

	"github.com/joho/godotenv"
)

func TestActorGetMock(t *testing.T) {

	err := db.Open(os.Getenv("DEV_API_DB_DSN"), config.DEBUG_MODE, config.USE_MOCK_CONNECTION)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	var actorCrud = MockActorCrud{}

	actorInfo, err := actorCrud.Get(db.DB(), 1)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if actorInfo.ID != 1 {
		t.Log("actor id is not 1")
		t.FailNow()
	}

	if actorInfo.FirstName != "Pepito" {
		t.Log("actor name is not Pepito")
		t.FailNow()
	}

	t.Log(actorInfo)
}

func BenchmarkActorGetMock(b *testing.B) {

	_ = godotenv.Load()
	err := db.Open(os.Getenv("DEV_API_DB_DSN"), config.DEBUG_MODE, config.USE_MOCK_CONNECTION)

	if err != nil {
		b.Log(err)
		b.FailNow()
	}

	var actorCrud = MockActorCrud{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = actorCrud.Get(db.DB(), 1)
	}
}
