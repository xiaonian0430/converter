//Author:xiaonian0430
//Email:xiaonian0430@163.com
//Date:2020-05-29
package main

import (
	"os"

	"fmt"

	"github.com/xiaonian0430/converter/converter"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("错误：缺少指定的json配置文件")
	} else {
		if converter, err := converter.NewConverter(args[0]); err != nil {
			fmt.Println(err.Error())
		} else {
			if err = converter.Convert(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
