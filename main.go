package main

import (
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
)

const (
	// 어플리케이션에서 사용할 세션의 키 정보
	sessionKey    = "simple_cat_session"
	sessionSecret = "simple_chat_session_secret"

	socketBufferSize = 1024
)

var (
	mongoSession *mgo.Session
	renderer     = render.New()
	upgrader     = &websocket.Upgrader{
		ReadBufferSize:  socketBufferSize,
		WriteBufferSize: socketBufferSize,
	}
)

func init() {
	s, err := mgo.Dial("mongodb://192.168.99.100:32768")
	if err != nil {
		panic(err)
	}

	// 세션 스위치. 단조로운 행동으로 ?
	s.SetMode(mgo.Monotonic, true)
	mongoSession = s
}

func main() {
	// 라우터 생성
	router := httprouter.New()

}
