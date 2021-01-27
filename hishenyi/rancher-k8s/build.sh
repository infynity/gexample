docker run   --rm  \
-v /Users/infy/Desktop/go_project/goexample/factory/rancher-k8s:/app \
-v /Users/infy/Desktop/go_project/goexample/hishenyi:/go \
-w /app/src \
-e GOOS=darwin \
-e GOPROXY=https://goproxy.cn \
golang:1.14.4-alpine3.12 \
go build -o ../888 main.go