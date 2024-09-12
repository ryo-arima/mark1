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
		cmd := exec.Command("bash", "-c", "/opt/mark1/admin-client")
		output, err := cmd.CombinedOutput() // 標準出力と標準エラーをキャプチャ
		if err != nil {
			fmt.Printf("Failed to run admin-client: %v\n", err)
			fmt.Printf("Output: %s\n", output)
		}
	case "app":
		cmd := exec.Command("bash", "-c", "/opt/mark1/app-client")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to run app-client: %v\n", err)
			fmt.Printf("Output: %s\n", output)
		}
	case "anonymous":
		cmd := exec.Command("bash", "-c", "/opt/mark1/anonymous-client")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to run anonymous-client: %v\n", err)
			fmt.Printf("Output: %s\n", output)
		}
	default:
		fmt.Println("Invalid role. You can specify 'admin', 'app', or 'anonymous' in MARK1_HOME_DIR/role")
	}
}
