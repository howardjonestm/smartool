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
	"fmt"
	"os"

	"github.com/spf13/cobra"
    "github.com/olekukonko/tablewriter"
)

// fetchMarketsCmd represents the fetchMarkets command
var fetchMarketsCmd = &cobra.Command{
	Use:   "fetchMarkets",
	Short: "Retrieve markets for a given event",
	Long: `Provide event ID to retrieve markets`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fetchMarkets(args[0])
	},
}

func init() {
	rootCmd.AddCommand(fetchMarketsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchMarketsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchMarketsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


type Markets struct {
	Markets []Market `json:"markets"`
}

type Market struct {
	MarkType MartketType `json:"market_type"`
	BetDelay int `json:"bet_delay"`
	Category string `json:"category"`
	Complete bool `json:"complete"`
	ContractSelections string `json:"-"`
	Description string `json:"description"`
	DisplayOrder int `json:"display_order"`
	DisplayType string `json:"display_type"`
	EventId string `json:"event_id"`
	Hidden bool `json:"hidden"`
	Id string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	State string `json:"state"`
	WinnerCount int `json:"winner_count"`
}

type MartketType struct {
	Name string `json:"name"`
	Param string `json:"param"`
}

func fetchMarkets(eventId string){
	response := new(Markets)
	url := fmt.Sprintf("https://api.smarkets.com/v3/events/%s/markets/?sort=event_id%2Cdisplay_order&include_hidden=false", eventId)
    getJson(url, response)
	eventTable := [][]string{}
	
	fmt.Println(len(response.Markets))
    for i :=0; i<len(response.Markets); i++ {
        row := []string{
            response.Markets[i].Id,
            response.Markets[i].Name,
            response.Markets[i].Category,      
        }
        eventTable = append(eventTable,row)
    }

    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"Market Id", "Name", "Category"})
    table.SetBorder(false)                             
    table.AppendBulk(eventTable)                              
    table.Render()
}
