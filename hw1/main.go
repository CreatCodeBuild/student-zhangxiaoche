package main

import "fmt"
import "io/ioutil"
import "strings"
import "flag"

func main() {
    var  logs,filter,separator string //= "D:/MyProject/homework/testcase/ceshi.txt","傀儡","，"
    flag.StringVar(&logs,"logs","null","日志文件路径")
    flag.StringVar(&filter,"filter","傀儡","查找内容")
    flag.StringVar(&separator,"separator","\r","分隔符")
   //暂停获取参数
   flag.Parse()
    if logs == "null"{
     fmt.Println("未输入日志文件路径")
     return
    }
    if filter == "null"{
     fmt.Println("未输入查找内容")
     return
    }
    if separator == "\r"{
     fmt.Println("未输入分隔符，默认按，分割行")
    }
		//fmt.Println(logs,filter,separator)
        fmt.Println("日志查找-结果：")
        contents,err := ioutil.ReadFile(logs)
        if err == nil {
                    //因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
                    result := strings.Replace(string(contents),"\n","",1)
                    //按分割符分割字符串
                    section  := strings.SplitAfter(result, separator)
                    for i := 0; i< len(section); i ++ {
                        if strings.Contains(section[i], filter) {
                            fmt.Println("----------------------------------------------------------------------")
                            fmt.Println(section[i])
                        }
                    }
        }
  //Search("D:/MyProject/homework/test case/ceshi.txt","傀儡","，")
}
