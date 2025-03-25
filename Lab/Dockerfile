# Чтобы не компилировать файл в контейнере и не качать зависимости заново
# проще делать это на хосте через PowerShell и компилятор Golang'a
# $env:GOOS="linux"
# $env:GOARCH="amd64"
# go build -o lab.sh .

# Потом - копируем исполняемый файл в наш контейнер с приложением

# Приложение в контейнере с портом 
FROM alpine:3.21
ARG APP_SH=lab.sh
WORKDIR /myapp
COPY ${APP_SH} ./app
EXPOSE 6080
ENTRYPOINT ["/myapp/app"]