package xkcdindex

type Comic struct {
	Month      string `json:"month,omitempty"`
	Num        int    `json:"num,omitempty"`
	Link       string `json:"link,omitempty"`
	Year       string `json:"year,omitempty"`
	News       string `json:"news,omitempty"`
	SafeTitle  string `json:"safe_title,omitempty"`
	Transcript string `json:"transcript,omitempty"`
	Img        string `json:"img,omitempty"`
	Title      string `json:"title,omitempty"`
	Day        string `json:"day,omitempty"`
	URL        string
}
