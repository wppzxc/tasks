## build
### backend
```
# docker run --rm  -e GO111MODULE=off -v $PWD:/go/src/github.com/wppzxc/tasks golang:1.13.4 go build -o /go/src/github.com/wppzxc/tasks/src/build/docker/tasks /go/src/github.com/wppzxc/tasks/src/main.go
# docker build -t wppzxc/tasks:v1.0.0 .
```
### front
```
# npm run build
```