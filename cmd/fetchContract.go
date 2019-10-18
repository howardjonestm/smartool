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

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// fetchContractCmd represents the fetchContract command
var fetchContractCmd = &cobra.Command{
	Use:   "fetchContract",
	Short: "Fetch contract",
	Long:  `Fetch contract for a given market ID`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fetchContract(args[0])
	},
}

func init() {
	rootCmd.AddCommand(fetchContractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchContractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchContractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Contracts struct {
	Contracts []Contract `json:"contracts"`
}

type Contract struct {
	ContType         ContractType `json:"contract_type"`
	DisplayOrder     string       `json:"display_order"`
	Hidden           bool         `json:"hidden"`
	Id               string       `json:"id"`
	Info             string       `json:"-"`
	MarketId         string       `json:"market_id"`
	Name             string       `json:"name"`
	OutcomeTimestamp string       `json:"outcome_timestamp"`
	Slug             string       `json:"string"`
	StateOrOutcome   string       `json:"state_or_outcome"`
}

type ContractType struct {
	Name  string `json:"name"`
	Param string `json:"param"`
}

func fetchContract(marketId string) {
	response := new(Contracts)
	url := fmt.Sprintf("https://api.smarkets.com/v3/markets/%s/contracts/", marketId)
	getJson(url, response)
	contractTable := [][]string{}

	fmt.Println(len(response.Contracts))
	for i := 0; i < len(response.Contracts); i++ {
		row := []string{
			response.Contracts[i].ContType.Name,
			response.Contracts[i].Id,
			response.Contracts[i].Name,
		}
		contractTable = append(contractTable, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"type", "Id", "name"})
	table.SetBorder(false)
	table.AppendBulk(contractTable)
	table.Render()
}
