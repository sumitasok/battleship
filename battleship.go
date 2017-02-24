package battleship

import (
	"errors"
)

const (
	MIN_GRID_SIZE = 1
	MAX_GRID_SIZE = 9
)

var (
	ErrMinGridSize  error = errors.New("Grid size needs to be more than " + string(MIN_GRID_SIZE))
	ErrMaxGridSize  error = errors.New("Grid size needs to be less than " + string(MIN_GRID_SIZE))
	ErrMinShipCount error = errors.New("Bring you ships! ship count need to be more than 0")
	ErrMaxShipCount error = errors.New("Bring you ships! ship count cannot be more than half of the grid size")
)

type battleGround map[int]map[int]string

func NewBattleGround(m, s int, positions string) (battleGround, error) {
	return battleGround{}, nil
}

func validateGridRange(m int) error {
	if m < MIN_GRID_SIZE {
		return ErrMinGridSize
	}

	if m > MAX_GRID_SIZE {
		return ErrMaxGridSize
	}

	return nil
}

func validateShipCount(m, s int) error {
	if s == 0 {
		return ErrMinShipCount
	}

	if s > m/2 {
		return ErrMaxShipCount
	}

	return nil
}
