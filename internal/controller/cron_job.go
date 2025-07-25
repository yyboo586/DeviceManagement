package controller

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/service"
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type cronJobController struct{}

var CronJobController = &cronJobController{}

func (c *cronJobController) Add(ctx context.Context, req *v1.AddCronJobReq) (res *v1.AddCronJobRes, err error) {
	// 解析时间配置并转换为cron表达式
	cronExpression, err := c.parseTimeConfigToCron(req.TimeConfig)
	if err != nil {
		return nil, fmt.Errorf("时间配置解析失败: %w", err)
	}

	// 验证cron表达式的合法性
	if err := c.validateCronExpression(cronExpression); err != nil {
		return nil, fmt.Errorf("cron表达式验证失败: %w", err)
	}

	// 生成人类可读的任务执行时间描述
	description := c.generateCronDescription(cronExpression)

	id, err := service.CronJob().Add(ctx, req, cronExpression, description)
	if err != nil {
		return nil, err
	}

	res = &v1.AddCronJobRes{
		ID:             id,
		CronExpression: cronExpression,
		Description:    description,
	}
	return
}

func (c *cronJobController) Delete(ctx context.Context, req *v1.DeleteCronJobReq) (res *v1.DeleteCronJobRes, err error) {
	err = service.CronJob().Delete(ctx, req.ID)
	if err != nil {
		return
	}
	return
}

func (c *cronJobController) Edit(ctx context.Context, req *v1.EditCronJobReq) (res *v1.EditCronJobRes, err error) {
	// 解析时间配置并转换为cron表达式
	cronExpression, err := c.parseTimeConfigToCron(req.TimeConfig)
	if err != nil {
		return nil, fmt.Errorf("时间配置解析失败: %w", err)
	}
	// 验证cron表达式的合法性
	if err := c.validateCronExpression(cronExpression); err != nil {
		return nil, fmt.Errorf("cron表达式验证失败: %w", err)
	}

	// 生成人类可读的任务执行时间描述
	description := c.generateCronDescription(cronExpression)

	err = service.CronJob().Edit(ctx, req, cronExpression, description)
	if err != nil {
		return
	}
	return
}

func (c *cronJobController) Get(ctx context.Context, req *v1.GetCronJobReq) (res *v1.GetCronJobRes, err error) {
	out, err := service.CronJob().Get(ctx, req.ID)
	if err != nil {
		return
	}

	res = &v1.GetCronJobRes{
		CronJob: &v1.CronJob{
			ID:             out.ID,
			Name:           out.Name,
			Params:         out.Params,
			InvokeType:     out.InvokeType,
			CronExpression: out.CronExpression,
			Description:    out.Description,
			Enabled:        out.Enabled == model.CronJobStatusEnabled,
			LastExecuteAt:  out.LastExecuteAt,
			ExecuteCount:   out.ExecuteCount,
			SuccessCount:   out.SuccessCount,
			FailedCount:    out.FailedCount,
			Remark:         out.Remark,
		},
	}
	return
}

func (c *cronJobController) GetList(ctx context.Context, req *v1.GetCronJobListReq) (res *v1.GetCronJobListRes, err error) {
	out, total, err := service.CronJob().GetList(ctx, req)
	if err != nil {
		return
	}

	res = &v1.GetCronJobListRes{
		PageRes: model.PageRes{
			Total: total,
		},
	}
	for _, item := range out {
		res.List = append(res.List, &v1.CronJob{
			ID:             item.ID,
			Name:           item.Name,
			Params:         item.Params,
			InvokeType:     item.InvokeType,
			CronExpression: item.CronExpression,
			Description:    item.Description,
			Enabled:        item.Enabled == model.CronJobStatusEnabled,
			LastExecuteAt:  item.LastExecuteAt,
			ExecuteCount:   item.ExecuteCount,
			SuccessCount:   item.SuccessCount,
			FailedCount:    item.FailedCount,
			Remark:         item.Remark,
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
		})
	}

	return
}

func (c *cronJobController) ExecuteJob(ctx context.Context, req *v1.ExecuteJobReq) (res *v1.ExecuteJobRes, err error) {
	err = service.CronJob().ExecuteJob(ctx, req.JobID)
	if err != nil {
		return
	}

	return &v1.ExecuteJobRes{}, nil
}

