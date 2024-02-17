package pocket

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const RetrieveEndpoint string = "/get"

type RetrieveRequestParams struct {
	State       LinkState   `json:"state,omitempty"`
	Favorite    Bool        `json:"favorite,omitempty"`
	Tag         string      `json:"tag,omitempty"`
	ContentType ContentType `json:"contentType,omitempty"`
	Sort        SortOrder   `json:"sort,omitempty"`
	DetailType  DetailType  `json:"detailType,omitempty"`
	Search      string      `json:"search,omitempty"`
	Domain      string      `json:"domain,omitempty"`
	Since       *time.Time  `json:"-"`
	Count       uint        `json:"-"`
	Offset      uint        `json:"-"`
}

type retrieveRequestParams struct {
	*RequestAuthParams
	*RetrieveRequestParams
	SinceUnix    string `json:"since,omitempty"`
	CountString  string `json:"count,omitempty"`
	OffsetString string `json:"offset,omitempty"`
}

type RetrieveRequestResponse struct {
	List   RetrieveRequestResponseItems `json:"list"`
	Status ActionStatus                 `json:"status"`
}

func (r *RetrieveRequestResponse) UnmarshalJSON(data []byte) error {
	var v struct {
		List   interface{} `json:"list"`
		Status int         `json:"status"`
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	r.Status = ActionStatus(v.Status)
	r.List = RetrieveRequestResponseItems{}

	if reflect.ValueOf(v.List).Kind() == reflect.Map {
		var l struct {
			List RetrieveRequestResponseItems `json:"list"`
		}

		if err := json.Unmarshal(data, &l); err != nil {
			return err
		}

		r.List = l.List
	}

	return nil
}

type RetrieveRequestResponseItems map[string]RetrieveRequestResponseItem

type RetrieveRequestResponseItem struct {
	ItemID            string     `json:"item_id"`
	ResolvedID        string     `json:"resolved_id"`
	GivenURL          string     `json:"given_url"`
	ResolvedURL       string     `json:"resolved_url"`
	GivenTitle        string     `json:"given_title"`
	ResolvedTitle     string     `json:"resolved_title"`
	Favorite          Bool       `json:"favorite"`
	Status            LinkStatus `json:"status"`
	Excerpt           string     `json:"excerpt"`
	IsArticle         Bool       `json:"is_article"`
	IsIndex           Bool       `json:"is_index"`
	HasImage          Ternary    `json:"has_image"`
	HasVideo          Ternary    `json:"has_video"`
	WordCount         string     `json:"word_count"`
	Tags              Tags       `json:"tags"`
	Authors           Authors    `json:"authors"`
	Images            Images     `json:"images"`
	Videos            Videos     `json:"videos"`
	TimeAddedUnix     string     `json:"time_added"`
	TimeUpdatedUnix   string     `json:"time_updated"`
	TimeReadUnix      string     `json:"time_read"`
	TimeFavoritedUnix string     `json:"time_favorited"`
	SortID            uint       `json:"sort_id"`
	Lang              string     `json:"lang"`
	TopImageURL       string     `json:"top_image_url"`
}

func (r *RetrieveRequestResponseItem) TimeAdded() time.Time {
	s, _ := strconv.ParseInt(r.TimeAddedUnix, 0, 64)
	if s == 0 {
		return time.Time{}
	}
	return time.Unix(s, 0)
}

func (r *RetrieveRequestResponseItem) TimeUpdated() time.Time {
	s, _ := strconv.ParseInt(r.TimeUpdatedUnix, 0, 64)
	if s == 0 {
		return time.Time{}
	}
	return time.Unix(s, 0)
}

func (r *RetrieveRequestResponseItem) TimeRead() time.Time {
	s, _ := strconv.ParseInt(r.TimeReadUnix, 0, 64)
	if s == 0 {
		return time.Time{}
	}
	return time.Unix(s, 0)
}

func (r *RetrieveRequestResponseItem) TimeFavorited() time.Time {
	s, _ := strconv.ParseInt(r.TimeFavoritedUnix, 0, 64)
	if s == 0 {
		return time.Time{}
	}
	return time.Unix(s, 0)
}

func (c *Client) Retrieve(ctx context.Context, params *RetrieveRequestParams) (*RetrieveRequestResponse, error) {
	p := &retrieveRequestParams{
		RequestAuthParams:     c.Auth.RequestAuthParams,
		RetrieveRequestParams: params,
	}

	if params.Since != nil {
		p.SinceUnix = fmt.Sprintf("%d", params.Since.Unix())
	}

	if params.Count > 0 {
		p.CountString = fmt.Sprintf("%d", params.Count)
	}

	if params.Offset > 0 {
		p.OffsetString = fmt.Sprintf("%d", params.Offset)
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to marshal input body")
	}

	r, err := c.do(ctx, RetrieveEndpoint, b)
	if err != nil {
		return nil, err
	}

	response := &RetrieveRequestResponse{}
	err = json.Unmarshal(r, response)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal response body")
	}

	return response, nil
}
