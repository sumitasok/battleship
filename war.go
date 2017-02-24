package battleship

func Attack(allowedHits int, bg battleGround, attackString string) error {
	position, err := parsePositionString(attackString)
	if err != nil {
		return err
	}

	if len(position) > allowedHits {
		return ErrExceededAllowedHits
	}

	err = validateShipPositions(allowedHits, position)
	if err != nil {
		return err
	}

	for _, p := range position {
		bg.hit(p[0], p[1])
	}

	return nil
}
