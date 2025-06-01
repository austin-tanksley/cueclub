package data

type Mix struct {
	ID         int64  `json:"id,omitempty"`
	YTID       string `json:"ytid,omitempty"`
	URL        string `json:"url,omitempty"`
	Title      string `json:"title"`
	Creator    string `json:"author_name"`
	CreatorURL string `json:"author_url"`
	Version    int    `json:"-"`
}
