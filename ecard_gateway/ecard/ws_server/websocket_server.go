package ws_server

import (
    "github.com/alfred-zhong/wserver"
    "net/http"
)

var WsServer *wserver.Server

func StartWebsocketServer() {
    WsServer = wserver.NewServer(":3456")

    WsServer.WSPath = "/ws"
    WsServer.SendPath = "/send"

    WsServer.AuthToken = func(token string) (userID string, err error) {
       if token == "token" {
           return "userID", nil
       }
       return "", err
    }

    WsServer.SendAuth = func(r *http.Request) bool {
        return true
    }

    if err := WsServer.ListenAndServe(); err != nil {
        panic(err)
    }
}