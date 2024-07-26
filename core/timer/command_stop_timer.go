package timer

import (
	"container/heap"

	"github.com/kisa77/goserver.v3/core/basic"
)

type stopTimerCommand struct {
	h TimerHandle
}

func (stc *stopTimerCommand) Done(o *basic.Object) error {
	defer o.ProcessSeqnum()

	if v, ok := TimerModule.tq.ref[stc.h]; ok {
		heap.Remove(TimerModule.tq, v)
	}

	return nil
}

func StopTimer(h TimerHandle) bool {
	return TimerModule.SendCommand(&stopTimerCommand{h: h}, true)
}
