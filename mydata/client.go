package mydata

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	client          *http.Client
	prod            bool
	userID          string
	subscriptionKey string
	onInvalid       OnInvalidAction // what to do when an invalid invoice is encountered
}

// NewClient creates a new myDATA client
func NewClient(userID, subscriptionKey string, timeoutInSeconds int, prod bool) *Client {
	return &Client{
		client: &http.Client{
			Timeout: time.Duration(timeoutInSeconds) * time.Second,
		},
		prod:            prod,
		userID:          userID,
		subscriptionKey: subscriptionKey,
		onInvalid:       ErrorOnInvalid,
	}
}

// SetOnInvalidAction sets the action to take when an invalid invoice is encountered
func (c *Client) SetOnInvalidAction(action OnInvalidAction) {
	c.onInvalid = action
}

// authorize adds the necessary headers to the request to authorize it
func (c *Client) authorize(req *http.Request) {
	req.Header.Set("aade-user-id", c.userID)
	req.Header.Set("Ocp-Apim-Subscription-Key", c.subscriptionKey)

}

// getURL creates the url for the request
func (c *Client) getURL(path string, queryArgs map[string]string) string {
	u := url.URL{
		Scheme: "https",
	}
	if c.prod {
		u.Host = prodHost
		u.Path = "/myDATA" + path
	} else {
		u.Host = develHost
		u.Path = path
	}
	if queryArgs != nil {
		rq := u.Query()
		for k, v := range queryArgs {
			rq.Set(k, v)
		}
		u.RawQuery = rq.Encode()
	}
	return u.String()
}

// responseToString reads the response body and returns it as a string
func (c *Client) responseToString(r *http.Response) (string, error) {
	defer func() {
		_ = r.Body.Close()
	}()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ParseXMLResponse parses the response body as xml and returns the result
func ParseXMLResponse[T any](r *http.Response) (*T, error) {
	defer func() {
		_ = r.Body.Close()
	}()
	var result T
	err := xml.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
