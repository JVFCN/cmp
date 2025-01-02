package Methods

import "fmt"

// FormatTime 将秒数转换为分:秒格式
func FormatTime(seconds int) string {
	minutes := seconds / 60
	seconds = seconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
