package battleship

import (
	"testing"
)

func TestAttack(t *testing.T) {
	bg := battleGround{}
	attack := "1:1"
	Attack(bg, attack)
}
