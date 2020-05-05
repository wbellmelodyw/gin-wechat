FROM golang:alpine
WORKDIR /app
COPY release/gin-wechat /app/
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
RUN apk add --no-cache gcc musl-dev git
#ADD gotify-app /app/
EXPOSE 80
#ENTRYPOINT ["./gin-wechat"]
#standard_init_linux.go:211: exec user process caused "no such file or directory"