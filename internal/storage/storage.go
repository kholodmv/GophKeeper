package storage

import (
	"errors"
	"github.com/kholodmv/GophKeeper/internal/database"
	"github.com/kholodmv/GophKeeper/internal/models"
	"gorm.io/gorm"
)

// ErrDuplicate - error duplicate
var (
	ErrDuplicate = errors.New("duplicate key value violates unique")
)

type Storage struct {
	db *gorm.DB
}

// NewStorage - storage constructor
func NewStorage(conf string) *Storage {
	return &Storage{
		db: database.InitDB(conf),
	}
}

// CreateUser - create new user on db
func (s *Storage) CreateUser(user *models.User) error {
	result := s.db.Create(user)

	if result.Error != nil && errors.Is(result.Error, ErrDuplicate) {
		return ErrDuplicate
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}

// ReadUser - return user model by login
func (s *Storage) ReadUser(login string) (*models.User, error) {
	var user *models.User

	result := s.db.First(&user, "login = ?", login)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// CreateSecret - create new user secret
func (s *Storage) CreateSecret(secret *models.Secret) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(secret).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil && errors.Is(err, ErrDuplicate) {
		return ErrDuplicate
	} else if err != nil {
		return err
	}
	return nil
}

// ReadSecret - read user secret
func (s *Storage) ReadSecret(title string, uid uint) (*models.Secret, error) {
	var secret *models.Secret

	result := s.db.Where("title = ? AND user_id = ?", title, uid).Find(&secret)

	if result.Error != nil {
		return nil, result.Error
	}

	return secret, nil
}
