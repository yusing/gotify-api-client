package gotify

import (
	"net/http"
	"net/url"

	httptransport "github.com/go-openapi/runtime/client"
	api "github.com/yusing/gotify-api-client/v2/client"
)

func NewClient(url *url.URL, client *http.Client) *api.GotifyREST {
	runtime := httptransport.NewWithClient(url.Host, url.Path, []string{url.Scheme}, client)
	return api.New(runtime, nil)
}
