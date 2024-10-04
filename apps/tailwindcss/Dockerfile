FROM golang AS base

FROM base AS build
WORKDIR /go_docker_server
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cmd ./app
RUN ls -R

FROM scratch AS run
COPY --from=build /go_docker_server/cmd /go-docker-server
CMD ["./go-docker-server"]
