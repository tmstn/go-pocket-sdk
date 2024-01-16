package pocket

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const ModifyEndpoint string = "/send"

type ActionMethod string

func (a ActionMethod) Add() bool {
	return a == AddMethod
}

func (a ActionMethod) Archive() bool {
	return a == ArchiveMethod
}

func (a ActionMethod) Readd() bool {
	return a == ReaddMethod
}

func (a ActionMethod) Favorite() bool {
	return a == FavoriteMethod
}

func (a ActionMethod) Unfavorite() bool {
	return a == UnfavoriteMethod
}

func (a ActionMethod) Delete() bool {
	return a == DeleteMethod
}

func (a ActionMethod) AddTags() bool {
	return a == AddTagsMethod
}
func (a ActionMethod) ReplaceTags() bool {
	return a == ReplaceTagsMethod
}

func (a ActionMethod) ClearTags() bool {
	return a == ClearTagsMethod
}

func (a ActionMethod) RenameTag() bool {
	return a == RenameTagMethod
}

func (a ActionMethod) DeleteTag() bool {
	return a == DeleteTagMethod
}

const (
	AddMethod         ActionMethod = "add"
	ArchiveMethod     ActionMethod = "archive"
	ReaddMethod       ActionMethod = "readd"
	FavoriteMethod    ActionMethod = "favorite"
	UnfavoriteMethod  ActionMethod = "unfavorite"
	DeleteMethod      ActionMethod = "delete"
	AddTagsMethod     ActionMethod = "tags_add"
	RemoveTagsMethod  ActionMethod = "tags_remove"
	ReplaceTagsMethod ActionMethod = "tags_replace"
	ClearTagsMethod   ActionMethod = "tags_clear"
	RenameTagMethod   ActionMethod = "tag_rename"
	DeleteTagMethod   ActionMethod = "tag_delete"
)

type ModifyRequestParams struct {
	AddActions         []*AddAction         `json:"-"`
	ArchiveActions     []*ArchiveAction     `json:"-"`
	ReaddActions       []*ReaddAction       `json:"-"`
	FavoriteActions    []*FavoriteAction    `json:"-"`
	UnfavoriteActions  []*UnfavoriteAction  `json:"-"`
	DeleteActions      []*DeleteAction      `json:"-"`
	AddTagsActions     []*AddTagsAction     `json:"-"`
	RemoveTagsActions  []*RemoveTagsAction  `json:"-"`
	ReplaceTagsActions []*ReplaceTagsAction `json:"-"`
	ClearTagsActions   []*ClearTagsAction   `json:"-"`
	RenameTagActions   []*RenameTagAction   `json:"-"`
	DeleteTagActions   []*DeleteTagAction   `json:"-"`
}

func (m *ModifyRequestParams) Add(a *AddAction) *ModifyRequestParams {
	m.AddActions = append(m.AddActions, a)
	return m
}

func (m *ModifyRequestParams) Archive(a *ArchiveAction) *ModifyRequestParams {
	m.ArchiveActions = append(m.ArchiveActions, a)
	return m
}

func (m *ModifyRequestParams) Readd(a *ReaddAction) *ModifyRequestParams {
	m.ReaddActions = append(m.ReaddActions, a)
	return m
}

func (m *ModifyRequestParams) Favorite(a *FavoriteAction) *ModifyRequestParams {
	m.FavoriteActions = append(m.FavoriteActions, a)
	return m
}

func (m *ModifyRequestParams) Unfavorite(a *UnfavoriteAction) *ModifyRequestParams {
	m.UnfavoriteActions = append(m.UnfavoriteActions, a)
	return m
}

func (m *ModifyRequestParams) Delete(a *DeleteAction) *ModifyRequestParams {
	m.DeleteActions = append(m.DeleteActions, a)
	return m
}

func (m *ModifyRequestParams) AddTags(a *AddTagsAction) *ModifyRequestParams {
	m.AddTagsActions = append(m.AddTagsActions, a)
	return m
}

func (m *ModifyRequestParams) RemoveTags(a *RemoveTagsAction) *ModifyRequestParams {
	m.RemoveTagsActions = append(m.RemoveTagsActions, a)
	return m
}

