package controller

import "github.com/shoriwe/routes-service/models"

func (c *Controller) CreateUser(user *models.User) error {
	return c.DB.Create(user).Error
}

func (c *Controller) DeleteUser(userUUID string) error {
	return c.DB.Where("uuid = ?", userUUID).Delete(&models.User{}).Error
}

func (c *Controller) UpdateUser(user *models.User) error {
	return c.DB.Updates(user).Error
}

type UserFilter struct {
	Page     int64            `json:"page"`
	PageSize int64            `json:"pageSize"`
	Username *string          `json:"username,omitempty"`
	Type     *models.UserType `json:"userType,omitempty"`
}

func (c *Controller) QueryUsers(filter *UserFilter) (*Result[*models.User], error) {
	if filter.PageSize == 0 {
		filter.PageSize = DefaultPageSize
	}
	var totalResults int64
	count := c.DB.Model(&models.User{})
	if filter.Username != nil {
		count = count.Where("username LIKE ?", *filter.Username)
	}
	if filter.Type != nil {
		count = count.Where("type = ?", *filter.Type)
	}
	qErr := count.Count(&totalResults).Error
	if qErr != nil {
		return nil, qErr
	}
	if totalResults == 0 {
		return &Result[*models.User]{
			Page:       filter.Page,
			TotalPages: 0,
		}, nil
	}
	result := &Result[*models.User]{
		Page:       filter.Page,
		TotalPages: totalResults / filter.PageSize,
	}
	if result.TotalPages == 0 {
		result.TotalPages = 1
	}
	query := c.DB.Offset(int(filter.Page - 1)).Limit(DefaultPageSize)
	if filter.Username != nil {
		query = query.Where("username LIKE ?", *filter.Username)
	}
	if filter.Type != nil {
		query = query.Where("type = ?", *filter.Type)
	}
	sErr := query.Find(&result.Results).Error
	return result, sErr
}
