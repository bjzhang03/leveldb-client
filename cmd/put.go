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
	"strings"

	"github.com/bjzhang03/leveldb-cli/dboperation"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "put key,value into the leveldb",
	Long: `You can put the key,value into the db! 
Multiply (key,value) pair can by split by space For example:

	put one,one
	put one,one two,two`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Using config file:", viper.ConfigFileUsed())
		log.Info("The args is ", args)
		// 统一出错处理
		defer func() {
			recover := recover()
			if recover != nil {
				log.Errorf("Got error when put %s into db! Error :%s", args, recover)
			}
		}()
		for _, key := range args {
			params := strings.Split(key, ",")
			if len(params) == 2 {
				err := dboperation.Put(params[0], params[1], viper.GetString("db.path"))
				if err != nil {
					log.Error("Put the [ %s, %s] into db failed!", params[0], params[1])
				}
			} else {
				log.Errorf("The format of param [ %s ] is error!", key)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(putCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
