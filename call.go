package facebook

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type (
	// Call is a place where Facebook Graph API
	// are being called. This is the HIT-man!
	Call struct {
		HTTPClient  http.Client
		HTTPRequest *http.Request
		HTTPHeader  map[string]string
		Method      string
		URL         string
		Body        io.Reader
	}

	// CallInterface contains the common method
	// to be setup before submitting request to
	// facebook graph api
	CallInterface interface {
		New() CallInterface

		SetMethod(method string) CallInterface
		SetURL(url string) CallInterface
		SetBody(body io.Reader) CallInterface
		AddHeader(key string, value string) CallInterface

		Submit() (interface{}, error)
	}
)

func (c *Call) newClient() http.Client {
	client := http.Client{}

	// Since there's something wrong with `http.Request`
	// when setting up a header for Authorization using Bearer,
	// then we need to make it right :)
	client.CheckRedirect = func(req *http.Request, reqs []*http.Request) error {
		req.Header.Add("Authorization", reqs[0].Header.Get("Authorization"))
		fmt.Println(reqs[0].Header.Get("Authorization"))
		return nil
	}

	return client
}

// New is use to initialize the Call methods
func (c *Call) New() CallInterface {
	return &Call{
		HTTPClient: c.newClient(),
		Body:       nil,
	}
}

// SetMethod is use to set the http method
func (c *Call) SetMethod(method string) CallInterface {
	c.Method = method
	return c
}

// SetURL is use to set the target URL
func (c *Call) SetURL(url string) CallInterface {
	c.URL = url
	return c
}

// SetBody is use to set the request body
func (c *Call) SetBody(body io.Reader) CallInterface {
	c.Body = body
	return c
}

// AddHeader is use to add request header
func (c *Call) AddHeader(key string, value string) CallInterface {
	if c.HTTPHeader == nil {
		c.HTTPHeader = make(map[string]string)
	}
	c.HTTPHeader[key] = value
	return c
}

// Submit is where the target URL is being called
func (c *Call) Submit() (interface{}, error) {
	req, errReq := http.NewRequest(c.Method, c.URL, c.Body)
	if errReq != nil {
		return nil, errReq
	}

	for key, val := range c.HTTPHeader {
		req.Header.Add(key, val)
	}

	resp, errResp := c.HTTPClient.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	respBytes, errRespBytes := ioutil.ReadAll(resp.Body)
	if errRespBytes != nil {
		return nil, errRespBytes
	}

	return string(respBytes), nil
}
