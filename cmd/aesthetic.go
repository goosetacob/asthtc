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

// aestheticCmd represents the aesthetic command
var aestheticCmd = &cobra.Command{
	Use:   "aesthetic",
	Short: "make a statement with ~ a e s t h e t i c ~",
	RunE: func(cmd *cobra.Command, args []string) error {
		aestheticStatement, err := makeItAesthetic(args)
		if err != nil {
			return err
		}

		fmt.Println(aestheticStatement)
		return nil
	},
}

func makeItAesthetic(unaesthetic []string) (string, error) {
	if len(unaesthetic) == 0 {
		return "", fmt.Errorf("need statement to make aesthetic")
	}
	aesthetic := []rune{'~', ' '}
	for _, word := range unaesthetic {
		for _, letter := range word {
			aesthetic = append(aesthetic, letter, ' ')
		}
	}
	aesthetic = append(aesthetic, '~')

	return string(aesthetic), nil
}

func init() {
	rootCmd.AddCommand(aestheticCmd)
}
