package crack

import "sync"

func GenerateBruteforcePayloads(chars string) []string {
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

func (w Word) next() bool {
	var left, right int

	left = len(w) - 2
	for w[left] >= w[left+1] && left >= 1 {
		left--
	}

	if left == 0 && w[left] >= w[left+1] {
		return false
	}

	right = len(w) - 1
	for w[left] >= w[right] {
		right--
	}

	w[left], w[right] = w[right], w[left]

	left++
	right = len(w) - 1

	for left < right {
		w[left], w[right] = w[right], w[left]
		left++
		right--
	}

	return true
}
