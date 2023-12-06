package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/lenismtho/area/pkg/core/gormModels"
	"github.com/lenismtho/area/pkg/core/models"
	"github.com/lenismtho/area/pkg/core/providers"
)

type DB interface {
	/*** USER ***/
	CreateUser(user models.User) error
	GetUserByEmail(email string) (*models.User, error)
	UpdateUserToken(email string, token string) error
	GetUserByToken(token string) (*models.User, error)
	UpdateUserMiddlewareTokenByEmail(email string, token string) error
	GetUserAreasByUserID(userID uint) ([]models.Area, error)
	UpdateUserUsername(username string, userID uint) error
	UpdateUserEmail(email string, userID uint) error
	UpdateUserPassword(password string, userID uint) error
	DeleteUser(userID uint) error
	UpdateUserFirstName(firstName string, userID uint) error
	UpdateUserLastName(lastName string, userID uint) error
	/*** AUTH ***/
	UpdateUserDiscordToken(externalId uint, discordToken string) error
	/*** SERVICES ***/
	GetUserByID(id uint) (*models.User, error)
	SetUserService(service models.Service) error
	GetUserServiceTokenCredentialsByUserID(userID uint) (providers.TokenCredentials, error)
	UpdateServiceTokenCredentialsByUserID(credentials providers.TokenCredentials, userID uint) error
	GetAServiceAccessTokenByUserID(userID uint, service string) (string, error)
	GetServicesByUserID(userID uint) ([]models.Service, error)
	DeleteServicesByUserID(userId uint) error
	DeleteServiceByUserIDAndName(userId uint, name string) error
	/*** AREAS ***/
	CreateArea(area models.Area) (uint, error)
	GetAreaByID(id uint) (*models.Area, error)
	DeleteAreaByID(areaID string) error
}

type database struct {
	db *gorm.DB
}

func NewDB() (DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = db.AutoMigrate(&gormModels.User{}, &gormModels.Service{}, &gormModels.Area{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &database{
		db: db,
	}, nil
}
