package repository

import (
	"errors"

	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/jinzhu/gorm"
)

func (s *AuthRepoSuite) Test_authRepo_Create() {
	expectInsertQuery := `INSERT .+`
	forceErr := errors.New("force err")

	// Need to mock time as well
	// s.Run("#Case1: Positif", func() {
	// 	db, mock := initMockGorm()
	// 	defer func(db *gorm.DB) {
	// 		_ = db.Close()
	// 	}(db)

	// 	mock.ExpectBegin().WillReturnError(nil)
	// 	mock.ExpectExec(expectInsertQuery).
	// 		WithArgs(time.Now(), time.Now(), time.Now(), "0812222222222", "test", "test", "test").
	// 		WillReturnError(nil)

	// 	data, err := NewRepository(db).Create(models.Auth{})
	// 	s.NotNil(data)
	// 	s.Nil(err)
	// })

	s.Run("#Case1: Negative", func() {
		db, mock := initMockGorm()
		defer func(db *gorm.DB) {
			_ = db.Close()
		}(db)

		mock.ExpectBegin().WillReturnError(nil)
		mock.ExpectExec(expectInsertQuery).WillReturnError(forceErr)

		data, err := NewRepository(db).Create(models.Auth{})
		s.Nil(data)
		s.NotNil(err)
		s.Equal(err, forceErr)
	})
}
