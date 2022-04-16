package repository

import (
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func (s *AuthRepoSuite) Test_authRepo_GetByPhone() {
	expectSelectQuery := `SELECT .+`
	forceErr := errors.New("force err")

	rowsKey := []string{"test"}
	rowsData := sqlmock.NewRows(rowsKey).AddRow("test")

	s.Run("#Case1: Positif", func() {
		db, mock := initMockGorm()
		defer func(db *gorm.DB) {
			_ = db.Close()
		}(db)

		mock.ExpectQuery(expectSelectQuery).WillReturnRows(rowsData)

		data, err := NewRepository(db).GetByPhone("081281288822")
		s.NotNil(data)
		s.Nil(err)
	})

	s.Run("#Case1: Negative", func() {
		db, mock := initMockGorm()
		defer func(db *gorm.DB) {
			_ = db.Close()
		}(db)

		mock.ExpectQuery(expectSelectQuery).WillReturnError(forceErr)

		data, err := NewRepository(db).GetByPhone("081281288822")
		s.Nil(data)
		s.NotNil(err)
		s.Equal(err, forceErr)
	})
}
