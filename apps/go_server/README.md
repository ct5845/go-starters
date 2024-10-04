# Vanilla GO Http Server Live Reload and Environment Variables

This will run up a very basic http server with two routes. Preview shows how it can load a different PORT based on the `.env` file it picks.

  - [x] Go
  - [x] Live-Reload ([AIR](https://github.com/air-verse/air))
  - [x] Environment Variables ([DotEnv](https://github.com/joho/godotenv))
  - [x] Http Server


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

### Preview
Runs the build, and then runs the app, defaults to `production` environment.
```bash
make preview
```