FROM golang:1.22.5 as base

WORKDIR /app

COPY /go.mod ./
COPY /go.sum ./
COPY ./.env ./

RUN go mod download

COPY . .

RUN go build -o main .

# making a distro less image

FROM gcr.io/distroless/base

COPY --from=base /app/main .
COPY --from=base /app/.env .

EXPOSE 8080

CMD [ "./main" ]