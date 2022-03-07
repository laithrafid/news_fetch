package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/laithrafid/news_fetch/domain"
)

func NewGNewsClient(hClient *http.Client) *GClient {
	gClient := &GClient{
		HttpClient: hClient,
	}

	return gClient
}

type GClient struct {
	HttpClient *http.Client
}

type Apis interface {
	GetSources() (*domain.NewsSource, error)
	GetHeadlines() (*domain.TopHeadline, error)
	GetEverything() (*domain.Everything, error)
}

func prepareUrl(base string, qp map[string][]string) string {
	var url strings.Builder
	url.WriteString(base)
	url.WriteString("?")
	for idx, val := range qp {
		for _, value := range val {
			url.WriteString(fmt.Sprintf("%s=%s", idx, value))
			url.WriteString("&")
		}
	}
	return strings.TrimRight(url.String(), "&")
}

func (c *GClient) GetSources(qp map[string][]string) (*domain.NewsSource, error) {
	url := prepareUrl("http://newsapi.org/v2/sources", qp)
	if req, err := http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	} else {
		req.Header.Add("X-Api-Key", "17d9a468d74748d3a39175d524747e95")
		resp, err := c.HttpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		ns := new(domain.NewsSource)
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		} else {
			if err := json.Unmarshal(body, &ns); err != nil {
				return nil, err
			} else {
				return ns, nil
			}
		}
	}
}

func (c *GClient) GetHeadlines(qp map[string][]string) (*domain.TopHeadline, error) {
	url := prepareUrl("http://newsapi.org/v2/top-headlines", qp)
	if req, err := http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	} else {
		req.Header.Add("X-Api-Key", "17d9a468d74748d3a39175d524747e95")
		resp, err := c.HttpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		ns := new(domain.TopHeadline)
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		} else {
			err := json.Unmarshal(body, &ns)
			if err != nil {
				return nil, err
			} else {
				return ns, nil
			}
		}
	}
}

func (c *GClient) GetEverything(qp map[string][]string) (*domain.Everything, error) {
	url := prepareUrl("http://newsapi.org/v2/everything", qp)
	if req, err := http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	} else {
		req.Header.Add("X-Api-Key", "17d9a468d74748d3a39175d524747e95")
		resp, err := c.HttpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		ns := new(domain.Everything)
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		} else {
			err := json.Unmarshal(body, &ns)
			if err != nil {
				return nil, err
			} else {
				return ns, nil
			}
		}
	}
}
