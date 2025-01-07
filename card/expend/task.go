package expend

import (
	"context"
	"time"
)

type ICalcTask interface {
	StartTask()
	doTask()
	startClear()
}

// 启动定时任务
func (t *extend) StartTimer() {
	go t.doTask()
}
func (t *extend) doTask() {

	err := t.startClear()
	if err != nil {
		// TODO add log
	}
	// 忽视错误，添加下一个定时任务

	// 获得下一个零点时间
	now := time.Now()
	todayMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 下一个零点时间执行自己
	time.AfterFunc(todayMidnight.Sub(now), func() {
		t.doTask()
	})
}

// 开始清算
func (t *extend) startClear() error {
	ctx := context.Background()
	ormClient := t.getClient()
	cacheMap := t.getCache()

	cacheMap.Range(func(key, value any) bool {
		cardCode, _ := key.(string)
		balance, _ := value.(int64)
		clearCardExpend(ctx, ormClient, cardCode, balance)
		return true
	})

	// 清空缓存
	cacheMap.Range(func(key, value any) bool {
		cacheMap.Delete(key)
		return true
	})

	return nil
}
