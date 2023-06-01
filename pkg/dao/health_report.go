package dao

import (
	"errors"
	"github.com/tanlay/cleanner-mysql-data/pkg/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HealthReportDao struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewHealthReportDao(db *gorm.DB) *HealthReportDao {
	return &HealthReportDao{
		db:     db,
		logger: zap.L().Named("health_report_dao"),
	}
}

// GetHealthReportCount 查询数量
func (d *HealthReportDao) GetHealthReportCount(ts int64) (int, error) {
	var count int
	err := d.db.Select("count(*)").Where("report_time < ?", ts).
		Table("health_report").Find(&count).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			d.logger.Error("get health_report count err ", zap.Error(err))
			return 0, err
		}
	}
	return count, nil
}

// GetHealthReportIds 获取id
func (d *HealthReportDao) GetHealthReportIds(ts int64, limit, offset int) ([]int64, error) {
	ids := make([]int64, 0)
	err := d.db.Select("id").Where("report_time < ?", ts).
		Table("health_report").Limit(limit).Offset(offset).Find(&ids).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			d.logger.Error("get health_report ids err ", zap.Error(err))
			return nil, err
		}
	}
	return ids, nil
}

// DeleteHealthReportById 通过id删除
func (d *HealthReportDao) DeleteHealthReportById(ids []int64) error {
	err := d.db.Where("id in ?", ids).Delete(model.HealthReport{}).Error
	if err != nil {
		d.logger.Error("delete health_report err ", zap.Error(err))
		return err
	}
	return nil
}
