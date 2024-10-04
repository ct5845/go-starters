package main

import (
	"fmt"
	"go_livereload_dotenv/app/utils"
	"os"
)

func main() {
	utils.Env("APP_ENV")

	fmt.Println(utils.Colors["Blue"] + "ENV[APP_ENV]: " + utils.Colors["Reset"] + os.Getenv("APP_ENV"))
	fmt.Println(utils.Colors["Blue"] + "ENV[MYVARIABLE]: " + utils.Colors["Reset"] + os.Getenv("MYVARIABLE"))
}
