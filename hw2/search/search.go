package search

import (
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func Search(logs io.ReadCloser, separator string, filter StrFilter) (resultlist list.List) {

	contents, err := ioutil.ReadAll(logs)
	if string(contents) == "" {
		fmt.Println("未输入日志文件路径")
		return
	}
	if (filter.addList == nil) && (filter.orList == nil) {
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
			fmt.Println(section[i])
			if(filter.testSubstring(section[i])){
				//通过过滤器条件
				resultlist.PushBack(section[i])
			}
			
		}
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}
	return
}
/////////////////////////////////////////////////////////////////
type StrFilter struct {
	var addList list.List
	var orList list.List
}

func Contain(substring string) condition {
	return func(input string) bool {
		return strings.Contains(input, substring)
	}
}

func NotContain(substring string) condition {
	return func(input string) bool {
		return !strings.Contains(input, substring)
	}
}
//条件通过为true
type condition func(input string)bool
//添加条件函数到与列表返回过滤器
func (sf *StrFilter) add(c condition) *StrFilter{\
	if(sf.addList == nil){
		sf.addList = list.New()
	}
	sf.addList.PushBack(c)
    return sf
}
//添加条件函数到或列表返回过滤器
func (sf StrFilter) or(c condition) *StrFilter{
	if(sf.List == nil){
		sf.orList= list.New()
	}
	sf.orList.PushBack(c)
	return sf
}
//检测传入值是否通过过滤器
func (sf StrFilter) testSubstring(str string) bool{
	addListTest := true
	orListTest := false

	for a := 0; a < len(sf.addList); a++ {
		cc = sf.addList.Front()
		if(!cc(str)){
			//如果不满足条件立即退出循环
			addListTest = false
			break
		}
		
	}
	for o := 0; o < len(sf.orList); o++ {
		cc = sf.orList.Front()
		cc(str)
		if(cc(str)){
			//如果满足条件一次立即退出循环
			orListTest = true
			break 
		}
	}
	return addListTest||orListTest
}