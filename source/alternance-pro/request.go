package alternancepro

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"scrapperjaltup/model"
	"scrapperjaltup/util"
	"scrapperjaltup/util/cities"
	"strconv"
	"strings"
	"time"

	_slug "github.com/gosimple/slug"
	"golang.org/x/net/html"
)

const (
	// APIURL is the official API Endpoint.
	APIURL = "https://www.alternance-professionnelle.fr"
	// APIUserAgent identifies this library with the API.
	APIUserAgent = "Jaltup"

	ValidDuration   = 30 * 24 * time.Hour
	PremiumDuration = 7 * 24 * time.Hour
)

// AlternancePro represents Client connection API.
type AlternancePro struct {
	timeout time.Duration
	client  *http.Client
}

// New creates a new API client.
func New() *AlternancePro {
	return NewWithClient(http.DefaultClient)
}

// NewWithClient creates a new API client with a specific http client.
func NewWithClient(httpClient *http.Client) *AlternancePro {
	return &AlternancePro{
		timeout: httpClient.Timeout,
		client:  httpClient,
	}
}

func (thiz *AlternancePro) Companies() ([]*model.Company, error) {
	companies := []*model.Company{}

	req := fmt.Sprintf("%s/%s", APIURL, "entreprises")
	resp, err := thiz.doRequest(http.MethodGet, req, "")
	if err != nil {
		return companies, err
	}

	doc, err := html.Parse(bytes.NewReader(resp))
	if err != nil {
		return companies, err
	}

	body := findNodeByTagName(doc, "body")
	if body == nil {
		return companies, errors.New("can't find body node")
	}
	list := findNodeByClassName(body, "mhh-entreprises")
	if list == nil {
		return companies, errors.New("can't find list node")
	}
	for child := list.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode {
			link, ok := getNodeLink(child)
			if !ok {
				continue
			}
			u, err := url.Parse(link)
			if err != nil {
				continue
			}
			fields := strings.Split(strings.TrimRight(u.Path, "/"), "/")
			slug := fields[len(fields)-1]

			company, err := thiz.Company(slug)
			if err == nil && company != nil {
				companies = append(companies, company)
			}
		}
	}

	return companies, nil
}

func (thiz *AlternancePro) Company(slug string) (*model.Company, error) {
	var (
		name, logo, email, phone, linkURL string
	)
	req := fmt.Sprintf("%s/%s/%s", APIURL, "entreprises", slug)
	resp, err := thiz.doRequest(http.MethodGet, req, "")
	if err != nil {
		return nil, err
	}
	linkURL = req

	doc, err := html.Parse(bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}
	body := findNodeByTagName(doc, "body")
	if body == nil {
		return nil, errors.New("can't find body node")
	}
	details := findNodeByClassName(body, "entreprises-profile-detail")
	if details == nil {
		return nil, errors.New("can't find details node")
	}
	title := findNodeByTagName(details, "h1")
	if title != nil {
		name, _ = getTextContent(title)
		name = util.Truncate(name, 120)
	} else {
		return nil, errors.New("can't find details title")
	}
	linkNode := findNodeByClassName(details, "internet")
	if linkNode != nil {
		href, ok := getNodeAttr(linkNode.FirstChild, "href")
		if ok {
			href = util.CleanURL(href)
			if href != "" {
				linkURL = href
			}
		}
	}
	image := findNodeByTagName(details, "img")
	if image != nil {
		logo, _ = getNodeAttr(image, "src")
		logo = util.Truncate(logo, 255)
	}
	emailNode := findNodeByClassName(body, "email")
	if emailNode != nil {
		email, _ = getTextContent(emailNode)
		email = util.Truncate(util.CleanEmail(email), 120)
	}
	phoneNode := findNodeByClassName(body, "telephone")
	if phoneNode != nil {
		phone, _ = getTextContent(phoneNode)
		phone = util.Truncate(phone, 20)
	}

	company := createCompany(name)
	company.ContactEmail = email
	company.PhoneNumber = phone
	company.WebSiteURL = linkURL
	company.Logo = logo

	return company, nil
}

