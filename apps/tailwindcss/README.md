# Tailwindcss

This project essentially adds Tailwindcss to the project, but it covers a bit more ground then that;
(1) Building the tailwindcss assets, in dev (with live-reloading covered), in prod (minified).
(2) Serving via the embed, and how to deal with static routes.
(3) Refactoring Makefile to use scripts, bit more flexibility.

- Go
- Live-Reload ([AIR](https://github.com/air-verse/air))
- Environment Variables ([DotEnv](https://github.com/joho/godotenv))
- Http Server
- Docker
- [x] Tailwindcss
- [x] GO Embed & Static Routes


### Install
```bash
make install
```

### Run
The live-reload app, defaults to `dev` environment. Builds to `/tmp`.
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
```

### Docker Run
Runs a new docker build and then runs the container
```bash
make docker-run
```

### Preview
Runs the build, and then runs the app, defaults to `prod` environment.
```bash
make preview
```