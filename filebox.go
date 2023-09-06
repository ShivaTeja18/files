package swaggerFiles

import (
	"github.com/ShivaTeja18/files/webdav"
)

func NewHandler() *webdav.Handler {
	return &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}
}
