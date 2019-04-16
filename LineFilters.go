package main 

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		upperCase := strings.ToUpper(scanner.Text())
		fmt.Fprintf(os.Stdout,"%s\n",upperCase)
	}

	if err := scanner.Err() ; err != nil {
		fmt.Fprintf(os.Stderr,"%s",err)
	}
	
}

// 1. go 读取文件需要先创建scanner，scanner需要指定文件描述符，这里是标准输入
// 2. 字符串操作需要strings包，这里用到了strings包的ToUpper函数，用于将小写字母转成大写字母
// 3. fmt.Fprintf 可以向指定的文件描述符打印
// 4. 捕获scanner的错误可以使用scanner.Err()函数
// 5. scanner.Scan()函数可以扫描输入文件描述符，外接一个for循环表示不停的扫描，scanner.Text() 可以获取扫描的结果

