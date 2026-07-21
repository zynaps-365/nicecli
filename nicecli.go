package nicecli

import (
	"fmt"
	"strings"
)

func Print(s string) {
	var border strings.Builder

	for i := 0; i < len([]rune(s)); i++ {
		border.WriteString("*")
	}

	fmt.Println(border.String())
	fmt.Println(s)
	fmt.Println(border.String())
}
