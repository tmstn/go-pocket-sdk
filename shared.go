package pocket

type ActionStatus int

func (a ActionStatus) Success() bool {
	return a == Success
}

func (a ActionStatus) Failure() bool {
	return a == Failure
}

const (
	Success ActionStatus = 1
	Failure ActionStatus = 0
)

type Ternary string

func (t Ternary) None() bool {
	return t == None
}

func (t Ternary) Has() bool {
	return t == Has
}

func (t Ternary) Is() bool {
	return t == Is
}

const (
	None Ternary = "0"
	Has  Ternary = "1"
	Is   Ternary = "2"
)

type Bool string

const (
	Yes Bool = "1"
	No  Bool = "0"
)

func (b Bool) Yes() bool {
	return b == Yes
}

func (b Bool) No() bool {
	return b == No
}

type LinkState string

func (l LinkState) Unread() bool {
	return l == Unread
}

func (l LinkState) Archive() bool {
	return l == Archive
}

func (l LinkState) All() bool {
	return l == All
}

const (
	Unread  LinkState = "unread"
	Archive LinkState = "archive"
	All     LinkState = "all"
)

type LinkStatus string

func (l LinkStatus) Active() bool {
	return l == Active
}

func (l LinkStatus) Archived() bool {
	return l == Archived
}

func (l LinkStatus) Delete() bool {
	return l == Delete
}

const (
	Active   LinkStatus = "0"
	Archived LinkStatus = "1"
	Delete   LinkStatus = "2"
)

type ContentType string

func (c ContentType) Article() bool {
	return c == ArticleType
}

func (c ContentType) Video() bool {
	return c == VideoType
}

func (c ContentType) Image() bool {
	return c == ImageType
}

const (
	ArticleType ContentType = "article"
	VideoType   ContentType = "video"
	ImageType   ContentType = "image"
)

type SortOrder string

func (s SortOrder) Newest() bool {
	return s == Newest
}

func (s SortOrder) Oldest() bool {
	return s == Oldest
}

func (s SortOrder) Title() bool {
	return s == Title
}

func (s SortOrder) Site() bool {
	return s == Site
}

const (
	Newest SortOrder = "newest"
	Oldest SortOrder = "oldest"
	Title  SortOrder = "title"
	Site   SortOrder = "site"
)

type DetailType string

func (d DetailType) Simple() bool {
	return d == Simple
}

func (d DetailType) Complete() bool {
	return d == Complete
}

const (
	Simple   DetailType = "simple"
	Complete DetailType = "complete"
)

const Untagged string = "_untagged_"

type Tags map[string]Tag

type Tag struct {
	ItemID   string `json:"item_id"`
	AuthorID string `json:"author_id"`
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
}

type Authors map[string]Author

type Author struct {
	ItemID   string `json:"item_id"`
	AuthorID string `json:"author_id"`
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
}

type Images map[string]Image

type Image struct {
	ItemID  string `json:"item_id"`
	ImageID string `json:"image_id"`
	Src     string `json:"src"`
	Width   string `json:"width,omitempty"`
	Height  string `json:"height,omitempty"`
	Credit  string `json:"credit,omitempty"`
	Caption string `json:"caption,omitempty"`
}

type Videos map[string]Video

type Video struct {
	ItemID  string `json:"item_id"`
	VideoID string `json:"video_id"`
	Name    string `json:"name,omitempty"`
	Src     string `json:"src,omitempty"`
	Width   string `json:"width,omitempty"`
	Height  string `json:"height,omitempty"`
	Type    string `json:"type,omitempty"`
	Vid     string `json:"vid,omitempty"`
}
