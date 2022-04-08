package game

type ModIcon struct {
}

//判断Icon是否存在，把非法请求拦截在外

func (self *ModIcon) IsHasIcon(icon int) bool {
	return true //暂时没有数据，先这样处理
}
