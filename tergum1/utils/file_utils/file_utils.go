package file_utils

import (
	"time"

	"github.com/sikalabs/tergum/tergum1/utils/rand_utils"
)

func GetFileName(prefix string, suffix string) string {
	return prefix + "." + time.Now().UTC().Format("2006-01-02_15-04-05") + "_" + rand_utils.GetRandString(3) + "." + suffix
}
