package battleship

import (
	"errors"
)

const (
	MIN_GRID_SIZE = 1
	MAX_GRID_SIZE = 9

	GROUND_BLANK = "_"
)

var (
	ErrMinGridSize  error = errors.New("Grid size needs to be more than " + string(MIN_GRID_SIZE))
	ErrMaxGridSize  error = errors.New("Grid size needs to be less than " + string(MIN_GRID_SIZE))
	ErrMinShipCount error = errors.New("Bring you ships! ship count need to be more than 0")
	ErrMaxShipCount error = errors.New("Bring you ships! ship count cannot be more than half of the grid size")
)

type battleGround struct {
	positions map[int]map[int]string
	m         int
	s         int
}

// init return an initialised battle ground with all blank grid.
func (b *battleGround) init(m int) error {
	b.m = m
	b.positions = make(map[int]map[int]string)

	for i := 0; i < m; i++ {
		b.positions[i] = make(map[int]string)
		for j := 0; j < m; j++ {
			b.positions[i][j] = GROUND_BLANK
		}
	}

	return nil
}

func (b *battleGround) String() string {
	layout := ""
	for i := 0; i < b.m; i++ {
		for j := 0; j < b.m; j++ {
			layout += b.positions[i][j]
		}
		layout += "\n"
	}
	return layout
}

// NewBattleGround creates a new ballte ground with the grid size and ship count specified
// if conditions are matched, and ship positions are added.
func NewBattleGround(m, s int, positions string) (battleGround, error) {
	if err := validateGridRange(m); err != nil {
		return battleGround{}, err
	}

	if err := validateShipCount(m, s); err != nil {
		return battleGround{}, err
	}

	bg := battleGround{}
	err := bg.init(m)
	if err != nil {
		return battleGround{}, err
	}

	return bg, nil
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
