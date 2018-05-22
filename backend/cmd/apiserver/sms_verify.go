package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/maxmcd/c4/backend/pkg/config"
	"github.com/maxmcd/c4/backend/pkg/models"
)

func verifyRequest(url string, values url.Values) error {
	values.Add("api_key", config.ApiKeys.TwilioVerifyKey)
	resp, err := http.PostForm(url, values)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	fmt.Println(err)
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Unexpected status code from verify api %d", resp.StatusCode))
	}
	return nil
}

func sendVerificationRequest(user models.User) error {
	if user.Phone == nil || user.CountryCode == nil {
		return errors.New("User has no valid phone number")
	}
	values := url.Values{
		"country_code": {fmt.Sprintf("%d", user.CountryCode.Int64)},
		"phone_number": {fmt.Sprintf("%d", user.Phone.Int64)},
		"via":          {"sms"},
	}
	return verifyRequest(
		"https://api.authy.com/protected/json/phones/verification/start",
		values,
	)
}

func confirmVerificationRequest(user models.User, code string) error {
	if user.Phone == nil || user.CountryCode == nil {
		return errors.New("User has no valid phone number")
	}
	values := url.Values{
		"country_code": {fmt.Sprintf("%d", user.CountryCode.Int64)},
		"phone_number": {fmt.Sprintf("%d", user.Phone.Int64)},
		"code":         {code},
	}

	return verifyRequest(
		"https://api.authy.com/protected/json/phones/verification/check",
		values,
	)
}
