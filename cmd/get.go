/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/bjzhang03/leveldb-cli/dboperation"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the value by the key",
	Long: `Get key from leveldb by the key from the given path! 
Multiply key can be split by space For example: 

	get one two three
	get - (get all the data in leveldb)`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Using config file:", viper.ConfigFileUsed())
		log.Info("The args is ", args)
		// 统一出错处理
		defer func() {
			recover := recover()
			if recover != nil {
				log.Errorf("Got error when get value from db! Error :%s", recover)
			}
		}()
		if len(args) == 1 && args[0] == "-" {
			result := dboperation.GetAll(viper.GetString("db.path"))
			log.Infof("Got all the data! %s ", result)
		} else {
			for _, key := range args {
				result := dboperation.Get(key, viper.GetString("db.path"))
				log.Infof("Get the result [%s, %s]", key, result)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
