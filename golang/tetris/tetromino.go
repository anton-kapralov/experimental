package tetris

type rotationDirection int

const (
	clockwise        = 1
	counterClockwise = -1
)

type point2D [2]int

func (t *point2D) rotate(center point2D, direction rotationDirection) {
	t[0] -= center[0]
	t[1] -= center[1]

	t[0], t[1] = int(direction)*t[1], int(direction)*-t[0]

	t[0] += center[0]
	t[1] += center[1]
}

type tetromino struct {
	tiles  [4]point2D
	center *point2D
}

func (t *tetromino) move(dr, dc int) {
	for i := range t.tiles {
		t.tiles[i][0] += dr
		t.tiles[i][1] += dc
	}
}

func (t *tetromino) rotate(direction rotationDirection) {
	for i := range t.tiles {
		t.tiles[i].rotate(*t.center, direction)
	}
}
