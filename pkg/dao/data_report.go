package dao

import (
	"errors"
	"github.com/tanlay/cleanner-mysql-data/pkg/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DataReportDao struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewDataReportDao(db *gorm.DB) *DataReportDao {
	return &DataReportDao{
		db:     db,
		logger: zap.L().Named("data_report_dao"),
	}
}

// GetDataReportCount 查询数量
func (d *DataReportDao) GetDataReportCount(ts int64) (int, error) {
	var count int
	err := d.db.Select("count(*)").Where("timestamp < ?", ts).
		Table("data_report").Find(&count).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			d.logger.Error("get data_report count err ", zap.Error(err))
			return 0, err
		}
	}
	return count, nil
}

// GetDataReportIds 获取id
func (d *DataReportDao) GetDataReportIds(ts int64, limit, offset int) ([]int64, error) {
	ids := make([]int64, 0)
	err := d.db.Select("id").Where("timestamp < ?", ts).
		Table("data_report").Limit(limit).Offset(offset).Find(&ids).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			d.logger.Error("get data_report ids err ", zap.Error(err))
			return nil, err
		}
	}
	return ids, nil
}

// DeleteDataReportById 通过id删除
func (d *DataReportDao) DeleteDataReportById(ids []int64) error {
	err := d.db.Where("id in ?", ids).Delete(model.DataReport{}).Error
	if err != nil {
		d.logger.Error("delete data_report err ", zap.Error(err))
		return err
	}
	return nil
}
