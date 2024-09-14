package models

type Url struct {
	Url string `json:"url"`
}

type UrlDB struct {
	Url        string `json:"url"`
	CreatedURL string `json:"created_url"`
}