func (m *ModifyRequestParams) ReplaceTags(a *ReplaceTagsAction) *ModifyRequestParams {
	m.ReplaceTagsActions = append(m.ReplaceTagsActions, a)
	return m
}

func (m *ModifyRequestParams) ClearTags(a *ClearTagsAction) *ModifyRequestParams {
	m.ClearTagsActions = append(m.ClearTagsActions, a)
	return m
}

func (m *ModifyRequestParams) RenameTag(a *RenameTagAction) *ModifyRequestParams {
	m.RenameTagActions = append(m.RenameTagActions, a)
	return m
}

func (m *ModifyRequestParams) DeleteTag(a *DeleteTagAction) *ModifyRequestParams {
	m.DeleteTagActions = append(m.DeleteTagActions, a)
	return m
}

type modifyRequestParams struct {
	*RequestAuthParams
	Actions []interface{} `json:"actions"`
}

type AddAction struct {
	// ItemID    string    `json:"item_id"`
	TwitterID string    `json:"ref_id,omitempty"`
	Tags      []string  `json:"-"`
	Time      time.Time `json:"-"`
	Title     string    `json:"title,omitempty"`
	URL       string    `json:"url"`
}

type addAction struct {
	*AddAction
	Action     ActionMethod `json:"action"`
	TagsString string       `json:"tags,omitempty"`
	TimeUnix   int64        `json:"time,omitempty"`
}

type ArchiveAction struct {
	ItemID string    `json:"item_id"`
	Time   time.Time `json:"-"`
}

type archiveAction struct {
	*ArchiveAction
	Action   ActionMethod `json:"action"`
	TimeUnix int64        `json:"time,omitempty"`
}

type ReaddAction struct {
	ItemID string    `json:"item_id"`
	Time   time.Time `json:"-"`
}

type readdAction struct {
	*ReaddAction
	Action   ActionMethod `json:"action"`
	TimeUnix int64        `json:"time,omitempty"`
}

type FavoriteAction struct {
	ItemID string    `json:"item_id"`
	Time   time.Time `json:"-"`
}

type favoriteAction struct {
	*FavoriteAction
	Action   ActionMethod `json:"action"`
	TimeUnix int64        `json:"time,omitempty"`
}

type UnfavoriteAction struct {
	ItemID string    `json:"item_id"`
	Time   time.Time `json:"-"`
}

type unfavoriteAction struct {
	*UnfavoriteAction
	Action   ActionMethod `json:"action"`
	TimeUnix int64        `json:"time,omitempty"`
}

type DeleteAction struct {
	ItemID string    `json:"item_id"`
	Time   time.Time `json:"-"`
}

type deleteAction struct {
	*DeleteAction
	Action   ActionMethod `json:"action"`
	TimeUnix int64        `json:"time,omitempty"`
}

type AddTagsAction struct {
	ItemID string    `json:"item_id"`
	Tags   []string  `json:"-"`
	Time   time.Time `json:"-"`
}

type addTagsAction struct {
	*AddTagsAction
	Action     ActionMethod `json:"action"`
	TagsString string       `json:"tags"`
	TimeUnix   int64        `json:"time,omitempty"`
}

type RemoveTagsAction struct {
	ItemID string    `json:"item_id"`
	Tags   []string  `json:"tags"`
	Time   time.Time `json:"-"`
}

type removeTagsAction struct {
	*RemoveTagsAction
	Action     ActionMethod `json:"action"`
	TagsString string       `json:"tags"`
	TimeUnix   int64        `json:"time,omitempty"`
}

type ReplaceTagsAction struct {
	ItemID string    `json:"item_id"`
	Tags   []string  `json:"tags"`
	Time   time.Time `json:"-"`
}

type replaceTagsAction struct {
	*ReplaceTagsAction
	Action     ActionMethod `json:"action"`
	TagsString string       `json:"tags"`
	TimeUnix   int64        `json:"time,omitempty"`
}

type ClearTagsAction struct {
	ItemID string    `json:"item_id"`
	Time   time.Time `json:"-"`
}

type clearTagsAction struct {
	*ClearTagsAction
	Action   ActionMethod `json:"action"`
	TimeUnix int64        `json:"time,omitempty"`
}

type RenameTagAction struct {
	OldTag string    `json:"old_tag"`
	NewTag string    `json:"new_tag"`
	Time   time.Time `json:"-"`
}

