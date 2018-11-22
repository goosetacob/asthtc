// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/goosetacob/asthtc/backend"
	"github.com/spf13/cobra"
)

// deBruijnCmd represents the deBruijn command
var deBruijnCmd = &cobra.Command{
	Use:   "deBruijn",
	Short: "computes the de bruijn sequence given alphabet k and sub-sequence size n",
	RunE: func(cmd *cobra.Command, args []string) error {
		alphabet, err := cmd.Flags().GetString("alphabet")
		if err != nil {
			return err
		}
		subSequenceSize, err := cmd.Flags().GetInt("subSequenceSize")
		if err != nil {
			return err
		}
		deBruijnSequence, err := tool.MakeItDeBruijn(alphabet, subSequenceSize)
		if err != nil {
			return err
		}

		fmt.Println(deBruijnSequence)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deBruijnCmd)
	deBruijnCmd.Flags().StringP("alphabet", "a", "", "alphabet for de bruijn sequence")
	deBruijnCmd.MarkFlagRequired("alphabet")
	deBruijnCmd.Flags().IntP("subSequenceSize", "s", 0, "size of subsequences in de bruijn sequence")
	deBruijnCmd.MarkFlagRequired("subSequenceSize")
}
