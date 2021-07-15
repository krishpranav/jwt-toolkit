package crack

import "sync"

func GenerateBruteForcePaylods(chars string) []string {
	var payloads []string
	for str := range generate(chars) {
		payloads = append(payloads, str)
	}
	return payloads
}

func generate(alphabet string) <-chan string {
	c := make(chan string, len(alphabet))

	go func() {
		defer close(c)

		if len(alphabet) == 0 {
			return
		}
	}()
}

type Word []rune
