package migrate

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

/*========================================================================*/

//统计区域当日终端上线数量的集合key
const 	RedisKeyGetOnlinePcToRedisSet = "areaid:%d:day:%s:OnlinePc"

type SimpleArea struct {
	ID int
	Name string
}

//SavePcOnlineNumsToDb 持久化当日登录pc数量至mysql  23 59 * * *
func SavePcOnlineNumsToDb(){
	logger.Info("Start to SavePcOnlineNumsToDb")

	//批处理
	var areas []SimpleArea
	model.GetSlaveDB().Model(model.Area{}).Select("id,name").Scan(&areas)


	currentTimeData:=time.Now().Format("2006-01-02")
	for _,v :=range areas{

		adminUser := &model.StatAreapcnumDay{
			Areaid: v.ID,
			Date: currentTimeData,
			Pcnum: model.GetAreaPCCount(v.ID),
			OnlinePcnum: GetOnlinePcnumsByAreaid(v.ID),
			Areaname: v.Name,
		}

		err := model.GetMasterDB().Create(adminUser).Error
		if err!=nil {
			logger.Warning("Failed to create sql:", adminUser,"err:",err)
		}
	}
	logger.Info("Finish to SavePcOnlineNumsToDb",areas,len(areas))

}


//GetOnlinePcToRedisSet 定时任务将当前在线的pc记入redis集合  key=areaid+nowdate  */10 * * * *
func SetOnlinePcToRedisSet(){
	//http://10.20.22.138:10082/online_client_ids
	onlineMachineIDs := GetOnlineClients()

	currentTimeData:=time.Now().Format("2006-01-02")
	RedisOp,_:=model.InitMasterRedis()
	defer RedisOp.Close()
	tmpHash := make(map[string]int,0)
	for _,v:=range onlineMachineIDs{
		//machine_id :="28eb55eb91b54e5b81ee1509e4164b97"
		machine_id :=v.(string)
		areaid,_ :=GetPcAreaIDByMachineID(machine_id)  //优化  批量读取
		RedisKey :=fmt.Sprintf(RedisKeyGetOnlinePcToRedisSet,areaid,currentTimeData)
		if _, exited := tmpHash[RedisKey];!exited{
			tmpHash[RedisKey]=1
		}
		RedisOp.Conn.Send("SADD", RedisKey, machine_id)
	}
	for index,_:=range tmpHash {
		RedisOp.Conn.Send("EXPIRE", index, 3*86400)
	}

	RedisOp.Conn.Flush()

}


//GetStatofPcByPeriod 查询阶段内数据入口
func GetStatofPcByPeriod(adminId int,begin string,end string)(res []model.StatAreapcnumDay){
	areaids :=[]int{47,48,49}
	model.GetSlaveDB().Model(model.StatAreapcnumDay{}).
		Select("sum(pcnum) as pcnum,date,sum(online_pcnum) as online_pcnum").
		Where("areaid in (?)", areaids).
		Where("date <= (?)", end).
		Where("date >= (?)",begin).
		Group("date").
		Order("date asc").
		Scan(&res)
	return

}



func GetOnlinePcnumsByAreaid(areaid int)int{
	currentTimeData:=time.Now().Format("2006-01-02")
	RedisOp,_:=model.InitSlaveRedis()
	defer RedisOp.Close()
	RedisKey :=fmt.Sprintf(RedisKeyGetOnlinePcToRedisSet,areaid,currentTimeData)
	res, _ := redis.Int(RedisOp.Conn.Do("SCARD", RedisKey))
	return res
}

func StatOfPcOnlineExportCsv(filename string,c *gin.Context) (err error) {

	// 写数据到csv文件
	b := &bytes.Buffer{}

	w := csv.NewWriter(b)
	header := []string{"日期", "终端总数量", "上线终端数量", "离线终端数量", "上线率"} //标题
	data := [][]string{
		header,
	}
	for k, _ := range []int{1,2,3} {
		hostIsOnline := "sd"

		line := []string{
			"asd",
			"ss%",
			"zxcxczx%",
			"sdasdasdasdas"+strconv.Itoa(k),
			hostIsOnline,
		}
		data = append(data, line)
	}
	// WriteAll方法使用Write方法向w写入多条记录，并在最后调用Flush方法清空缓存。
	w.WriteAll(data)
	w.Flush()
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", filename))
	//tet, _ := UTF82GBK(b.Bytes())
	tet:=b.String()
	c.String(200, tet)
	return
}


func GetOnlinePcnumsByAreaids(areaids []int)(res map[int]int){
	currentTimeData:=time.Now().Format("2006-01-02")
	res = make(map[int]int,0)
	RedisOp,_:=model.InitSlaveRedis()
	defer RedisOp.Close()
	for areaid := range areaids{
		RedisKey :=fmt.Sprintf(RedisKeyGetOnlinePcToRedisSet,areaid,currentTimeData)
		res[areaid], _ = redis.Int(RedisOp.Conn.Do("SCARD", RedisKey))
	}

	return
}


//GetNowadayRealTimeStat   获取今日的实时报表数据   实时redis在线 + 实时mysql存量pc
func GetNowadayRealTimeStat(adminId int){

}
