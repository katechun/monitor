package v1

import (
    "github.com/gin-gonic/gin"
    "monitor/e"
    "monitor/model"
    "monitor/ssh"
    "net/http"
    "strings"
)

func GetIo(c *gin.Context) {
    cli := ssh.New("59.110.243.78","root","_qq123456",22)
    //读成功次数  读请求扇区数   写成功次数  写请求扇区数
    ioInfo1,err := cli.Run("cat /proc/diskstats |grep -vE 'dm-|sr0'|grep '     0'|awk '{ print $4\" \"$6\" \"$8\" \"$10 }'")
    ioInfo:= strings.Fields(ioInfo1)
    rSusessCnt := ioInfo[0]
    rSectorCnt := ioInfo[1]
    wSusessCnt := ioInfo[2]
    wSectorCnt := ioInfo[3]

    res := model.Io{RSusCnt:rSusessCnt,RSecCnt:rSectorCnt,WSusCnt:wSusessCnt,WSecCnt:wSectorCnt}

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
