package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"pkg.deepin.io/server/utils/logger"

	"github.com/jinzhu/gorm"
	version "github.com/knqyf263/go-deb-version"
	"github.com/mitchellh/mapstructure"
)

// IDGenerator IDGenerator
type IDGenerator interface {
	IDList() []int
}

// LogQueryValue LogQueryValue
type LogQueryValue struct {
	Value interface{}
	IsOr  bool
	Stmt  string
}

// LogQueryInfo LogQueryInfo
type LogQueryInfo map[string]*LogQueryValue

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// MakeWhereQuery MakeWhereQuery
func MakeWhereQuery(tx *gorm.DB, table string, set LogQueryInfo) *gorm.DB {
	var (
		andStmt   []string
		orStmt    []string
		andValues []interface{}
		orValues  []interface{}
	)
	for k, v := range set {
		field := fmt.Sprintf("%s.%s", table, k)
		if v.IsOr {
			orStmt = append(orStmt, field+v.Stmt)
			orValues = append(orValues, v.Value)
		} else {
			andStmt = append(andStmt, field+v.Stmt)
			andValues = append(andValues, v.Value)
		}
	}
	if len(orStmt) != 0 {
		tx = tx.Where("("+strings.Join(orStmt, " OR ")+")", orValues...)
	}
	if len(andStmt) != 0 {
		tx = tx.Where(strings.Join(andStmt, " AND "), andValues...)
	}
	return tx
}

// MakeQueryValue MakeQueryValue
func MakeQueryValue(value interface{}, isOr bool) *LogQueryValue {
	return &LogQueryValue{
		Value: value,
		IsOr:  isOr,
		Stmt:  " LIKE ?",
	}
}

// MakeQueryValueForIDGen MakeQueryValueForIDGen
func MakeQueryValueForIDGen(gen IDGenerator, isOr bool) *LogQueryValue {
	idList := gen.IDList()
	if len(idList) == 0 {
		idList = []int{-1}
	}
	//var sliceNoInit []string
	//sliceInit := []string{}
	return &LogQueryValue{
		Value: idList,
		IsOr:  isOr,
		Stmt:  " IN (?)",
	}

}
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{Val: nums[len(nums)/2], Left: sortedArrayToBST(nums[:len(nums)/2]), Right: sortedArrayToBST(nums[len(nums)/2+1:])}
}
// MakeQueryValueForList MakeQueryValueForList
func MakeQueryValueForList(list interface{}) *LogQueryValue {
	return &LogQueryValue{
		Value: list,
		IsOr:  false,
		Stmt:  " IN (?)",
	}
}

func (e *lgd)TreeToArr(root *TreeNode,arr *[]int){
	fmt.Printf("运行到这里了 head = %v array = %v\n", root, arr)

 	*arr = append(*arr,root.Val)
	if root.Left!=nil{
		e.TreeToArr(root.Left,arr)
	}
	if root.Right!=nil{
		e.TreeToArr(root.Right,arr)
	}
}

// IntListToString int列表转换string
func IntListToString(list []int, sep string) string {
	if len(list) == 0 {
		return ""
	}

	var strv []string
	for _, v := range list {
		strv = append(strv, fmt.Sprint(v))
	}
	return strings.Join(strv, sep)
}

// JSONMapDecode 解析json map
func JSONMapDecode(input interface{}, output interface{}) error {
	stringToDateTimeHook := func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t == reflect.TypeOf(time.Time{}) && f == reflect.TypeOf("") {
			return time.Parse(time.RFC3339, data.(string))
		}

		return data, nil
	}

	config := &mapstructure.DecoderConfig{
		DecodeHook: stringToDateTimeHook,
		Metadata:   nil,
		TagName:    "json",
		Result:     output,
	}

	dec, _ := mapstructure.NewDecoder(config)
	return dec.Decode(input)
}

