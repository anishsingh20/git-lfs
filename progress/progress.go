// Package progress provides common progress monitoring / display features
// NOTE: Subject to change, do not rely on this package from outside git-lfs source
package progress

import "github.com/git-lfs/git-lfs/git/githistory/log"

type Meter interface {
	log.Task

	Start()
	Pause()
	Add(int64)
	Skip(size int64)
	StartTransfer(name string)
	TransferBytes(direction, name string, read, total int64, current int)
	FinishTransfer(name string)
	Finish()
}
