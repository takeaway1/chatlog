//go:build !windows

package windows

import (
	"github.com/shirou/gopsutil/v4/process"
	"github.com/takeaway1/chatlog/internal/wechat/model"
)

func initializeProcessInfo(p *process.Process, info *model.Process) error {
	return nil
}
