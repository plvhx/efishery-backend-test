FROM golang:1.17

WORKDIR /usr/src/app

COPY ./golang-serv ./serv
COPY ./cache ./cache
COPY ./.jwt_pass ./.jwt_pass

RUN ls -al
RUN cd ./serv && go mod download && go mod verify
RUN cd ./serv && go build main.go

EXPOSE 9000/tcp

WORKDIR ./serv

CMD ["./main"]
