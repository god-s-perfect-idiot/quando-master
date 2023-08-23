package analyser

type Cursor struct {
	line   int
	column int
}

func NewCursor() *Cursor {
	return &Cursor{
		line:   0,
		column: 0,
	}
}

func (c *Cursor) advanceLine() {
	c.line++
	c.column = 0
}

func (c *Cursor) advanceColumn() {
	c.column++
}

func (c *Cursor) getL() int {
	return c.line
}

func (c *Cursor) getC() int {
	return c.column
}
