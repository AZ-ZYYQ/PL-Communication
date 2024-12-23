时间：2022-09-29  06:50
标签： #工具 
引用：
[Typora 的 Markdown 语法](https://support.typoraio.cn/zh/Markdown-Reference/))

##### 标题
`#`  一级标题
`##`  二级标题
`###`  三级标题
obsidian最高支持6级标题
notion最高支持3级标题

---
##### 序列
###### 有序列
输入格式：`数字` + `点`+ `空格`
1. 例

###### 无序列
输入格式：星星(`*`) 或 短横杠(`-`)或 加号(`+`) + `空格`
- 星星
- 短横杠
- 加号

---
##### 代办
语法：`-`+`空格`+`[`+`空格/X`+`]`+`空格`+文本
- [ ] 代办1
- [x] 代办完成1

x或X都可以

---
##### 单分隔线
输入格式：三个短横杠(`-`)，或者三个连续的星号(`*`)；中间除了空格不能有其他东西

Notion只支持三个连续的短横杠表示分隔符
obdian两种方式都支持

---
##### 引用
输入格式：大于号(`>`) + 空格

Notion中竖横( `|` )替代大于号(`>`)，且不支持多级引用
obdian中原生支持`>` + 空格的方式，且支持多级引用

---
##### 代码块
单个元素突出显示：\`突出内容\`
`这是一种突出显示方式`，适用于一些概念或者名词的突出显示。

单行代码块：`空一行`+`table(四个空格键)`
单行代码块可以多行叠加使用，但效果不如用多行代码块语法好看。

	printf("Word Hello！");

多行代码块：
	\`\`\` + 语言名
	代码
	\`\`\`
常见编程语言：C，C++，Python，Java，mermaid，Html

Html代码块显示方式：开头`<XXX>`，结尾`</XXX>`
<html>
哈哈哈哈？
<pre>这又是啥</pre>
</html>

---
##### 缩进
html缩进语法
<table>
&emsp;全角缩进<br>
&ensp;半角缩进

</table>

印象笔记的markdown原生支持html语法
notion和obsidian的markdown原生不支持html语法，但语法可以在拓展语言代码块中实现。

四个空格等于一个table缩进，引用和缩进还是有区别的。

---
##### 特殊字符转义表示
用`\`+特殊字符来表示特殊字符。

如：\`

---
##### 文字链接
1. 行内式
> 即链接跟在文本后面

语法：`[显示文本]`后跟`(链接)`
这是一个[必应搜索](http://biying.com)，点击即可跳转

2. 参考式
> 即用一个数字代替网址，单独一行进行解释。

语法：
`[]`跟`[代号]`
`[代号]:链接`

obdisian和Notion不支持这个用法，印象笔记的markdown支持。

---
##### 自动连接转化
>将本文标记为链接

语法：`<链接内容>`

obsidian和notion不支持这种语法，支持自动识别常见的格式。
印象笔记的markdown语法支持如：<123456789.com>

---
##### 插入附件
>图片、视频、附件等等

###### 行内式
语法：
`![Alt text](本地路径“提示文本”)`，或
`![[附件名称]]`
![](markdown语法测试附件1.jpg)
obsidian不支持文本提示，即括号中不应存在引号内容，`[]`可为空，但必须存在
![[markdown语法测试附件2.png]]

###### 参考式
语法：
`![id]`
`[id]:网络路径/本地路径`

显然obsidian和notion不支持这种语法，印象笔记的markdown语法支持这种方式。

---
##### 文本样式
###### 加粗
输入格式： `**`文本 `**`或者 `__`文本 `__`（文本前后两个单下划线）
**星星**，
__双单下划线__

###### 斜体强调
输入格式：`*`文本`*`或者`_`文本`_`
*星号斜体*
_单下划线斜体_

###### 高亮显示
语法：\==高亮文本\==
==高亮显示==

###### 删除线
语法：`~~`删除文本`~~`
~~被删除对象~~

---

待补充。。。