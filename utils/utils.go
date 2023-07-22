package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/blackestwhite/zwrapper/entity"
)

func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Post(wp entity.WebhookPayload) (err error) {
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(wp)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", wp.URL, b)
	if err != nil {
		return err
	}
	req.Header = http.Header{
		"Content-Type":               {"application/json"},
		"x-oremote-api-access-token": {wp.Key},
	}

	client := http.Client{}
	_, err = client.Do(req)
	return
}
