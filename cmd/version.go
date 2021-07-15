package cmd

import (
	"fmt"

	"github.com/krishpranav/jwt-toolkit/pkg/printing"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(printing.VERSION)
	},
}
