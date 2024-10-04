# GO Http Server Live Reload and Environment Variables and Docker steps

This takes the previous 2 examples [go_server](../go_server) and [go_docker](../go_docker) and combines them into one project.

  - [x] Go
  - [x] Live-Reload ([AIR](https://github.com/air-verse/air))
  - [x] Environment Variables ([DotEnv](https://github.com/joho/godotenv))
  - [x] Http Server
  - [x] Docker


### Install
```bash
make install
```

### Run
The live-reload app, defaults to `development` environment. Builds to `./build`.
```bash
make run
```

### Build
Builds to app, and copes environment files to `/build`
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

### Docker Run
Runs a new docker build and then runs the container
```bash
make docker-run
```

### Preview
Runs the build, and then runs the app, defaults to `production` environment.
```bash
make preview
```