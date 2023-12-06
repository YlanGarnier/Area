package database

import (
	"errors"
	"fmt"

	"github.com/lenismtho/area/pkg/core/authentication"
	"github.com/lenismtho/area/pkg/core/gormModels"
	"github.com/lenismtho/area/pkg/core/models"
	"gorm.io/gorm"
)

func (d *database) UpdateUserDiscordToken(externalId uint, discordToken string) error {
	panic("not implemented")
}

func (d *database) CreateUser(user models.User) error {
	newUser := gormModels.User{
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Kind:      user.Kind,
		Token:     user.Token,
	}
	tx := d.db.Where("email = ? AND kind = ?", user.Email, user.Kind).First(&gormModels.User{})
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		tx = d.db.Create(&newUser)
		if tx.Error != nil {
			return fmt.Errorf("failed to create user: %w", tx.Error)
		}
	} else {
		return fmt.Errorf("user already exists: %w", tx.Error)
	}
	return nil
}

func (d *database) getUserByEmail(email string) (*gormModels.User, error) {
	var user gormModels.User
	tx := d.db.Where("email = ?", email).First(&user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to find user: %w", tx.Error)
	}
	return &user, nil
}

func (d *database) GetUserByEmail(email string) (*models.User, error) {
	user, err := d.getUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Username:  user.Username,
		Token:     user.Token,
		Kind:      user.Kind,
	}, nil
}

func (d *database) UpdateUserToken(email string, token string) error {
	tx := d.db.Model(&gormModels.User{}).Where("email = ?", email).Update("token", token)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to find user: %w", tx.Error)
	} else if tx.Error != nil {
		return fmt.Errorf("failed to update user token: %w", tx.Error)
	}
	return nil
}

func (d *database) GetUserByToken(token string) (*models.User, error) {
	hashToken, err := authentication.HashOAT(token)
	if err != nil {
		return nil, fmt.Errorf("failed to hash OAT: %w", err)
	}
	var user gormModels.User
	tx := d.db.Where("token = ?", hashToken).First(&user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to find user: %w", tx.Error)
	} else if tx.Error != nil {
		return nil, fmt.Errorf("failed to find user: %w", tx.Error)
	}
	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Username:  user.Username,
		Token:     user.Token,
		Kind:      user.Kind,
	}, nil
}

func (d *database) getUserGormByID(id uint) (*gormModels.User, error) {
	var user gormModels.User
	tx := d.db.Where("id = ?", id).First(&user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to find user: %w", tx.Error)
	} else if tx.Error != nil {
		return nil, fmt.Errorf("failed to find user: %w", tx.Error)
	}
	return &user, nil
}

func (d *database) GetUserByID(id uint) (*models.User, error) {
	user, err := d.getUserGormByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Username:  user.Username,
		Token:     user.Token,
		Kind:      user.Kind,
	}, nil
}

func (d *database) UpdateUserMiddlewareTokenByEmail(email string, token string) error {
	err := d.db.Model(&gormModels.User{}).Where("email = ?", email).Update("token", token).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to find user: %w", err)
	} else if err != nil {
		return fmt.Errorf("failed to update user token: %w", err)
	}
	return nil
}

func (d *database) GetUserAreasByUserID(userID uint) ([]models.Area, error) {
	var areas []gormModels.Area
	tx := d.db.Where("user_id = ?", userID).Find(&areas)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to find areas: %w", tx.Error)
	} else if tx.Error != nil {
		return nil, fmt.Errorf("failed to find areas: %w", tx.Error)
	}
	var areasModels []models.Area
	for _, area := range areas {
		areasModels = append(areasModels, models.Area{
			ID:          area.ID,
			Name:        area.Name,
			ActService:  area.ActService,
			RouteAction: area.RouteAction,
			ReaServices: area.ReaService,
			Route:       area.Route,
			UserID:      area.UserID,
		})
	}
	return areasModels, nil
}

func (d *database) UpdateUserUsername(username string, userID uint) error {
	user, err := d.getUserGormByID(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to find user: %w", err)
	}
	err = d.db.Model(&user).Update("username", username).Error
	if err != nil {
		return fmt.Errorf("failed to update user username: %w", err)
	}
	return nil
}

func (d *database) UpdateUserEmail(email string, userID uint) error {
	user, err := d.getUserGormByID(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to find user: %w", err)
	}
	err = d.db.Model(&user).Update("email", email).Error
	if err != nil {
		return fmt.Errorf("failed to update user email: %w", err)
	}
	return nil
}

func (d *database) UpdateUserPassword(password string, userID uint) error {
	user, err := d.getUserGormByID(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to find user: %w", err)
	}
	err = d.db.Model(&user).Update("password", password).Error
	if err != nil {
		return fmt.Errorf("failed to update password email: %w", err)
	}
	return nil
}

func (d *database) DeleteUser(userID uint) error {
	user, err := d.getUserGormByID(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to find user: %w", err)
	}
	err = d.db.Delete(&user).Error
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (d *database) UpdateUserFirstName(firstName string, userID uint) error {
	user, err := d.getUserGormByID(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to find user: %w", err)
	}
	err = d.db.Model(&user).Update("first_name", firstName).Error
	if err != nil {
		return fmt.Errorf("failed to update user first name: %w", err)
	}
	return nil
}

func (d *database) UpdateUserLastName(lastName string, userID uint) error {
	user, err := d.getUserGormByID(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to find user: %w", err)
	}
	err = d.db.Model(&user).Update("last_name", lastName).Error
	if err != nil {
		return fmt.Errorf("failed to update user last name: %w", err)
	}
	return nil
}
