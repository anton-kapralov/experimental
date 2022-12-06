package console

import (
	"os"
)
import "golang.org/x/term"

const (
	Sigint = 0x03
	Up     = 0x1b5b41
	Down   = 0x1b5b42
	Right  = 0x1b5b43
	Left   = 0x1b5b44
)

type Console struct {
	oldState *term.State
}

func New() (*Console, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}
	return &Console{oldState: oldState}, nil
}

func (c *Console) Close() {
	term.Restore(int(os.Stdin.Fd()), c.oldState)
}

func (c *Console) Read() (int32, error) {
	b := make([]byte, 4)
	n, err := os.Stdin.Read(b)
	if err != nil {
		return 0, err
	}
	var code int32
	for i := 0; i < n; i++ {
		code <<= 8
		code |= int32(b[i])
	}
	return code, nil
}
