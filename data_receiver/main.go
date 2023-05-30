package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/keselj-strahinja/toll-calculator/types"
)

func main() {
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
}

type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
	prod  DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {

	p, err := NewKafkaProducer()
	if err != nil {
		return nil, err
	}
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
		prod:  p,
	}, nil
}

func (dr *DataReceiver) producData(data types.OBUData) error {
	return dr.prod.ProduceData(data)
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}

	conn, err := u.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	dr.conn = conn

	go dr.wsReceiveLoop()

}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("new OBU client connected !")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error: ", err)
			continue
		}
		// fmt.Printf("receuved OBU data from [%d] :: <lat%.2f, long %.2f> \n", data.OBUID, data.Lat, data.Lon)
		if err := dr.producData(data); err != nil {
			fmt.Println("kafka produce error", err)
		}

	}
}
