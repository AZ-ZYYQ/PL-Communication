---
时间：2022-11-05  23:21
标签： #
钩子： 概述：

---
##### 问题
参考：[vim编辑器配置](https://blog.csdn.net/Liukairui/article/details/107392243)
1. 解决在linux上使用gvim显示cannot open display的问题；
问题链接：["E233: cannot open display" when starting gvim on root shell (linuxquestions.org)](https://www.linuxquestions.org/questions/ubuntu-63/e233-cannot-open-display-when-starting-gvim-on-root-shell-751248/)
解决方法：
在终端输入
`export DISPLAY=':0.0'`

2. vim配置大全
[(15条消息) Vim使用全指南 (环境配置，插件推荐，美化) (C++,Python,MarkDown,R...)_Liukairui的博客-CSDN博客](https://blog.csdn.net/Liukairui/article/details/107392243)

3. 未知配置
" main函数生成快捷方式
```.vimrc
map mf i#include <stdio.h><Esc>o<Esc>oint main(void)<Esc>o{<Esc>oreturn 0;<Esc>o}<Esc>2ko
```

4. gvim打开文件，出现# E576: Error while reading ShaDa file: there is an item at position 270498 that must not be there: Missing itemsare for internal uses only #6875
解决方法：进入到`~/.local/nvim/share/Shada`，删除里面所有的文件。
参考链接：[E576：读取 ShaDa 文件时出错：位置 270498 处有一个项目不能在那里：缺少的项目仅供内部使用 ·问题 #6875 ·尼奥维姆/尼奥维姆 (github.com)](https://github.com/neovim/neovim/issues/6875)

5. 解决vim编辑器中不同模式下光标显示相同的问题：
参考链接：[How to change the cursor between Normal and Insert modes in Vim? - Stack Overflow](https://stackoverflow.com/questions/6488683/how-to-change-the-cursor-between-normal-and-insert-modes-in-vim)
- 确定teminal的类型
- 在vim的配置文件`.vimrc`中加入如下命令
```
let &t_SI = "\e[6 q"
let &t_EI = "\e[2 q"

" reset the cursor on start (for older versions of vim, usually not required)
augroup myCmds
au!
autocmd VimEnter * silent !echo -ne "\e[2 q"
augroup END
```
Other options (replace the number after `\e[`):
```
Ps = 0  -> blinking block.
Ps = 1  -> blinking block (default).
Ps = 2  -> steady block.
Ps = 3  -> blinking underline.
Ps = 4  -> steady underline.
Ps = 6  -> steady bar (xterm)
```


---
##### 模式切换
正常模式(normal)、命令模式、插入模式、替换模式、可视化模式
-   正常模式：按`Esc` 键。
-   命令模式：正常模式下，按`:` 键，vim的左下角会有`-- INSERT --`字样。
-   插入模式：正常模式下
	`i`：在当前字符前进入插入模式。
	`a`：在当前字符后进入插入模式。
	`A`：在当前行最末尾进入插入模式。
	`o`：在当前行以下另起一行并进入插入模式。  
	`O`：在当前行以上另起一行并进入插入模式。
	`c + motion`：删除从当前字符直到指定位移的地方，同时进入插入模式。
	`ce`、`c + ctrl + ]`...
-   替换模式：正常模式下，按`R` 键，vim的左下角会有`-- REPLACE --`字样。
-   可视化模式：正常模式下，按`v` 键，vim的左下角会有`-- VISUAL --`字样。

---
##### 移动
只有正常模式、可视化模式、替换模式才会使用到移动快捷键。

使用小写`hjkl` 四个键来**左下上右**式地移动一个字符的位置。
   k         提示：h键在左边，可以向左移动。  
h        l    l 键在右边，可以向右移动。  
    j         j 键可以向下移动（k键向上移动）。 

1. 词移动
`w` ：移动到下一个单词开始位置。  
`5w/W`：多倍位移。

`e` ：移动到当前/下一个单词的末尾。 
`4e`：多倍位移。

`b`：移动到上一个字的第一个字母
`ge`：移动到上一个词的末尾
**ps**：上面几句话都是胡扯，我没理解移动的规则。
一个词以非单词字符结尾，例如`.`、`,`

**单词的定义**
- 连续的同类型的字符串(不包括空格)视为一个单词**
	类型：连续的中文字母数字，符号(中英文)；
	空格隔开两个字符，和上一个类型的字符组成同一个单词。
- 测试
    zhongwenzh ongwen中文中文中文    字符：。.##$#￥@2134哈哈哈jjj...

2. 子串移动
`E`：移动到子串的末尾。
`W`：移动到下一个类型块开始位置
`B`：方向同上
`gE`：方向同上

3. 字符匹配移动
`fx`：向前查找本行中字符x
`Fx`：向左查找本行中字符x

`tx`：向前把光标移动到目标字符的前一个字符上。
`Tx`：向左把光标移动到目标字符的前一个字符上。
使用`;`命令重复，使用`,`命令反向重复

`%` ： 在正常模式下，直接输入`%` ，vim会搜索并跳转到与当前行第一个括号相匹配的下一个括号位置（可能匹配点在之后几行），包括大括号，小括号，{}， （）， []。

4. 行移动
`n$`或`end键` ：移动到当前行/下n行的末尾。

`0`：移动到当前行的第一个字符。
`^`：移动到光标所在行的第一个非空白字符

`I`：移动到光标所在行行首，并进入插入模式
`A`：移动到光标所在行行尾，并进入插入模式

`n-`：移动到上一行开头第一个非空白字
`n+`：移动到下一行开头第一个非空白字符
ps：可以翻页

`shift + [`：跳转到上一个非连续的空行。 `shift + ]`：跳转到下一个非连续的空行。

`shift + 9`，`(`：跳转到上一个文本块的首行
`shift + 0`，`)`：跳转到下一个文本块的首行

5. 页面滚动
`ctrl + e`：上滚单行(Obsidian命令冲突，显示预览模式)
`ctrl + y`：下滚单行

`ctrl + u`：光标向下移动半页，所以文字向下移动半屏（obsidian中无用）
`ctrl + d`：光标向下移动一页，所以文字向上移动半屏（obsidian中是查找）

`ctrl + f`：正向整屏滚动（Obsidain命令冲突，查找）
`ctrl + B`：向后滚动（Obsidian命令冲突，加粗）

`zz`：将光标所在行移动到页面中间
`zb`：将光标所在行移动到页面底部
`zt`：将光标所在行移动到页面顶部

6. 页面移动
`nH`：移动到当前页面的行首（或相对行首第n行）
`M`：光标移动到当前屏幕的中间第一行
`nL`：移动到当前页面的行尾（或相对行尾第n行）
ps：n超出当前页面行数，会往上或往下翻页

`n%`：跳转到页面的第n%的位置

`[[`：跳转到文件第一行第一个字符。
`gg` ： 跳转到文件第一行第一个字符
`ngg`：跳转到文件的第n行

`]]`：跳转到文件最后一行第一个字符。
`G` ： 跳转到文件最后一行第一个字符。  
`nG`：移动到最后一行或第n行`nG`：移动到最后一行或第n行
`:n`：移动到第n行

7. 跳转与标记
`ctrl + [`：跳转到文本链接的主题
`ctrl + o`：回跳操作
`"`：跳转后回跳
`ctrl + i`
`:jumps`：输出调往的位置列表
`m + <a-z>`：用`<a-z>`标志当前光标的位置，用`'<a-z>`跳转到定义的标记
`:marks`：列出所有的标记

8. 定量位移
`#motion` ：在移位字符前加上一个数字，表示移位指定数目次。
`5w`、`4 + ctrl + [`、`4k`

---
##### 操作
1. 剪切或删除
>删除可在正常模式、可视化模式进行操作。
>操作命令 + 位移命令

`nx`：剪切光标所在单词
`nX`：剪切光标上一个单词

`D`：剪切当前行光标选中单词及后半部分的所有内容。
`ctrl + d`：剪切当前行（具体我也不是很清楚）
`dd`：剪切当前行
`#dd`：dd前加一个数字，剪切指定数量行。

`d + motion`：剪切位移内容，位移参考如上。
`d2w`、`d3k`、`d$`

`c$`：剪切从光标所在的字符到行尾的所有字符，并进入到插入模式
`cw`、`cb`、`c$`，
`s`：删除光标所在字符，并进入插入模式
`S`：删除光标所在行，并进入插入模式

`shift+j`：把两行连起来，即删除两行间的换行符

2. 复制粘贴
`p` ： 将刚刚用`y` 保存的文本粘贴插入到当前光标处，但不进入插入模式。      
`y` ： 
- 直接使用`y+motion`
    `nyb`、`nyw`、`y$`等
- 进入可视化模式，选择一些文本，按`y` 保存选择的文本，并退出可视化模式到正常模式。  
`nyy`：默认复制改行，复制光标自所在行开始，向下的n行
ps：剪切和复制都会把内容复制到粘贴板中。


3. 撤销
>常用语正常模式下

`u` ：撤销上一次操作。  
`ctrl + z`：撤销上一次操作
`U` ：把当前行恢复到修改前（再按就整体恢复到修改后），是一个新的操作，并不是撤销操作。  
`CTRL+r` ：重做被撤销的命令。

---
##### 替换
1. 查找 
这些字符\*\[\]\^\%\/\\\?\~\$具有特殊作用，要查找需使用转义符
`/`<要查找的字符>：向下查找
`?`<要查找的字符>：反向查找
一般是进行字符串匹配，若是要限定某个关键词，则可以使用`\<关键词\>`，其中`\<`标记关键词开头，`\>`标记关键词结尾。
按`<enter>`输入查找命令结束，按`n`查找下一个关键词，即向后匹配，按`N`查找上一个关键词，即向前匹配
`set incsearch`：搜索时忽略大小写，若要强制大小写则后面加`\C`，不区分`\c`
`:noh`：取消高亮

`*`：向下查找光标所定位单词
`#`：反向查找
仅支持英文，支持`n*`模式，默认是整词匹配，如果要部分匹配，则需使用`g*`或`g#`

**正则字符含义**
`^`：字符匹配行首，例`^include`仅匹配每一行开始的`include`
`$`：字符匹配行尾，规则同理
`.`：匹配任何字符，例`c.m`可以匹配任意字符串中包含`c`和`m`的字符串，如`computer`、`became`


**查找记录**
`/`+`<up>`或`/`+`<down>`：向上或向下浏览查找历史记录
`/关键词`+`<up>`：匹配和关键词相关的历史查找记录

2. 替换模式
`rx` ： 把当前字符替换为x，但是不进入插入模式。

3. 命令替换
`:rang s/str1/str2/gc`
rang：要替换的范围
`%`：表示所有行
`.`：表示当前行
`$`：表示最末行
`#,#`：表示特定两行之间
举例：
10,30：第10行到第20行
2,$-5：第2行到第倒数第5行
.,$-5：当前行到倒数第5行
同理可以用于复制和粘贴操作：10,30d表示剪切第10行和第20行

s：转入替换模式
string1：这是要查找的一个正则表达式
string2：这是希望把匹配串变成的模式的正则表达式
g：可选标志，带这个标志表示替换将针对行中每个匹配的串进行，否则则只替换行中第一个匹配串
c：可选标志，表示替换前询问

---
##### 命令
1. 显示所在文件的位置
`ctrl + G`：

2. 命令移动
`3a!<esc>`：在normal模式下执行!，把后面的命令执行三次
`3x`：删除三个字符
其他同理

3. 保存文件
`:w` ： 保存当前文件。  
`:w!`：强制保存文件
`:w [filename]` ： 将当前文件内容另存为filename命令的文件。

4. 退出
`:q`：退出
`:q!`：强制退出

5. 保存并退出
`:wq`：保存并退出


6. 调用外部命令
`:!CMD` ：在命令行终端里执行CMD命令，并且挂起vim，显示命令结果，按回车返回vim。

7. 提取文件内容与合并文件
`:r TEST` ： 将TEST文件内容拷贝到当前光标处。  
`:r !ls` ： 将当前目录下所有文件名称一个文件名一行地插入到当前光标下。
 ##### 退出vim
`q!` ： 命令模式下输入`:q!` 强制退出vim。（这一招会强制退出vim到命令行终端，同时放弃所有没有保存的修改）

8. 查看帮助文档
`<HELP>` 键，`<F1>` 键，`:help` 命令都可以查看帮助。  
`CTRL+w CTRL+w` ： 在vim中不同的窗口间跳转。  
`:q` : 退出当前vim窗口。  
`:help subject` ： 查看subject的帮助，如`:help w` ，`:help c_CTRL-D`，`:help insert-index` 。

8. 创建启动脚本
`~/.vimrc` ： 当前用户默认的vim启动脚本，可以在脚本里启用或停用特定功能。  
`$VIMRUNTIME/vimrc_example.vim` ： 参考启动脚本。  
`:help vimrc-intro` : 查看启动脚本的帮助。

9. 自动完成
`CTRL+d` ： 自动将可能的命令或参数在vim界面最下方临时显示出来。  
`<TAB>` ： 按顺序将这些可能性补全在当前命令上。

10. 光标位置和文件状态
`CTRL+g` ： 显示文件名，光标在文件中的位置以及文件状态。  

11. 编辑新文件
`vi filename` ： 在命令行终端中，输入命令`vi filename` 并回车，将进入vim，打开文件名为filename的文件，同时进入正常模式。
`vs filename`：分屏显示，并创建名为filename的文件

12. 设置
`:set option` ： 设置vim的配置选项。

常用的option配置选项有：
- `hls` 高亮显示搜索命中
- `ic` 忽略大小写
- `is` 增量搜索
配置选项前加上no表示取消此配置选项，例如：`/ignore\c`_

13. 从文件中读取可执行命令
`:so %`: so是source的缩写

14. 刷新当前文档
`:e`：刷新当前文档即其内容

15. 分屏操作
参考链接: [5分钟学会 Vim 分屏操作 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/337157587#:~:text=ctrl%2Bw,%E4%B8%8A%E4%B8%8B%E5%B7%A6%E5%8F%B3%E9%94%AE)
`ctrl + W + H J K L`：移动分屏
`ctrl + w + h j k l`：光标移动


##### Vim编辑器-插件使用
参考：[junegunn/vim-plug: Minimalist Vim Plugin Manager (github.com)](https://github.com/junegunn/vim-plug)
1. vim-plug
自动管理vim编辑器的插件，包含：自动安装、更新、清除、查看状态等才做，可以指定在某些特殊环境下使用某些插件
- 下载plug.vim，将其放置在vim编辑器的配置文件夹中
liux系统：`/usr/share/vim/vim90/autoload`
- 安装git软件
- 使用`PlugInstall`命令

2. 常用插件


3. 一种安装插件的方式
//插件: Tab键一键补全
- 导入 supertab.vmb
- 输入：vim supertab.vmb
- 运行 “：so %”，其中vim中的so命令是source命令的缩写，它的作用是从文件中读取可执行命令（shell命令）来执行
- 在底行模式在输入 UseVimball ~/.vim
- vi ~/.vimrc
- 文件中加入以下这行：let g:SuperTabDefaultCompletionType="context"
- 当出现：so %的时候出现一堆红色字体，执行下面的操作：
//	ls -al //找到.vim文件夹 查看文件夹的名字是不是root，如果是执行下面的操作 sudo chown farsight:farsight .vim //将文件的拥有者换成我们自己 重新执行上面的操作






