package cmd

import (
	crack "github.com/krishpranav/jwt-toolkit/pkg/crack"
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
			log.Error("e.g jwt-toolkit crack {JWT_CODE} -w {WORDLIST}")
		}
	},
}

func init() {
	rootCmd.AddCommand(crackCmd)

	crackCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dict", "cracking mode, you can use 'dict' or 'brute'")
	crackCmd.PersistentFlags().StringVarP(&wordlist, "wordlist", "w", "", "wordlist file / only dictionary attack")
	crackCmd.PersistentFlags().StringVar(&chars, "chars", "abcdefghijklmnopqrstuvwxyz0123456789", "char list / only bruteforce")
	crackCmd.PersistentFlags().IntVarP(&conc, "concurrency", "c", 100, "number of concurrency")
	crackCmd.PersistentFlags().IntVar(&max, "max", 6, "max length / only bruteforce")
	crackCmd.PersistentFlags().BoolVar(&power, "power", false, "Used all CPU your computer")
	crackCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Show testing log")

}
