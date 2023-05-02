package utils

import (
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"strconv"
	"strings"
)

func GenerateUUIDFromEmailAndPhoneNumber(email, phoneNumber string) (string, error) {
	// Generate users parent id
	phoneNumberGenCode := phoneNumber
	emailCode := strings.Split(email, "@")

	if strings.HasPrefix(phoneNumberGenCode, "62") {
		phoneNumberGenCode = phoneNumberGenCode[2:]
	}

	phoneNumberAsInt, err := strconv.ParseInt(phoneNumberGenCode, 10, 64)
	if err != nil {
		return "", customerror.GetError(customerror.InternalServer, err)
	}

	phoneNumberEncode := strings.ToUpper(strconv.FormatInt(phoneNumberAsInt, 36))
	data := []string{emailCode[0], phoneNumberEncode}
	return strings.Join(data, "-"), nil
}
