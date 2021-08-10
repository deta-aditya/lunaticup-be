package tournaments

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	MethodSingleElim          = "SINGLE_ELIM"
	MethodPrelimArrangedGroup = "PRELIM_ARRANGED_GROUP"

	ErrInvalidName    = errors.New("name should not be empty")
	ErrInvalidPlayers = errors.New("tournament should contain at least 2 players")
	ErrInvalidMethod  = fmt.Errorf("method should be either %s or %s with appropriate config", MethodSingleElim, MethodPrelimArrangedGroup)

	errMap = map[string]error{
		"IsNameValid":    ErrInvalidName,
		"IsPlayersValid": ErrInvalidPlayers,
		"IsMethodValid":  ErrInvalidMethod,
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

func (tr TournamentRaw) IsMethodValid() (valid bool) {
	if tr.Method == MethodSingleElim {
		valid = true
	}

	if tr.Method == MethodPrelimArrangedGroup {
		conf, ok := tr.MethodConfig.(map[string]interface{})
		if !ok {
			return
		}

		ppgRaw, ok := conf["players_per_group"]
		if !ok {
			return
		}

		ppg, ok := ppgRaw.(int)
		if !ok {
			return
		}

		ppgAboveMin := ppg >= 3
		ppgBeMultipleOfPlayersCount := len(tr.Players)%ppg == 0

		valid = ppgAboveMin && ppgBeMultipleOfPlayersCount
	}

	return
}
