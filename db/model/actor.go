package model

import (
	"restexample/db"
	"time"
)

// IActor representa una cosa con la capacidad de operar sobre la entidad actor
type IActor interface {

	// Get obtiene un actor por su id
	Get(db.QueryAble, int64) (Actor, error)

	// Los siguientes métodos están comentados porque aun no los implementaremos
	// Fetch(QueryAble, ActorFilter) ([]Actor, error)
	// Save(QueryAble, *Actor) error
}

// Actor mantiene los datos de una entidad actor
type Actor struct {
	ID         int64     `json:"id"`         // actor_id
	FirstName  string    `json:"firstName"`  // first_name
	LastName   string    `json:"lastName"`   // last_name
	LastUpdate time.Time `json:"lastUpdate"` // last_update
}

var _ IActor = ActorCrud{}

// ActorCrud encapsula métodos capaces de operar sobre la entidad actor. Implementa a IActor
type ActorCrud struct{}

// Get recupera un actor por su id
func (a ActorCrud) Get(q db.QueryAble, ID int64) (actor Actor, err error) {

	// Si no se ha inicializado la conexión, devuelve ErrNoConection
	if q == nil {
		err = db.ErrNoConnection
		return
	}

	err = q.QueryRow(sqlGetActorByID, ID).Scan(
		&actor.ID,
		&actor.FirstName,
		&actor.LastName,
		&actor.LastUpdate,
	)

	return
}

var _ IActor = MockActorCrud{}

// MockActorCrud encapsula métodos para operar sobre la entidad actor simulando una conexión. Implementa a IActor
type MockActorCrud struct{}

// Get recupera un actor por su id
func (a MockActorCrud) Get(q db.QueryAble, ID int64) (actor Actor, err error) {
	// Como MockActorCrud solo simula la conexión, la obligamos a devolver datos de actor en duro
	// para simular la obtención de una entidad actor
	actor.ID = ID
	actor.FirstName = "Pepito"
	actor.LastName = "Grillo"
	actor.LastUpdate, _ = time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Feb 4, 2014 at 6:05pm (PST)")
	return
}

// NewActorCrud retorna algo que implementa a IActor.
// Si mock es true, retornara una instancia de MockActorCrud
// para simular la conexión.
func NewActorCrud(mock bool) IActor {
	if mock {
		return MockActorCrud{}
	}

	return ActorCrud{}
}

const (
	USE_MOCK_CONNECTION = true
	NO_MOCK_CONNECTION  = false
	DEBUG_MODE          = true
	NO_DEBUG_MODE       = false
	sqlGetActorByID     = `SELECT actor_id, first_name, last_name, last_update FROM actor WHERE actor_id = ?`
)
