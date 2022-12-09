FROM golang:1.19
#토커 컨테이너 안에 
WORKDIR /server/
COPY ./go.mod /server/
COPY ./go.sum /server/
RUN go mod tidy

COPY . /server/

CMD go run ./cmd
