package controller

import (
	"fmt"
	"github.com/panjf2000/ants"
	"github.com/tanlay/cleanner-mysql-data/conf"
	"github.com/tanlay/cleanner-mysql-data/pkg/dao"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
	"time"
)

type HealthReportController struct {
	db      *gorm.DB
	logger  *zap.Logger
	delTime time.Time
}

func NewHealthReportController(db *gorm.DB, delTime time.Time) *HealthReportController {
	return &HealthReportController{
		db:      db,
		logger:  zap.L().Named("health_report_controller"),
		delTime: delTime,
	}
}

func (d *HealthReportController) DeleteHealthReportController(pool *ants.Pool) error {
	drDao := dao.NewHealthReportDao(d.db)
	ts := time.Now().UnixMilli()
	total, err := drDao.GetHealthReportCount(d.delTime.UnixMilli())
	if err != nil {
		d.logger.Error("get health_report count err ", zap.Error(err))
		return err
	}
	d.logger.Info(fmt.Sprintf("health_report需要删除的ids数量：%d, 查询条件：%v(秒级时间戳：%d)",
		total, d.delTime.Format("2006-01-02 15:04:05"), d.delTime.Unix()))
	if total == 0 {
		return nil
	}
	//定义串行每一批的总数量，如：符合条件的数量有2亿条，定义串行每次删除20万条，
	perTotalCount := conf.C.PerTotalCount
	//需要串行执行多少批次
	totalBatch := total / perTotalCount
	if total%perTotalCount != 0 {
		totalBatch = totalBatch + 1
	}
	wg := &sync.WaitGroup{}

	delNum := 0

	for i := 0; i < totalBatch; i++ {
		//串行查询id，查出来后并发删除
		ids, err := drDao.GetHealthReportIds(ts, perTotalCount, 0)
		if err != nil {
			d.logger.Error("get health_report ids err ", zap.Error(err))
			return err
		}
		delNum = delNum + len(ids)
		//并发执行的数量
		batchCount := conf.C.BatchCount
		//每一批串行的数据，需要执行batch次
		batch := len(ids) / batchCount
		if len(ids)%batchCount != 0 {
			batch = batch + 1
		}
		for n := 0; n < batch; n++ {
			delIds := make([]int64, 0)
			if n == batch-1 {
				delIds = ids[n*batchCount:]
			} else {
				delIds = ids[n*batchCount : (n+1)*batchCount]
			}
			wg.Add(1)
			if err := pool.Submit(func() {
				defer wg.Done()
				d.logger.Info(fmt.Sprintf("delete health_report, ids: %v", delIds))
				if err := drDao.DeleteHealthReportById(delIds); err != nil {
					d.logger.Error("delete health_report err ", zap.Error(err))
					return
				}
				//每次提交ids后休眠IntervalTime秒
				time.Sleep(time.Duration(conf.C.IntervalTime) * time.Second)
			}); err != nil {
				d.logger.Error("submit health_report ids err ", zap.Error(err))
				return err
			}
		}
		wg.Wait()
		d.logger.Info(fmt.Sprintf("删除health_report第%d批数据完成，共%d次", i+1, totalBatch))
	}
	d.logger.Info(fmt.Sprintf("health_report删除完成，共删除数据%d条", delNum))
	return nil
}
