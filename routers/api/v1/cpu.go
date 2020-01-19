package v1

import (
    "github.com/gin-gonic/gin"
    "monitor/e"
    "monitor/model"
    "monitor/ssh"
    "net/http"
)

func GetCpu(c *gin.Context) {
    cli := ssh.New("59.110.243.78","root","_qq123456",22)
    num,err := cli.Run("cat /proc/cpuinfo | grep processor | wc -l")
    name,err := cli.Run("cat /proc/cpuinfo | grep 'model name'|tail -n 1|awk -F':' '{ print $2 }'|sed 's/^ //'")
    res := model.Cpu{Num:num,Name:name}

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
