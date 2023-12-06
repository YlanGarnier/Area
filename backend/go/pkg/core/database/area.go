package database

import (
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/lenismtho/area/pkg/core/gormModels"
	"github.com/lenismtho/area/pkg/core/models"
	"github.com/lenismtho/area/pkg/protogen"
)

func (d *database) CreateArea(area models.Area) (uint, error) {
	raw, err := proto.Marshal(area.Base)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal base: %w", err)
	}

	base := base64.StdEncoding.EncodeToString(raw)
	newArea := gormModels.Area{
		Name:        area.Name,
		ActService:  area.ActService,
		RouteAction: area.RouteAction,
		ReaService:  area.ReaServices,
		Route:       area.Route,
		Base:        base,
		UserID:      area.UserID,
	}

	tx := d.db.Create(&newArea)
	if tx.Error != nil {
		return 0, fmt.Errorf("failed to create area: %w", tx.Error)
	}
	return newArea.ID, nil
}

func (d *database) GetAreaByID(id uint) (*models.Area, error) {
	var area gormModels.Area
	tx := d.db.First(&area, id)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to get area by id: %w", tx.Error)
	}
	raw, err := base64.StdEncoding.DecodeString(area.Base)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base: %w", err)
	}
	var base protogen.Base
	err = proto.Unmarshal(raw, &base)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal base: %w", err)
	}
	return &models.Area{
		Name:        area.Name,
		ActService:  area.ActService,
		RouteAction: area.RouteAction,
		ReaServices: area.ReaService,
		Route:       area.Route,
		Base:        &base,
		UserID:      area.UserID,
	}, nil
}

func (d *database) DeleteAreaByID(areaID string) error {
	var area gormModels.Area
	tx := d.db.Where("id = ?", areaID).Delete(&area)
	if tx.Error != nil {
		return fmt.Errorf("failed to delete area by id: %w", tx.Error)
	}
	return nil
}
