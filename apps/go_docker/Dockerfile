FROM golang AS base

FROM base AS build
WORKDIR /go_docker
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o cmd ./app
RUN ls -R

FROM scratch AS run
COPY --from=build /go_docker/cmd /go-docker
CMD ["./go-docker"]
