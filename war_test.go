package battleship

import (
	"testing"
)

func TestAttack(t *testing.T) {
	bg, err := NewBattleGround(9, 4, "1:1,2:3,3:4,2:4")
	if err != nil {
		panic(err)
	}

	attack := "1:1,2:1"
	err = Attack(1, &bg, attack)
	if err == nil {
		t.Errorf("Error adhering to attach count")
	}

	err = Attack(2, &bg, attack)
}

func TestWinner(t *testing.T) {
	bg1, _ := NewBattleGround(9, 4, "1:1,2:3,3:4,2:4")
	bg1.shipDown = 1

	bg2, _ := NewBattleGround(9, 4, "1:1,2:3,3:4,2:4")
	bg2.shipDown = 2

	if Winner(&bg1, &bg2) != 1 {
		t.Errorf("Error finding the winner")
	}

	bg1, _ = NewBattleGround(9, 4, "1:1,2:3,3:4,2:4")
	bg1.shipDown = 2

	bg2, _ = NewBattleGround(9, 4, "1:1,2:3,3:4,2:4")
	bg2.shipDown = 1

	if Winner(&bg1, &bg2) != 2 {
		t.Errorf("Error finding the winner")
	}

	bg1, _ = NewBattleGround(9, 4, "1:1,2:3,3:4,2:4")
	bg1.shipDown = 2

	bg2, _ = NewBattleGround(9, 4, "1:1,2:3,3:4,2:4")
	bg2.shipDown = 2

	if Winner(&bg1, &bg2) != 0 {
		t.Errorf("Error finding the tie")
	}
}