// parseTimeConfigToCron 将时间配置解析为cron表达式
func (c *cronJobController) parseTimeConfigToCron(timeConfig *v1.TaskTimeConfig) (string, error) {
	if timeConfig == nil {
		return "", fmt.Errorf("时间配置不能为空")
	}

	second, err := c.parseTimeUnit(timeConfig.Second, 0, 59, "秒")
	if err != nil {
		return "", err
	}

	// 解析各个时间单位
	minute, err := c.parseTimeUnit(timeConfig.Minute, 0, 59, "分钟")
	if err != nil {
		return "", err
	}

	hour, err := c.parseTimeUnit(timeConfig.Hour, 0, 23, "小时")
	if err != nil {
		return "", err
	}

	day, err := c.parseTimeUnit(timeConfig.Day, 1, 31, "日")
	if err != nil {
		return "", err
	}

	month := "*"

	week, err := c.parseTimeUnit(timeConfig.Week, 0, 6, "周")
	if err != nil {
		return "", err
	}

	// 组合成cron表达式：分钟 小时 日 月 周
	cronExpression := fmt.Sprintf("%s %s %s %s %s %s", second, minute, hour, day, month, week)
	return cronExpression, nil
}

// parseTimeUnit 解析单个时间单位
func (c *cronJobController) parseTimeUnit(unit *v1.TimeUnitConfig, min, max int, unitName string) (string, error) {
	if unit == nil {
		return "*", nil // 如果没有配置，默认为所有值
	}

	switch unit.Type {
	case "all":
		return "*", nil

	case "specific":
		if len(unit.SpecificValues) == 0 {
			return "", fmt.Errorf("%s配置错误: 特定值列表不能为空", unitName)
		}
		// 验证值范围
		for _, val := range unit.SpecificValues {
			if val < min || val > max {
				return "", fmt.Errorf("%s配置错误: 值%d超出范围[%d,%d]", unitName, val, min, max)
			}
		}
		// 排序并转换为字符串
		sort.Ints(unit.SpecificValues)
		values := make([]string, len(unit.SpecificValues))
		for i, val := range unit.SpecificValues {
			values[i] = strconv.Itoa(val)
		}
		return strings.Join(values, ","), nil

	case "range":
		if unit.Start < min || unit.Start > max || unit.End < min || unit.End > max {
			return "", fmt.Errorf("%s配置错误: 范围值超出限制[%d,%d]", unitName, min, max)
		}
		if unit.Start > unit.End {
			return "", fmt.Errorf("%s配置错误: 起始值不能大于结束值", unitName)
		}
		return fmt.Sprintf("%d-%d", unit.Start, unit.End), nil

	case "interval":
		if unit.Step <= 0 {
			return "", fmt.Errorf("%s配置错误: 步长必须大于0", unitName)
		}
		if unit.Start < min || unit.Start > max {
			return "", fmt.Errorf("%s配置错误: 起始值超出范围[%d,%d]", unitName, min, max)
		}
		if unit.End > 0 && (unit.End < min || unit.End > max) {
			return "", fmt.Errorf("%s配置错误: 结束值超出范围[%d,%d]", unitName, min, max)
		}
		if unit.End > 0 && unit.Start > unit.End {
			return "", fmt.Errorf("%s配置错误: 起始值不能大于结束值", unitName)
		}

		if unit.Start == min {
			// 从最小值开始的间隔，如 */5
			return fmt.Sprintf("*/%d", unit.Step), nil
		} else if unit.End > 0 {
			// 指定范围的间隔，如 1-10/2
			return fmt.Sprintf("%d-%d/%d", unit.Start, unit.End, unit.Step), nil
		} else {
			// 从指定值开始的间隔，如 1/5
			return fmt.Sprintf("%d/%d", unit.Start, unit.Step), nil
		}
	default:
		return "", fmt.Errorf("%s配置错误: 不支持的类型%s", unitName, unit.Type)
	}
}

// validateCronExpression 验证cron表达式的合法性
func (c *cronJobController) validateCronExpression(cronExpression string) error {
	if cronExpression == "" {
		return fmt.Errorf("cron表达式不能为空")
	}

	parts := strings.Fields(cronExpression)
	if len(parts) != 5 {
		return fmt.Errorf("cron表达式格式错误: 必须包含5个字段(分钟 小时 日 月 周)")
	}

	// 验证各个字段
	fields := []struct {
		name string
		min  int
		max  int
	}{
		{"分钟", 0, 59},
		{"小时", 0, 23},
		{"日", 1, 31},
		{"月", 1, 12},
		{"周", 0, 6},
	}

	for i, field := range fields {
		if err := c.validateCronField(parts[i], field.name, field.min, field.max); err != nil {
			return err
		}
	}

	return nil
}

