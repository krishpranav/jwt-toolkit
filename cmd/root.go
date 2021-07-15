package cmd

import (
	"fmt"
	"os"

	printing "github.com/krishpranav/jwt-toolkit/pkg/printing"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "jwt-toolkit",
	Short: "Hack the JWT(JSON Web Token) | by @hahwul | " + printing.VERSION,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	printing.Banner()
}

func initConfig() {

}
