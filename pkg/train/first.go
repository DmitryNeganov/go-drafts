package train

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func DoFirstTest() {
	var inputNumb string

	fmt.Fscanln(reader, &inputNumb)

	if len(inputNumb) < 2 {
		fmt.Fprintln(writer, 0)
		return
	}

	pos := 0
	min := int(inputNumb[pos] - '0')
	prev := int(inputNumb[pos] - '0')

	for i, digit := range inputNumb {
		cur := int(digit - '0')
		if cur > prev {
			pos = i - 1
			break
		}
		if min > cur {
			min = cur
			pos = i
			if cur == 0 {
				break
			}
		}
		prev = cur
	}

	var result strings.Builder
	for i, digit := range inputNumb {
		if i != pos {
			result.WriteRune(digit)
		}
	}

	fmt.Fprintln(writer, result.String())
}
