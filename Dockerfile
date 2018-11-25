FROM golang:1.10.1

RUN apt-get update

# install and configure Redis
WORKDIR /tmp
RUN apt-get install -y build-essential tcl
RUN curl -O http://download.redis.io/redis-stable.tar.gz
RUN tar xzvf redis-stable.tar.gz
WORKDIR /tmp/redis-stable
RUN make
RUN make install
RUN mkdir /etc/redis
RUN cp /tmp/redis-stable/redis.conf /etc/redis

# install and configure geniuslurker
WORKDIR /go/src/github.com/AlexanderYAPPO/geniuslurker
COPY . .
RUN go get ./...
WORKDIR /go/bin
RUN go build /go/src/github.com/AlexanderYAPPO/geniuslurker/geniuslurker_apps/telegram_bot/telegram_bot.go
ADD ./geniuslurker.yaml /etc/geniuslurker.yaml

# install and configure supervisor
RUN apt-get install -y supervisor
ADD ./supervisord.conf /etc/supervisord.conf

CMD ["supervisord", "-n"]