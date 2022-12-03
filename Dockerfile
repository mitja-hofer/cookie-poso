FROM golang:1.19-alpine
WORKDIR /app

ENV COOKIE_DB_HOST=cookie-db.mysql.database.azure.com
ENV COOKIE_DB_USER=mitja
ENV COOKIE_DB_PASSWORD=Sivalni.Stroj
ENV TERM xterm-256color

COPY . ./
RUN go mod download

RUN go build -o bin/CookiePoso

EXPOSE 8080
#CMD ["/CookiePoso"]
ENTRYPOINT [ "/bin/sh"]