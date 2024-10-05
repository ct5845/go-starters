FROM node:lts AS build-css
WORKDIR /go_tailwindcss
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN make build-css
#RUN rm -rf node_modules
#RUN ls -R

FROM golang:1.23 AS build-go
WORKDIR /go_tailwindcss
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 make build-go
#RUN ls -R

FROM golang:1.23 AS dist
WORKDIR /go_tailwindcss
COPY --from=build-go /go_tailwindcss/build .
COPY --from=build-css /go_tailwindcss/build .
#RUN ls -R

FROM scratch AS run
COPY --from=dist /go_tailwindcss .
CMD ["./main"]