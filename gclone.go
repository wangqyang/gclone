package main

import (
	"github.com/rclone/rclone/cmd"
	_ "github.com/rclone/rclone/cmd/all" // import all commands
	"github.com/rclone/rclone/fs"
	_ "github.com/rclone/rclone/lib/plugin"     // import plugins
	_ "github.com/wangqyang/gclone/backend/all" // import all backends
	_ "github.com/wangqyang/gclone/cmd/copy"
)

func main() {
	fs.Version = fs.Version+"-mod"
	cmd.Main()
}