// validateCronField 验证单个cron字段
func (c *cronJobController) validateCronField(field, fieldName string, min, max int) error {
	if field == "*" {
		return nil
	}

	// 处理包含逗号的字段，如 1,3,5
	if strings.Contains(field, ",") {
		parts := strings.Split(field, ",")
		for _, part := range parts {
			if err := c.validateCronField(part, fieldName, min, max); err != nil {
				return err
			}
		}
		return nil
	}

	// 处理间隔，如 */5
	if strings.HasPrefix(field, "*/") {
		stepStr := strings.TrimPrefix(field, "*/")
		step, err := strconv.Atoi(stepStr)
		if err != nil || step <= 0 {
			return fmt.Errorf("%s字段错误: 无效的间隔值%s", fieldName, stepStr)
		}
		return nil
	}

	// 处理间隔范围，如 1-10/2
	if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		if len(parts) != 2 {
			return fmt.Errorf("%s字段错误: 无效的间隔格式%s", fieldName, field)
		}

		rangePart := parts[0]
		stepStr := parts[1]

		// 验证步长
		step, err := strconv.Atoi(stepStr)
		if err != nil || step <= 0 {
			return fmt.Errorf("%s字段错误: 无效的步长值%s", fieldName, stepStr)
		}

		// 验证范围部分
		if rangePart == "*" {
			return nil
		}

		if strings.Contains(rangePart, "-") {
			rangeParts := strings.Split(rangePart, "-")
			if len(rangeParts) != 2 {
				return fmt.Errorf("%s字段错误: 无效的范围格式%s", fieldName, rangePart)
			}

			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return fmt.Errorf("%s字段错误: 起始值无效%s", fieldName, rangeParts[0])
			}

			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return fmt.Errorf("%s字段错误: 结束值无效%s", fieldName, rangeParts[1])
			}

			if start < min || start > max || end < min || end > max {
				return fmt.Errorf("%s字段错误: 范围值超出限制[%d,%d]", fieldName, min, max)
			}

			if start > end {
				return fmt.Errorf("%s字段错误: 起始值不能大于结束值", fieldName)
			}
		} else {
			// 单个值
			val, err := strconv.Atoi(rangePart)
			if err != nil {
				return fmt.Errorf("%s字段错误: 值无效%s", fieldName, rangePart)
			}
			if val < min || val > max {
				return fmt.Errorf("%s字段错误: 值超出范围[%d,%d]", fieldName, min, max)
			}
		}

		return nil
	}

	// 处理范围，如 1-5
	if strings.Contains(field, "-") {
		parts := strings.Split(field, "-")
		if len(parts) != 2 {
			return fmt.Errorf("%s字段错误: 无效的范围格式%s", fieldName, field)
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("%s字段错误: 起始值无效%s", fieldName, parts[0])
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("%s字段错误: 结束值无效%s", fieldName, parts[1])
		}

		if start < min || start > max || end < min || end > max {
			return fmt.Errorf("%s字段错误: 范围值超出限制[%d,%d]", fieldName, min, max)
		}

		if start > end {
			return fmt.Errorf("%s字段错误: 起始值不能大于结束值", fieldName)
		}

		return nil
	}

	// 处理单个值
	val, err := strconv.Atoi(field)
	if err != nil {
		return fmt.Errorf("%s字段错误: 值无效%s", fieldName, field)
	}

	if val < min || val > max {
		return fmt.Errorf("%s字段错误: 值超出范围[%d,%d]", fieldName, min, max)
	}

	return nil
}

