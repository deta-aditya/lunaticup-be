package tournaments

import (
	"fmt"
	"testing"
)

var (
	nameValid   = "Test Name"
	nameInvalid = ""

	methodValid   = "SINGLE_ELIM"
	methodInvalid = "RANDOM_STR"

	playersValid = []string{
		"Test Player 1",
		"Test Player 2",
	}
	playersInvalid = []string{
		"Test Player 1",
	}
)

func TestTournamentRaw_IsValid(t *testing.T) {
	tests := []struct {
		name   string
		tr     TournamentRaw
		hasErr bool
	}{
		{
			name: "it should return empty errors when valid",
			tr: TournamentRaw{
				Name:    nameValid,
				Method:  methodValid,
				Players: playersValid,
			},
			hasErr: false,
		},
		{
			name: "it should return some errors when invalid",
			tr: TournamentRaw{
				Name:    nameInvalid,
				Method:  methodInvalid,
				Players: playersInvalid,
			},
			hasErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.IsValid()
			if len(got) > 0 != tt.hasErr {
				t.Errorf("TournamentRaw.IsValid() has error %v, want %v", len(got), tt.hasErr)
			}
		})
	}
}

func TestTournamentRaw_IsNameValid(t *testing.T) {
	tests := []struct {
		name   string
		trName string
		want   bool
	}{
		{
			name:   "it should return true if valid",
			trName: nameValid,
			want:   true,
		},
		{
			name:   "it should return false if invalid",
			trName: nameInvalid,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TournamentRaw{
				Name: tt.trName,
			}
			if got := tr.IsNameValid(); got != tt.want {
				t.Errorf("TournamentRaw.IsNameValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTournamentRaw_IsPlayersValid(t *testing.T) {
	tests := []struct {
		name    string
		players []string
		want    bool
	}{
		{
			name:    "it should return true if valid",
			players: playersValid,
			want:    true,
		},
		{
			name:    "it should return false if invalid",
			players: playersInvalid,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TournamentRaw{
				Players: tt.players,
			}
			if got := tr.IsPlayersValid(); got != tt.want {
				t.Errorf("TournamentRaw.IsPlayersValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTournamentRaw_IsMethodValid(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		config       interface{}
		playersCount int
		want         bool
	}{
		{
			name:   "it should return true if method is SINGLE_ELIM and any config",
			method: MethodSingleElim,
			want:   true,
		},
		{
			name:   "it should return true if method is PRELIM_ARRANGED_GROUP and config has valid players_per_group",
			method: MethodPrelimArrangedGroup,
			config: map[string]interface{}{
				"players_per_group": 3,
			},
			playersCount: 6,
			want:         true,
		},
		{
			name:   "it should return false if method is PRELIM_ARRANGED_GROUP and no config",
			method: MethodPrelimArrangedGroup,
			want:   false,
		},
		{
			name:   "it should return false if method is PRELIM_ARRANGED_GROUP and config has players_per_group less than 3",
			method: MethodPrelimArrangedGroup,
			config: map[string]interface{}{
				"players_per_group": 2,
			},
			playersCount: 3,
			want:         false,
		},
		{
			name:   "it should return false if method is PRELIM_ARRANGED_GROUP and config has players_per_group not multiple of players count",
			method: MethodPrelimArrangedGroup,
			config: map[string]interface{}{
				"players_per_group": 3,
			},
			playersCount: 8,
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TournamentRaw{
				Players:      generatePlayers(tt.playersCount),
				Method:       tt.method,
				MethodConfig: tt.config,
			}
			if got := tr.IsMethodValid(); got != tt.want {
				t.Errorf("TournamentRaw.IsMethodValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generatePlayers(count int) (p []string) {
	for i := 0; i < count; i++ {
		p = append(p, fmt.Sprintf("Test Player %d", i))
	}
	return
}
