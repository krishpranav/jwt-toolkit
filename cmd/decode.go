package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	jwtInterface "github.com/hahwul/jwt-hack/pkg/jwt"

	//. "github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode [JWT Token]",
	Short: "Decode JWT to JSON",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			var token *jwt.Token
			var log = logrus.New()
			var jdata map[string]interface{}

			log.Out = os.Stdout
			token = jwtInterface.JWTdecode(args[0])
			header, _ := json.Marshal(token.Header)
			log.WithFields(logrus.Fields{
				"method": token.Method,
				"header": string(header),
			}).Info("Decoded data(claims)")

			data, _ := json.Marshal(token.Claims)
			json.Unmarshal([]byte(data), &jdata)

		}

	},
}
