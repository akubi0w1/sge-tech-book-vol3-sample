package closer

type Closer struct {
	fns []func()
}

func (c *Closer) Add(fn func()) {
	c.fns = append(c.fns, fn)
}

func (c *Closer) Close() {
	for i := len(c.fns) - 1; i >= 0; i-- {
		c.fns[i]()
	}
}
