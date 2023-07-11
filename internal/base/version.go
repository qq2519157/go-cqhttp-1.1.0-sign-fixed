package base

import "runtime/debug"

// Version go-cqhttp的版本信息，在编译时使用ldflags进行覆盖
var Version = "1.1.0.fix"

func init() {
	if Version != "unknown" {
		return
	}
	info, ok := debug.ReadBuildInfo()
	if ok {
		Version = info.Main.Version
	}
}
