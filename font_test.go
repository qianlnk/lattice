package lattice

import (
	"fmt"
	"testing"
)

func TestFont(t *testing.T) {
	font := NewFont(24, "", "a", " ")
	buf, err := font.GetBitmap("fontlib", "你")
	if err != nil {
		fmt.Println(err)
		return
	}

	font.Print(buf)
}
