FROM golang:1.19-alpine as builder

WORKDIR /app
COPY . .

RUN go mod download

RUN go build -o "myprojectname" ./cmd

FROM alpine:3.16

WORKDIR /app

RUN apk update \
    && apk add --no-cache ca-certificates bash gcc \
    && update-ca-certificates --fresh \
    && rm -rf /var/cache/apk/*

RUN addgroup app_user && adduser -S app_user -u 1000 -G app_user

COPY --chown=app_user:app_user --from=builder /app/myprojectname .
COPY --chown=app_user:app_user --from=builder /app/.env .

# ignore if you turned off file logging
RUN mkdir logs
RUN chown app_user ./logs

RUN chmod +x ./myprojectname

USER app_user

EXPOSE 8080
ENTRYPOINT ["./myprojectname"]