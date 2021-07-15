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
