package battleship

import (
	"errors"
	"strconv"
	"strings"
)

const (
	MIN_GRID_SIZE = 1
	MAX_GRID_SIZE = 9

	GROUND_BLANK      = "_"
	ALIVE_BATTLESHIPS = "B"
)

var (
	ErrMinGridSize             error = errors.New("Grid size needs to be more than " + string(MIN_GRID_SIZE))
	ErrMaxGridSize             error = errors.New("Grid size needs to be less than " + string(MIN_GRID_SIZE))
	ErrMinShipCount            error = errors.New("Bring you ships! ship count need to be more than 0")
	ErrMaxShipCount            error = errors.New("Bring you ships! ship count cannot be more than half of the grid size")
	ErrIncorrectPositionString error = errors.New("Position string incorrect")
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

func (b *battleGround) plotShips(s int, positions [][]int) error {
	if err := validateShipPositions(s, positions); err != nil {
		return err
	}

	for _, ship := range positions {
		b.positions[ship[0]][ship[1]] = ALIVE_BATTLESHIPS
	}

	return nil
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

	position, err := parsePositionString(positions)
	if err != nil {
		return battleGround{}, err
	}

	bg := battleGround{}
	err = bg.init(m)
	if err != nil {
		return battleGround{}, err
	}

	bg.plotShips(s, position)

	return bg, nil
}

func parsePositionString(position string) ([][]int, error) {
	points := [][]int{}

	pointStrArray := strings.Split(position, ",")

	for _, pointStr := range pointStrArray {
		pointS := strings.Split(pointStr, ":")

		if len(pointS) != 2 {
			return [][]int{}, ErrIncorrectPositionString
		}

		x, err := strconv.Atoi(pointS[0])
		if err != nil {
			return [][]int{}, ErrIncorrectPositionString
		}

		y, err := strconv.Atoi(pointS[1])
		if err != nil {
			return [][]int{}, ErrIncorrectPositionString
		}

		points = append(points, []int{x, y})
	}

	return points, nil
}

func validateShipPositions(s int, position [][]int) error {
	if len(position) != s {
		return ErrIncorrectPositionString
	}

	return nil
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
