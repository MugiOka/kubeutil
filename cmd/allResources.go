/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"os/exec"

	"github.com/mattn/go-pipeline"
	"github.com/spf13/cobra"
)

type getOptions struct {
	namespace string
}

func newAllResourcesCmd() *cobra.Command {
	var (
		o = &getOptions{}
	)
	cmd := &cobra.Command{
		Use:   "all-resources",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			runAllResourcesCmd(o)
		},
	}
	cmd.Flags().StringVarP(&o.namespace, "namespace", "n", "", "Get resources in a specific namespace.")
	return cmd
}

func runAllResourcesCmd(o *getOptions) {
	if o.namespace != "" {
		getAllResourcesIn(o.namespace)
	} else {
		getAllResourcesInAllNamespace()
	}
}

func getAllResourcesInAllNamespace() {
	out1, err1 := pipeline.Output(
		[]string{"kubectl", "api-resources", "--namespaced=true", "--verbs=list", "-o=name"},
		[]string{"tr", "\n", ","},
		[]string{"sed", "-e", "s/,$//g"},
	)
	if err1 != nil {
		log.Fatal(err1)
	}
	out2, err2 := exec.Command("kubectl", "get", string(out1), "-A").Output()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("%s", out2)
}

func getAllResourcesIn(n string) {

	out1, err1 := pipeline.Output(
		[]string{"kubectl", "api-resources", "--namespaced=true", "--verbs=list", "-o=name"},
		[]string{"tr", "\n", ","},
		[]string{"sed", "-e", "s/,$//g"},
	)
	if err1 != nil {
		log.Fatal(err1)
	}
	out2, err2 := exec.Command("kubectl", "get", string(out1), "-n", n).Output()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("%s", out2)
}

func init() {
	getCmd.AddCommand(newAllResourcesCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allResourcesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allResourcesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
