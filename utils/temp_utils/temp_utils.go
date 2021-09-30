package temp_utils

import (
	"os"
	"path/filepath"

	"github.com/sikalabs/tergum/utils/rand_utils"
)

func GetTempFileName() string {
	return filepath.Join(
		os.TempDir(),
		"tergum-tmp-"+rand_utils.GetRandString(10),
	)
}
