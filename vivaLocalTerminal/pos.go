package vivaLocalTerminal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SalesURL            = "/pos/v1/sale"
	RefundURL           = "/pos/v1/refund"
	GetSessionURL       = "/pos/v1/sessions/%s"
	AbortTransactionURL = "/pos/v1/abort"
)

func requester[Request any, Response any](method string, path string, request Request) (Response, error) {
	var zeroResponse Response // Create a zero value of Response type
	requestJSON, err := json.Marshal(request)
	if err != nil {

		return zeroResponse, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, path, bytes.NewBuffer(requestJSON))
	if err != nil {
		return zeroResponse, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return zeroResponse, err
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return zeroResponse, err
	}
	return response, nil
}

func requesterGet[Response any](path string) (Response, error) {
	var zeroResponse Response // Create a zero value of Response type
	resp, err := http.Get(path)
	if err != nil {
		return zeroResponse, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&zeroResponse); err != nil {
		return zeroResponse, err
	}
	return zeroResponse, nil
}

type Pos struct {
	ip   string
	port string
}

func NewPos(ip string, port string) Pos {
	return Pos{ip: ip, port: port}
}

func (p Pos) getURL(path string) string {
	return "http://" + p.ip + ":" + p.port + path
}

func (p Pos) Sales(salesRequest SalesRequest) (SalesResponse, error) {
	return requester[SalesRequest, SalesResponse](http.MethodPost, p.getURL(SalesURL), salesRequest)
}

func (p Pos) Refund(refundRequest SaleRefundRequest) (SaleRefundResponse, error) {
	return requester[SaleRefundRequest, SaleRefundResponse](http.MethodPost, p.getURL(RefundURL), refundRequest)
}

func (p Pos) GetSession(sessionId string) (GetSessionResponse, error) {
	return requesterGet[GetSessionResponse](p.getURL(fmt.Sprintf(GetSessionURL, sessionId)))
}

func (p Pos) AbortTransaction(abortTransactionRequest AbortTransactionRequest) (AbortTransactionResponse, error) {
	return requester[AbortTransactionRequest, AbortTransactionResponse](http.MethodPost, p.getURL(AbortTransactionURL), abortTransactionRequest)
}
