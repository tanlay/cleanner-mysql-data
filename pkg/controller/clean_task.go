package controller

import (
	"context"
	"fmt"
	"github.com/panjf2000/ants"
	"github.com/tanlay/cleanner-mysql-data/conf"
	"github.com/tanlay/cleanner-mysql-data/pkg/lib/gormx"
	"github.com/tanlay/cleanner-mysql-data/pkg/lib/panichandler"
	"go.uber.org/zap"
	"sync"
	"time"
)

type CleanTask struct {
	ctx    context.Context
	logger *zap.Logger
	conf   conf.Conf
}

func NewCleanTask(ctx context.Context, conf conf.Conf) *CleanTask {
	return &CleanTask{
		ctx:    ctx,
		logger: zap.L().Named("controller_tasks"),
		conf:   conf,
	}
}

func (c *CleanTask) CleanDataTask(date string) error {
	c.logger.Info("CleanTask.CleanDataTask")
	var now time.Time
	if date != "" {
		localTime, err := time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
		if err != nil {
			c.logger.Error("日期格式错误")
			return err
		}
		now = localTime
	}
	pool, err := ants.NewPool(conf.C.GoCount, ants.WithMaxBlockingTasks(conf.C.GoCount))
	if err != nil {
		c.logger.Error("new pool err ", zap.Error(err))
		return err
	}
	defer pool.Release()
	wg := &sync.WaitGroup{}
	//判断配置文件如果所有表都未开启，则不进行后续操作
	if !conf.C.Database.DataReportEnable {
		c.logger.Info("未删除任何数据，需开启相应的表配置项")
		return nil
	}

	if conf.C.Database.DataReportEnable {
		wg.Add(1)
		c.logger.Info("data_report表配置项data_report_enable已开启")
		go c.CleanDataReportDataTask(wg, now, pool)
	}
	wg.Wait()
	c.logger.Info("删除完成")
	return nil
}

func (c *CleanTask) CleanDataReportDataTask(wg *sync.WaitGroup, now time.Time, pool *ants.Pool) {
	defer panichandler.ZapHandler(zap.L()).Handle()
	defer wg.Done()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	controller := NewDataReportController(gormx.JudgeryDB, now)
	sinceNow := time.Now()
	if err := controller.DeleteDataReportController(pool); err != nil {
		c.logger.Error("delete data_report err ", zap.Error(err))
		return
	}
	c.logger.Info(fmt.Sprintf("清理data_report耗时%v", time.Since(sinceNow)))
	c.logger.Info("清理data_report完成")
}