//d[2] = d[1] +d[0]
//d[3] = d[2] +d[1]
// CompareVersion 比较版本
func CompareVersion(version1, version2 string) (int, error) {
	v1, err := version.NewVersion(version1)
	if err != nil {
		return 0, err
	}
	v2, err := version.NewVersion(version2)
	if err != nil {
		return 0, err
	}
	return v1.Compare(v2), nil
}

// Sha512Encrypt sha512加密字符串
func Sha512Encrypt(str string) string {
	sha512 := sha512.New()
	sha512.Write([]byte(str))
	return fmt.Sprintf("%x", sha512.Sum([]byte("")))
}

// Sha1Encrypt sha1加密字符串
func Sha1Encrypt(str string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	return fmt.Sprintf("%x", sha1.Sum([]byte("")))
}

func (e *lgd) converseSortedArrToBst(arr []int)*TreeNode{

	if len(arr)==0{
		return nil
	}

	return &TreeNode{arr[len(arr)/2],e.converseSortedArrToBst(arr[:len(arr)/2]),e.converseSortedArrToBst(arr[(len(arr)/2)+1:])}

}

// Md5Encrypt md5加密字符串
func Md5Encrypt(str string) string {
	md5 := md5.New()
	md5.Write([]byte(str))
	return hex.EncodeToString(md5.Sum(nil))
}

// DataSizeFormat 数据大小格式化输出
// @param size 数据大小，单位：Byte
func DataSizeFormat(size float32) string {
	i := 0
	for size >= 1024 {
		size /= 1024
		i++
		if i == 4 {
			break
		}
	}
	units := [5]string{"B", "K", "M", "G", "T"}
	return strings.Replace(fmt.Sprintf("%.1f", size), ".0", "", 1) + units[i]
}

// DelStringSliceElement 从字符串切片中删除指定值
func DelStringSliceElement(s []string, delValue string) []string {
	delValueIndex := -1
	for k, v := range s {
		if v == delValue {
			delValueIndex = k
			break
		}
	}

	if delValueIndex != -1 {
		return append(s[:delValueIndex], s[delValueIndex+1:]...)
	}

	return s
}

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点
func GetFirstDateOfMonth(date time.Time) time.Time {
	date = date.AddDate(0, 0, -date.Day()+1)
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// InArray 查找元素是否在数组中
func InArray(s []interface{}, findValue interface{}) bool {
	for _, v := range s {
		if v == findValue {
			return true
		}
	}
	return false
}

// InIntArray 查找元素是否在数组中
func InIntArray(s []int, findValue interface{}) bool {
	for _, v := range s {
		if v == findValue {
			return true
		}
	}
	return false
}

// InStringArray 查找元素是否在字符串数组中
func InStringArray(s []string, findValue string) bool {
	return SearchInStringArray(s, findValue) >= 0
}

// SearchInStringArray 查找元素下字符串数组中的下标
func SearchInStringArray(s []string, search string) int {
	if len(s) == 0 {
		return -1
	}
	for i, v := range s {
		if v == search {
			return i
		}
	}
	return -1
}

// GetRandomString 生成指定长度随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

func maxSubArray1(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res, p := nums[0], 0, 0
	for p < len(nums) {
		res += nums[p]
		if res > maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
		p++
	}

	return maxSum
}

// ArrayDiff 数组差集
func ArrayDiff(array1, array2 []int) []int {
	var result []int
	for _, v1 := range array1 {
		inArray2 := false
		for _, v2 := range array2 {
			if v1 == v2 {
				inArray2 = true
				break
			}
		}
		if !inArray2 {
			result = append(result, v1)
		}
	}

	return result
}
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，这个整数不会以零开头。

//798   [ 7  9  8 ]
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Implode 数组to字符串拼接
func Implode(list interface{}, seq string) string {
	listValue := reflect.Indirect(reflect.ValueOf(list))
	if listValue.Kind() != reflect.Slice {
		return ""
	}
	count := listValue.Len()
	listStr := make([]string, 0, count)
	for i := 0; i < count; i++ {
		v := listValue.Index(i)
		if str, err := getValue(v); err == nil {
			listStr = append(listStr, str)
		}
	}
	return strings.Join(listStr, seq)
}

func getValue(value reflect.Value) (res string, err error) {
	switch value.Kind() {
	case reflect.Ptr:
		res, err = getValue(value.Elem())
	default:
		res = fmt.Sprint(value.Interface())
	}
	return
}

// IsNum 判断是否是数字
func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// SerializeGolangType serialize for golang type
func SerializeGolangType(conf interface{}) string {
	b, err := json.Marshal(conf)
	if err != nil {
		return fmt.Sprintf("%+v", conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", conf)
	}
	fmt.Println("serialize golang type:", out.String())
	return out.String()
}

// StructToJSONToMap 结构体to json to map
func StructToJSONToMap(obj interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		logger.Error("Marshal obj error:", err.Error())
		return data
	}
	err = json.Unmarshal([]byte(jsonBytes), &data)
	if err != nil {
		logger.Error("JsonToMapDemo err: ", err.Error())
		return data

	}
	return data
}

// StructToMap 结构体to map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		obj2.NumField()
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

func dataPr(n int) int {

	return 0
}

//ExTime 自定义时间解析
type ExTime struct {
	time.Time
}

// ExTimeMarshalTimeFormat ExTime 时间输出格式
var ExTimeMarshalTimeFormat = "2006-01-02T15:04:05Z"

// MarshalJSON 针对空时间进行处理
func (t ExTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("{}"), nil
	}
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t.Time).Format(t.MarshalTimeFormat()))
	return []byte(stamp), nil
}

