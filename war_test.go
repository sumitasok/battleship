package battleship

import (
	"testing"
)

func TestAttack(t *testing.T) {
	bg, err := NewBattleGround(9, 4, "1:1,2:3,3:4,2:4")
	if err != nil {
		panic(err)
	}

	if bg.String() != "_________\n_B_______\n___BB____\n____B____\n_________\n_________\n_________\n_________\n_________\n" {
		t.Errorf("plotting error")
	}

	attack := "1:1,2:1"
	err = Attack(1, bg, attack)
	if err == nil {
		t.Errorf("Error adhering to attach count")
	}

	err = Attack(2, bg, attack)

	if bg.String() != "_________\n_X_______\n_O_BB____\n____B____\n_________\n_________\n_________\n_________\n_________\n" {
		t.Errorf("plotting error")
	}
}
