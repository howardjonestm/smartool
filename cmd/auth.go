/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Write a token to current directory",
	Long:  `Run "smartool [email] [password]" to write a local token`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		execute(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Response struct {
	Factor string `json:"factor"`
	Stop   string `json:"stop"`
	Token  string `json:"token"`
	Verify bool   `json:"verify"`
}

func execute(email, password string) {

	token := retrieveToken(email, password)
	writeToFile(token)
}

func retrieveToken(username, password string) string {

	url := "https://api.smarkets.com/v3/sessions"

	requestBody := fmt.Sprintf(`{"password": "%s","remember": true,"username": "%s"}`, password, username)

	resp := httpGet(url, requestBody)

	response := new(Response)
	reader := strings.NewReader(string(resp))

	json.NewDecoder(reader).Decode(response)

	return response.Token
}

func httpGet(url, requestBody string) []byte {

	requestBodyByte := []byte(requestBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyByte))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	checkErr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	return body

}

func writeToFile(token string) {
	homeDir := os.Getenv("HOME")
	directory := fmt.Sprintf("%s/.SMARKETS_TOKEN", homeDir)
	os.Remove(directory)

	file, err := os.OpenFile(directory, os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)

	file.WriteString(token)
	file.Sync()
}

func checkErr(e error) {
	if e != nil {
		log.Println(e)
	}
}
