package mailgun

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var ApiEndpoint, ApiKey string

type Message struct {
	FromName    string
	FromAddress string
	ToAddress   string
	Subject     string
	TextBody    string
	HtmlBody    string
}

func (m Message) From() string {
	return fmt.Sprintf("%s <%s>", m.FromName, m.FromAddress)
}

func Send(message Message) ([]byte, error) {
	client := &http.Client{}

	values := make(url.Values)
	values.Set("from", message.From())
	values.Set("to", message.ToAddress)
	values.Set("subject", message.Subject)

	if len(message.HtmlBody) > 0 {
		values.Set("html", message.HtmlBody)
	} else {
		values.Set("text", message.TextBody)
	}

	request, _ := http.NewRequest("POST", ApiEndpoint, strings.NewReader(values.Encode()))
	request.Header.Set("content-type", "application/x-www-form-urlencoded")
	request.SetBasicAuth("api", ApiKey)

	response, e1 := client.Do(request)
	if e1 != nil {
		log.Println(e1)
		return nil, e1
	}
	defer response.Body.Close()

	body, e2 := ioutil.ReadAll(response.Body)
	if e2 != nil {
		log.Println(e2)
		return nil, e2
	}

	return body, nil
}
