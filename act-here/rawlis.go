package main

import (
	"bytes"
	"container/list"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func printList(coll *list.List) {
	for e := coll.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func printListR(coll *list.List) {
	for e := coll.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func main() {
	coll := list.New()

	coll.PushBack(1)
	coll.PushBack("Gopher")

	coll.PushFront("Cynhard")
	two := coll.PushFront(2)
	printList(coll)

	before2 := coll.InsertBefore("Before2", two)
	printList(coll)

	after2 := coll.InsertAfter("After2", two)
	printList(coll)

	fmt.Println("---------------")
	coll.MoveAfter(before2, two)
	printList(coll)

	coll.MoveBefore(after2, two)
	printList(coll)

	coll.MoveToFront(before2)
	printList(coll)
	coll.MoveToBack(after2)
	printList(coll)

	fmt.Println("===================")
	coll2 := list.New()
	coll2.PushBack(3)
	coll2.PushFront("Tomcat")

	coll.PushBackList(coll2)
	coll.PushFrontList(coll2)

	printList(coll)

	fmt.Println("~~~~~~~~~~~~~~~~~~")
	fmt.Println(coll.Front().Value,"----")
	fmt.Println(coll.Back().Value,"====")

	fmt.Println(coll.Len(),"+++++")

	coll.Remove(two)

	printList(coll)

	coll.Init()
	fmt.Println(123)
	printList(coll)
}


func v103stat(){
	model.GetSlaveDB().
		Model(model.Pc{}).
		Order("id asc").
		Offset(10*i).
		Limit(10).
		Scan(&pcs)
	var ids []int
	for _, pc := range pcs {
		fmt.Println(pc)
		ids = append(ids, pc.ID)
	}
	areaMap, _ := model.GetPcsAreaMap(ids)

	//fmt.Println(pcs)
	fmt.Printf("%+v",areaMap)

	for _, pc := range pcs {
		if val,ok:=areaMap[pc.ID];ok{
			fmt.Printf("%+v",val	)

			//存在区域的pc进行时长统计
			e.GetOnePcOnlineDuration(pc.ID,val.ID)

		}

		fmt.Println()

	}
}


package service

import (
"bytes"
"container/list"
"dcmc-cms/model"
"encoding/csv"
"encoding/json"
"fmt"
"github.com/garyburd/redigo/redis"
"github.com/gin-gonic/gin"
"pkg.deepin.com/service/utils/logger"
"strconv"
"time"
)

//RedisKeyGetOnlinePcToRedisSet 统计区域当日终端上线数量的集合key
const RedisKeyGetOnlinePcToRedisSet = "areaid:%d:day:%s:OnlinePc"

type SimpleArea struct {
	ID   int
	Name string
}

type StatPcSvc struct {
	AdminID int
	Nowdate string
}

func InitStatPc(adminID int) (statPc *StatPcSvc) {
	statPc = &StatPcSvc{
		AdminID: adminID,
		Nowdate: time.Now().Format("2006-01-02"),
	}
	return
}

//SavePcOnlineNumsToDb 持久化当日登录pc数量至mysql  59 23 * * *
func (e *StatPcSvc) SavePcOnlineNumsToDb() {
	currentTimeData := time.Now().Format("2006-01-02")

	logger.Info("Start to SavePcOnlineNumsToDb")
	model.GetMasterDB().Where("date = ?", currentTimeData).Delete(&model.StatAreapcnumDay{})

	//批处理
	var areas []SimpleArea
	model.GetMasterDB().Model(model.Area{}).Select("id,name").Scan(&areas)

	for _, v := range areas {

		adminUser := &model.StatAreapcnumDay{
			Areaid:      v.ID,
			Date:        currentTimeData,
			Pcnum:       model.GetAreaPCCount(v.ID),
			OnlinePcnum: e.GetOnlinePcnumsByAreaid(v.ID),
			Areaname:    v.Name,
		}

		err := model.GetMasterDB().Create(adminUser).Error
		if err != nil {
			logger.Warning("Failed to create sql:", adminUser, "err:", err)
		}
	}
	logger.Info("Finish to SavePcOnlineNumsToDb", areas, len(areas))
}

//SetOnlinePcToRedisSet 定时任务将当前在线的pc记入redis集合  key=areaid+nowdate  */10 * * * *   间隔越短越精确
func (e *StatPcSvc) SetOnlinePcToRedisSet() {
	//http://10.20.22.138:10082/online_client_ids
	logger.Info("Start to SetOnlinePcToRedisSet")

	onlineMachineIDs := GetOnlineClients()

	currentTimeData := time.Now().Format("2006-01-02")
	RedisOp, _ := model.InitMasterRedis()
	defer RedisOp.Close()
	tmpHash := make(map[string]int, 0)
	//优化  todo 协程处理
	for _, v := range onlineMachineIDs {
		//machine_id :="28eb55eb91b54e5b81ee1509e4164b97"
		machineID := v.(string)
		areaid, _ := GetPcAreaIDByMachineID(machineID) //优化  批量读取
		if areaid != 0 {
			RedisKey := fmt.Sprintf(RedisKeyGetOnlinePcToRedisSet, areaid, currentTimeData)
			if _, exited := tmpHash[RedisKey]; !exited {
				tmpHash[RedisKey] = 1
			}
			RedisOp.Conn.Send("SADD", RedisKey, machineID)
		}

	}
	for index := range tmpHash {
		RedisOp.Conn.Send("EXPIRE", index, 3*86400)
	}

	RedisOp.Conn.Flush()
	logger.Info("end to SetOnlinePcToRedisSet data:", onlineMachineIDs, "do expire:", tmpHash)

}

//GetStatofPcByPeriod 查询阶段内数据入口
func (e *StatPcSvc) GetStatofPcByPeriod(begin string, end string) []model.StatAreapcnumDay {

	var res []model.StatAreapcnumDay
	rawModel := model.GetSlaveDB().Model(model.StatAreapcnumDay{}).
		Select("sum(pcnum) as pcnum,date,sum(online_pcnum) as online_pcnum")
	admin := &model.Admin{ID: e.AdminID}
	if !admin.IsSuperAdmin() {
		areaids := admin.GetAdminAreaIDsPost()
		rawModel = rawModel.Where("areaid in (?)", areaids)
	}
	rawModel.
		Where("date <= (?)", end).
		Where("date >= (?)", begin).
		Group("date").
		Order("date asc").
		Scan(&res)

	//当天的数据需要实时获取
	if end >= e.Nowdate {
		nowDayStat := e.GetNowadayRealTimeStat()
		res = append(res, nowDayStat)
	}

	allDates := GetBetweenDates(begin, end)

	result := make([]model.StatAreapcnumDay, 0)

	var hashMap = make(map[string]model.StatAreapcnumDay, 0)
	for i, v := range res {
		res[i].Rate = 0
		if v.Pcnum != 0 {
			if v.OnlinePcnum > v.Pcnum {
				res[i].OnlinePcnum = v.Pcnum
			}

			res[i].OfflinePcnum = v.Pcnum - res[i].OnlinePcnum

			res[i].Rate, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", float64(res[i].OnlinePcnum)/float64(v.Pcnum)), 64)

		}
		hashMap[v.Date] = res[i]
	}

	for _, v := range allDates {

		if val, ok := hashMap[v]; ok {
			result = append(result, val)
		} else {
			result = append(result, model.StatAreapcnumDay{Date: v, Rate: 0})

		}

	}

	return result

}




func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

//StatOfPcOnlineExportCsv 导出csv方法
func (e *StatPcSvc) StatOfPcOnlineExportCsv(filename string, c *gin.Context, raw interface{}) {

	// 写数据到csv文件
	b := &bytes.Buffer{}


	w := csv.NewWriter(b)
	var  data = [][]string{}
	if rawd,ok :=raw.([]model.StatAreapcnumDay);ok{
		header := []string{"日期", "终端总数量", "上线终端数量", "离线终端数量", "上线率"} //标题
		data = [][]string{
			header,
		}
		for _, v := range rawd {
			line := []string{
				v.Date,
				strconv.Itoa(v.Pcnum),
				strconv.Itoa(v.OnlinePcnum),
				strconv.Itoa(v.OfflinePcnum),
				fmt.Sprintf("%.2f%%", v.Rate*100),
			}
			data = append(data, line)
		}
	}
	if rawd,ok :=raw.([]model.a);ok{

		header := []string{"日期", "终端使用时长", "平均使用时长"} //标题
		data = [][]string{
			header,
		}
		for _, v := range rawd {
			line := []string{
				v.Date,
				fmt.Sprintf("%.2f",v.Totalduration),
				fmt.Sprintf("%.2f",v.Avgduration),
			}
			data = append(data, line)
		}
	}


	// WriteAll方法使用Write方法向w写入多条记录，并在最后调用Flush方法清空缓存。
	w.WriteAll(data)
	w.Flush()
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", filename))
	//tet, _ := UTF82GBK(b.Bytes())
	tet := b.String()
	c.String(200, tet)

}




//GetOnlinePcnumsByAreaid 根据areaid获取今日在线pc数量
func (e *StatPcSvc) GetOnlinePcnumsByAreaid(areaid int) int {
	currentTimeData := time.Now().Format("2006-01-02")
	RedisOp, _ := model.InitSlaveRedis()
	defer RedisOp.Close()
	RedisKey := fmt.Sprintf(RedisKeyGetOnlinePcToRedisSet, areaid, currentTimeData)
	res, _ := redis.Int(RedisOp.Conn.Do("SCARD", RedisKey))
	return res
}

func (e *StatPcSvc) GetOnlinePcnumsByAreaids(areaids []int) (res map[int]int) {
	currentTimeData := time.Now().Format("2006-01-02")
	res = make(map[int]int, 0)
	RedisOp, _ := model.InitSlaveRedis()
	defer RedisOp.Close()
	for _, areaid := range areaids {
		RedisKey := fmt.Sprintf(RedisKeyGetOnlinePcToRedisSet, areaid, currentTimeData)
		res[areaid], _ = redis.Int(RedisOp.Conn.Do("SCARD", RedisKey))
	}

	return
}

//GetNowadayRealTimeStat   获取今日的实时报表数据   实时redis在线 + 实时mysql存量pc
func (e *StatPcSvc) GetNowadayRealTimeStat() (res model.StatAreapcnumDay) {
	admin := &model.Admin{ID: e.AdminID}
	areaids := admin.GetAdminAllAreaIds()

	var count int
	model.GetSlaveDB().Model(&model.AreaPC{}).Where("area_id in (?)", areaids).Count(&count)

	res.Pcnum = count
	res.Date = time.Now().Format("2006-01-02")
	res.OnlinePcnum = 0
	tmp := e.GetOnlinePcnumsByAreaids(areaids)
	for _, v := range tmp {
		res.OnlinePcnum += v
	}

	return
}

//*********v1.0.3 pc在线时长统计**********

func MakeListOfLogDetails() *list.List {
	coll := list.New()
	return coll
}

const CARCULATE_DAY_NUM = -1

type SimplePC struct {
	MachineID string
	AreaID    int
	PCID      int
	ID        int
}
type OnlineDetail struct {
	Type string
	Time    time.Time
}
func (e *StatPcSvc) SavePcOnlineDurationToDb() {

	logger.Info("Start to SavePcOnlineDurationToDb")
	//model.GetMasterDB().Where("date = ?", currentTimeData).Delete(&model.StatAreapcnumDay{})

	i := 0

	for {
		var pcs []SimplePC
		//get all pc within belong area

		model.GetSlaveDB().Table("dcmc_area_pc").Select("dcmc_area_pc.area_id,dcmc_area_pc.id, dcmc_pc.machine_id, dcmc_area_pc.pc_id").Joins("JOIN dcmc_pc ON dcmc_area_pc.pc_id = dcmc_pc.id").
			Order("id asc").
			Offset(10 * i).
			Limit(10).
			Scan(&pcs)

		currentTime := time.Now()
		zeroTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
			0, 0, 0, 0,time.UTC)

		lastDay := zeroTime.AddDate(0,0,CARCULATE_DAY_NUM).Format("2006-01-02")

		for _, pc := range pcs {


			onlineDuration, intimes, outtimes, details := e.GetOnePcOnlineDuration(pc.PCID, pc.AreaID,zeroTime)

			fmt.Printf("%+v", pc)
			fmt.Println()
			duration := &model.a{
				Areaid:         pc.AreaID,
				Machineid:      pc.MachineID,
				Onlineduration: onlineDuration,
				Pcid:           pc.PCID,
				Logintimes:          intimes,
				Logouttimes:          outtimes,
				Todaydetail:          details,
				Date:           lastDay,
			}

			logger.Info(duration)
			fmt.Println("===========")
			err := model.GetMasterDB().Create(duration).Error
			if err != nil {
				logger.Warning("Failed to create sql:", duration, "err:", err)
			}

		}

		i++
		if len(pcs) < 10 {
			break
		}
	}
}

func (e *StatPcSvc) GetOnePcOnlineDuration(pcid int, areaid int,zeroTime time.Time) (secs int,intimes int,outtimes int, details string){

	oldtt:=zeroTime.AddDate(0, 0, CARCULATE_DAY_NUM)
	old1Time := oldtt.Format("2006-01-02 15:04:05")


	datett := oldtt.Add(time.Second*86400)
	dateTime := datett.Format("2006-01-02 15:04:05")

	fmt.Println(dateTime, old1Time)
	var logs []model.LogHumanPc
	model.GetMasterDB().Model(&model.LogHumanPc{}).
		Order("created_at asc").
		Where("`pc_id` = ?", pcid).
		Where("`area_id` = ?", areaid).
		Where("`created_at` >= ?", old1Time).
		Where("`created_at` < ?", dateTime).
		Scan(&logs)

	logger.Dump(logs)
	if len(logs)==0{
		//当日无登录登出  查找之前是否有登录记录  todo

		var log model.LogHumanPc
		model.GetMasterDB().Model(&model.LogHumanPc{}).
			Order("created_at desc").
			Where("`pc_id` = ?", pcid).
			Where("`area_id` = ?", areaid).
			Where("`created_at` < ?", old1Time).
			Limit(1).
			First(&log)
		if log.Operation==model.LogOperationLogin{
			return 86400,0,0,""
		}else{
			return 0,0,0,""
		}
	}

	var todayDuration float64
	var logintimes,logoutimes int

	detail:=make([]OnlineDetail,0)
	for k,v:=range logs{

		if v.Operation == model.LogOperationLogout {
			logoutimes++
		}else{
			logintimes++
		}

		detail = append(detail,OnlineDetail{Time: v.CreatedAt,Type: v.Operation.String()})
		if k==0{
			//首次记录事件 若为登出  当天第一次记录为登出，则将00：00：01秒视为登录时间；
			if v.Operation==model.LogOperationLogout {
				todayDuration+=v.CreatedAt.Sub(oldtt).Seconds()
			}
		}

		if k==len(logs)-1{
			//最后一次记录事件 若为登录  当天最后一次记录为登录，则将23：59：59秒视为登出时间；
			if v.Operation==model.LogOperationLogin {
				todayDuration+=datett.Sub(v.CreatedAt).Seconds()
			}
		}

		if k-1>=0 {
			//存在前一次的登录记录  more than only one event
			if v.Operation==model.LogOperationLogin  && logs[k-1].Operation==model.LogOperationLogout{
				continue
			}
			todayDuration+=v.CreatedAt.Sub(logs[k-1].CreatedAt).Seconds()
		}

	}

	st,_:=json.Marshal(detail)
	fmt.Println(string(st),logintimes,logoutimes,"@@@@@@@@@@@@@@@@@@@@@")

	return int(todayDuration),logintimes,logoutimes,string(st)
}


func (e *StatPcSvc) GetPcDurationByPeriod(begin string, end string) []model.a {

	var res []model.a
	rawModel := model.GetSlaveDB().Model(model.a{}).
		Select("sum(onlineduration) as onlineduration,date,count(1) as pcnum")

	admin := &model.Admin{ID: e.AdminID}

	if !admin.IsSuperAdmin() {
		areaids := admin.GetAdminAreaIDsPost()
		rawModel = rawModel.Where("areaid in (?)", areaids)
	}
	rawModel.
		Where("date <= (?)", end).
		Where("date >= (?)", begin).
		Group("date").
		Order("date asc").
		Scan(&res)


	var hashMap = make(map[string]model.a, 0)
	for i, v := range res {
		res[i].Totalduration,_=strconv.ParseFloat(fmt.Sprintf("%.2f",float64(v.Onlineduration)/float64(3600)),64)

		res[i].Avgduration,_=strconv.ParseFloat(fmt.Sprintf("%.2f",float64(res[i].Totalduration/float64(v.Pcnum))),64)
		hashMap[v.Date] = res[i]
	}

	allDates := GetBetweenDates(begin, end)
	result := make([]model.a, 0)

	for _, v := range allDates {
		if val, ok := hashMap[v]; ok {
			result = append(result, val)
		} else {
			result = append(result, model.a{Date: v, Totalduration: 0,Avgduration: 0})

		}
	}
	logger.Dump(result,"***************************")

	return result
}