type renameTagAction struct {
	*RenameTagAction
	Action   ActionMethod `json:"action"`
	TimeUnix int64        `json:"time,omitempty"`
}

type DeleteTagAction struct {
	Tag  string    `json:"tag"`
	Time time.Time `json:"-"`
}

type deleteTagAction struct {
	*DeleteTagAction
	Action   ActionMethod `json:"action"`
	TimeUnix int64        `json:"time,omitempty"`
}

type ModifyRequestResponse struct {
	ActionResults []string     `json:"action_results"`
	Status        ActionStatus `json:"status"`
}

func (c *Client) Modify(ctx context.Context, params *ModifyRequestParams) (*ModifyRequestResponse, error) {
	p := &modifyRequestParams{
		RequestAuthParams: c.Auth.RequestAuthParams,
	}

	for _, v := range params.AddActions {
		p.Actions = append(p.Actions, &addAction{
			Action:     AddMethod,
			AddAction:  v,
			TagsString: strings.Join(v.Tags, ","),
			TimeUnix:   v.Time.Unix(),
		})
	}

	for _, v := range params.ArchiveActions {
		p.Actions = append(p.Actions, &archiveAction{
			Action:        ArchiveMethod,
			ArchiveAction: v,
			TimeUnix:      v.Time.Unix(),
		})
	}

	for _, v := range params.ReaddActions {
		p.Actions = append(p.Actions, &readdAction{
			Action:      ReaddMethod,
			ReaddAction: v,
			TimeUnix:    v.Time.Unix(),
		})
	}

	for _, v := range params.FavoriteActions {
		p.Actions = append(p.Actions, &favoriteAction{
			Action:         FavoriteMethod,
			FavoriteAction: v,
			TimeUnix:       v.Time.Unix(),
		})
	}

	for _, v := range params.UnfavoriteActions {
		p.Actions = append(p.Actions, &unfavoriteAction{
			Action:           UnfavoriteMethod,
			UnfavoriteAction: v,
			TimeUnix:         v.Time.Unix(),
		})
	}

	for _, v := range params.DeleteActions {
		p.Actions = append(p.Actions, &deleteAction{
			Action:       DeleteMethod,
			DeleteAction: v,
			TimeUnix:     v.Time.Unix(),
		})
	}

	for _, v := range params.AddTagsActions {
		p.Actions = append(p.Actions, &addTagsAction{
			Action:        AddTagsMethod,
			AddTagsAction: v,
			TagsString:    strings.Join(v.Tags, ","),
			TimeUnix:      v.Time.Unix(),
		})
	}

	for _, v := range params.RemoveTagsActions {
		p.Actions = append(p.Actions, &removeTagsAction{
			Action:           RemoveTagsMethod,
			RemoveTagsAction: v,
			TagsString:       strings.Join(v.Tags, ","),
			TimeUnix:         v.Time.Unix(),
		})
	}

	for _, v := range params.ReplaceTagsActions {
		p.Actions = append(p.Actions, &replaceTagsAction{
			Action:            ReplaceTagsMethod,
			ReplaceTagsAction: v,
			TagsString:        strings.Join(v.Tags, ","),
			TimeUnix:          v.Time.Unix(),
		})
	}

	for _, v := range params.ClearTagsActions {
		p.Actions = append(p.Actions, &clearTagsAction{
			Action:          ClearTagsMethod,
			ClearTagsAction: v,
			TimeUnix:        v.Time.Unix(),
		})
	}

	for _, v := range params.RenameTagActions {
		p.Actions = append(p.Actions, &renameTagAction{
			Action:          RenameTagMethod,
			RenameTagAction: v,
			TimeUnix:        v.Time.Unix(),
		})
	}

	for _, v := range params.DeleteTagActions {
		p.Actions = append(p.Actions, &deleteTagAction{
			Action:          DeleteTagMethod,
			DeleteTagAction: v,
			TimeUnix:        v.Time.Unix(),
		})
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to marshal input body")
	}

	r, err := c.post(ctx, ModifyEndpoint, b)
	if err != nil {
		return nil, err
	}

	response := &ModifyRequestResponse{}
	err = json.Unmarshal(r, response)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal response body")
	}

	return response, nil
}
