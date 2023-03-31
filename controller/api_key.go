package controller

import "github.com/shoriwe/routes-service/models"

func (c *Controller) CreateAPIKey(apiKey *models.APIKey) error {
	return c.DB.Create(apiKey).Error
}

func (c *Controller) DeleteAPIKey(apiKeyUUID string) error {
	return c.DB.Where("uuid = ?", apiKeyUUID).Delete(&models.APIKey{}).Error
}

type APIKeyFilter struct {
	Page        int64   `json:"page"`
	PageSize    int64   `json:"pageSize"`
	VehicleUUID *string `json:"vehicleUUID"`
}

func (c *Controller) QueryAPIKeys(filter *APIKeyFilter) (*Result[*models.APIKey], error) {
	if filter.PageSize == 0 {
		filter.PageSize = DefaultPageSize
	}
	var totalResults int64
	count := c.DB.Model(&models.APIKey{})
	if filter.VehicleUUID != nil {
		count = count.Where("vehicle_uuid = ?", *filter.VehicleUUID)
	}
	qErr := count.Count(&totalResults).Error
	if qErr != nil {
		return nil, qErr
	}
	if totalResults == 0 {
		return &Result[*models.APIKey]{
			Page:       filter.Page,
			TotalPages: 0,
		}, nil
	}
	result := &Result[*models.APIKey]{
		Page:       filter.Page,
		TotalPages: totalResults / filter.PageSize,
	}
	if result.TotalPages == 0 {
		result.TotalPages = 1
	}
	query := c.DB.Offset(int(filter.Page - 1)).Limit(DefaultPageSize)
	if filter.VehicleUUID != nil {
		query = query.Where("vehicle_uuid = ?", *filter.VehicleUUID)
	}
	sErr := query.Find(&result.Results).Error
	return result, sErr
}
