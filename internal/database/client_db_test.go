package database

import (
	"database/sql"
	"testing"

	"github.com/davifs92/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	client := &entity.Client{
		Id:    "1",
		Name:  "Test",
		Email: "j@j.com",
	}

	err := s.clientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("John", "j@j.com")
	err := s.clientDB.Save(client)
	s.Nil(err)

	retrievedClient, err := s.clientDB.Get("1")
	s.Nil(err)
	s.Equal(client.Id, retrievedClient.Id)
	s.Equal(client.Name, retrievedClient.Name)
	s.Equal(client.Email, retrievedClient.Email)
}
