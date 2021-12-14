package router

import (
    "ecard_backend/conf"
    "ecard_backend/controller"
    "ecard_backend/log"
    "ecard_backend/midlleware"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "os"
)

func Init() {
    gin.SetMode(gin.DebugMode)
    r := gin.Default()

    r.Use(cors.New(cors.Config{
       AllowOrigins: []string{"*"},
       AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
       AllowHeaders: []string{"Origin", "Authorization", "Content-type", "User-Agent"},
    }))

    api := r.Group("/api")
    api.POST("/login", controller.UserLogin)

    api.Use(midlleware.Auth())
    {
        api.POST("/device", controller.AddDevice)
        api.GET("/device/count", controller.GetDeviceCount)

        api.POST("/virtual-device", controller.AddVirtualDevice)
        api.GET("/virtual-device", controller.GetVirtualDevice)

        api.POST("/student", controller.AddStudent)
        api.GET("/student/count", controller.GetStudentCount)
        api.GET("/student/info", controller.GetStudentInfo)

        api.GET("/alarm/count", controller.GetAlarmCount)
        api.GET("/alarm/info", controller.GetAlarmInfo)
        api.GET("/alarm/rank", controller.GetAlarmRank)


        api.GET("/school/info", controller.GetSchoolInfo)

        api.GET("/location/info", controller.GetLocationInfo)
        api.GET("/track/info", controller.GetTrackInfo)

        api.POST("/alarm", controller.Alarm)

        api.POST("/fence", controller.SetFence)
        api.GET("/fence", controller.GetFence)

        api.POST("/late-time", controller.SetLateTime)
        api.GET("/late-time", controller.GetLateTime)

        api.GET("/late-situation", controller.GetLateSituation)
    }

    err := r.Run(conf.HTTP_PORT)
    if err != nil {
        log.Error("Fail to start http server:%v", err)
        os.Exit(1)
    }
}
