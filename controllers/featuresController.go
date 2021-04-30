package controllers

import (
	"net/url"

	"github.com/backend/database"
	"github.com/backend/models"
	"github.com/gofiber/fiber/v2"
)

// Features Method
func Features(c *fiber.Ctx) error {
	var feature []models.Feature

	database.DB.Find(&feature)
	return c.JSON(feature)
}

// CreateFeature Method
func CreateFeature(c *fiber.Ctx) error {
	var data models.Feature

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	feature := models.Feature{
		Name:               data.Name,
		IsActive:           data.IsActive,
		Description:        data.Description,
		DeactivationReason: data.DeactivationReason,
	}

	result := database.DB.Create(&feature)

	return c.JSON(result)
}

// UpdateFeature Method
func UpdateFeature(c *fiber.Ctx) error {
	id := c.Params("id")
	newState := c.Params("is_active")
	deactivationReason := c.Params("deactivation_reason")
	decodedDeactivationReason, _ := url.QueryUnescape(deactivationReason)

	var feature models.Feature

	database.DB.Table("features").Raw("UPDATE features SET is_active = " + newState + ", deactivation_reason = \"" + decodedDeactivationReason + "\" WHERE id = " + id + ";").Scan(&feature)

	return c.JSON(feature)
}
