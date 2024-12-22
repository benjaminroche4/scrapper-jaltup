package labonnealternance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// APIURL is the official API Endpoint.
	APIURL = "https://labonnealternance-recette.apprentissage.beta.gouv.fr/api"
	// APIVersion the api current version.
	APIVersion = "v1"
	// APIUserAgent identifies this library with the API.
	APIUserAgent = "Jaltup"
)

// LaBonneAlternance represents Client connection API.
type LaBonneAlternance struct {
	timeout time.Duration
	client  *http.Client
}

// New creates a new API client.
func New() *LaBonneAlternance {
	return NewWithClient(http.DefaultClient)
}

// NewWithClient creates a new API client with a specific http client.
func NewWithClient(httpClient *http.Client) *LaBonneAlternance {
	return &LaBonneAlternance{
		timeout: httpClient.Timeout,
		client:  httpClient,
	}
}

func (thiz *LaBonneAlternance) JobFormations(romes []string) (*JobFormationsResponse, error) {
	params := url.Values{
		"romes":  {strings.Join(romes, ",")},
		"caller": {"jeanbaptiste@a26k.ch"},
	}
	url := fmt.Sprintf("%s/%s/%s", APIURL, APIVersion, "jobsEtFormations")
	resp, err := thiz.doRequest(http.MethodGet, url, params, &JobFormationsResponse{})
	if err != nil {
		return nil, err
	}

	return resp.(*JobFormationsResponse), nil
}

// doRequest executes a HTTP Request to the API and returns the result.
func (thiz *LaBonneAlternance) doRequest(
	method string,
	reqURL string,
	values url.Values,
	typ interface{},
) (interface{}, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	ctx, cancel := thiz.context()
	if cancel != nil {
		defer cancel()
	}

	req, err = http.NewRequestWithContext(ctx, method, reqURL, nil)
	if err != nil {
		return nil, ErrNewRequest(err)
	}
	req.URL.RawQuery = values.Encode()

	req.Header.Add("User-Agent", APIUserAgent)

	// Execute request
	resp, err = thiz.client.Do(req)
	if err != nil {
		return nil, ErrDoRequest(err)
	}
	defer resp.Body.Close()

	// Read request
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrReadBody(method, reqURL, err)
	}

	if resp.StatusCode != 200 {
		return nil, ErrServerHTTPError(resp.StatusCode)
	}

	// Check empty
	if len(respBody) == 0 {
		return nil, ErrEmptyBody(method, reqURL)
	}

	// Parse request
	if typ != nil {
		jsonData := typ
		err = json.Unmarshal(respBody, &jsonData)
		if err != nil {
			return nil, ErrUnmarshal(method, reqURL, err)
		}
		return jsonData, nil
	}
	var jsonResp interface{}
	err = json.Unmarshal(respBody, &jsonResp)
	if err != nil {
		return nil, ErrUnmarshal(method, reqURL, err)
	}

	return jsonResp, nil
}

func (thiz *LaBonneAlternance) context() (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc
	ctx := context.Background()
	if thiz.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), thiz.timeout)
	}

	return ctx, cancel
}
