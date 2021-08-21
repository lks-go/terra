package fight_test

import (
	"fmt"
	"testing"

	"github.com/lks-go/terra/pkg/fight"
)

func TestFight_Status(t *testing.T) {

	tests := []struct {
		StatusName string
		StatusCode int
		IsError    bool
	}{
		{StatusName: "Unknown", StatusCode: 0, IsError: true},
		{StatusName: "Created", StatusCode: 1, IsError: false},
		{StatusName: "Going", StatusCode: 2, IsError: false},
		{StatusName: "Finished", StatusCode: 3, IsError: false},
		{StatusName: "One more unknown status", StatusCode: 2021, IsError: true},
	}

	f := fight.New(&fight.Config{}, nil)

	for _, tt := range tests {
		testName := fmt.Sprintf("Status name: %s, satus code %d", tt.StatusName, tt.StatusCode)
		t.Run(testName, func(t *testing.T) {
			if err := f.SetStatus(tt.StatusCode); err == nil && tt.IsError {
				t.Errorf("expected error, got nil")
			}

			// test must be failed before than escape the function
			if tt.IsError {
				return
			}

			gotStatus := f.Status()
			if gotStatus != tt.StatusCode {
				t.Errorf("expected status code %d, got %d", tt.StatusCode, gotStatus)
			}
		})
	}

}
