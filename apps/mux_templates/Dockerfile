FROM node:lts AS build-css
WORKDIR /mux_templates
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN make build-css
#RUN rm -rf node_modules
#RUN ls -R

FROM golang:1.23 AS build-go
WORKDIR /mux_templates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 make build-go
#RUN ls -R

FROM golang:1.23 AS dist
WORKDIR /mux_templates
COPY --from=build-go /mux_templates/build .
COPY --from=build-css /mux_templates/build .
#RUN ls -R

FROM scratch AS run
COPY --from=dist /mux_templates .
CMD ["./main"]