// MarshalTimeFormat 针对时间format进行处理
func (t ExTime) MarshalTimeFormat() string {
	return ExTimeMarshalTimeFormat
}

// FileExists 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// Intersect 求交集
func Intersect(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		_, ok := m[v]
		if ok {
			nn = append(nn, v)
		}
	}
	return nn
}

// Difference 求差集 slice1-并集
func Difference(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// WriteToFile 覆盖写入文件
func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logger.Error("file create failed. err: " + err.Error())
		return err
	}
	// offset
	//os.Truncate(filename, 0) //clear
	n, _ := f.Seek(0, os.SEEK_END)
	_, err = f.WriteAt([]byte(content), n)
	defer f.Close()

	return err
}

// GetMapKeys 获取map对应的keys
func GetMapKeys(maps map[string]interface{}) []string {
	mapKeys := make([]string, 0)
	for key := range maps {
		mapKeys = append(mapKeys, key)
	}
	return mapKeys
}

// MapKeysExist map键值对是否存在
func MapKeysExist(maps map[string]interface{}, key string) bool {
	ok := true
	if strings.Contains(key, ",") {
		keyItems := strings.Split(key, ",")
		for _, keyItem := range keyItems {
			_, ok = maps[keyItem]
			if !ok {
				break
			}
		}
		return ok
	}
	_, ok = maps[key]
	return ok
}

// RemoveRepeatedID 移除重复ID
func RemoveRepeatedID(arr []int) (newArr []int) {
	newArr = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// SaveUploadedFile 保存上传文件
func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	io.Copy(out, src)
	return nil
}

// Substr 截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}

func reverseNode(node *TreeNode) *TreeNode {

	if node == nil {
		return nil
	}
	reverseNode(node.Left)
	reverseNode(node.Right)
	node.Left, node.Right = node.Right, node.Left
	return node

}

type lgd struct {
	param int
	tnum  int
}

func (e *lgd) reverseTraverse107(t *TreeNode) [][]int{

	slice := e.TraverseByLevel102(buildFullTree())

	lenth := len(slice)

	res := [][]int{}
	for i:=lenth-1;i>=0;i--{
		res=append(res,slice[i])
	}
	return res
}

