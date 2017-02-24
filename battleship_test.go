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

func TestValidateShipCount(t *testing.T) {
	if validateShipCount(10, 0) == nil {
		t.Errorf("validation failed to catch ship minimum size")
	}

	if validateShipCount(10, 6) == nil {
		t.Errorf("validation failed to catch maximum allowed ship")
	}

	if validateShipCount(10, 5) != nil {
		t.Errorf("validation `false negative` caught allowed ship count to be wrong")
	}

	if validateShipCount(9, 5) == nil {
		t.Errorf("validation failed to catch maximum allowed ship")
	}

	if validateShipCount(9, 4) != nil {
		t.Errorf("validation `false negative` caught allowed ship count to be wrong")
	}
}
