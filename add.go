package pocket

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

const AddEndpoint string = "/add"

type AddRequestParams struct {
	URL     string     `json:"url"`
	Title   string     `json:"title,omitempty"`
	Tags    string     `json:"tags,omitempty"`
	TweetID string     `json:"tweet_id,omitempty"`
	Time    *time.Time `json:"-"`
}

type addRequestParams struct {
	*RequestAuthParams
	*AddRequestParams
	TimeUnix int64 `json:"time,omitempty"`
}

type AddRequestResponse struct {
	Item   []AddRequestResponseItem `json:"item"`
	Status ActionStatus             `json:"status"`
}

type AddRequestResponseItem struct {
	ItemID         string `json:"item_id"`
	NormalURL      string `json:"normal_url"`
	ResolvedID     string `json:"resolved_id"`
	ResolvedURL    string `json:"resolved_url"`
	DomainID       string `json:"domain_id"`
	OriginDomainID string `json:"origin_domain_id"`
	// ResponseCode string   `json:"response_code"`
	// MimeType      string `json:"mime_type"`
	// ContentLength string `json:"content_length"`
	// Encoding      string `json:"encoding"`
	// DateResolved  string `json:"date_resolved"`
	// DatePublished string `json:"date_published"`
	Title     string  `json:"title"`
	Excerpt   string  `json:"excerpt"`
	WordCount string  `json:"word_count"`
	HasImage  Ternary `json:"has_image"`
	HasVideo  Ternary `json:"has_video"`
	IsIndex   Bool    `json:"is_index"`
	IsArticle Bool    `json:"is_article"`
	Authors   Authors `json:"authors"`
	Images    Images  `json:"images"`
	Videos    Videos  `json:"videos"`
}

func (c *Client) Add(ctx context.Context, params *AddRequestParams) (*AddRequestResponse, error) {
	p := &addRequestParams{
		RequestAuthParams: c.Auth.RequestAuthParams,
		AddRequestParams:  params,
	}

	if params.Time != nil {
		p.TimeUnix = params.Time.Unix()
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to marshal input body")
	}

	r, err := c.do(ctx, AddEndpoint, b)
	if err != nil {
		return nil, err
	}

	response := &AddRequestResponse{}
	err = json.Unmarshal(r, response)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal response body")
	}

	return response, nil
}
