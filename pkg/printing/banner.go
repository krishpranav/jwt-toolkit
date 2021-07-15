package printing

import (
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora"
)

func Banner() {
	out("-----------")
	out("JWT-TOOLKIT")
	out("-----------")
}

func out(text string) {
	fmt.Fprintln(os.Stderr, Cyan(text))
}
