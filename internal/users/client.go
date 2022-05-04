package users

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

)

type Client interface {
	 Get(ctx context.Context,name, id string) (*map[string]interface{}, error)
}
type httpClient struct {
	client	http.Client
}
func NewHttpClient(client *http.Client) Client {
	return &httpClient{
		client: *client,
	}
}
func (c *httpClient) Get(ctx context.Context,name, id string) (*map[string]interface{}, error){

resp, err :=c.client.Do(&http.Request{
	Method: "GET",
	URL: &url.URL{
		Scheme: "http",
		Host: "localhost:8088",
		Path: "/"+name+"/"+id,
	},
	Header: http.Header{
		"Content-Type": []string{"application/json"},
		"Accept": []string{"application/json"},
		"Authorization": []string{"Bearer 12345"},
		"X-Commerce": []string{"Commerce"},
		"X-Country": []string{"Country"},
	},
})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
