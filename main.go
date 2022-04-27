package main

import (
	"flag"
	"fmt"
)

var (
	input   string
	key     string
	encode  bool
	decode  bool
	embed   bool
	extract bool
	clear   bool
	show    bool
)

func init() {
	flag.StringVar(&input, "i", "", "输入文件的名称")
	flag.StringVar(&key, "k", "", "AES算法加密和解密的密钥")
	flag.BoolVar(&encode, "e", false, "执行AES加密操作")
	flag.BoolVar(&decode, "d", false, "执行AES解密操作")
	flag.BoolVar(&embed, "m", false, "执行PNG嵌入操作")
	flag.BoolVar(&extract, "x", false, "执行PNG提取操作")
	flag.BoolVar(&clear, "c", false, "清除图片中嵌入的数据")
	flag.BoolVar(&show, "s", false, "显示PNG结构信息")

	flag.CommandLine.Usage = func() {
		fmt.Println("gdt: GDT数据处理工具")
		fmt.Println("基础指令:")
		fmt.Println("\t加密数据: gdt -e -k \"password\" < input.txt > output.txt")
		fmt.Println("\t解密数据: gdt -d -k \"password\" < input.txt > output.txt")
		fmt.Println("\t嵌入数据: gdt -m -i base.png < msg.txt")
		fmt.Println("\t提取数据: gdt -x -i base.png > msg.txt")
		fmt.Println("\t查看图片: gdt -s -i base.png")
		fmt.Println("\t清除数据: gdt -c -i base.png")
		fmt.Println("组合指令:")
		fmt.Println("\t加密并嵌入数据: gdt -e -k \"password\" < input.txt | gdt -m -i base.png")
		fmt.Println("\t提取并解密数据: gdt -x -i base.png | gdt -d -k \"password\" > output.txt")
		fmt.Println("补充说明:")
		fmt.Println("\t可以直接使用终端输入和输出结果 Windows平台输入Ctrl+Z Linux平台输入Ctrl+D 表示EOF")

	}
}

func main() {
	flag.Parse()

	if encode {
		Encode(key)
	}

	if decode {
		Decode(key)
	}

	if embed {
		Embed(input)
	}

	if extract {
		Extract(input)
	}

	if show {
		Show(input)
	}

	if clear {
		Clear(input)
	}

}
