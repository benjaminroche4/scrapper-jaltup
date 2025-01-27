package hellowork

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"regexp"
	"scrapperjaltup/model"
	_html "scrapperjaltup/util/html"
	"strings"
	"time"

	_slug "github.com/gosimple/slug"
	"golang.org/x/net/html"
)

const (
	// APIURL is the official API Endpoint.
	APIURL = "https://www.hellowork.com/fr-fr"
	// APIUserAgent identifies this library with the API.
	APIUserAgent = "Jaltup"
)

// HelloWork represents Client connection API.
type HelloWork struct {
	timeout time.Duration
	client  *http.Client
}

// New creates a new API client.
func New() *HelloWork {
	return NewWithClient(http.DefaultClient)
}

// NewWithClient creates a new API client with a specific http client.
func NewWithClient(httpClient *http.Client) *HelloWork {
	return &HelloWork{
		timeout: httpClient.Timeout,
		client:  httpClient,
	}
}

func (thiz *HelloWork) Companies() ([]*model.Company, error) {
	nameClean := regexp.MustCompile(`(?i)\s*Recrutement\s*`)

	companies := []*model.Company{}

	req := fmt.Sprintf("%s/%s", APIURL, "entreprise.html")
	resp, err := thiz.doRequest(http.MethodGet, req, "")
	if err != nil {
		return companies, err
	}

	doc, err := html.Parse(bytes.NewReader(resp))
	if err != nil {
		return companies, err
	}

	body := _html.FindNodeByTagName(doc, "body")
	if body == nil {
		return companies, errors.New("can't find body node")
	}

	list := _html.FindNodeByClassName(body, "grid-3")
	if list == nil {
		return companies, errors.New("can't find list node")
	}

	for child := _html.NextChildNodeElement(list); child != nil; child = _html.NextNodeElement(child) {
		header := _html.FindNodeByClassName(child, "header")
		infos := _html.FindNodeByClassName(child, "infos")
		images := _html.FindAllNodesByTagName(header, "img")
		if len(images) < 1 {
			continue
		}
		logo, ok := _html.GetNodeAttr(images[1], "src")
		if !ok {
			continue
		}
		filename := filepath.Base(logo)
		publicID := strings.ReplaceAll(filename, filepath.Ext(filename), "")
		name, ok := _html.GetTextContent(_html.FindNodeByClassName(infos, "name"))
		if !ok {
			continue
		}

		name = nameClean.ReplaceAllString(name, "")

		companies = append(companies, &model.Company{
			ID:           0,
			PublicID:     publicID,
			Name:         name,
			Siret:        "",
			ContactEmail: "",
			PhoneNumber:  "",
			WebSiteURL:   "",
			Logo:         logo,
			CreatedAt:    time.Now(),
			Slug:         _slug.Make(name),
			Verified:     false,
		})
	}

	return companies, nil
}

func (thiz *HelloWork) doRequest(method, reqURL, body string) ([]byte, error) {
	var (
		req    *http.Request
		resp   *http.Response
		err    error
		reader io.Reader
	)

	ctx, cancel := thiz.context()
	if cancel != nil {
		defer cancel()
	}
	if method != http.MethodGet && body != "" {
		reader = strings.NewReader(body)
	}
	req, err = http.NewRequestWithContext(ctx, method, reqURL, reader)
	if err != nil {
		return nil, ErrNewRequest(err)
	}

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

	return respBody, nil
}

// unexported functions

func (thiz *HelloWork) context() (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc
	ctx := context.Background()
	if thiz.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), thiz.timeout)
	}

	return ctx, cancel
}
