// Copyright © 2018 Gustavo Blanco gustavo.jr.blanco@gmail.com
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

	"github.com/spf13/cobra"
)

// vowelessCmd represents the voweless command
var vowelessCmd = &cobra.Command{
	Use:   "voweless",
	Short: "make a statement without vowls",
	RunE: func(cmd *cobra.Command, args []string) error {
		aestheticStatement, err := makeItVoweless(args)
		if err != nil {
			return err
		}

		fmt.Println(aestheticStatement)
		return nil
	},
}

func makeItVoweless(unaesthetic []string) (string, error) {
	if len(unaesthetic) == 0 {
		return "", fmt.Errorf("need statement to make voweless")
	}

	var voweless []rune
	for _, word := range unaesthetic {
		for _, letter := range word {
			switch letter {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				continue
			default:
				voweless = append(voweless, letter)
			}
		}
	}

	return string(voweless), nil
}

func init() {
	rootCmd.AddCommand(vowelessCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vowelessCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vowelessCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
