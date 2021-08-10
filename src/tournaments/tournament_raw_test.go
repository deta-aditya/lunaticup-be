package tournaments

import (
	"testing"
)

var (
	nameValid   = "Test Name"
	nameInvalid = ""

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
		name     string
		tr       TournamentRaw
		wantErrs []error
	}{
		{
			name: "it should return empty errors when valid",
			tr: TournamentRaw{
				Name:    nameValid,
				Players: playersValid,
			},
			wantErrs: []error{},
		},
		{
			name: "it should return some errors when invalid",
			tr: TournamentRaw{
				Name:    nameInvalid,
				Method:  "SINGLE_ELIM",
				Players: playersInvalid,
			},
			wantErrs: []error{
				ErrInvalidName,
				ErrInvalidPlayers,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.IsValid()
			for i, v := range got {
				if tt.wantErrs[i] != v {
					t.Errorf("TournamentRaw.IsValid()[%d] = %v, want %v", i, got[i], tt.wantErrs[i])
				}
			}
		})
	}
}

func TestTournamentRaw_IsNameValid(t *testing.T) {
	tests := []struct {
		name string
		tr   TournamentRaw
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.IsNameValid(); got != tt.want {
				t.Errorf("TournamentRaw.IsNameValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTournamentRaw_IsPlayersValid(t *testing.T) {
	tests := []struct {
		name string
		tr   TournamentRaw
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.IsPlayersValid(); got != tt.want {
				t.Errorf("TournamentRaw.IsPlayersValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
