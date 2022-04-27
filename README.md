GDT数据处理工具
==================

GDT实现了对数据的AES加密和解密功能, 并提供了将数据嵌入到PNG图片的功能.


### 基础指令

- 加密数据: gdt -e -k "password" < input.txt > output.txt
- 解密数据: gdt -d -k "password" < input.txt > output.txt
- 嵌入数据: gdt -m -i base.png < msg.txt
- 提取数据: gdt -x -i base.png > msg.txt
- 查看图片: gdt -s -i base.png
- 清除数据: gdt -c -i base.png

### 组合指令

- 加密并嵌入数据: gdt -e -k "password" < input.txt | gdt -m -i base.png
- 提取并解密数据: gdt -x -i base.png | gdt -d -k "password" > output.txt
