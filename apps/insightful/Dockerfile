FROM golang:1.23 AS build-go
WORKDIR /insightful
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build-go
#RUN ls -R

FROM golang:1.23 AS dist
WORKDIR /insightful
COPY --from=build-go /insightful/build .
#RUN ls -R

FROM gcr.io/distroless/base AS run
COPY --from=dist /insightful .
CMD ["./main"]