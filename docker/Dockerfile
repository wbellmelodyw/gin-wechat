FROM golang:alpine
WORKDIR /app
COPY release/gin-wechat /app/
COPY gin.toml /app/
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
RUN apk add --no-cache gcc musl-dev git
#ADD gotify-app /app/
EXPOSE 80
ENTRYPOINT ["./gin-wechat"]
#go build -ldflags '-linkmode external -w -extldflags "-static"' -o release/gin-wechat
#docker build -f="docker/Dockerfile" -t gin-wechat .
#docker run -d -t -i -p 80:80 gin-wechat:latest  #docker-compose后台运行
#清理无用容器
#清理无用镜像