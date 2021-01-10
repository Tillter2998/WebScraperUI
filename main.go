package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	var linkList []string

	for {
		fmt.Println("Web Scrapper GUI")
		fmt.Println("---------------------")

		fmt.Print("Enter URL: ")
		text, _ := reader.ReadString('\n')

		response, err := http.Get("http://localhost:7171/search?url=" + text)
		if err != nil {
			fmt.Printf("An error occured: %s", err)
			reader.ReadString('\n')
			continue
		}

		links, _ := ioutil.ReadAll(response.Body)
		_ = json.Unmarshal(links, &linkList)

		clearScreen()
		if strings.Contains(linkList[0], "Missing") {
			fmt.Println("No Links Found")
		} else {
			fmt.Println("Links Found")
		}
		fmt.Println("---------------------")

		for _, link := range linkList {
			fmt.Printf(string(link) + "\n")
		}
		reader.ReadString('\n')
		clearScreen()

	}

}
