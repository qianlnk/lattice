package main

import (
	"fmt"
	"os"

	"gopkg.in/iconv.v1"
)

func main() {
	hzk16, err := os.Open("fontlib/HZK14")
	if err != nil {
		fmt.Println("open HZK16 fail")
		return
	}

	cd, err := iconv.Open("gb2312", "utf-8")
	if err != nil {
		fmt.Println("open iconv fail")
		return
	}
	defer cd.Close()

	word := []byte(cd.ConvString("敢"))

	fmt.Println(word)
	key := []byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}
	ph := int64(word[0]) - int64(0xa0)
	pl := int64(word[1]) - int64(0xa0)
	offset := ((ph-1)*94 + (pl - 1)) * 28
	ret, err := hzk16.Seek(int64(offset), os.SEEK_SET)
	fmt.Println(ret, err, offset, ph, pl, word[0], word[1], 0xa0)

	buffer := make([]byte, 28)
	_, err = hzk16.Read(buffer)
	if err != nil {
		fmt.Println("read HZK16 fail")
	}

	for i := 0; i < 28; i++ {
		fmt.Printf("%02x ", buffer[i])
	}
	fmt.Println()
	for i := 0; i < 14; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 8; k++ {
				flag := buffer[i*2+j] & key[k]
				if flag != 0 {
					fmt.Printf("●")
				} else {
					fmt.Printf("○")
				}
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

	hzk16.Close()
}
