package repository

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/laithrafid/news_fetch/src/cassandra"
	"github.com/laithrafid/news_fetch/src/domain"
)

func NewUserRepo() UserRepo {
	return &userrepo{}
}

type UserRepo interface {
	GetUserByTgDetils(ID int) (domain.UserDetails, error)
	InsertUser(m domain.UserDetails) error
	GetAllUser() ([]domain.UserDetails, error)
}

type userrepo struct {
}

func (u *userrepo) GetUserByTgDetils(ID int) (domain.UserDetails, error) {
	m := map[string]interface{}{}
	query := fmt.Sprintf("SELECT ID, name from user where ID = %d ", ID)
	fmt.Println(query)
	iter := cassandra.GetSession().Query(query).Consistency(gocql.One).Iter()
	var subscriber domain.UserDetails
	for iter.MapScan(m) {
		if ID, ok := m["id"].(int); ok {
			subscriber = domain.UserDetails{
				ID:   int64(ID),
				Name: fmt.Sprintf("%v", m["name"]),
			}
		}
		m = map[string]interface{}{}
	}
	return subscriber, nil
}

func (u *userrepo) InsertUser(m domain.UserDetails) error {
	query := "insert into user(uid, name) values (?,?)"
	if err := cassandra.GetSession().Query(query, m.ID, m.Name).Consistency(gocql.One).Exec(); err != nil {
		fmt.Errorf("Error encountered : %s", err.Error())
		return err
	}
	return nil
}

func (u *userrepo) GetAllUser() ([]domain.UserDetails, error) {
	m := map[string]interface{}{}
	query := "SELECT uid, name from user"
	iter := cassandra.GetSession().Query(query).Consistency(gocql.One).Iter()
	var subscribers []domain.UserDetails
	for iter.MapScan(m) {
		if id, ok := m["uid"].(int); ok {
			subscribers = append(subscribers, domain.UserDetails{
				ID:   int64(id),
				Name: fmt.Sprintf("%v", m["name"]),
			})
			m = map[string]interface{}{}
		}
	}
	return subscribers, nil
}
