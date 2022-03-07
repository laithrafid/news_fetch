package domain

import "time"

type NewsResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
}
type NewsBySource struct {
	SourceId          string    `json:"sid"`
	CreatedAt         time.Time `json:"created_at"`
	TitleHash         string    `json:"title_hash"`
	NewsAuthor        string    `json:"nauthor"`
	NewsContent       string    `json:"ncontent"`
	NewsDescription   string    `json:"ndesc"`
	NewsPublishedAt   string    `json:"npublished_at"`
	NewsTitle         string    `json:"ntitle"`
	NewsUrl           string    `json:"nurl"`
	NewsUrlToImage    string    `json:"nurl_to_image"`
	SourceCategory    string    `json:"scategory"`
	SourceDescription string    `json:"sdesc"`
	SourceCountry     string    `json:"scountry"`
	SourceLanguage    string    `json:"slang"`
	SourceName        string    `json:"sname"`
	SourceUrl         string    `json:"surl"`
}

type Source struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

type SourceList struct {
	Sources []Source `json:"sources"`
}

type ArticleSource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	Source      ArticleSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Url         string        `json:"url"`
	UrlToImage  string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
}

type ArticleList struct {
	Articles []Article `json:"articles"`
}

type TopHeadline struct {
	NewsResponse
	ArticleList
}

type Everything struct {
	NewsResponse
	ArticleList
}

type NewsSource struct {
	NewsResponse
	SourceList
}
