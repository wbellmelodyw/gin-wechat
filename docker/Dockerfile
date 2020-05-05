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
#docker build -f="docker/Dockerfile" -t gin-wechat .