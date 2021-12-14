package server

import (
    "crypto/tls"
    "ecard_gateway/conf"
    "ecard_gateway/ecard/handler"
    "ecard_gateway/ecard/packet"
    "ecard_gateway/log"
    "ecard_gateway/utils"
    "encoding/hex"
    "net"
    "sync"
    "time"
)

var (
    tcpServer *TcpServer
    mu        sync.Mutex
    notifyMap   sync.Map
)

type TcpServer struct {
    sendMsgChannel chan []byte
}

// 单例模式
func GetTcpServer() *TcpServer {
    if tcpServer == nil {
        mu.Lock()
        defer mu.Unlock()
        if tcpServer == nil {
            tcpServer = &TcpServer{
                sendMsgChannel: make(chan []byte, 1000),
            }
        }
    }
    return tcpServer
}

func (t *TcpServer) Start() {
    cert, err := tls.LoadX509KeyPair("./conf/server.pem", "./conf/server.key")
    if err != nil {
        panic("Error load tls file: " + err.Error())
    }
    config := &tls.Config{Certificates: []tls.Certificate{cert}}
    listener, err := tls.Listen("tcp", conf.TCP_SERVER_PORT, config)
    if err != nil {
        panic("Error listening: " + err.Error())
    }
    defer listener.Close()
    log.Info("-----------TCP Goroutine-Server listen on port" + conf.TCP_SERVER_PORT + "-----------")
    
    // 循环监听来自客户端的连接，给每个连接开一个协程 goroutine-per-connection
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Error("Error accepting:%s", err.Error())
        }

        // 维护连接心跳，当15分钟总内连接没有事件，则关闭连接
        notify := make(chan struct{})
        notifyMap.Store(conn, notify)
        go heartBeating(conn, notify)

        // 将该连接的后续工作交给协程处理
        go doServerStuff(conn)
    }
}

func doServerStuff(conn net.Conn) {
    //程序崩溃时恢复，并记录崩溃原因
    utils.RecoverAndLog()
    for {
        pkt, err := packet.GetPacket(conn)
        if err != nil {
            log.Error("get packet error:%v", err)
            return
        }
        log.Info("From Client %s:%s", conn.RemoteAddr(), hex.EncodeToString(pkt.PacketBuf))

        result := handler.AnalyzePacket(pkt)
        if result != nil {
            log.Info("From Server:%s", hex.EncodeToString(result))
            if err := utils.SendAll(conn, result); err != nil {
                log.Error("Error server writing: " + err.Error())
            }
        }

    }
}

func heartBeating(conn net.Conn, notify chan struct{}) {
    for {
        select {
        case <-notify:
            break
        case <-time.After(15 * time.Minute):    //15分钟内没连接没有事件发生，则关闭连接
            //if err := conn.Close(); err != nil {
            //    log.Error("conn close error:%s", err.Error())
            //}
            //notifyMap.Delete(conn)
        }
    }
}
