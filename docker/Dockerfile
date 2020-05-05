FROM golang:alpine
WORKDIR /app
COPY . /app/
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
#RUN apk add --no-cache gcc musl-dev git && go build -v
#ADD gotify-app /app/
#EXPOSE 80
#ENTRYPOINT ["./gin-wechat"]