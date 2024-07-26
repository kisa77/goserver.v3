package task

import (
	"github.com/kisa77/goserver.v3/core/basic"
	"github.com/kisa77/goserver.v3/core/utils"
)

type taskResCommand struct {
	t *Task
}

func (trc *taskResCommand) Done(o *basic.Object) error {
	defer o.ProcessSeqnum()
	defer utils.DumpStackIfPanic("taskExeCommand")
	trc.t.n.Done(<-trc.t.r, trc.t)
	return nil
}

func SendTaskRes(o *basic.Object, t *Task) bool {
	if o == nil {
		return false
	}
	return o.SendCommand(&taskResCommand{t: t}, true)
}
