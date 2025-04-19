package adapter

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"fmt"

	"golang.org/x/xerrors"
)

type HackerNewsCilent interface {
	GetTopStories() ([]int, error)
	GetItem(id int) (Item, error)
}

type hackerNewsCilentClient struct {
	baseURL string
	client  *http.Client
}

func NewHackerNewsClient() HackerNewsCilent {
	return &hackerNewsCilentClient{
		baseURL: "https://hacker-news.firebaseio.com/v0",
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *hackerNewsCilentClient) GetTopStories() ([]int, error) {
	query := url.Values{}
	query.Add("print", "pretty")
	encodeQuery := query.Encode()
	resp, err := c.client.Get(c.baseURL + "/topstories.json?" + encodeQuery)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("failed to close response body", "err", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.New("failed to get top stories")
	}
	var ids []int
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, xerrors.New(err.Error())
	}
	return ids, nil
}

func (c *hackerNewsCilentClient) GetItem(id int) (Item, error) {
	query := url.Values{}
	query.Add("print", "pretty")
	encodeQuery := query.Encode()
	resp, err := c.client.Get(c.baseURL + "/item/" + fmt.Sprint(id) + ".json?" + encodeQuery)
	if err != nil {
		return Item{}, xerrors.New(err.Error())
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("failed to close response body", "err", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return Item{}, xerrors.New("failed to get item")
	}
	var item Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return Item{}, xerrors.New(err.Error())
	}
	return item, nil
}

type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Id          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}
