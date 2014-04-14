package main

import (
    "os"
    "log"
    "net/http"

    "go-socket.io"
)

func main() {
    log.Println("log")
    sio := socketio.NewSocketIOServer(&socketio.Config{})
    sio.On("hi", func (ns *socketio.NameSpace) {
        log.Println("got hi")
        ns.Emit("hi", "test")
    })
    log.Println(os.Getenv("ZUUL_PORT"))
    http.ListenAndServe(":" + os.Getenv("ZUUL_PORT"), sio)
}
