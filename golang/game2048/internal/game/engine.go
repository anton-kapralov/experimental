package game

import "math/rand"

type Direction int

const (
	DirectionUnknown Direction = iota
	DirectionLeft
	DirectionUp
	DirectionRight
	DirectionDown
)

type Status int

const (
	StatusUnknown Status = iota
	StatusPlaying
	StatusOver
)

type State struct {
	Status Status    `json:"status"`
	Score  int       `json:"score"`
	Board  [4][4]int `json:"board"`
}

func New(rng *rand.Rand) *State {
	s := &State{
		Status: StatusPlaying,
	}
	for i := 0; i < 2; i++ {
		v := 2
		if rng.Int()%4 == 0 {
			v = 4
		}
		idx := rng.Intn(16)
		for s.Board[idx/4][idx%4] != 0 {
			idx = rng.Intn(16)
		}
		s.Board[idx/4][idx%4] = v
	}
	return s
}

func (s *State) Move(direction Direction, rng *rand.Rand) *State {
	newState := &State{
		Status: s.Status,
		Score:  s.Score,
	}
	for i := 0; i < 4; i++ {
		copy(newState.Board[i][:], s.Board[i][:])
	}
	switch direction {
	case DirectionLeft:
		for i := 0; i < 4; i++ {
			shiftLeft(&newState.Board[i])
			newState.Score += collapseToLeft(&newState.Board[i])
			shiftLeft(&newState.Board[i])
		}
	case DirectionUp:
		for i := 0; i < 4; i++ {
			shiftUp(i, &newState.Board)
			newState.Score += collapseToTop(i, &newState.Board)
			shiftUp(i, &newState.Board)
		}
	case DirectionRight:
		for i := 0; i < 4; i++ {
			shiftRight(&newState.Board[i])
			newState.Score += collapseToRight(&newState.Board[i])
			shiftRight(&newState.Board[i])
		}
	case DirectionDown:
		for i := 0; i < 4; i++ {
			shiftDown(i, &newState.Board)
			newState.Score += collapseToBottom(i, &newState.Board)
			shiftDown(i, &newState.Board)
		}
	}

	var emptyCells []int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if newState.Board[i][j] != 0 {
				continue
			}
			idx := i*4 + j
			emptyCells = append(emptyCells, idx)
		}
	}

	if len(emptyCells) > 0 {
		v := 2
		if rng.Int()%4 == 0 {
			v = 4
		}
		idx := emptyCells[rng.Intn(len(emptyCells))]
		newState.Board[idx/4][idx%4] = v
	}

	if len(emptyCells) == 1 {
		canMove := false
		for i := 0; i < 4; i++ {
			for j := 1; j < 4; j++ {
				if newState.Board[i][j-1] == newState.Board[i][j] ||
					newState.Board[j-1][i] == newState.Board[j][i] {
					canMove = true
					break
				}
			}
		}
		if !canMove {
			newState.Status = StatusOver
		}
	}
	return newState
}

func shiftLeft(row *[4]int) {
	l := 0
	for ; l < 4; l++ {
		if row[l] == 0 {
			break
		}
	}
	r := l + 1
	for ; l < 4; l, r = l+1, r+1 {
		for ; r < 4; r++ {
			if row[r] != 0 {
				break
			}
		}
		if r >= 4 {
			break
		}
		row[l] = row[r]
	}
	for ; l < 4; l++ {
		row[l] = 0
	}
}

func shiftRight(row *[4]int) {
	r := 3
	for ; r >= 0; r-- {
		if row[r] == 0 {
			break
		}
	}
	l := r - 1
	for ; r >= 0; l, r = l-1, r-1 {
		for ; l >= 0; l-- {
			if row[l] != 0 {
				break
			}
		}
		if l < 0 {
			break
		}
		row[r] = row[l]
	}
	for ; r >= 0; r-- {
		row[r] = 0
	}
}

func shiftUp(col int, board *[4][4]int) {
	u := 0
	for ; u < 4; u++ {
		if board[u][col] == 0 {
			break
		}
	}
	l := u + 1
	for ; u < 4; u, l = u+1, l+1 {
		for ; l < 4; l++ {
			if board[l][col] != 0 {
				break
			}
		}
		if l >= 4 {
			break
		}
		board[u][col] = board[l][col]
	}
	for ; u < 4; u++ {
		board[u][col] = 0
	}
}

func shiftDown(col int, board *[4][4]int) {
	l := 3
	for ; l >= 0; l-- {
		if board[l][col] == 0 {
			break
		}
	}
	u := l - 1
	for ; l >= 0; u, l = u-1, l-1 {
		for ; u >= 0; u-- {
			if board[u][col] != 0 {
				break
			}
		}
		if u < 0 {
			break
		}
		board[l][col] = board[u][col]
	}
	for ; l >= 0; l-- {
		board[l][col] = 0
	}
}

func collapseToLeft(row *[4]int) int {
	score := 0
	for i := 1; i < 4; i++ {
		if row[i-1] == 0 || row[i-1] != row[i] {
			continue
		}
		row[i-1] *= 2
		score += row[i-1]
		row[i] = 0
		i++
	}
	return score
}

func collapseToTop(col int, board *[4][4]int) int {
	score := 0
	for i := 1; i < 4; i++ {
		if board[i-1][col] == 0 || board[i-1][col] != board[i][col] {
			continue
		}
		board[i-1][col] *= 2
		score += board[i-1][col]
		board[i][col] = 0
		i++
	}
	return score
}

func collapseToRight(row *[4]int) int {
	score := 0
	for i := 2; i >= 0; i-- {
		if row[i+1] == 0 || row[i+1] != row[i] {
			continue
		}
		row[i+1] *= 2
		score += row[i+1]
		row[i] = 0
		i--
	}
	return score
}

func collapseToBottom(col int, board *[4][4]int) int {
	score := 0
	for i := 2; i >= 0; i-- {
		if board[i+1][col] == 0 || board[i+1][col] != board[i][col] {
			continue
		}
		board[i+1][col] *= 2
		score += board[i+1][col]
		board[i][col] = 0
		i--
	}
	return score
}
