# GO with Live Reload and Docker

Get a simple "Hello World" with a minimal docker build. Has AIR for live-reload when running the app locally.

  - [x] Go
  - [x] Live-Reload ([AIR](https://github.com/air-verse/air))
  - [x] Docker

### Run
The live-reload app, Builds to `./build`.
```bash
make run
```

### Build
Builds the app to `/build`.
```bash
make build
```

### Docker Build
Runs a new docker build.
```bash
make docker-build
# or
make docker-build FLAGS=--no-cache
```

### Preview
Runs the build, and then runs the app.
```bash
make preview
```