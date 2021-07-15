package crack

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	jwtInterface "github.com/hahwul/jwt-hack/pkg/jwt"
	color "github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Crack(mode, token, data string, concurrency, max int, power bool, verbose bool) {
	log.Out = os.Stdout
	fmt.Println("[*] Start " + mode + " cracking mode")
	if mode == "brute" {
		bf := GenerateBruteforcePayloads(data)
		RunTestingJWT(token, bf, concurrency, verbose)
	} else {
		var words []string
		ff, err := readLinesOrLiteral(data)
		_ = err
		for _, word := range ff {
			words = append(words, word)
		}

		words = unique(words)
		log.WithFields(logrus.Fields{
			"size": len(words),
		}).Info("Loaded words (remove duplicated)")
		RunTestingJWT(token, words, concurrency, verbose)
	}
}

func RunTestingJWT(token string, lists []string, concurrency int, verbose bool) {
	wordlists := make(chan string)
	lenWordlist := len(lists)
	nowLine := 0
	found := false
	secret := ""
	// Add go routine job
	var wg sync.WaitGroup
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
	mutex := &sync.Mutex{}
	if !verbose {
		percent := float64(nowLine / lenWordlist)
		str := fmt.Sprintf(" Cracking.. [%d / %d][%f]", nowLine, lenWordlist, percent)
		s.Color("red", "bold")
		s.Prefix = " "
		s.Suffix = str
		s.Start()
	}
}