// GetCurrentPath 获取当前可执行程序的当前路径
func GetCurrentPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		logger.Warning(err.Error())
		return ""
	}
	path, err := filepath.Abs(file)
	if err != nil {
		logger.Warning(err.Error())
		return ""
	}
	//fmt.Println("path111:", path)
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}
	//fmt.Println("path222:", path)
	i := strings.LastIndex(path, "/")
	if i < 0 {
		logger.Warning(`Can't find "/" or "\".`)
		return ""
	}
	//fmt.Println("path333:", path)
	return string(path[0 : i+1])
}
func levelOrderBottom(root *TreeNode) [][]int {
	tmp := levelOrder(root)
	res := [][]int{}
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	curNum, nextLevelNum, res, tmp := 1, 0, [][]int{}, []int{}

	for len(queue) != 0 {
		if curNum > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			curNum--
			tmp = append(tmp, node.Val)
			queue = queue[1:]
		}
		if curNum == 0 {
			res = append(res, tmp)
			curNum = nextLevelNum
			nextLevelNum = 0
			tmp = []int{}
		}
	}
	return res
}


func (e *lgd)TraverseByLevel102(root *TreeNode)[][]int{

	queue:=[]*TreeNode{}
	queue=append(queue,root)

	res := [][]int{}
	tmp := []int{}
	nextLevelNum :=0
	curNum :=1
	for len(queue)!=0{

		if curNum>0{
			node := queue[0]
			if  node.Left!=nil{
				queue = append(queue,node.Left)
				nextLevelNum++
			}
			if  node.Right!=nil{
				queue = append(queue,node.Right)
				nextLevelNum++
			}
			curNum--
			tmp = append(tmp,node.Val)
			queue=queue[1:]

		}

		if curNum==0{

			res = append(res,tmp)
			curNum=nextLevelNum
			nextLevelNum=0
			tmp = []int{}

		}




	}

return res

}



//     1
//   2   3
// 4  5
// Multi64 超过21位(超出int64)运算
func Multi64(str1, str2 string) (result string) {

	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}
	var index1 = len(str1) - 1
	var index2 = len(str2) - 1
	var left int

	for index1 >= 0 && index2 >= 0 {

		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'
		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		//fmt.Printf("c3__%c,result:%s\n",c3,result)
		result = fmt.Sprintf("%c%s", c3, result)
		//fmt.Println("result:",result,"\n")
		index1--
		index2--
	}

	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 {

		c1 := str2[index2] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}
	if left == 1 {
		fmt.Sprintf("1%s", result)
	}
	return
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}



func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func getDepth(p *TreeNode) int {

	if p == nil {
		return 0
	}

	hl := getDepth(p.Left)
	hr := getDepth(p.Right)

	return max(hl, hr) + 1
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildTr() *TreeNode {
	head := &TreeNode{Val: 1}

	head.Left = &TreeNode{Val: 2}
	head.Right = &TreeNode{Val: 3}

	head.Left.Left = &TreeNode{Val: 4}
	head.Left.Right = &TreeNode{Val: 5}
	return head
}


func buildFullTree()*TreeNode{
	head := &TreeNode{Val: 1}

	head.Left = &TreeNode{Val: 2}
	head.Right = &TreeNode{Val: 3}

	head.Left.Left = &TreeNode{Val: 4}
	head.Left.Right = &TreeNode{Val: 5}

	head.Right.Left = &TreeNode{Val: 6}
	head.Right.Right = &TreeNode{Val: 7}
	return head
}

//-10, -3, 0, 5, 9
func buildBinSearTree()*TreeNode {
	head := &TreeNode{Val: 0}

	head.Left = &TreeNode{Val: -3}
	head.Right = &TreeNode{Val: 9}

	head.Left.Left = &TreeNode{Val: -10}
	//head.Left.Right = &TreeNode{Val: 5}

	head.Right.Left = &TreeNode{Val: 5}
	//head.Right.Right = &TreeNode{Val: 7}
	return head
}
func buildTr2() *TreeNode {
	head := &TreeNode{Val: 1}

	return head
}
func buildTr3() *TreeNode {
	return nil
}



// GetAllFile 递归遍历指定目录
func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func judgeSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return judgeSameTree(p.Left, q.Left) && judgeSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