func (thiz *AlternancePro) Jobs(page int, jobtype string) ([]*model.Offer, int, error) {
	jobResponse := JobResponse{}
	offers := []*model.Offer{}

	pageNum := page
	if pageNum <= 0 {
		pageNum = 1
	}
	req := fmt.Sprintf("%s/%s", APIURL, "jm-ajax/get_listings")
	params := url.Values{
		"lang":            {""},
		"search_keywords": {""},
		"search_location": {""},
		"filter_job_type": {"cdd", "cdi", "stage"},
		"page":            {strconv.Itoa(page)},
		"per_page":        {"10"},
		"orderby":         {"featured"},
		"order":           {"DESC"},
		"show_pagination": {"false"},
	}
	if jobtype != "" {
		params.Set("filter_job_type", jobtype)
	}

	payload := params.Encode()
	resp, err := thiz.doRequest(http.MethodPost, req, payload)
	if err != nil {
		return nil, 0, err
	}
	err = json.Unmarshal(resp, &jobResponse)
	if err != nil {
		return nil, 0, err
	}

	doc, err := html.Parse(strings.NewReader(jobResponse.HTML))
	if err != nil {
		return nil, 0, err
	}
	nodes := findAllNodesByTagName(doc, "a")
	for _, node := range nodes {
		link, ok := getNodeLink(node)
		if !ok {
			continue
		}
		u, err := url.Parse(link)
		if err != nil {
			continue
		}
		fields := strings.Split(strings.TrimRight(u.Path, "/"), "/")
		slug := fields[len(fields)-1]

		offer, err := thiz.Job(slug)
		if err == nil {
			offers = append(offers, offer)
		}
	}

	return offers, jobResponse.MaxNumPages, nil
}

func (thiz *AlternancePro) Job(slug string) (*model.Offer, error) {
	req := fmt.Sprintf("%s/%s/%s", APIURL, "offres-emploi", slug)
	resp, err := thiz.doRequest(http.MethodGet, req, "")
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}
	return parseJob(doc, req, slug)
}

func (thiz *AlternancePro) doRequest(method, reqURL, body string) ([]byte, error) {
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

func (thiz *AlternancePro) context() (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc
	ctx := context.Background()
	if thiz.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), thiz.timeout)
	}

	return ctx, cancel
}

func createCompany(name string) *model.Company {
	return &model.Company{
		ID:           0,
		PublicID:     util.GenerateUniqueID(10),
		Name:         name,
		Siret:        "",
		ContactEmail: "",
		PhoneNumber:  "",
		WebSiteURL:   "",
		Logo:         "",
		CreatedAt:    time.Now().UTC(),
		Slug:         _slug.Make(name),
		Verified:     false,
	}
}

func createPlace(name string) *model.Place {
	zipCode := ""
	latitude := 0.0
	longitude := 0.0
	city := cities.FindCity(name)
	if city != nil {
		zipCode = city.ZipCode
		latitude = city.Latitude
		longitude = city.Longitude
	}

	return &model.Place{
		FullAddress: name,
		City:        name,
		ZipCode:     zipCode,
		Latitude:    latitude,
		Longitude:   longitude,
	}
}

