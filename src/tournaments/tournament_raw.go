package tournaments

import (
	"errors"
	"reflect"
)

var (
	ErrInvalidName    = errors.New("name should not be empty")
	ErrInvalidPlayers = errors.New("tournament should contain at least 2 players")

	errMap = map[string]error{
		"IsNameValid":    ErrInvalidName,
		"IsPlayersValid": ErrInvalidPlayers,
	}
)

type TournamentRaw struct {
	Name         string      `json:"name"`
	Method       string      `json:"method"`
	MethodConfig interface{} `json:"method_config"`
	Players      []string    `json:"players"`
}

func (tr TournamentRaw) IsValid() (errs []error) {
	for method, err := range errMap {
		result := reflect.ValueOf(tr).MethodByName(method).Call(nil)
		valid := result[0].Bool()

		if !valid {
			errs = append(errs, err)
		}
	}

	return
}

func (tr TournamentRaw) IsNameValid() bool {
	return len(tr.Name) > 0
}

func (tr TournamentRaw) IsPlayersValid() bool {
	return len(tr.Players) >= 2
}
