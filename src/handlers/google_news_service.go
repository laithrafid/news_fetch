package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/laithrafid/news_fetch/src/clients"
	"github.com/laithrafid/news_fetch/src/domain"
	"github.com/laithrafid/news_fetch/src/repository"
	"github.com/laithrafid/news_fetch/src/validations"
)

type GNewsService struct {
	Client *clients.GClient
}

func NewGNews(hClient *http.Client, cassession *gocql.Session) *NewsService {
	return &NewsService{
		Service: &GNewsService{
			Client:   clients.NewGNewsClient(hClient),
			DbClient: repository.(cassession),
		},
	}
}

func (g *GNewsService) GetSources(c *gin.Context) {
	qp := c.Request.URL.Query()
	ok, err := validations.RequestQParams(qp, "source")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": err.Error()})
	}
	if sources, ok := g.GetNewsSources(qp); !ok {
		fmt.Printf("Error encountered : %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": err})
	} else {
		c.JSON(http.StatusOK, sources)
	}
}

func (g *GNewsService) GetHeadlines(c *gin.Context) {
	qp := c.Request.URL.Query()
	ok, err := validations.RequestQParams(qp, "headlines")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": err.Error()})
	}
	if s, err := g.Client.GetHeadlines(qp); err != nil {
		fmt.Printf("Error encountered : %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": err})
	} else {
		var newsBySourceList []domain.NewsBySource
		for _, val := range s.ArticleList.Articles {
			tHash := md5.Sum([]byte(val.Description))
			newsBySourceList = append(newsBySourceList, domain.NewsBySource{
				SourceId:        val.Source.Id,
				TitleHash:       hex.EncodeToString(tHash[:]),
				NewsAuthor:      val.Author,
				NewsContent:     val.Content,
				NewsDescription: val.Description,
				NewsTitle:       val.Description,
				NewsUrl:         val.Url,
				NewsUrlToImage:  val.UrlToImage,
				NewsPublishedAt: val.PublishedAt,
				SourceName:      val.Source.Name,
			})
		}
		for _, val := range newsBySourceList {
			if err := g.DbClient.InsertTopNews(val); err != nil {
				log.Println(err.Error())
			} else {
				c.JSON(http.StatusOK, newsBySourceList)
			}
		}
		// c.JSON(http.StatusOK, newsBySourceList)
	}
}

func (g *GNewsService) GetNews(c *gin.Context) {
	qp := c.Request.URL.Query()
	ok, err := validations.RequestQParams(qp, "all")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": err.Error()})
		return
	}
	if s, err := g.Client.GetEverything(qp); err != nil {
		fmt.Printf("Error encountered : %v", err.Error())
		fmt.Println(s)
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": err})
		return
	} else {
		if lmt, err := strconv.Atoi(qp["top"][0]); err != nil {
			lmt = 3
		} else {
			//persist in db
			var newsBySourceList []domain.NewsBySource
			for _, val := range s.ArticleList.Articles {
				tHash := md5.Sum([]byte(val.Description))
				newsBySourceList = append(newsBySourceList, domain.NewsBySource{
					SourceId:        val.Source.Id,
					TitleHash:       hex.EncodeToString(tHash[:]),
					NewsAuthor:      val.Author,
					NewsContent:     val.Content,
					NewsDescription: val.Description,
					NewsTitle:       val.Description,
					NewsUrl:         val.Url,
					NewsUrlToImage:  val.UrlToImage,
					NewsPublishedAt: val.PublishedAt,
					SourceName:      val.Source.Name,
				})
			}
			for _, val := range newsBySourceList {
				if err := g.DbClient.InsertTopNews(val); err != nil {
					log.Println(err.Error())
					return
				}
			}

			// fetch from db and return response
			if resp, err := g.GetTopNewsBySourceFromDb(qp, lmt); err == nil {
				c.JSON(http.StatusOK, resp)
			}
		}
	}
}

func (g *GNewsService) GetNewsSources(qp map[string][]string) (*domain.NewsSource, bool) {
	if s, err := g.Client.GetSources(qp); err != nil {
		fmt.Printf("Error encountered : %v", err.Error())
		return nil, false
	} else {
		var newsBySourceList []domain.NewsBySource
		for _, val := range s.SourceList.Sources {
			newsBySourceList = append(newsBySourceList, domain.NewsBySource{
				SourceId:          val.Id,
				SourceName:        val.Name,
				SourceDescription: val.Description,
				SourceUrl:         val.Url,
				SourceCategory:    val.Category,
				SourceLanguage:    val.Language,
				SourceCountry:     val.Country,
			})
		}
		for _, val := range newsBySourceList {
			if err := g.DbClient.InsertSources(val); err != nil {
				log.Println(err.Error())
				return nil, false
			}
		}
		return s, true
	}
}

func (g *GNewsService) GetTopNewsBySourceFromDb(qp map[string][]string, lim int) ([]domain.NewsBySource, error) {
	srcs := ""
	for _, val := range qp["sources"] {
		sVal := strings.Split(val, ",")
		for _, v := range sVal {
			srcs = srcs + "'" + v + "',"
		}
	}
	srcs = strings.TrimRight(srcs, ",")
	if dbResp, err := g.DbClient.GetTopNewsBySource(srcs, lim); err == nil {
		return dbResp, nil
	} else {
		return nil, err
	}
}
