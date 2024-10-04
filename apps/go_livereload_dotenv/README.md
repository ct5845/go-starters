# Vanilla GO with Live Reload and Environment Variables

I consider live-reloading to be a baked in neccesity for web (any?)development. I also consider being able to load a .env file crucial for later builds, so these are the only two extra dependencies brought in.

  - [x] Go
  - [x] Live-Reload ([AIR](https://github.com/air-verse/air))
  - [x] Environment Variables ([DotEnv](https://github.com/joho/godotenv))


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