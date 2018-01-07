FROM golang
RUN git clone https://github.com/bwmarrin/discordgo /go/src/github.com/bwmarrin/discordgo
RUN git clone https://github.com/gorilla/websocket /go/src/github.com/gorilla/websocket
RUN go get golang.org/x/crypto/nacl/secretbox
COPY ./app/src/ /home/root/
RUN go build -a /home/root/Bot.go
RUN chmod a+x /go/Bot
CMD ["/go/Bot"]
