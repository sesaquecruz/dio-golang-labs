package game

func PipPan(value int) *string {
	if value%3 != 0 && value%5 != 0 {
		return nil
	}

	var text string

	if value%3 != 0 {
		text = "Pan"
		return &text
	}

	if value%5 != 0 {
		text = "Pin"
		return &text
	}

	text = "PinPan"
	return &text
}
