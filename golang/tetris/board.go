package tetris

const (
	height = 20
	width  = 10
)

type Board [height][width]int

func (b *Board) removeFullRows() {
	offset := 0
	for i := height - 1; i >= 0; i-- {
		full := true
		for j := 0; j < width; j++ {
			if b[i][j] == 0 {
				full = false
				break
			}
		}
		if full {
			offset++
			continue
		}
		if offset == 0 {
			continue
		}
		// Move row down by offset.
		for j := 0; j < width; j++ {
			b[i+offset][j] = b[i][j]
		}
	}

	if offset == 0 {
		return
	}

	// Wipe moved rows.
	for i := 0; i < offset; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = 0
		}

	}

	// Drop "flying" bricks.
	for j := 0; j < width; j++ {
		if b[height-1][j] != 0 {
			continue
		}
		drop := 1
		for i := height - 2; i >= 0; i-- {
			if b[i][j] != 0 {
				break
			}
			drop++
		}

		for i := height - 1 - drop; i >= 0; i-- {
			b[i+drop][j] = b[i][j]
			b[i][j] = 0
		}
	}
}
