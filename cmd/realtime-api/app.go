package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	messageHttpHandlerV1 "github.com/naufalridho/realtime-api/handler/v1/message/http"
	messageWebSocketHandlerV1 "github.com/naufalridho/realtime-api/handler/v1/message/websocket"
	messageInMemoryRepository "github.com/naufalridho/realtime-api/repository/message/inmemory"
	messageDefaultService "github.com/naufalridho/realtime-api/service/message/default"
)

func main() {
	mr := messageInMemoryRepository.New()
	ms := messageDefaultService.New(mr)
	mh := messageHttpHandlerV1.New(ms)
	mws := messageWebSocketHandlerV1.New()

	r := httprouter.New()
	r.POST("/v1/message/send", mh.SendMessage)
	r.GET("/v1/message/get", mh.GetAllMessages)
	r.GET("/v1/message/realtime", mws.ReadMessageWebSocket)
	go mws.WriteMessageWebSocket()

	log.Println("service is running...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
