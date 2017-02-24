package battleship

import (
	"testing"
)

func TestValidateGridRange(t *testing.T) {
	if validateGridRange(0) == nil {
		t.Errorf("validation failed to catch the grid size below allowed value")
	}

	if validateGridRange(10) == nil {
		t.Errorf("validation failed to catch the grid size above allowed value")
	}
}
