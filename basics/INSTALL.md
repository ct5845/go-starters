# Development Environment

You'll need/probably want some basic things installed to even begin.

## Installs

- **[Go](https://go.dev/doc/install)**
- **[Powershell](https://aka.ms/powershell-release?tag=lts)**
- **[MAKE](https://sourceforge.net/projects/gnuwin32/)**

## Downloads

- **[Tailwind](https://github.com/tailwindlabs/tailwindcss/releases/latest)**
  ```bash
  $ curl -L --ssl-no-revoke https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-windows-x64.exe --output scripts/cmd/tailwindcss.exe
  ```
  Replacing `tailwindcss-windows-x64.exe` with the correct version for your machine.

## Environment Variables

View [Powershell](../scripts/powershell/README.md) to find out more, but run this script to set-up from helpfull system environment paths.

This adds;

- `touch`
- `make`

To the windows terminal commands available (to match/sync with Linux)

```cmd
.\scripts\powershell\env-vars.ps1
```
