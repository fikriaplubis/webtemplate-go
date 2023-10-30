package constant

import "os"

var path = map[string]string{
	"DirectTransferOri": "FolderDirectransferOri",
}

func GetPath(key string) string {
	return os.Getenv(path[key])
}
