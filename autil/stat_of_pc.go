package service

import (
	"bytes"
	"dcmc-cms/model"
	"encoding/csv"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"pkg.deepin.com/service/utils/logger"
	"strconv"
	"time"
)
//RedisKeyGetOnlinePcToRedisSet 统计区域当日终端上线数量的集合key
const 	RedisKeyGetOnlinePcToRedisSet = "areaid:%d:day:%s:OnlinePc"

type SimpleArea struct {
	ID int
	Name string
}

type StatPcSvc struct {
	AdminID int
	Nowdate string
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitStatPc(adminID int) (statPc *StatPcSvc) {
	statPc = &StatPcSvc{
		AdminID: adminID,
		Nowdate: time.Now().Format("2006-01-02"),
	}
	return
}


//SavePcOnlineNumsToDb 持久化当日登录pc数量至mysql  23 59 * * *
func (e *StatPcSvc)SavePcOnlineNumsToDb(){
	logger.Info("Start to SavePcOnlineNumsToDb")

	//批处理
	var areas []SimpleArea
	model.GetSlaveDB().Model(model.Area{}).Select("id,name").Scan(&areas)


	currentTimeData:=e.Nowdate
	for _,v :=range areas{

		adminUser := &model.StatAreapcnumDay{
			Areaid: v.ID,
			Date: currentTimeData,
			Pcnum: model.GetAreaPCCount(v.ID),
			OnlinePcnum: e.GetOnlinePcnumsByAreaid(v.ID),
			Areaname: v.Name,
		}

		err := model.GetMasterDB().Create(adminUser).Error
		if err!=nil {
			logger.Warning("Failed to create sql:", adminUser,"err:",err)
		}
	}
	logger.Info("Finish to SavePcOnlineNumsToDb",areas,len(areas))
}


//SetOnlinePcToRedisSet 定时任务将当前在线的pc记入redis集合  key=areaid+nowdate  */10 * * * *   间隔越短越精确
func (e *StatPcSvc)SetOnlinePcToRedisSet(){
	//http://10.20.22.138:10082/online_client_ids
	logger.Info("Start to SetOnlinePcToRedisSet")


	onlineMachineIDs := GetOnlineClients()

	currentTimeData:=e.Nowdate
	RedisOp,_:=model.InitMasterRedis()
	defer RedisOp.Close()
	tmpHash := make(map[string]int,0)
	//优化  协程处理 todo
	for _,v:=range onlineMachineIDs{
		//machine_id :="28eb55eb91b54e5b81ee1509e4164b97"
		machineID :=v.(string)
		areaid,_ :=GetPcAreaIDByMachineID(machineID)  //优化  批量读取
		if areaid!=0{
			RedisKey :=fmt.Sprintf(RedisKeyGetOnlinePcToRedisSet,areaid,currentTimeData)
			if _, exited := tmpHash[RedisKey];!exited{
				tmpHash[RedisKey]=1
			}
			RedisOp.Conn.Send("SADD", RedisKey, machineID)
		}

	}
	for index:=range tmpHash {
		RedisOp.Conn.Send("EXPIRE", index, 3*86400)
	}

	RedisOp.Conn.Flush()
	logger.Info("end to SetOnlinePcToRedisSet data:",onlineMachineIDs,"do expire:",tmpHash)

}



//GetStatofPcByPeriod 查询阶段内数据入口
func (e *StatPcSvc)GetStatofPcByPeriod(begin string,end string)[]model.StatAreapcnumDay{

	var res []model.StatAreapcnumDay
	rawModel  :=model.GetSlaveDB().Model(model.StatAreapcnumDay{}).
		Select("sum(pcnum) as pcnum,date,sum(online_pcnum) as online_pcnum")
	admin := &model.Admin{ID: e.AdminID}
	if !admin.IsSuperAdmin(){
		areaids := admin.GetAdminAreaIDsPost()
		rawModel=rawModel.Where("areaid in (?)", areaids)
	}
	rawModel.
	Where("date <= (?)", end).
	Where("date >= (?)",begin).
	Group("date").
	Order("date asc").
	Scan(&res)


	//当天的数据需要实时获取
	if end>=e.Nowdate{
		nowDayStat := e.GetNowadayRealTimeStat()
		res = append(res,nowDayStat)
	}

	allDates := GetBetweenDates(begin,end)


	result := make([]model.StatAreapcnumDay,0)


	var hashMap =make(map[string]model.StatAreapcnumDay,0)
	for i,v := range res{
		res[i].Rate="0.00%"
		if v.Pcnum!=0{

			res[i].OfflinePcnum = v.Pcnum-v.OnlinePcnum

			res[i].Rate =  fmt.Sprintf("%.2f%%", float64(v.OnlinePcnum)/float64(v.Pcnum)*100)
		}
		hashMap[v.Date]=res[i]
	}

	for _,v:= range allDates{

		if val,ok:=hashMap[v];ok{
			result=append(result,val)
		}else{
			result=append(result,model.StatAreapcnumDay{})

		}

	}



	return result

}

//StatOfPcOnlineExportCsv 导出csv方法
func (e *StatPcSvc)StatOfPcOnlineExportCsv(filename string,c *gin.Context,raw []model.StatAreapcnumDay)  {

	// 写数据到csv文件
	b := &bytes.Buffer{}

	w := csv.NewWriter(b)
	header := []string{"日期", "终端总数量", "上线终端数量", "离线终端数量", "上线率"} //标题
	data := [][]string{
		header,
	}
	for _, v := range raw{
		line := []string{
			v.Date,
			strconv.Itoa(v.Pcnum),
			strconv.Itoa(v.OnlinePcnum),
			strconv.Itoa(v.OfflinePcnum),
			v.Rate,
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

}

//GetOnlinePcnumsByAreaid 根据areaid获取今日在线pc数量
func (e *StatPcSvc)GetOnlinePcnumsByAreaid(areaid int)int{
	currentTimeData:=e.Nowdate
	RedisOp,_:=model.InitSlaveRedis()
	defer RedisOp.Close()
	RedisKey :=fmt.Sprintf(RedisKeyGetOnlinePcToRedisSet,areaid,currentTimeData)
	res, _ := redis.Int(RedisOp.Conn.Do("SCARD", RedisKey))
	return res
}

func (e *StatPcSvc)GetOnlinePcnumsByAreaids(areaids []int)(res map[int]int){
	currentTimeData:=e.Nowdate
	res = make(map[int]int,0)
	RedisOp,_:=model.InitSlaveRedis()
	defer RedisOp.Close()
	for areaid := range areaids{
		RedisKey :=fmt.Sprintf(RedisKeyGetOnlinePcToRedisSet,areaid,currentTimeData)
		res[areaid], _ = redis.Int(RedisOp.Conn.Do("SCARD", RedisKey))
	}

	return
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

//GetNowadayRealTimeStat   获取今日的实时报表数据   实时redis在线 + 实时mysql存量pc
func (e *StatPcSvc)GetNowadayRealTimeStat()(res model.StatAreapcnumDay){
	admin := &model.Admin{ID: e.AdminID}
	areaids := admin.GetAdminAllAreaIds()

	var count int
 	model.GetSlaveDB().Model(&model.AreaPC{}).Where("area_id in (?)", areaids).Count(&count)


	res.Pcnum = count
	res.Date  = e.Nowdate
	res.OnlinePcnum=0
	tmp :=e.GetOnlinePcnumsByAreaids(areaids)
	for _,v :=range tmp{
		res.OnlinePcnum+=v
	}

	return
}

func (e *Lgd2)JudgeAloopLink141(head *ListNode)(res bool){
	fast := head
	slow := head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false


}

/**
A和B两个链表长度可能不同，但是A+B和B+A的长度是相同的，所以遍历A+B和遍历B+A一定是同时结束。 如果A,B相交的话A和B有一段尾巴是相同的，所以两个遍历的指针一定会同时到达交点 如果A,B不相交的话两个指针就会同时到达A+B（B+A）的尾节点
 */
func (e *Lgd2)getIntersectionNode(headA, headB *ListNode) *ListNode {
	//boundary check
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	//if a & b have different len, then we will stop the loop after second iteration
	for a != b {
		//for the end of first iteration, we just reset the pointer to the head of another linkedlist
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		fmt.Printf("a = %v b = %v\n", a, b)
	}
	return a
}


type Lgd2 struct{

}

func (e *Lgd2)rob198(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// dp[i] 代表抢 nums[0...i] 房子的最大价值
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}





func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func (e *Lgd2)removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return newHead.Next
}


func DeleteNd(head *ListNode,val int)*ListNode{

	newhead := &ListNode{}

	pre := newhead
	cur :=head
	for cur!=nil{

		if cur.Val!=val{
			pre.Next = cur
			pre = pre.Next

		}
		cur = cur.Next
	}


	return newhead.Next
}

func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}



func ReverseLstRecusion(head *ListNode)*ListNode{
	if head.Next==nil{
		return head
	}

	newhead :=ReverseLstRecusion(head.Next)
	head.Next.Next=head
	head.Next=nil

	//head.Next=nil
	//newhead.Next=head
	return newhead
}



func ReverseLstTwoPtr(head *ListNode)*ListNode{

	cur := head

	var prev *ListNode

	for cur!=nil{
		tmp :=cur.Next
		cur.Next=prev
		prev=cur
		cur=tmp
	}

	return prev

}