package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/maxmcd/c4/backend/pkg/config"
	"github.com/maxmcd/c4/backend/pkg/models"
)

type nexmoResponse struct {
	RequestID string `json:"request_id,omitempty"`
	Status    string `json:"status,omitempty"`
	ErrorText string `json:"error_text,omitempty"`
}

func nexmoRequest(url string) (nexmoResponse, error) {
	var nr nexmoResponse
	resp, err := http.Get(url)
	if err != nil {
		return nr, err
	}

	if resp.StatusCode != http.StatusOK {
		return nr, errors.New(fmt.Sprintf("Unexpected status code from nexmo %d", resp.StatusCode))
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nr, err
	}

	err = json.Unmarshal(bodyBytes, &nr)
	if err != nil {
		return nr, err
	}
	if nr.Status != "0" {
		return nr, errors.New(nr.ErrorText)
	}
	return nr, nil
}

func sendVerificationRequest(user models.User) (string, error) {
	if user.NexmoRequestID != "" {
		cancelVerificationRequest(user.NexmoRequestID)
	}

	url := fmt.Sprintf(
		"https://api.nexmo.com/verify/json?api_key=%s&api_secret=%s&number=%s&brand=%s&pin_expiry=%d",
		config.ApiKeys.NexmoKey,
		config.ApiKeys.NexmoSecret,
		user.PhoneNumber(),
		"Rivet",
		60,
	)
	nr, err := nexmoRequest(url)
	if err != nil {
		return "", err
	}
	go func() {
		// kill the request if incomplete to avoid annoying call fallback
		time.Sleep(31 * time.Second)
		cancelVerificationRequest(nr.RequestID)
	}()
	return nr.RequestID, nil
}

func confirmVerificationRequest(request_id string, code string) error {
	url := fmt.Sprintf(
		"https://api.nexmo.com/verify/check/json?api_key=%s&api_secret=%s&request_id=%s&code=%s",
		config.ApiKeys.NexmoKey,
		config.ApiKeys.NexmoSecret,
		request_id,
		code,
	)
	fmt.Println(url)
	nr, err := nexmoRequest(url)
	glog.Infoln(nr.ErrorText)
	glog.Infoln(nr.Status)
	if err != nil {
		return err
	}
	return nil
}

func cancelVerificationRequest(request_id string) error {
	url := fmt.Sprintf(
		"https://api.nexmo.com/verify/control/json?api_key=%s&api_secret=%s&request_id=%s&cmd=cancel",
		config.ApiKeys.NexmoKey,
		config.ApiKeys.NexmoSecret,
		request_id,
	)
	nr, err := nexmoRequest(url)
	glog.Infoln(nr)
	glog.Infoln(err)
	if err != nil {
		return err
	}
	return nil
}
