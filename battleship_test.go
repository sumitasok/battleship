package battleship

import (
	"testing"
)

func TestValidateGridRange(t *testing.T) {
	if err := validateGridRange(0); err == nil || err != ErrMinGridSize {
		t.Errorf("validation failed to catch the grid size below allowed value")
	}

	if err := validateGridRange(10); err == nil || err != ErrMaxGridSize {
		t.Errorf("validation failed to catch the grid size above allowed value")
	}
}

func TestValidateShipCount(t *testing.T) {
	if err := validateShipCount(10, 0); err == nil || err != ErrMinShipCount {
		t.Errorf("validation failed to catch ship minimum size")
	}

	if err := validateShipCount(10, 6); err == nil || err != ErrMaxShipCount {
		t.Errorf("validation failed to catch maximum allowed ship")
	}

	if err := validateShipCount(10, 5); err != nil || err == ErrMaxShipCount || err == ErrMinShipCount {
		t.Errorf("validation `false negative` caught allowed ship count to be wrong")
	}

	if err := validateShipCount(9, 5); err == nil || err != ErrMaxShipCount {
		t.Errorf("validation failed to catch maximum allowed ship")
	}

	if err := validateShipCount(9, 4); err != nil || err == ErrMaxShipCount || err == ErrMinShipCount {
		t.Errorf("validation `false negative` caught allowed ship count to be wrong")
	}
}
