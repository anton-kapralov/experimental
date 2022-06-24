package csutil

type Deque struct {
	deque []interface{}
}

func (d *Deque) Len() int {
	return len(d.deque)
}

func (d *Deque) IsEmpty() bool {
	return len(d.deque) == 0
}

func (d *Deque) AddLast(v interface{}) {
	d.deque = append(d.deque, v)
}

func (d *Deque) RemoveLast() interface{} {
	last := len(d.deque) - 1
	if last < 0 {
		return nil
	}
	v := d.deque[last]
	d.deque = d.deque[:last]
	return v
}
