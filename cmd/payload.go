package cmd

import (
	"github.com/dgrijalva/jwt-go"
	jwtInterface "github.com/krishpranav/jwt-toolkit/pkg/jwt"
	jwtPayload "github.com/krishpranav/jwt-toolkit/pkg/payload"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var payloadCmd = &cobra.Command{
	Use:   "payload [JWT Token]",
	Short: "Generate JWT Attack payloads",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			var token *jwt.Token
			token = jwtInterface.JWTdecode(args[0])
			jwtPayload.GenerateAllPayloads(token)
		} else {
			log.Error("Arguments Error")
			log.Error("e.g jwt-toolkit payload {JWT_CODE}")
		}
	},
}

func init() {
	rootCmd.AddCommand(payloadCmd)
}
