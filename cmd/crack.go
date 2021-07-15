package cmd

import (
	crack "github.com/hahwul/jwt-hack/pkg/crack"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var wordlist, chars, mode string
var max, conc int
var power, verbose bool

var crackCmd = &cobra.Command{
	Use:   "crack [JWT Token]",
	Short: "Cracking JWT Token",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			if mode == "dict" {
				crack.Crack(mode, args[0], wordlist, conc, max, power, verbose)
			} else if mode == "brute" {
				crack.Crack(mode, args[0], chars, conc, max, power, verbose)
			}
		} else {
			log.Error("Arguments Error")
			log.Error("e.g jwt-hack crack {JWT_CODE} -w {WORDLIST}")
		}
	},
}
