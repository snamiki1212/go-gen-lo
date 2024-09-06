/*
Copyright © 2024 snamiki1212

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// TODO:
// cmd --entity=User  --slice=UserVals --lo=Filter,Map --rename=Map:Loop    --no-extend
// cmd --entity=*User --slice=UserPtrs --lo=Filter,Map --rename=Map:Loop    --no-extend
type arguments struct {
	// Target entity name
	entity string

	// TODO:
	isPtrEntity bool

	// Target slice name
	slice string

	// Input file name
	input string

	// Output file name
	output string

	// Method name of lo to generate
	lo []string

	// Mapping field name to accessor name
	rename map[string]string // key: lo method name, value: generated method name.

	// noExtend []string
}

var rename []string

var args = arguments{
	// rename: map[string]string{},
}

func (a *arguments) loadRename(as []string) error {
	container := make([]error, 0)
	for _, ac := range as {
		pair := strings.Split(ac, ":")
		if len(pair) != 2 {
			container = append(container, fmt.Errorf("invalid rename: %s", ac))
			continue
		}
		src, dst := pair[0], pair[1]
		args.rename[src] = dst
	}
	if len(container) != 0 {
		return fmt.Errorf("%v", container)
	}
	return nil
}

func loader() error {
	// Load arguments
	if err := args.loadRename(rename); err != nil {
		return fmt.Errorf("load accessor error: %w", err)
	}
	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-gen-lo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	RunE: func(cmd *cobra.Command, _ []string) error {
		// Load arguments
		if err := loader(); err != nil {
			return fmt.Errorf("loader error: %w", err)
		}

		// Parse source code
		dt, err := parse(args, reader)
		if err != nil {
			return fmt.Errorf("parse error: %w", err)
		}

		// Generate code
		txt, err := generate(dt)
		if err != nil {
			return fmt.Errorf("generate error: %w", err)
		}

		// Write to output file
		err = write(args.output, txt)
		if err != nil {
			return fmt.Errorf("write error: %w", err)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	// entity
	rootCmd.Flags().StringVarP(&args.entity, "entity", "e", "", "target entity name")
	_ = rootCmd.MarkFlagRequired("entity")

	// slice
	rootCmd.Flags().StringVarP(&args.slice, "slice", "s", "", "target slice name")
	_ = rootCmd.MarkFlagRequired("slice")

	// input
	rootCmd.Flags().StringVarP(&args.input, "input", "i", "", "input file name")
	_ = rootCmd.MarkFlagRequired("input")

	// output
	rootCmd.Flags().StringVarP(&args.output, "output", "o", "", "output file name")
	_ = rootCmd.MarkFlagRequired("output")

	// // rename
	// rootCmd.Flags().StringSliceVarP(&args.rename, "rename", "r", []string{}, "rename method e.g. --rename=Map:Loop")
}
