package lattice

import (
	"fmt"
	"testing"
)

func TestFont(t *testing.T) {
	font := NewFont(12, "", "方", "  ")
	buf, err := font.GetBitmap("fontlib", "方")
	if err != nil {
		fmt.Println(err)
		return
	}

	font.Print(buf)

	font = NewFont(24, "", "佳", "  ")
	buf, err = font.GetBitmap("fontlib", "佳")
	if err != nil {
		fmt.Println(err)
		return
	}

	font.Print(buf)

	font = NewFont(16, "", "丽", "  ")
	buf, err = font.GetBitmap("fontlib", "丽")
	if err != nil {
		fmt.Println(err)
		return
	}

	font.Print(buf)
}
