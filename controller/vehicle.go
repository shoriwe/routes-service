package controller

import "github.com/shoriwe/routes-service/models"

func (c *Controller) CreateVehicle(vehicle *models.Vehicle) error {
	return c.DB.Create(vehicle).Error
}

func (c *Controller) DeleteVehicle(vehicleUUID string) error {
	return c.DB.Where("uuid = ?", vehicleUUID).Delete(&models.Vehicle{}).Error
}

func (c *Controller) UpdateVehicle(vehicle *models.Vehicle) error {
	return c.DB.Updates(vehicle).Error
}

type VehicleFilter struct {
	Page     int64   `json:"page"`
	PageSize int64   `json:"pageSize"`
	Name     *string `json:"name,omitempty"`
	Plate    *string `json:"plate,omitempty"`
}

func (c *Controller) QueryVehicles(filter *VehicleFilter) (*Result[*models.Vehicle], error) {
	if filter.PageSize == 0 {
		filter.PageSize = DefaultPageSize
	}
	var totalResults int64
	count := c.DB.Model(&models.Vehicle{})
	if filter.Name != nil {
		count = count.Where("name LIKE ?", *filter.Name)
	}
	if filter.Plate != nil {
		count = count.Where("plate LIKE ?", *filter.Plate)
	}
	qErr := count.Count(&totalResults).Error
	if qErr != nil {
		return nil, qErr
	}
	if totalResults == 0 {
		return &Result[*models.Vehicle]{
			Page:       filter.Page,
			TotalPages: 0,
		}, nil
	}
	result := &Result[*models.Vehicle]{
		Page:       filter.Page,
		TotalPages: totalResults / filter.PageSize,
	}
	if result.TotalPages == 0 {
		result.TotalPages = 1
	}
	query := c.DB.Offset(int(filter.Page - 1)).Limit(DefaultPageSize)
	if filter.Name != nil {
		query = query.Where("name LIKE ?", *filter.Name)
	}
	if filter.Plate != nil {
		query = query.Where("plate LIKE ?", *filter.Plate)
	}
	sErr := query.Find(&result.Results).Error
	return result, sErr
}
