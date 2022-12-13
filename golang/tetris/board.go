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

	b.dropFlyingBricks()
}

var stickingDirections = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (b *Board) dropFlyingBricks() {
	//TODO: implement me.
}
