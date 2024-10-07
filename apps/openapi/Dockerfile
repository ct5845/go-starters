FROM golang:1.23 AS build-go
WORKDIR /open_api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 make build-go
#RUN ls -R

FROM golang:1.23 AS dist
WORKDIR /open_api
COPY --from=build-go /open_api/build .
#RUN ls -R

FROM scratch AS run
COPY --from=dist /open_api .
CMD ["./main"]