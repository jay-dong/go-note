package main
import(
	"fmt"
	"strings"
)

func main(){

	//判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）是否相同
	fmt.Println(strings.EqualFold("董","dong"))
	//字符串前缀判断
	fmt.Println(strings.HasPrefix("hello world","hello"))
	fmt.Println(strings.HasPrefix("hello world","world"))
	//字符串后缀判断
	fmt.Println(strings.HasSuffix("dongjiebo","dong"))
	fmt.Println(strings.HasSuffix("dongjiebo","jiebo"))
	//判断字符串是否包含指定字符串
	fmt.Println("判断字符串是否包含指定字符串",strings.Contains("xiaobaibai","i"))
	//判断字符串utf8码值
	fmt.Println("判断字符串utf8码值",strings.ContainsRune("xiaobai",8))
	//判断字符串是否包含指定字符串中任意一个字符
	fmt.Println("判断字符串是否包含指定字符串中任意一个字符",strings.ContainsAny("xiaobaibai","dong"))
	fmt.Println("判断字符串是否包含指定字符串中任意一个字符",strings.ContainsAny(" "," lll"))
	//返回字符串中指定字符串的数量 若为空值，则返回当前字符串的数量
	fmt.Println("返回字符串中指定字符串的数量",strings.Count("cheeseee","e"))
	fmt.Println("返回字符串中指定字符串的数量",strings.Count("helloll","he"))
	//返回字符串中指定字符串首次出现的位置,如果不存在则返回-1
	fmt.Println("返回字符串中指定字符串首次出现的位置",strings.Index("xiaobai","ao"))
	fmt.Println("返回字符串中指定字符串首次出现的位置",strings.Index("xiaobai","p"))
	//返回字符串中指定字符串最后一次出现的位置,如果不存在则返回-1
	fmt.Println("返回字符串中指定字符串最后一次出现的位置",strings.LastIndex("xiaobai","a"))
	fmt.Println("返回字符串中指定字符串最后一次出现的位置",strings.LastIndex("xiaobai","p"))
	//返回字符串指定字符串unicode码值出现的位置
	fmt.Println(strings.IndexRune("xiaobaibai",1))


}