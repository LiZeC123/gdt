package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	input   string
	output  string
	msg     string
	key     string
	encode  bool
	decode  bool
	embed   bool
	extract bool
	show    bool
)

func init() {
	flag.StringVar(&input, "i", "", "输入文件的名称")
	flag.StringVar(&output, "o", "", "输出文件的名称")
	flag.StringVar(&msg, "g", "", "数据文件的名称")
	flag.StringVar(&key, "k", "", "AES算法加密和解密的密钥")
	flag.BoolVar(&encode, "e", false, "执行AES加密操作")
	flag.BoolVar(&decode, "d", false, "执行AES解密操作")
	flag.BoolVar(&embed, "m", false, "执行PNG嵌入操作")
	flag.BoolVar(&extract, "x", false, "执行PNG提取操作")
	flag.BoolVar(&show, "s", false, "显示PNG结构信息")

	flag.CommandLine.Usage = func() {
		fmt.Println("gdt: GDT数据处理工具")
		fmt.Println("基础指令:")
		fmt.Println("加密数据: gdt -e -i input.txt -o output.txt -k \"password\"")
		fmt.Println("解密数据: gdt -d -i input.txt -o output.txt -k \"password\"")
		fmt.Println("嵌入数据: gdt -m -i base.png -g msg.txt")
		fmt.Println("提取数据: gdt -x -i base.png -g msg.txt")
		fmt.Println("查看图片: gdt -s -i base.png")
		fmt.Println("组合指令:")
		fmt.Println("加密并嵌入数据: gdt -e -m -i base.png -g msg.txt -k \"password\"")
		fmt.Println("提取并解密数据: gdt -x -d -i base.png -g msg.txt -k \"password\"")
	}
}

func main() {
	flag.Parse()

	if encode && embed {
		tmp := msg + ".tmp"
		Encode(msg, key, tmp)
		Embed(input, tmp, input)
		_ = os.Remove(tmp)
		return
	}

	if extract && decode {
		tmp := msg + ".tmp"
		Extract(input, tmp, input)
		Decode(tmp, key, msg)
		err := os.Remove(tmp)
		if err != nil {
			panic(err)
		}
		return
	}

	if encode {
		Encode(input, key, output)
	}

	if decode {
		Decode(input, key, output)
	}

	if embed {
		Embed(input, msg, input)
	}

	if extract {
		Extract(input, msg, input)
	}

	if show {
		Show(input)
	}

}
