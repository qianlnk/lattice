package lattice

import (
	"fmt"
	"runtime"

	"os"

	"strconv"

	"strings"

	"gopkg.in/iconv.v1"
)

type Font struct {
	Size       int
	Typeface   string
	Foreground string
	Background string
	Code       string
}

func NewFont(size int, typeface, fore, back string) *Font {
	var code string
	if runtime.GOOS == "linux" {
		code = "utf-8"
	} else {
		code = "gb2312"
	}

	return &Font{
		Size:       size,
		Typeface:   typeface,
		Foreground: fore,
		Background: back,
		Code:       code,
	}
}

func convString(word string) string {
	cd, err := iconv.Open("gb2312", "utf-8")
	if err != nil {
		return ""
	}
	defer cd.Close()

	return cd.ConvString(word)
}

func (f *Font) GetZM(path string, words []byte) ([]byte, error) {
	ascFile := strings.Trim(path, "/") + "/ASC" + strconv.Itoa(int(f.Size))
	asc, err := os.Open(ascFile)
	if err != nil {
		return nil, err
	}
	defer asc.Close()
	return nil, nil
}

func (f *Font) GetHZ(path string, words []byte) ([]byte, error) {
	hzkFile := strings.Trim(path, "/") + "/HZK" + strconv.Itoa(int(f.Size))
	hzk, err := os.Open(hzkFile)
	if err != nil {
		return nil, err
	}
	defer hzk.Close()

	ph := int(words[0]) - int(0xa0)
	pl := int(words[1]) - int(0xa0)
	var size int
	size = f.Size
	if f.Size <= 16 {
		size = 16
	}
	offset := ((ph-1)*94 + (pl - 1)) * f.Size * (size / 8)
	_, err = hzk.Seek(int64(offset), os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	fmt.Println("16*(f.Size/8)", f.Size*(size/8))
	res := make([]byte, f.Size*(size/8))
	_, err = hzk.Read(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (f *Font) GetBitmap(path, word string) ([]byte, error) {
	if f.Code == "utf-8" {
		word = convString(word)
	}

	words := []byte(word)
	fmt.Println("words", words)
	if len(words) < 2 {
		return f.GetZM(path, words)
	}

	return f.GetHZ(path, words)
}

func (f *Font) Print(buf []byte) {
	key := []byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}
	fmt.Println(len(buf))
	for i := 0; i < len(buf); i++ {
		fmt.Printf("%02x ", buf[i])
	}
	fmt.Println()
	var size int
	size = f.Size
	if f.Size <= 16 {
		size = 16
	}
	for i := 0; i < f.Size; i++ {
		for j := 0; j < size/8; j++ {
			for k := 0; k < 8; k++ {
				flag := buf[i*2+j] & key[k]
				if flag != 0 {
					fmt.Printf(f.Foreground)
				} else {
					fmt.Printf(f.Background)
				}
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
