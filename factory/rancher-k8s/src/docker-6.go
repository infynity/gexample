package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"log"
)

type Empty struct{}

func main() {
	cli, err := client.NewClient("tcp://49.234.65.113:2345", "v1.41", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	config := &container.Config{
		WorkingDir: "/app",
		ExposedPorts: nat.PortSet{
			"80/tcp": Empty{},
		},
		Image: "alpine:3.12",
		Cmd:   []string{"./kkk"},
	}
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"80/tcp": []nat.PortBinding{
				{HostPort: "80"}, //宿主机的端口
			},
		},
		Binds: []string{"/root:/app"},
	}
	ret, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, "tes2")
	if err != nil {
		log.Fatal(err)
	}
	err = cli.ContainerStart(ctx, ret.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("容器启动成功,ID是:", ret.ID)

}
