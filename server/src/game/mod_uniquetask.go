package game

import "sync"

// 唯一任务模块
// 对于玩家生涯来说，只能做一次

// TaskInfo 任务属性
type TaskInfo struct {
	TaskId int
	State  int
}

type ModUniqueTask struct {
	//保存任务的信息用map
	MyTaskInfo map[int]*TaskInfo
	Locker     *sync.RWMutex
}

func (self *ModUniqueTask) IsTaskFinish(taskId int) bool {
	if taskId == 10001 || taskId == 10002 {
		return true
	}

	task, ok := self.MyTaskInfo[taskId]
	if !ok {
		return false
	}
	return task.State == TaskStateFinish
}

// TaskStateInit 突破任务是无法联机的,用宏定义记录状态
const (
	TaskStateInit   = 0 //初始状态
	TaskStateDoing  = 1 //表明正在做的任务	，如果玩家进入这个副本中了，把这个任务视为进行中
	TaskStateFinish = 2 //完成之后置为2
)
