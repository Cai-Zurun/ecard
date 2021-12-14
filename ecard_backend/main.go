package main

import (
    "ecard_backend/db"
    "ecard_backend/late_situation"
    "ecard_backend/log"
    "ecard_backend/router"
    "ecard_backend/virtual_device"
)

func main() {
    log.Init("./log/ecard_backend.log", "Debug", 100, 1)

    db.Init()
    virtual_device.Daemon()
    late_situation.Daemon()
    router.Init()
}
