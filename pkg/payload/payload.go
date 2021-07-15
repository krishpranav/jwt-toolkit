package paylaod

import (
	b64 "encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func GenerateAllPayloads(token *jwt.Token) {
	log.Out = os.Stdout
	GenerateNonePayloads(token.Raw)
	GenerateUrlPayloads(token.Raw)
}
