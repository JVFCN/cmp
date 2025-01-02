package Methods

import (
	"regexp"
	"strconv"
)

func DurationToSeconds(duration string) int {
	// 定义正则表达式来匹配时长格式
	re := regexp.MustCompile(`Duration: (\d+):(\d+):(\d+)\.(\d+)`)
	matches := re.FindStringSubmatch(duration)
	if matches == nil {
		return 0
	}

	// 提取时、分、秒、毫秒
	hours, _ := strconv.Atoi(matches[1])
	minutes, _ := strconv.Atoi(matches[2])
	seconds, _ := strconv.Atoi(matches[3])

	// 将时、分、秒、毫秒转换为总秒数
	totalSeconds := hours*3600 + minutes*60 + seconds
	return totalSeconds
}
