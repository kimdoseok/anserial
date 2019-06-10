// anserial project main.go
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

const digits = "0123456789ABCDEF"

//const digits = "ABCDEF"
const prefix = ""

type Anserial struct {
	digits string
	prefix string
}

func (ans Anserial) getNext(anum string) string {
	dlen := utf8.RuneCountInString(digits)
	plen := utf8.RuneCountInString(ans.prefix)
	bnum := string([]rune(anum)[plen:])
	blen := utf8.RuneCountInString(bnum)
	//bnum := anum

	validnum := true
	for _, rv := range bnum {
		if strings.IndexRune(digits, rv) < 0 {
			validnum = false
			break
		}
	}
	if !validnum {
		return ans.prefix + string(digits[0])
	}

	zero := false
	if string(digits[0]) == "0" {
		zero = true
	}
	overflow := false
	for i, _ := range bnum {
		pos := blen - 1 - 1*i
		adigit := string([]rune(bnum)[pos])
		idx := strings.Index(digits, adigit)
		if idx < 0 {
			return ans.prefix + string(digits[0])
		}
		if idx == dlen-1 {
			overflow = true
		}
		//		fmt.Println(string(digits[0]))
		if pos == blen-1 {
			if idx == dlen-1 {
				bnum = string([]rune(bnum)[:pos]) + string([]rune(digits)[0])
				overflow = true
				continue
			} else {
				bnum = bnum[:pos] + string([]rune(digits)[idx+1])
				return ans.prefix + bnum

			}
		}
		if overflow {
			if idx == dlen-1 {
				bnum = string([]rune(bnum)[:pos]) + string([]rune(digits)[0]) + string([]rune(bnum)[pos+1:])
				overflow = true
				continue
			} else {
				bnum = string([]rune(bnum)[:pos]) + string([]rune(digits)[idx+1]) + string([]rune(bnum)[pos+1:])
				return ans.prefix + bnum
			}
		} else {
			return ans.prefix + bnum

		}

	}

	if overflow {
		if zero {
			bnum = string([]rune(digits)[1]) + bnum
		} else {
			bnum = string([]rune(digits)[0]) + bnum
		}
	}

	return ans.prefix + bnum
}

func main() {
	ans := Anserial{digits, prefix}
	num := prefix + string(digits[0])
	f, err := os.Create("anserial.txt")
	if err != nil {
		log.Println(err)
		f.Close()
		return
	}
	for i := 1; i < 10000; i++ {
		fmt.Println(i, num)
		_, err = f.WriteString(fmt.Sprintf("%d %s\n", i, num))
		if err != nil {
			log.Println(err)
			f.Close()
			return
		}
		num = ans.getNext(num)
	}
	err = f.Close()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Done")
}
