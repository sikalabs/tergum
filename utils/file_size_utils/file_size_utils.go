package file_size_utils

import "fmt"

func PrettyFileSize(size int64) string {
	switch {
	case size > 1024*1024*1024:
		return fmt.Sprintf("%.1fG", float64(size)/(1024*1024*1024))
	case size > 1024*1024:
		return fmt.Sprintf("%.1fM", float64(size)/(1024*1024))
	case size > 1024:
		return fmt.Sprintf("%.1fK", float64(size)/1024)
	default:
		return fmt.Sprintf("%d", size)
	}
}
