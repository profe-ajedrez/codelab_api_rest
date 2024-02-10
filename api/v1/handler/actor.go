package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"restexample/config"
	"restexample/db"
	"restexample/db/model"
	"restexample/logadapter"
	"strconv"
	"strings"
)

var (
	actorRXPWithID = regexp.MustCompile(`^/actors/[0-9]+$`)
)

type ActorHandler struct{}

func (a *ActorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && actorRXPWithID.MatchString(r.URL.Path):
		a.GetActorByID(w, r)
	default:
		return
	}
}

func (a *ActorHandler) GetActorByID(w http.ResponseWriter, r *http.Request) {
	matches := strings.Split(r.URL.Path, "/")

	if len(matches) < 2 {
		logadapter.Log.Error("couldnt parse actor id from query param")
		response := NewActorInternalServerErrorResponse()
		js, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	ID, err := strconv.ParseInt(matches[2], 10, 64)

	if err != nil {
		logadapter.Log.Error("couldnt parse actor id as integer value")
		response := NewActorBadRequestResponse()
		js, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	actorRepo := model.NewActorCrud(config.Mock())

	actor, err := actorRepo.Get(db.DB(), ID)

	if err != nil {
		if err == sql.ErrNoRows {
			logadapter.Log.Error("actor id not found")
			response := NewActorNotFoundResponse()
			js, _ := json.Marshal(response)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(js)
			return
		}

		logadapter.Log.Error("couldnt get actor by id, something was wrong")
		response := NewActorBadRequestResponse()
		js, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(js)
		return
	}

	response := NewActorResponse(actor)
	js, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)
}

type ActorResponse struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

func NewActorResponse(data interface{}) ActorResponse {
	return ActorResponse{
		Status: "ok",
		Code:   200,
		Data:   data,
	}
}

func NewActorNotFoundResponse() ActorResponse {
	return ActorResponse{
		Status: "fail",
		Code:   404,
		Data:   make(map[string]interface{}, 0),
	}
}

func NewActorBadRequestResponse() ActorResponse {
	return ActorResponse{
		Status: "fail",
		Code:   400,
		Data:   make(map[string]interface{}, 0),
	}
}

func NewActorInternalServerErrorResponse() ActorResponse {
	return ActorResponse{
		Status: "fail",
		Code:   500,
		Data:   make(map[string]interface{}, 0),
	}
}
