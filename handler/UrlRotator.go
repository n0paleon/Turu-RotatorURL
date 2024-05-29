package handler

import (
	"TuruSMM/internal/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type URLHandler struct {
	DB     *gorm.DB
	Log    *logrus.Logger
	Config *viper.Viper
}

func NewURLHandler(db *gorm.DB, log *logrus.Logger, config *viper.Viper) *URLHandler {
	return &URLHandler{
		DB:     db,
		Log:    log,
		Config: config,
	}
}

func (h *URLHandler) Rotator(ctx *fiber.Ctx) error {
	h.Log.Warnf("ini error")

	tx := h.DB.WithContext(ctx.UserContext()).Begin()
	defer tx.Rollback()

	var url entity.UrlList
	err := tx.Where("total_hit < target_hit AND hit_today < hit_per_day AND rotate = ?", true).
		Order("updated_at ASC").
		First(&url).Error
	if err != nil {
		return ctx.JSON(&fiber.Map{
			"message": "all task is done!",
		})
	}

	url.HitToday += 1
	url.TotalHit += 1

	tx.Save(&url)

	if err := tx.Commit().Error; err != nil {
		return ctx.JSON(&fiber.Map{
			"message": "error while commit transaction",
			"error":   err.Error(),
		})
	}

	return ctx.Redirect(url.TargetURL)
}

func (h *URLHandler) RotateByID(ctx *fiber.Ctx) error {
	tx := h.DB.WithContext(ctx.UserContext()).Begin()
	defer tx.Rollback()

	var url entity.UrlList
	err := tx.Where("total_hit < target_hit AND hit_today < hit_per_day AND order_id = ?", ctx.Params("id")).
		Order("updated_at ASC").
		First(&url).Error
	if err != nil {
		return ctx.JSON(&fiber.Map{
			"message": "the job is done!",
		})
	}

	url.HitToday += 1
	url.TotalHit += 1

	tx.Save(&url)

	if err := tx.Commit().Error; err != nil {
		return ctx.JSON(&fiber.Map{
			"message": "error while commit transaction",
			"error":   err.Error(),
		})
	}

	return ctx.Redirect(url.TargetURL)
}

func (h *URLHandler) CustomRotate(ctx *fiber.Ctx) error {
	tx := h.DB.WithContext(ctx.UserContext()).Begin()
	defer tx.Rollback()

	var url entity.RotatedUrl
	err := tx.Where("rotate_id = ?", ctx.Params("id")).
		Order("updated_at ASC").
		First(&url).Error
	if err != nil {
		return ctx.JSON(&fiber.Map{
			"message": "url rotate_id not found!",
		})
	}

	url.Hit += 1

	tx.Save(&url)

	if err := tx.Commit().Error; err != nil {
		return ctx.JSON(&fiber.Map{
			"message": "error while commit transaction",
			"error":   err.Error(),
		})
	}

	return ctx.Redirect(url.TargetURL)
}
