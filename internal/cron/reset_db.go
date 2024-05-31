package cron

import (
	"TuruSMM/internal/entity"
	"context"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type CronJobs struct {
	DB     *gorm.DB
	Log    *logrus.Logger
	Config *viper.Viper
}

func NewCronJobs(db *gorm.DB, log *logrus.Logger, config *viper.Viper) *CronJobs {
	return &CronJobs{
		DB:     db,
		Log:    log,
		Config: config,
	}
}

func (c *CronJobs) Schedule() {
	task := cron.New()
	defer task.Stop()

	task.AddFunc(c.Config.GetString("database.cron"), func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		tx := c.DB.WithContext(ctx).Begin()
		defer tx.Rollback()

		err := tx.Session(&gorm.Session{
			AllowGlobalUpdate: true,
		}).Model(&entity.UrlList{}).Update("hit_today", 0).Error
		if err != nil {
			c.Log.Errorf("failed to run cron job: %+v", err)
		}

		if err = tx.Commit().Error; err != nil {
			c.Log.Errorf("failed to commit transaction: %+v", err)
		} else {
			c.Log.Println("success run cron jobs")
		}
	})

	go task.Start()
}
