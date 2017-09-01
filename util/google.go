package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../service"
	"github.com/kataras/iris/core/errors"
)

type GoogleIDTokenInfo struct {
	ISS               string  `json:"iss"`
	SUB               string  `json:"sub"`
	AZP               string  `json:"azp"`
	AUD               string  `json:"aud"`
	IAT               string  `json:"iat"`
	EXP               string  `json:"exp"`

	Email             string  `json:"email"`
	EmailVerified     string  `json:"email_verified"`
	Name              string  `json:"name"`
	Picture           string  `json:"picture"`
	GivenName         string  `json:"given_name"`
	FamilyName        string  `json:"family_name"`
	Locale            string  `json:"locale"`

	ErrorDescription  string  `json:"error_description"`
}

func VerifyGoogleIDToken(token string) (info GoogleIDTokenInfo, err error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=" + token)

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	err = json.Unmarshal(body, &info)

	if err != nil {
		return
	}

	if info.SUB == "" {
		err = errors.New(info.ErrorDescription)
	} else if info.AUD != service.Config.GoogleClientID {
		err = errors.New("Google Client ID mismatch!")
	}
	/*
	// Google has already checked them
	else if info.ISS != "accounts.google.com" && info.ISS != "https://accounts.google.com" {
		err = errors.New("Token is not issued by Google!")
	} else if exp, _ := strconv.Atoi(info.EXP); time.Now().Unix() > int64(exp) {
		err = errors.New("Token expired!")
	}
	*/

	return
}
