package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// New construct db mocking
func initMockGorm() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		panic(err)
	}
	return gormDB, mock
}

type AuthRepoSuite struct {
	suite.Suite
	*require.Assertions
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(AuthRepoSuite))
}

func (s *AuthRepoSuite) SetupTest() {
	s.Assertions = require.New(s.T())
}

func (s *AuthRepoSuite) TestNewRepository() {
	s.Run("#Case1", func() {
		repo := NewRepository(&gorm.DB{})
		s.NotNil(repo)
	})
}
