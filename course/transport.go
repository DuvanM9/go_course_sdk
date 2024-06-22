package course

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/DuvanM9/gocourse_domain/domain"
	"github.com/ncostamagna/go_http_client/client"
)

type (
	DataResponse struct {
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
	}

	Transport interface {
		Get(id string) (*domain.Course, error)
	}

	clientHTTP struct {
		client client.Transport
	}
)

// Get implements Transport.
func (c *clientHTTP) Get(id string) (*domain.Course, error) {
	// ahora data response es de tipo domain Course
	dataResponse := DataResponse{Data: &domain.Course{}}

	//sirve para agregar al path query params
	u := url.URL{}
	u.Path += fmt.Sprintf("/courses/%s", id)
	resp := c.client.Get(u.String())
	if resp.Err != nil {
		return nil, resp.Err
	}

	if resp.StatusCode == 404 {
		return nil, ErrNotFound{Message: resp.String()}
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("%s", resp)
	}

	// rellenar dataResponse con la respuesta de resp
	if err := resp.FillUp(&dataResponse); err != nil {
		return nil, err
	}

	return dataResponse.Data.(*domain.Course), nil
}

func NewHttpClient(baseUrl, token string) Transport {
	header := http.Header{}

	if token != "" {
		header.Set("Authorization", token)
	}

	return &clientHTTP{
		client: client.New(header, baseUrl, 5000*time.Millisecond, true),
	}
}
