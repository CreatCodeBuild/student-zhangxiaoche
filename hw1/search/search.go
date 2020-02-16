package search

import (
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// Search x
func Search(logs io.ReadCloser, filter, separator string) (resultlist list.List) {

	contents, err := ioutil.ReadAll(logs)
	if string(contents) == "" {
		fmt.Println("未输入日志文件路径")
		return
	}
	if filter == "" {
		fmt.Println("未输入查找内容")
		return
	}
	if separator == "" {
		separator = "，"
		fmt.Println("未输入分隔符，默认按，分割行")
	}
	//fmt.Println(logs,filter,separator)
	fmt.Println("日志查找-结果：")
	if err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		//按分割符分割字符串
		section := strings.SplitAfter(result, separator)
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		for i := 0; i < len(section); i++ {
			if strings.Contains(section[i], filter) {
				fmt.Println("-------------------------------------")
				fmt.Println(section[i])
				resultlist.PushBack(section[i])
			}
		}
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}
	return
}
