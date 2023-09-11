package mailer

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// Mailer is a mailer that sends emails.
type Nfty struct {
	*Config
}

// SendNotification sends a notification.
func (m *Nfty) SendNotification(topic string, subject, body string, url string) error {

	client := &http.Client{}

	bo := make(map[string]interface{})

	action := make(map[string]interface{})

	var a []map[string]interface{}

	action["action"] = "view"
	action["label"] = "Logs"
	action["url"] = url
	action["clear"] = false

	a = append(a, action)

	bo["topic"] = topic
	bo["message"] = body
	bo["title"] = subject
	bo["actions"] = a

	b, err := json.Marshal(bo)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", m.Host, strings.NewReader(string(b)))

	auth := base64.StdEncoding.EncodeToString([]byte(m.Username + ":" + m.Password))

	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "text/markdown")

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	if strings.HasPrefix(res.Status, "20") {
		return nil
	} else {
		return errors.New(res.Status)
	}

}
