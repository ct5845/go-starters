FROM golang:1.23 AS build-go
WORKDIR /insightful
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 make build-go
#RUN ls -R

FROM golang:1.23 AS dist
WORKDIR /insightful
COPY --from=build-go /insightful/build .
#RUN ls -R

FROM scratch AS run
COPY --from=dist /insightful .
CMD ["./main"]