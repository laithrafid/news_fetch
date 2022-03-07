# Dockerfile.deploy

FROM golang:1.17.6 as builder

ENV APP_USER app
ENV APP_HOME github.com/laithrafid/news_fetch/

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

WORKDIR $APP_HOME
USER $APP_USER
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o news_fetch

FROM debian:buster

ENV APP_USER app
ENV APP_HOME github.com/laithrafid/

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY --chown=0:0 --from=builder $APP_HOME/news_fetch $APP_HOME/news_fetch
COPY --chown=0:0 --from=builder $APP_HOME/app.env $APP_HOME/app.env

EXPOSE 8080
USER $APP_USER
CMD ["./news_fetch"]