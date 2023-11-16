package cursor

import (
	"errors"
)

//nolint:gochecknoglobals // Intentional global variable
var Nil = Cursor{Left: "", Right: ""}

type Cursor struct {
	Left  string `json:"left_cursor"`
	Right string `json:"right_cursor"`
}

func newCursor(lc, rc string) (Cursor, error) {
	if lc == "" && rc == "" {
		return Cursor{}, errors.New("pagin: invalid cursor")
	}

	return Cursor{
		Left:  lc,
		Right: rc,
	}, nil
}

func NewLeft(lc string) (Cursor, error) {
	return newCursor(lc, "")
}

func NewRight(rc string) (Cursor, error) {
	return newCursor("", rc)
}

func NewBounded(lc, rc string) (Cursor, error) {
	if lc == "" || rc == "" {
		return Cursor{}, errors.New("pagin: invalid bounded cursor")
	}

	return newCursor(lc, rc)
}

func (c *Cursor) IsEmpty() bool {
	return c.Left == "" && c.Right == ""
}

// Cursor returns can return both Right and Left cursor if those are populated.
// In case if none of those fields are populated, empty string will be returned.
func (c *Cursor) Cursor() string {
	if c.IsRight() {
		return c.Right
	}
	if c.IsLeft() {
		return c.Left
	}

	return ""
}

func (c *Cursor) HasLeft() bool {
	return c.Left != ""
}

func (c *Cursor) HasRight() bool {
	return c.Right != ""
}

func (c *Cursor) IsLeft() bool {
	return c.Left != "" && c.Right == ""
}

func (c *Cursor) IsRight() bool {
	return c.Left == "" && c.Right != ""
}

func (c *Cursor) IsBounded() bool {
	return c.Left != "" && c.Right != ""
}
