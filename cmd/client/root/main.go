package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	MARK1_HOME_DIR := os.Getenv("MARK1_HOME_DIR")
	fmt.Println("MARK1_HOME_DIR:", MARK1_HOME_DIR)

	filePath := MARK1_HOME_DIR + "/role"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to read file: MARK1_HOME_DIR/role")
		fmt.Println("please check 'MARK1_HOME_DIR' environment variable")
		return
	}
	re := regexp.MustCompile(`\s+`)
	role := re.ReplaceAllString(string(content), "")

	switch role {
	case "admin":
		if err := exec.Command("bash", "-c", "/opt/mark1/admin-client").Run(); err != nil {
			fmt.Printf("Failed to run admin-client: %v\n", err)
		}
	case "app":
		if err := exec.Command("bash", "-c", "/opt/mark1/app-client").Run(); err != nil {
			fmt.Printf("Failed to run app-client: %v\n", err)
		}
	case "anonymous":
		if err := exec.Command("bash", "-c", "/opt/mark1/anonymous-client").Run(); err != nil {
			fmt.Printf("Failed to run anonymous-client: %v\n", err)
		}
	default:
		fmt.Println("Invalid role. You can specify 'admin', 'app', or 'anonymous' in MARK1_HOME_DIR/role")
	}
}
