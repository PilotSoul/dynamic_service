package controllers

import (
	"PilotSoul/dynamic_service/src/domain"
	"PilotSoul/dynamic_service/src/infrastructure"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// CreateSegment func creates a segment.
// @Description Создание сегмента.
// @Accept json
// @Param input body domain.Segment true "Название сегмента"
// @Success 200 {object} domain.Segment
// @Router /create_segment [post]
func CreateSegment(c *fiber.Ctx) error {
	// Создание сегмента
	fmt.Println("Hellllo")
	segment := new(domain.Segment)
	if err := c.BodyParser(segment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if segment.Name == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Name cannot be empty",
		})
	}
	segment_db := infrastructure.DB.Db.Where("Name = ?", segment.Name).First(&segment)
	if segment_db.RowsAffected > 0 {
		return fmt.Errorf("segment with name %s already exist", segment.Name)
	}
	infrastructure.DB.Db.Create(&segment)
	return c.Status(200).JSON(segment)
}

// DeleteSegment func deletes a segment.
// @Description Удаление сегмента.
// @Accept json
// @Param input body domain.Segment true "Название сегмента"
// @Success 200
// @Router /delete_segment [post]
func DeleteSegment(c *fiber.Ctx) error {
	// Удаление сегмента
	segment := new(domain.Segment)
	if err := c.BodyParser(segment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	segment_name := segment.Name
	db := infrastructure.DB.Db.Where("Name = ?", segment_name).Delete(&segment)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected < 1 {
		return fmt.Errorf("row with name=%s cannot be deleted because it doesn't exist", segment_name)
	}
	return c.Status(200).JSON("Segment deleted")
}

type UserSegments struct {
	UserID   int      `json:"user_id"`
	Segments []string `json:"segments"`
}

// AddSegmentыfunc add segments to user.
// @Description Добавление сегментов пользователю.
// @Accept json
// @Param input body UserSegments true "Пользовательские сегменты"
// @Success 200
// @Router /add_user_to_segment [post]
func AddSegments(c *fiber.Ctx) error {
	// Добавление сегментов пользователю
	userSegments := new(UserSegments)
	if err := c.BodyParser(&userSegments); err != nil {
		return err
	}
	user := new(domain.User)
	user_db := infrastructure.DB.Db.Where("ID = ?", userSegments.UserID).First(&user)
	if user_db.RowsAffected < 1 {
		return fmt.Errorf("user with id=%d doesn't exist", userSegments.UserID)
	} else if user_db.Error != nil {
		return user_db.Error
	}
	for i := 0; i < len(userSegments.Segments); i++ {
		segment := new(domain.Segment)

		segment_name := userSegments.Segments[i]
		db := infrastructure.DB.Db.Where("Name = ?", segment_name).First(&segment)
		if db.RowsAffected < 1 {
			return fmt.Errorf("segment %s cannot be added to user because it doesn't exist", segment_name)
		} else if db.Error != nil {
			return db.Error
		}
		userSegment := []domain.UserSegment{{UserID: user.ID, SegmentID: segment.ID}}
		// db.Model(&user).Association("Segments").Append(&segment_name)
		infrastructure.DB.Db.Create(&userSegment)

	}
	return c.Status(200).JSON("Segments added")
}

type UserId struct {
	ID int `json:"ID"`
}

// ShowUserSegments func show user's segments.
// @Description Вывод списка активных сегментов у пользователя.
// @Accept json
// @Param input body UserId true "Пользователь"
// @Success 200
// @Router /show_segments [get]
func ShowUserSegments(c *fiber.Ctx) error {
	// Вывод списка активных сегментов у пользователя
	user := new(domain.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	infrastructure.DB.Db.Preload("Segments").Find(&user, "id = ?", user.ID)
	userSegments := new(UserSegments)
	userSegments.UserID = user.ID
	for i := 0; i < len(user.Segments); i++ {
		userSegments.Segments = append(userSegments.Segments, user.Segments[i].Name)
	}
	return c.Status(200).JSON(userSegments)
}

// DeleteSegments func delete user's segments.
// @Description Удаление сегментов у пользователя.
// @Accept json
// @Param input body UserSegments true "Пользовательские сегменты"
// @Success 200
// @Router /delete_user_from_segment [post]
func DeleteSegments(c *fiber.Ctx) error {
	// Удаление сегмента у пользователя
	userSegments := new(UserSegments)
	if err := c.BodyParser(&userSegments); err != nil {
		return err
	}
	user := new(domain.User)
	user_db := infrastructure.DB.Db.Where("ID = ?", userSegments.UserID).First(&user)
	if user_db.RowsAffected < 1 {
		return fmt.Errorf("user with id=%d doesn't exist", userSegments.UserID)
	} else if user_db.Error != nil {
		return user_db.Error
	}
	for i := 0; i < len(userSegments.Segments); i++ {
		for i := 0; i < len(userSegments.Segments); i++ {
			segment := new(domain.Segment)

			segment_name := userSegments.Segments[i]
			db := infrastructure.DB.Db.Where("Name = ?", segment_name).First(&segment)
			if db.RowsAffected < 1 {
				return fmt.Errorf("segment %s cannot be added to user because it doesn't exist", segment_name)
			} else if db.Error != nil {
				return db.Error
			}
			userSegment := []domain.UserSegment{{UserID: user.ID, SegmentID: segment.ID}}
			infrastructure.DB.Db.Delete(&userSegment)
		}
	}
	return c.Status(200).JSON("OK")
}
