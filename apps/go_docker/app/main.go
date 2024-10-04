package main

import (
	"fmt"
	"go_docker/app/utils"
)

func main() {
	fmt.Printf("%sHello %sW%so%sr%sl%sd\n",
		utils.Colors["Blue"],
		utils.Colors["Red"],
		utils.Colors["Yellow"],
		utils.Colors["Green"],
		utils.Colors["Magenta"],
		utils.Colors["Cyan"])
}
