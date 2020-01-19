package v1

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "monitor/e"
    "monitor/model"
    "monitor/ssh"
    "net/http"
    "strconv"
    "strings"
)

func GetNet(c *gin.Context) {
    var (
        res []model.Net
    )
    cli := ssh.New("59.110.243.78","root","_qq123456",22)
    //网卡名称 接收的包  发送的包
    netCnt1,err := cli.Run("cat /proc/net/dev|grep -vE 'lo|docker|veth|Inter|face'|awk '{ print $1\" \"$3\" \"$11 }'|grep -v ^$|wc -l")
    netCnt2 := strings.Replace(netCnt1,"\n","",-1)
    netCnt,_ := strconv.Atoi(netCnt2)
    for i := 1; i <= netCnt; i++ {
        info := fmt.Sprintf("cat /proc/net/dev|grep -vE 'lo|docker|veth|Inter|face'|awk '{ print $1\" \"$3\" \"$11 }'|sed -n '%dp'", i)
        netInfo1,_  := cli.Run(info)
        netInfo := strings.Fields(netInfo1)
        netName := netInfo[0]
        netResev := netInfo[1]
        netSend := netInfo[2]
        res = append(res,model.Net{Name:netName,Resev:netResev,Send:netSend})
    }

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
