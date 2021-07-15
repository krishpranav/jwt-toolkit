package cmd

import (
	"encoding/json"
	"fmt"

	jwtInterface "github.com/hahwul/jwt-hack/pkg/jwt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var secret, algo string

var encodeCmd = &cobra.Command{
	Use:   "encode [JSON]",
	Short: "Encode json to JWT",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			mapInterface := []byte(args[0])
			var raw map[string]interface{}
			if err := json.Unmarshal(mapInterface, &raw); err != nil {
				// err
				log.Error("JSON Unmarshal Error")
				panic(0)
			}
			log.WithFields(log.Fields{
				"algorithm": algo,
			}).Info("Encoded result")
			fmt.Println(jwtInterface.JWTencode(raw, secret, algo))
		} else {
			log.Error("Arguments Error")
			log.Error("e.g jwt-hack encode {JWT_CODE} --secret={YOUR_SECRET}")
		}
	},
}