func parseJob(doc *html.Node, req, slug string) (*model.Offer, error) {
	body := findNodeByTagName(doc, "body")
	if body == nil {
		return nil, errors.New("can't find body node")
	}
	title, err := parseJobTitle(body)
	if err != nil {
		return nil, err
	}
	jobDetails, err := parseJobDetails(body)
	if err != nil {
		return nil, err
	}
	description := parseJobDescription(body)
	company := createCompany(jobDetails.CompanyName)
	city := util.CleanCityName(jobDetails.City)
	createdAt := time.Now().Truncate(24 * time.Hour)
	endPremiumAt := createdAt.Add(PremiumDuration)
	premium := false
	if time.Now().Unix() < endPremiumAt.Unix() {
		premium = true
	}
	return &model.Offer{
		ID:       0,
		Company:  *company,
		PublicID: util.GenerateUniqueID(10),
		Title:    util.CleanTitle(title),
		Place:    *createPlace(city),
		Job: model.Job{
			Description:  description,
			ContractType: jobDetails.ContractType,
			Duration:     8,
			Remote:       false,
			StudyLevel:   jobDetails.StudyLevel,
			StartDate:    jobDetails.StartDate,
		},
		URL:          util.Truncate(req, 255),
		Tag:          []string{},
		Status:       "published",
		CreatedAt:    createdAt,
		EndAt:        createdAt.Add(ValidDuration),
		EndPremiumAt: endPremiumAt,
		Slug:         _slug.Make(title),
		Premium:      premium,
		ExternalID:   util.Truncate(slug, 255),
		ServiceName:  "alternance-professionnelle",
		Categories:   []model.Category{},
	}, nil
}

func parseJobTitle(body *html.Node) (string, error) {
	var title string

	detail := findNodeByClassName(body, "entreprises-profile-detail")
	if detail == nil {
		return "", errors.New("can't find detail node")
	}
	titleNode := findNodeByTagName(detail, "h1")
	if titleNode != nil {
		title, _ = getTextContent(titleNode)
		title = util.Truncate(title, 120)
	} else {
		return "", errors.New("can't find details title")
	}

	return title, nil
}

func parseJobDescription(body *html.Node) string {
	description := ""
	main := findNodeByClassName(body, "entreprises-main")
	if main != nil {
		descriptionNode := findNodeByTagName(main, "h2")
		child := nextNodeElement(descriptionNode)
		for child != nil {
			text, ok := getTextContent(child)
			if ok {
				description += text + "\n"
			}
			child = nextNodeElement(child)
		}
	}

	return description
}

func parseJobDetails(body *html.Node) (*JobDetail, error) {
	var (
		company, city, contractType, studyLevel string
		startDate                               time.Time
	)
	companyNode := findNodeByClassName(body, "entreprises-sidebar")
	if companyNode == nil {
		return nil, errors.New("can't find company node")
	}
	child := nextChildNodeElement(companyNode)
	if child != nil {
		child = nextChildNodeElement(child)
	}
	for child != nil {
		section, _ := getTextContent(child)
		if child.Data == "h3" {
			section = strings.ToLower(section)
			if strings.Contains(section, "entreprise") {
				child = nextNodeElement(child)
				if child != nil && child.Data == "p" {
					company, _ = getTextContent(child)
					company = util.Truncate(company, 120)
				}
			}
			if strings.Contains(section, "lieu") {
				child = nextNodeElement(child)
				if child != nil && child.Data == "p" {
					city, _ = getTextContent(child)
				}
			}
			if strings.Contains(section, "date de début") {
				child = nextNodeElement(child)
				if child != nil && child.Data == "p" {
					date, _ := getTextContent(child)
					startDate, _ = time.Parse("02/01/2006", date)
				}
			}
			if strings.Contains(section, "type de contrat") {
				child = nextNodeElement(child)
				if child != nil && child.Data == "p" {
					contractType, _ = getTextContent(child)
					contractType = strings.ToUpper(contractType)
				}
			}
			if strings.Contains(section, "niveau de formation") {
				child = nextNodeElement(child)
				if child != nil && child.Data == "p" {
					studyLevel, _ = getTextContent(child)
					studyLevel = strings.ToUpper(studyLevel)
				}
			}
		}
		child = nextNodeElement(child)
	}

	return &JobDetail{
		CompanyName:  company,
		City:         city,
		StartDate:    startDate,
		ContractType: contractType,
		StudyLevel:   studyLevel,
	}, nil
}