// generateCronDescription 根据cron表达式生成人类可读的描述
func (c *cronJobController) generateCronDescription(cronExpression string) string {
	parts := strings.Fields(cronExpression)
	if len(parts) != 5 {
		return "无效的时间配置"
	}

	minute, hour, day, month, week := parts[0], parts[1], parts[2], parts[3], parts[4]

	description := "任务将在"

	// 解析分钟
	minuteDesc := c.parseMinuteDescription(minute)
	description += minuteDesc

	// 解析小时
	hourDesc := c.parseHourDescription(hour)
	description += hourDesc

	// 解析日期（只有当小时不是*时才添加）
	if day != "*" {
		dayDesc := c.parseDayDescription(day)
		description += dayDesc
	}

	// 解析月份（只有当日期不是*时才添加）
	if month != "*" {
		monthDesc := c.parseMonthDescription(month)
		description += monthDesc
	}

	// 解析星期（只有当星期不是*时才添加）
	if week != "*" {
		weekDesc := c.parseWeekDescription(week)
		description += weekDesc
	}

	description += "执行"

	return description
}

// parseMinuteDescription 解析分钟描述
func (c *cronJobController) parseMinuteDescription(minute string) string {
	if minute == "*" {
		return "每分钟的"
	}

	if strings.HasPrefix(minute, "*/") {
		step := strings.TrimPrefix(minute, "*/")
		if step == "1" {
			return "每分钟的"
		}
		return fmt.Sprintf("每%s分钟的", step)
	}

	if strings.Contains(minute, ",") {
		parts := strings.Split(minute, ",")
		if len(parts) == 2 {
			return fmt.Sprintf("第%s分钟和第%s分钟的", parts[0], parts[1])
		}
		return fmt.Sprintf("第%s分钟的", minute)
	}

	if strings.Contains(minute, "/") {
		rangeParts := strings.Split(minute, "/")
		if strings.Contains(rangeParts[0], "-") {
			rangeRange := strings.Split(rangeParts[0], "-")
			return fmt.Sprintf("第%s到第%s分钟每隔%s分钟的", rangeRange[0], rangeRange[1], rangeParts[1])
		}
		return fmt.Sprintf("从第%s分钟开始每隔%s分钟的", rangeParts[0], rangeParts[1])
	}

	if strings.Contains(minute, "-") {
		parts := strings.Split(minute, "-")
		return fmt.Sprintf("第%s到第%s分钟的", parts[0], parts[1])
	}

	return fmt.Sprintf("第%s分钟的", minute)
}

// parseHourDescription 解析小时描述
func (c *cronJobController) parseHourDescription(hour string) string {
	if hour == "*" {
		return "每小时"
	}

	if strings.HasPrefix(hour, "*/") {
		step := strings.TrimPrefix(hour, "*/")
		if step == "1" {
			return "每小时"
		}
		return fmt.Sprintf("每%s小时", step)
	}

	if strings.Contains(hour, ",") {
		parts := strings.Split(hour, ",")
		if len(parts) == 2 {
			return fmt.Sprintf("第%s小时和第%s小时", parts[0], parts[1])
		}
		return fmt.Sprintf("第%s小时", hour)
	}

	if strings.Contains(hour, "/") {
		rangeParts := strings.Split(hour, "/")
		if strings.Contains(rangeParts[0], "-") {
			rangeRange := strings.Split(rangeParts[0], "-")
			return fmt.Sprintf("第%s到第%s小时每隔%s小时", rangeRange[0], rangeRange[1], rangeParts[1])
		}
		return fmt.Sprintf("从第%s小时开始每隔%s小时", rangeParts[0], rangeParts[1])
	}

	if strings.Contains(hour, "-") {
		parts := strings.Split(hour, "-")
		return fmt.Sprintf("第%s到第%s小时", parts[0], parts[1])
	}

	return fmt.Sprintf("第%s小时", hour)
}

// parseDayDescription 解析日期描述
func (c *cronJobController) parseDayDescription(day string) string {
	if day == "*" {
		return "的每一天"
	}

	if strings.HasPrefix(day, "*/") {
		step := strings.TrimPrefix(day, "*/")
		if step == "1" {
			return "的每一天"
		}
		return fmt.Sprintf("每隔%s天", step)
	}

	if strings.Contains(day, ",") {
		parts := strings.Split(day, ",")
		if len(parts) == 2 {
			return fmt.Sprintf("的第%s天和第%s天", parts[0], parts[1])
		}
		return fmt.Sprintf("的第%s天", day)
	}

	if strings.Contains(day, "/") {
		rangeParts := strings.Split(day, "/")
		if strings.Contains(rangeParts[0], "-") {
			rangeRange := strings.Split(rangeParts[0], "-")
			return fmt.Sprintf("的第%s到第%s天每隔%s天", rangeRange[0], rangeRange[1], rangeParts[1])
		}
		return fmt.Sprintf("从第%s天开始每隔%s天", rangeParts[0], rangeParts[1])
	}

	if strings.Contains(day, "-") {
		parts := strings.Split(day, "-")
		return fmt.Sprintf("的第%s到第%s天", parts[0], parts[1])
	}

	return fmt.Sprintf("的第%s天", day)
}

