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
	"encoding/json"
	"net/http"
	"os"
	"time"

	// "fmt"
	// "io/ioutil"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// fetchFootballCmd represents the fetchFootball command
var fetchFootballCmd = &cobra.Command{
	Use:   "fetchEvents",
	Short: "Retrieve upcoming football markets",
	Long: `
	This is a first attempt at using the tool to retrieve upcoming football markets, it can be expected that this tool will be deprecated in the future
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fetchEvents()
	},
}

func init() {
	rootCmd.AddCommand(fetchFootballCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchFootballCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchFootballCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Events struct {
	Events     []Event `json:"events"`
	Pagination string  `json:"-"`
}


type Event struct {
	Bettable      bool      `json:"bettable"`
	Created       string    `json:"created"`
	Description   string    `json:"description"`
	DisplayOrder  int       `json:"display_order"`
	EndDate       string    `json:"end_date"`
	FullSlug      string    `json:"full_slug"`
	Hidden        bool      `json:"hidden"`
	Id            string    `json:"id"`
	InplayEnabled bool      `json:"inplay_enabled"`
	Modified      string    `json:"modified"`
	Name          string    `json:"name"`
	ParentId      string    `json:"parent_id"`
	ShortName     string    `json:"short_name"`
	Slug          string    `json:"slug"`
	SpecialRules  string    `json:"special_rules"`
	StartDate     string    `json:"start_date"`
	StartDateTime string    `json:"start_datetime"`
	State         string    `json:"upcoming"`
	EventType     EventType `json:"type"`
}

type EventType struct {
	Domain string `json:"domain"`
	Scope  string `json:"scope"`
}

func fetchEvents() {
	response := new(Events)
	getJson("https://api.smarkets.com/v3/events/?state=new&state=upcoming&state=live&type=football_match&type_domain=football&type_scope=single_event&with_new_type=true&sort=start_datetime%2Cid&limit=100&include_hidden=false", response)

	eventTable := [][]string{}

	for i := 0; i < len(response.Events); i++ {
		row := []string{
			response.Events[i].Id,
			response.Events[i].Name,
			response.Events[i].StartDateTime,
		}
		eventTable = append(eventTable, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Event Id", "Name", "Start Time"})
	table.SetBorder(false)
	table.AppendBulk(eventTable)
	table.Render()
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}