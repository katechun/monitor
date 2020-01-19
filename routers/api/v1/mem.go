package v1

import (
    "github.com/gin-gonic/gin"
    "monitor/e"
    "monitor/model"
    "monitor/ssh"
    "net/http"
)

func GetMem(c *gin.Context) {
    cli := ssh.New("59.110.243.78","root","_qq123456",22)
    memTotal,err := cli.Run("cat /proc/meminfo|grep MemTotal|awk '{ print $2 }'")
    memFree,err := cli.Run("cat /proc/meminfo|grep MemFree|awk '{ print $2 }'")
    memAvailable,err := cli.Run("cat /proc/meminfo|grep MemAvailable|awk '{ print $2 }'")
    res := model.Mem{Total:memTotal,Free:memFree,Available:memAvailable}

    if err != nil {
        c.JSON(e.ERROR,gin.H{
            "status":e.ERROR,
            "msg": e.GetMsg(e.ERROR),
        })
    }else {
        c.JSON(http.StatusOK,gin.H{
            "status":e.SUCCESS,
            "msg":e.GetMsg(e.SUCCESS),
            "data":res,
        })
    }
}