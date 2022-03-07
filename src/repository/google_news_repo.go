package repository

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/laithrafid/bookstore_utils-go/errors_utils"
	"github.com/laithrafid/news_fetch/src/cassandra"
	"github.com/laithrafid/news_fetch/src/domain"
)

const (
	queryGetTopNewsBySource = "SELECT sid, created_at, title_hash, nauthor, ncontent, ndesc, npublished_at, ntitle, nurl, nurl_to_image, sname, scategory, sdesc, scountry, slang, surl from news_by_source where sid in (%s) limit ?"
	queryInsterTopNews      = "insert into news_by_source(sid, created_at, title_hash, nauthor, ncontent, ndesc, sname, npublished_at) values (?,?,?,?,?,?,?,?)"
)

func NewNewsRepo() NewsRepo {
	return &newsrepo{}
}

type NewsRepo interface {
	GetTopNewsBySource(sid string, lim int) ([]domain.NewsBySource, error)
	InsertTopNews(m domain.NewsBySource) error
}

type newsrepo struct {
}

func (r *newsrepo) GetTopNewsBySource(sid string, lim int) ([]domain.NewsBySource, error) {
	var newsEnt []domain.NewsBySource
	m := map[string]interface{}{}
	query := fmt.Sprintf(queryGetTopNewsBySource, sid)
	iter := cassandra.GetSession().Query(query, lim).Consistency(gocql.One).Iter()
	for iter.MapScan(m) {
		newsEnt = append(newsEnt, domain.NewsBySource{
			SourceId: fmt.Sprintf("%v", m["sid"]),
			// CreatedAt:       m["created_at"].(time.Time),
			NewsAuthor:        fmt.Sprintf("%v", m["nauthor"]),
			NewsContent:       fmt.Sprintf("%v", m["ncontent"]),
			NewsDescription:   fmt.Sprintf("%v", m["ndesc"]),
			NewsPublishedAt:   fmt.Sprintf("%v", m["npublished_at"]),
			NewsTitle:         fmt.Sprintf("%v", m["ntitle"]),
			NewsUrl:           fmt.Sprintf("%v", m["nurl"]),
			NewsUrlToImage:    fmt.Sprintf("%v", m["nurl_to_image"]),
			SourceName:        fmt.Sprintf("%v", m["sname"]),
			SourceCategory:    fmt.Sprintf("%v", m["ncategory"]),
			SourceDescription: fmt.Sprintf("%v", m["sdesc"]),
			SourceCountry:     fmt.Sprintf("%v", m["scountry"]),
			SourceLanguage:    fmt.Sprintf("%v", m["slang"]),
			SourceUrl:         fmt.Sprintf("%v", m["surl"]),
		})
		m = map[string]interface{}{}
	}
	return newsEnt, nil
}

func (r *newsrepo) InsertTopNews(m domain.NewsBySource) error {
	if err := cassandra.GetSession().Query(queryInsterTopNews,
		m.SourceId,
		gocql.TimeUUID(),
		m.TitleHash,
		m.NewsAuthor,
		m.NewsContent,
		m.NewsDescription,
		m.SourceName,
		m.NewsPublishedAt).Consistency(gocql.One).Exec(); err != nil {
		return errors_utils.NewInternalServerError("error when trying to save access token in database", err)
	}
	return nil
}
