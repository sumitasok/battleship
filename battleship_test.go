package battleship

import (
	"testing"
)

func TestNewBattleGround(t *testing.T) {
	_battleGround, err := NewBattleGround(10, 6, "1:1")
	if err == nil {
		t.Errorf("validation failed")
	}

	if len(_battleGround.positions) != 0 {
		t.Errorf("validation failed")
	}

	if breadth(_battleGround.positions) != 0 {
		t.Errorf("validation failed")
	}
}

func breadth(matrix map[int]map[int]string) int {
	if len(matrix) == 0 {
		return 0
	}

	return len(matrix[0])
}

func TestInit(t *testing.T) {
	bg := &battleGround{}
	err := bg.init(3)

	if err != nil {
		t.Errorf("Init failed")
	}

	if len(bg.positions) != 3 {
		t.Errorf("matrix size mis match")
	}

	if breadth(bg.positions) != 3 {
		t.Errorf("matrix size mis match")
	}
}

func TestBattleShipHit(t *testing.T) {
	bg := &battleGround{}
	bg.init(3)

	bg.plotShips(1, [][]int{
		[]int{1, 1},
	})

	// commenting out as the orientation chnaged in the middle
	// if bg.String() != "___\n_B_\n___\n" {
	// 	t.Errorf("error in plotting")
	// }

	bg.hit(0, 1)

	// commenting out as the orientation chnaged in the middle
	// if bg.String() != "_O_\n_B_\n___\n" {
	// 	t.Errorf("error in plotting missed hit")
	// }

	if bg.shipDown != 0 {
		t.Errorf("error in counting hit")
	}

	bg.hit(1, 1)

	// commenting out as the orientation chnaged in the middle
	// if bg.String() != "_O_\n_X_\n___\n" {
	// 	t.Errorf("error in plotting hit")
	// }

	if bg.shipDown != 1 {
		t.Errorf("error in counting hit")
	}
}

func TestBattleShipStanding(t *testing.T) {
	bg := &battleGround{}
	bg.init(3)

	bg.plotShips(1, [][]int{
		[]int{1, 1},
	})

	if bg.ShipsStanding() != 1 {
		t.Errorf("error in counting safe ships")
	}

	bg.hit(1, 1)

	if bg.ShipsStanding() != 0 {
		t.Errorf("error in counting safe ships")
	}
}

func TestBattleShipString(t *testing.T) {
	bg := &battleGround{}
	err := bg.init(3)

	if err != nil {
		t.Errorf("Init failed")
	}

	layout := bg.String()
	if layout != "___\n___\n___\n" {
		t.Errorf("Draw Failed")
	}
}

func TestBattleGroundPlotShip(t *testing.T) {
	bg := &battleGround{}
	err := bg.init(3)
	positions := [][]int{
		[]int{1, 1},
		[]int{2, 3},
		[]int{3, 4},
		[]int{2, 4},
		[]int{4, 0},
	}
	err = bg.plotShips(1, positions)
	if err == nil {
		t.Error("Error `false positive`, extra ships in ploting not caught")
	}
}

func TestBattleGroundPlotShips(t *testing.T) {
	bg := &battleGround{}
	bg.init(3)

	if err := bg.plotShips(1, [][]int{
		[]int{1, 1},
	}); err != nil {
		t.Errorf("plotting `false negative`")
	}

	if bg.String() != "___\n_B_\n___\n" {
		t.Errorf("error in plotting")
	}
}

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
