package mydata

import (
	"context"
	"encoding/xml"
	"errors"
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

type ToMap interface {
	ToMap() (map[string]string, error)
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
		u.Host = productionHost
		u.Path = path
	} else {
		u.Host = developmentHost
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

func Requester[P ToMap, T any](c *Client, params P, urlPath string) (int, *T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.client.Timeout)
	defer cancel()
	queryArgs, err := params.ToMap()
	if err != nil {
		return InternalErrorCode, nil, ErrorQueryURLCreation
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getURL(urlPath, queryArgs), nil)
	if err != nil {
		return InternalErrorCode, nil, ErrorRequestCreation
	}
	c.authorize(request)
	response, err := c.client.Do(request)
	if err != nil {
		return InternalErrorCode, nil, ErrorGettingResponse
	}

	//b, _ := c.responseToString(response)
	//fmt.Println(b)

	result, err := ParseXMLResponse[T](response)
	if err != nil {
		return InternalErrorCode, nil, errors.Join(ErrorXMLParsingResponse, err)
	}
	return response.StatusCode, result, nil
}
