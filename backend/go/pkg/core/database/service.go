package database

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/lenismtho/area/pkg/core/gormModels"
	"github.com/lenismtho/area/pkg/core/models"
	"github.com/lenismtho/area/pkg/core/providers"
	"gorm.io/gorm"
)

func (d *database) SetUserService(service models.Service) error {
	user, err := d.getUserGormByID(service.UserID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	err = d.db.Model(&gormModels.Service{}).Where("identifier = ? AND name = ?", service.Identifier, service.Name).First(&gormModels.Service{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Services = append(user.Services, &gormModels.Service{
			Name:         service.Name,
			Identifier:   service.Identifier,
			Token:        service.AccessToken,
			RefreshToken: service.RefreshToken,
			UserID:       service.UserID,
		})
		return d.db.Save(&user).Error
	}
	return fmt.Errorf("service already exists: %w", err)
}

func (d *database) GetUserServiceTokenCredentialsByUserID(userID uint) (providers.TokenCredentials, error) {
	var service gormModels.Service
	err := d.db.Model(&gormModels.Service{}).Where("user_id = ?", userID).First(&service).Error
	if err != nil {
		return providers.TokenCredentials{}, fmt.Errorf("failed to get service: %w", err)
	}
	return providers.TokenCredentials{
		AccessToken:  service.Token,
		RefreshToken: service.RefreshToken,
	}, nil
}

func (d *database) UpdateServiceTokenCredentialsByUserID(credentials providers.TokenCredentials, userID uint) error {
	var service gormModels.Service
	err := d.db.Model(&gormModels.Service{}).Where("user_id = ?", strconv.Itoa(int(userID))).First(&service).Error
	if err != nil {
		return fmt.Errorf("failed to get service: %w", err)
	}
	service.Token = credentials.AccessToken
	service.RefreshToken = credentials.RefreshToken
	return d.db.Save(&service).Error
}

func (d *database) GetAServiceAccessTokenByUserID(userID uint, service string) (string, error) {
	var serviceDB gormModels.Service
	err := d.db.Model(&gormModels.Service{}).Where("user_id = ? AND name = ?", userID, service).First(&serviceDB).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", fmt.Errorf("failed to get service, service not found: %w", err)
	} else if err != nil {
		return "", fmt.Errorf("failed to get service: %w", err)
	}
	return serviceDB.Token, nil
}

func (d *database) GetServicesByUserID(userID uint) ([]models.Service, error) {
	var services []gormModels.Service
	err := d.db.Model(&gormModels.Service{}).Where("user_id = ?", userID).Find(&services).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get services: %w", err)
	}
	var servicesToReturn []models.Service
	for _, service := range services {
		servicesToReturn = append(servicesToReturn, models.Service{
			Name:         service.Name,
			Identifier:   service.Identifier,
			AccessToken:  service.Token,
			RefreshToken: service.RefreshToken,
			UserID:       service.UserID,
		})
	}
	return servicesToReturn, nil
}

func (d *database) DeleteServicesByUserID(userId uint) error {
	err := d.db.Model(&gormModels.Service{}).Where("user_id = ?", userId).Delete(&gormModels.Service{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete services: %w", err)
	}
	return nil
}

func (d *database) DeleteServiceByUserIDAndName(userId uint, name string) error {
	err := d.db.Model(&gormModels.Service{}).Where("user_id = ? AND name = ?", userId, name).Delete(&gormModels.Service{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete services: %w", err)
	}
	return nil
}
