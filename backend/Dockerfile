FROM golang:1.18-alpine


#RUN apk update && apk add tzdata
#ENV TZ=Europe/Moscow

#RUN apk add --no-cache alpine-conf && \
    #setup-timezone -z Europe/Moscow


COPY ./ /go/src/goserver
WORKDIR /go/src/goserver

ENV GOPATH=/go/src/backend/.go
ENV GIN_MODE=release


#RUN go mod init

#COPY go.mod go.sum ./
RUN go mod download
RUN go build -o server .



EXPOSE 8081

CMD [ "./server" ]