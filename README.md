GDT数据处理工具
==================

### 基础指令

加密数据: gdt -e -i input.txt -o output.txt -k "password"
解密数据: gdt -d -i input.txt -o output.txt -k "password"
嵌入数据: gdt -m -i base.png -g msg.txt
提取数据: gdt -x -i base.png -g msg.txt
查看图片: gdt -s -i base.png

### 组合指令

加密并嵌入数据: gdt -e -m -i base.png -g msg.txt -k "password"
提取并解密数据: gdt -x -d -i base.png -g msg.txt -k "password"
