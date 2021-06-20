package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	"github.com/labstack/echo"
)

var (
	// Timeout http request
	Timeout = 30 * time.Second
)

type (
	// httpResponse model
	httpResponse struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Meta    interface{} `json:"meta,omitempty"`
		Data    interface{} `json:"data,omitempty"`
		Errors  interface{} `json:"errors,omitempty"`
	}

	// Meta model
	Meta struct {
		Page         int `json:"page"`
		TotalRecords int `json:"totalRecords"`
	}
)

// NewMeta func
func NewMeta(page, totalRecords int) Meta {
	return Meta{
		Page: page, TotalRecords: totalRecords,
	}
}

// HTTPResponse abstract interface
type HTTPResponse interface {
	JSON(c echo.Context) error
}

// HTTPRequest function for getting json generally
func HTTPRequest(ctx context.Context, method, url string, body io.Reader, target interface{}, header map[string]string) (err error) {
	respBody, err := HTTPRequestWithResponse(ctx, method, url, body, header)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(respBody, target); err != nil {
		return fmt.Errorf("Error: %s ::: \nContent: %s", err.Error(), string(respBody))
	}

	return nil
}

// HTTPRequestWithResponse ...
func HTTPRequestWithResponse(ctx context.Context, method, url string, body io.Reader, header map[string]string) (respBody []byte, err error) {
	var req *http.Request
	var resp *http.Response

	req, err = http.NewRequest(method, url, body)
	defer func() {
		if req != nil {
			req.Close = true
			if req.Body != nil {
				req.Body.Close()
			}
		}
	}()
	if err != nil {
		return nil, err
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{Timeout: Timeout}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, _ = ioutil.ReadAll(resp.Body)
	defer func() {
		resp.Close = true
		client.CloseIdleConnections()
		resp.Body.Close()
	}()

	return
}

// NewHTTPResponse for create common response, data must in first params and meta in second params
func NewHTTPResponse(code int, message string, params ...interface{}) HTTPResponse {
	commonResponse := new(httpResponse)

	for _, param := range params {
		// get value param if type is pointer
		refValue := reflect.ValueOf(param)
		if refValue.Kind() == reflect.Ptr {
			refValue = refValue.Elem()
		}
		param = refValue.Interface()
		switch param.(type) {
		case Meta:
			commonResponse.Meta = param
		default:
			commonResponse.Data = param
		}
	}
	if code < http.StatusBadRequest {
		commonResponse.Success = true
	} else {
		commonResponse.Success = false
	}
	commonResponse.Code = code
	commonResponse.Message = message
	return commonResponse
}

// JSON for set http JSON response (Content-Type: application/json)
func (resp *httpResponse) JSON(c echo.Context) error {
	return c.JSON(resp.Code, resp)
}
