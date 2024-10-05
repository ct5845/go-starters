# MUX Templates

Here we're using the std go `mux` to do routing, and adding in `templates`
I add a basic "User List" with "User Detail" screen to allow a simple CRUD update.
This is "in memory" nothing gets actually saved.

- [x] Go Templates
- Live-Reload ([AIR](https://github.com/air-verse/air))
- Environment Variables ([DotEnv](https://github.com/joho/godotenv))
- Docker
- Tailwindcss

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