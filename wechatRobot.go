/*
* @Author: bear
* @Date:   2018-10-20 23:19:09
* @Last Modified by:   bear
* @Last Modified time: 2018-10-20 23:19:18
 */

package main

import (
	"github.com/qianlnk/gobot"
)

func main() {
	cfg := gobot.Load()
	rebot, err := gobot.NewWecat(cfg)
	if err != nil {
		panic(err)
	}

	rebot.Start()
}
