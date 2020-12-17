package lib

import (
	"encoding/json"
	"github.com/hashicorp/raft"
	"io"
)

type MyFSM struct {

}
//真正 持久化数据
func(this *MyFSM) Apply(log *raft.Log) interface{}{
	req:=NewCacheRequest()
	err:=json.Unmarshal(log.Data,req)
	Set(req.Key,req.Value)  //真正执行  数据保存
	return err
}
func(this *MyFSM) Snapshot() (raft.FSMSnapshot, error){
	return nil,nil
}
func(this *MyFSM) Restore(io.ReadCloser) error{
	return nil
}
