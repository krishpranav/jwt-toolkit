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

		var wg sync.WaitGroup
		wg.Add(len(alphabet))

		for i := 1; i <= len(alphabet); i++ {
			go func(i int) {
				Word(alphabet[:i]).Permute(c)

				wg.Done()
			}(i)
		}

		wg.Wait()
	}()

	return c
}

type Word []rune

func (w Word) Permute(out chan<- string) {
	if len(w) <= 1 {
		out <- string(w)
		return
	}

	out <- string(w)

	for w.next() {
		out <- string(w)
	}

}
