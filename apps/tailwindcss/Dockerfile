FROM node AS node-build
WORKDIR /go_tailwindcss
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npx tailwindcss -i app/static/styles.css -o static/styles.min.css --minify

FROM golang AS build
WORKDIR /go_tailwindcss
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=node-build /go_tailwindcss/static/. /go_tailwindcss/build/static/
RUN CGO_ENABLED=0 GOOS=linux go build -o build/cmd ./app
RUN ls -R

# swap this out to FROM scratch
FROM build AS run
COPY --from=build /go_tailwindcss/build .
RUN ls -R
CMD ["./cmd"]
