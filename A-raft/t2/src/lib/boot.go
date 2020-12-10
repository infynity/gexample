package lib

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var RaftNode *raft.Raft
func BootStrap(path string ) error  {
	sysConfig,err:=LoadConfig(path)
	if err!=nil{
		return err
	}
	fmt.Println(sysConfig)
	config:=raft.DefaultConfig()
	config.LocalID=raft.ServerID(sysConfig.ServerID)
	config.Logger=hclog.New(&hclog.LoggerOptions{
		Name:  sysConfig.ServerName,
		Level: hclog.LevelFromString("DEBUG"),
		Output:os.Stderr,
	})


	//logStore保存配置
	dir,_:=os.Getwd()
	root:=strings.Replace(dir,"\\","/",-1)
	log_store,err:=raftboltdb.NewBoltStore(root+sysConfig.LogStore)
	if err!=nil{
		return err
	}

	//保存节点信息
	stable_store,err:=raftboltdb.NewBoltStore(root+sysConfig.StableStore)
	if err!=nil{
		return err
	}
	//不存储快照
	snapshotStore:=raft.NewDiscardSnapshotStore()

	// 节点之间的通信
	addr,err:=net.ResolveTCPAddr("tcp",sysConfig.Transport)
	transport,err:=raft.NewTCPTransport(addr.String(),addr,5,time.Second*10,os.Stdout)
	if err!=nil{
		log.Fatal(err)
	}
	fsm:=&MyFSM{}

	RaftNode,err=raft.NewRaft(config,fsm,log_store,stable_store,snapshotStore,transport)
	if err!=nil{
		return err
	}
	configuration := raft.Configuration{
		Servers: sysConfig.Servers,
	}

	RaftNode.BootstrapCluster(configuration)
	return nil
}