// parseMonthDescription 解析月份描述
func (c *cronJobController) parseMonthDescription(month string) string {
	if month == "*" {
		return "的每个月"
	}

	if strings.HasPrefix(month, "*/") {
		step := strings.TrimPrefix(month, "*/")
		if step == "1" {
			return "的每个月"
		}
		return fmt.Sprintf("每隔%s个月", step)
	}

	if strings.Contains(month, ",") {
		parts := strings.Split(month, ",")
		if len(parts) == 2 {
			return fmt.Sprintf("的第%s月和第%s月", parts[0], parts[1])
		}
		return fmt.Sprintf("的第%s月", month)
	}

	if strings.Contains(month, "/") {
		rangeParts := strings.Split(month, "/")
		if strings.Contains(rangeParts[0], "-") {
			rangeRange := strings.Split(rangeParts[0], "-")
			return fmt.Sprintf("的第%s到第%s月每隔%s个月", rangeRange[0], rangeRange[1], rangeParts[1])
		}
		return fmt.Sprintf("从第%s月开始每隔%s个月", rangeParts[0], rangeParts[1])
	}

	if strings.Contains(month, "-") {
		parts := strings.Split(month, "-")
		return fmt.Sprintf("的第%s到第%s月", parts[0], parts[1])
	}

	return fmt.Sprintf("的第%s月", month)
}

// parseWeekDescription 解析星期描述
func (c *cronJobController) parseWeekDescription(week string) string {
	if week == "*" {
		return ""
	}

	weekNames := map[string]string{
		"0": "周日", "1": "周一", "2": "周二", "3": "周三",
		"4": "周四", "5": "周五", "6": "周六",
	}

	if strings.HasPrefix(week, "*/") {
		step := strings.TrimPrefix(week, "*/")
		if step == "1" {
			return "的每周"
		}
		return fmt.Sprintf("每隔%s周", step)
	}

	if strings.Contains(week, ",") {
		parts := strings.Split(week, ",")
		weekDescs := make([]string, 0, len(parts))
		for _, part := range parts {
			if weekName, exists := weekNames[part]; exists {
				weekDescs = append(weekDescs, weekName)
			} else {
				weekDescs = append(weekDescs, fmt.Sprintf("第%s天", part))
			}
		}
		if len(weekDescs) == 2 {
			return fmt.Sprintf("的%s和%s", weekDescs[0], weekDescs[1])
		}
		return fmt.Sprintf("的%s", strings.Join(weekDescs, "、"))
	}

	if strings.Contains(week, "-") {
		parts := strings.Split(week, "-")
		startWeek := weekNames[parts[0]]
		endWeek := weekNames[parts[1]]
		if startWeek != "" && endWeek != "" {
			return fmt.Sprintf("的%s到%s", startWeek, endWeek)
		}
		return fmt.Sprintf("的第%s到第%s天", parts[0], parts[1])
	}

	if strings.Contains(week, "/") {
		rangeParts := strings.Split(week, "/")
		if strings.Contains(rangeParts[0], "-") {
			rangeRange := strings.Split(rangeParts[0], "-")
			startWeek := weekNames[rangeRange[0]]
			endWeek := weekNames[rangeRange[1]]
			if startWeek != "" && endWeek != "" {
				return fmt.Sprintf("的%s到%s每隔%s周", startWeek, endWeek, rangeParts[1])
			}
			return fmt.Sprintf("的第%s到第%s天每隔%s周", rangeRange[0], rangeRange[1], rangeParts[1])
		}
		startWeek := weekNames[rangeParts[0]]
		if startWeek != "" {
			return fmt.Sprintf("从%s开始每隔%s周", startWeek, rangeParts[1])
		}
		return fmt.Sprintf("从第%s天开始每隔%s周", rangeParts[0], rangeParts[1])
	}

	if weekName, exists := weekNames[week]; exists {
		return fmt.Sprintf("的%s", weekName)
	}

	return fmt.Sprintf("的第%s天", week)
}
