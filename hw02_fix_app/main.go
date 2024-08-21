package init

import (
	"fmt"
	"os"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
)

func init() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	_, err := fmt.Scanln(&path)
	if err != nil {
		fmt.Printf("Error reading data file: %s\n", err)
		os.Exit(1)
	}

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err := reader.ReadJSON(path)
	if err != nil {
		fmt.Printf("Error reading data file: %s\n", err)
		os.Exit(1)
	}

	printer.PrintStaff(staff)
}
