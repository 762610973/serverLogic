package game

import (
	"fmt"
	"math/rand"
	"serverLogic/server/src/csvs"
	"sync"
	"time"
)

type Server struct {
	sync.WaitGroup
}

var server *Server

func GetServer() *Server {
	if server == nil {
		server = new(Server)
	}
	return server
}

func (s *Server) Start() {
	csvs.CheckLoadCsv()
	fmt.Println("数据测试----start")
	rand.Seed(time.Now().UnixNano())
	//需要进行服务器的配置
	// 启动一个违禁词goroutine

	go GetManageBanWord().Run()

	playerGM := NewTestPlayer()
	go playerGM.Run()
	for {

	}
	fmt.Println("服务器已经关闭成功")
}

func (s *Server) Close() {

}
func (s *Server) AddGoroutine() {

}

func (s *Server) SubGoroutine() {

}
