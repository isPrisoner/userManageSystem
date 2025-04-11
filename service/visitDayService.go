package service

import (
	"strconv"
	"time"
	"userManageSystem/models"
	"userManageSystem/utils"
)

// VisitData 生成模拟数据
type VisitData struct {
	Days   []string `json:"days"`
	Counts []int    `json:"counts"`
}

func GenerateMockData(date string) *utils.Message {
	days, _ := strconv.Atoi(date)

	// 创建日期标签
	labels := make([]string, days)
	// 创建访问数量数据
	counts := make([]int, days)

	// 今天是截至时间
	today := time.Now()

	for i := 0; i < days; i++ {
		// 计算日期
		day := today.AddDate(0, 0, -i)
		// 从后向前存储
		labels[days-i-1] = day.Format("2006-01-02")
		res := models.QueryTime(day.Format("2006-01-02"))
		var count int
		for res.Next() {
			res.Scan(&count)
		}
		counts[days-i-1] = count
	}
	return utils.Ok(VisitData{
		Days:   labels,
		Counts: counts,
	}, "获取成功")
}
