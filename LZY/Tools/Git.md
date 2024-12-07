---
时间：2023-02-24  18:55
钩子：
概述：
---
git是一个开源的分布式版本控制系统，可以保存本地仓库里的代码历史，也可以将本地代码同步至远程仓库，实现多人协助。
保存本地仓库里的代码历史是指：每次在本地仓库里修改代码时，git都会记录修改内容，提交时间，提交人，和提交说明，方便查看代码的变化过程，也可以回退到某个历史版本。
常见的远程仓库就是github里的仓库。
##### 问题
1. # error: src refspec xxx does not match any / error: failed to push some refs to；在一个库里面新建一个分支，连接到本地一个库。
[Git 常见错误 之 error: src refspec xxx does not match any / error: failed to push some refs to 简单解决方法_error:src refspec_仙魁XAN的博客-CSDN博客](https://blog.csdn.net/u014361280/article/details/109703556)

2. 如何连接添加本地库连接远程库？
添加ssh，具体看下方介绍。
使用`git clone 地址`命令
`push`前先要使用`pull`.

3. 当两个本地对同一个文件进行了修改，出现：# Please commit your changes or stash them before you merge
参考链接：[Git冲突：Please commit your changes or stash them before you merge_serialhelper.cs_1024_Byte的博客-CSDN博客](https://blog.csdn.net/DDD4V/article/details/118896307)
出现这个问题的原因是其他人修改了xxx.php并提交到版本库中去了，而你本地也修改了xxx.php，  
这时候你进行git pull操作就好出现冲突了，解决方法，在上面的提示中也说的很明确了。
```c
//保存修改
git stash
git pull
git stash pop
```
通过git stash将工作区恢复到上次提交的内容，同时备份本地所做的修改，之后就可以正常git pull了，git pull完成后，执行git stash pop将之前本地做的修改应用到当前工作区。
git stash: 备份当前的工作区的内容，从最近的一次提交中读取相关内容，让工作区保证和上次提交的内容一致。同时，将当前的工作区内容保存到Git栈中。
git stash pop: 从Git栈中读取最近一次保存的内容，恢复工作区的相关内容。由于可能存在多个Stash的内容，所以用栈来管理，pop会从最近的一个stash中读取内容并恢复。
git stash list: 显示Git栈内的所有备份，可以利用这个列表来决定从那个地方恢复。
git stash clear: 清空Git栈。此时使用gitg等图形化工具会发现，原来stash的哪些节点都消失了。
```c
//放弃一方本地库的修改，使用如下命令
git reset --hard
git pull
```

4. 同时链接两个远程仓库
参考链接: [Git：解决本地库同时关联GitHub和Gitee_git](https://blog.csdn.net/Bertil/article/details/119464439#:~:text=%E9%A6%96%E5%85%88%E5%9C%A8GitHub%E5%92%8CGitee%E5%88%86%E5%88%AB%E6%96%B0%E5%BB%BA%E4%B8%80%E4%B8%AA%E4%BB%93%E5%BA%93gitTest%20%E7%84%B6%E5%90%8E%EF%BC%8C%E5%9C%A8%E6%9C%AC%E5%9C%B0%E6%96%B0%E5%BB%BA%E4%B8%80%E4%B8%AA%E6%96%87%E4%BB%B6%E5%A4%B9gitTest%EF%BC%8C%E5%8F%B3%E9%94%AE%E7%82%B9%E5%87%BBgit%20bash%E5%90%8E%E8%BE%93%E5%85%A5%20git%20init,%E5%AF%B9%E4%BB%93%E5%BA%93%E8%BF%9B%E8%A1%8C%E5%88%9D%E5%A7%8B%E5%8C%96%20%E6%8E%A5%E7%9D%80%EF%BC%8C%E6%88%91%E4%BB%AC%E5%9C%A8%E6%9C%AC%E5%9C%B0%E5%BA%93%E4%B8%8A%E4%BD%BF%E7%94%A8%E5%91%BD%E4%BB%A4%20git%20remote%20add%20%E6%8A%8A%E5%AE%83%E5%90%8C%E6%97%B6%E5%92%8CGithub%E3%80%81Gitee%E7%9A%84%E8%BF%9C%E7%A8%8B%E5%BA%93%E5%85%B3%E8%81%94%E8%B5%B7%E6%9D%A5)
- 设置权限
首先找到github 和 gitee配置公钥的地方，把本地`~/.ssh/id.rsa.pub`中的内容拷贝进去。

- 添加远程库链接
`git remote add gitee 链接`
`git remote add github 链接`

- 查看关联的远程状态
`git remote -v`

- 统一添加本地库文件至缓冲区，并提交
`git add .`
`git commit -m "comment"`

- 推送至远程库
`git push github master`
`git push gitee master`


- [ ] 怎么修改远程仓库名，怎么删除远程仓库?

---
##### Git 实现原理
通常用repo代表一个git仓库，一个git仓库由工作区(working tree)，暂存区(index/stage)，本地仓库，远程仓库(可选)组成。文件在工作区中创建、删除、修改，使用`git add`提交到暂存区，最后使用`git commit`创建一个commit。

git通过`.git`文件管理`repo`中所有文件的各个版本，每一次commit都会生成一份`repo`的完整快照。其中，只有发生修改的文件才会创建一份完整的文件，而未修改的文件，快照中只包含一个指向上一个commit的指针，以此类推，所以commit具有前后的继承关系。

所有的快照都保存在`.git`中，以时间为轴往前递进，除非删除该`.git`文件夹 或者 使用`git rm`，否则commit就不会丢失，每个commit都对应一个`commit id`，只要知道commit id，就可以随时切换到该commit id对应的文件状态。

分支的本质是指向commit的可变指针，也代表着一组通过相同指针进行提交的commit的集合。

当使用`git clone`拉取远程仓库 或 使用`git push -u`第一次和远程仓库建立联系时，在本地会创建用于与远程仓库同步的==远程追踪指针==，在每次使用`git fetch` 或 `git pull`时，会将指定的远程仓库所有分支内容拉取到本地，并更新本地==远程追踪指针==的指向。

==远程追踪指针== 仅用于与远程仓库同步，只有在同步时它们的指向才会发生改变，否则就可以认为它们是不可变得。

远程仓库也有一个HEAD指针指向默认的分支，暂且称其为远程HEAD指针

当需要创建本地分支时，使用`git branch xx` 或 `git checkout -b xx`会创建一个本地分支指针，本地指针默认指向本地HEAD指针，如果没有本地仓库，本地HEAD指针 和 远程HEAD指针指向一致。

本地HEAD指针一般指向当前分支的最后一个commit。

如果需要创建指向远程特定仓库的分支指针，可以使用`git checkout xx`，其中`xx`代表远程仓库的分支名，通过本地的==远程追踪指针==可以获取。

---
##### 合并算法
合并会产生冲突，冲突需要自行解决。
1. 快速合并
指针的移动

2. 三路归并合并


3. 递归三路归并合并算法

---
##### 全局配置
1. 设置账户名和邮箱
```shell
#设置账户名和邮箱
git config --global user.name="LZY"
git config --global user.email="1981418038@qq.com"

#添加全局代理配置
git config --global http.proxy http://127.0.0.1:7890 #或
git config --global http.https://github.com.proxy http:/127.0.0.1:7890
```
`--global`是全局配置选项，全局设置的本质是修改Git的全局配置文件`~/.gitconfig`中的内容，而不是各个仓库里的配置文件`.git/config`，全局配置文件中的设置对所有仓库有效

2. 查看配置
```shell
#查看全局配置
git config --global -l #也可使用--list选项
#查看所有配置
git config -l  #或
git config --list
```

3. 清空配置
- 打开配置文件直接删除相关内容
windows: `~/.gitconfig`
![[windows下的git配置文件路径.png]]
linux: `~/.gitconfig`
![[gitconfig配置文件.png]]
```shell
git config --global --edit
#默认以vim编辑器打开
```
- 使用命令删除
```shell
#删除账号密码全局配置
git config --global --unset user.name
git config --global --unset user.email

#删除代理配置
git config --global --unset http.proxy
git config --global --unset https.proxy
```

---
##### Git代理
当在克隆或从远程仓库获取数据时，很可能因为网络状况不佳遇到很慢甚至超时的情况，比如从`github`上拉取仓库，这种情况下就需要配置代理。
参考链接: [一文让你了解如何为 Git 设置代理 - Eric (ericclose.github.io)](https://ericclose.github.io/git-proxy-config.html)
1. 使用HTTP/HTTPS协议代理
```shell
#针对所有域名的仓库
git config --global http.proxy <protocol>://<host>:<port>
#针对特定域名的仓库
git config --global http.<url>.proxy <protocol>://<host>:<port>
```
针对所有域名可以理解为设置该配置后，不管是访问`gitee`还是访问`github`都走代理模式
`--global`: 如上
`<protocol>`: 指的是代理协议，如http https sock5等
`<host>`: 为代理主机，如使用本地代理主机 127.0.0.1 或 localhost等
`<port>`: 为代理端口号，如clash使用的7890或7891等
`http.proxy`: 只能使用http，https是错误用法
`<url>`: 指的是你需要代理的远程仓库，支持HTTP / HTTPS 传输协议的格式
- 举例
```shell
#针对所有域名的代理配置
git config --global http.proxy http://127.0.0.1:7890
#针对特定域名的代理配置
git config --global http.https://github.com.proxy http://127.0.0.1:7890
```
`http://127.0.0.1`: 查看clash代理主机地址
`7890`: 查看clash软件的代理端口号

---
##### Git基本命令
新建一个文件夹，使用`git bash`打开
1. git初始化
`git init`：初始化该文件夹，会在文件夹中创建一个`.git`文件，可以理解为用git追踪该文件夹.
注意：新建文件夹的==所有者==应该和 `.git`文件夹的拥有者相同，否则就会出现警告，可以在`文件夹-属性-安全-高级`中修改所有者。

2. 删除git初始化
- 直接删除`.git`文件
- 使用命令
```shell
find . -name ".git" | xargs rm -Rf #以前这个命令看不懂，哈哈哈哈，现在能看懂了
```

3. 查看文件状态
```shell
git status
```

4. 添加文件到缓存区
```shell
git add 文件名 #添加某个文件到缓存区
git add .  #.表示当前文件夹，默认情况，需要解析.gitignore文件内容
git add *  #*可以理解为通配符，就是指当前文件夹中的所有文件，不解析.gitignore文件内容
```

5. 提交文件到本地仓库
```shell
git commit -m "注释" #单引号和双引号有区别，单引号会作为注释内容
```

6. 创建分支
```shell
git branch debug
#创建debug分支，不切换
git checkout -b debug
#创建debug分支，并切换到debug分支
```

7. 查看分支
```shell
git branch  #查看本地分支
git branch -r #查看远程分支
git remote -v #查看链接的远程库
git branch -a #查看本地和远程的所有分支
```
![[Pasted image 20230805131228.png|600]]
其中`*`号表示的是当前所在的分支

8. 切换分支
```shell
git checkout debug
#从当前分支，切换到debug分支
```

9. 删除分支
```shell
git branch -d <branchname>  #删除本地分支
git branch -d -r <branchname> #删除远程分支，删除后还需推送到服务器
git push origin -delete :<远程分支名>
```

10. 分支重命名
```shell
git branch -m <oldbranch> <newbranch>  #重命名本地分支
```

命令选项
`-d` --delete 删除
`-D` --delete --force强制删除
`-m` --move 移动或重命名
`-M` --move --force 强制
`-r` --remote 远程
`-a` --all 所有

11. 打标签
标签有两种，一种是轻量标签，可以理解为某个特定commit的引用，是确定不变的。另一种是附注标签，它是git仓库中的一个完整的对象，可以被校验，此外，它包含一些附注信息，如：打标签者的名字，邮件地址，日期时间，标签信息等。
- 创建标签
```shell
git tag v1.0  #给当前分支打上轻量标签
git tag -a v2.0 -m "标签信息" #给当前分支打上附注标签
git tag -a v3.0 commit_id #对指定的commit打上标签
```
- 显示标签信息/列表
```shell
git tag (-l --list)  #以字母顺序列出标签，选项后面可以使用带通配符的字符串，如下
git tag -l "*-v65"
```
也可以使用`git show`查看特定的标签
```shell
git show <tagname>
```
- 共享标签
一般情况下，标签是仅存于本地的，如果想让标签在远程仓库中同步，可以使用`git push`命令
```shell
git push origin <tagname> #推送单个标签到远程
git push origin --tags    #将所有标签推送到远程仓库
```
- 删除标签
```shell
git tag -d v1.0  #删除本地标签v1.0
git push <remote> :refs/tags/<tagname>  
#本地删除标签后，将信息同步到远程可以删除远程仓库中的标签
git push <remote> --delete <tagname>
```
第二条命令的本质是将冒号前面的空值推送到远程标签名，从而高效的删除它

- 切换标签
由于标签所指示的commit是不可变的，因此将HEAD指针指向标签，会导致HEAD指针处于"detached HEAD"的状态，和本地==远程追踪指针==类似，在这种状态下，任何提交修改，标签都不会发生变化，而且新的提交不会属于任何分支，并且将无法访问，除非通过commit id访问。因此，如果需要基于标签进行修改，通常需要创建一个新的分支
```shell
git checkout -b version2 v2.0.0 #从标签v2.0.0上创建一个新的分支并跳转
git checkout
```

---
##### Git log
1. git log
`git log`用于查询当前分支上的commit记录
```shell
git log  #查看分支上所有的commit日志
git log --pretty=oneline #单行显示日志
git log show commit_id #查看指定commit的log
```


2. git reflog
`git reflog`用于查询本地HEAD指针的指向记录，HEAD指针一般用于指向当前commit，其中ORIG_HEAD指向HEAD指针的上一次指向。
```shell
git reflog
```

---
##### 撤销命令
工作区新建文件，修改后，再没有add到暂存区中，无法通过git操作撤销==部分修改==，但是可以使用`git clean`或 `git reset --hard`将新增的还未追踪的文件全部清空。

当这些文件add到暂存区后，可以理解为暂存区保存了这些文件的一个状态，可以使用`git checkout .`撤销这些文件新增修改，将文件恢复到暂存区中的状态。

对于已经追踪的文件，在没有add到暂存区之前，同样可以使用`git checkout .`将文件恢复到修改之前的状态。
如果已经add到暂存区，可以理解为暂存区保存了这些文件的最新状态，后续新增修改，可以使用`git checkout .`进行撤销。

对于add到暂存区但还未commit的内容，可以使用`git reset`撤销暂存区的新增。

上面就是针对工作区 和 暂存区的 撤销动作

下面是针对commit的相关撤销

1. `git clean`


2. `git reset`
根据命令选项，来执行不同程度的重置操作，如果后面接了路径，则跳过改变HEAD指向的分支，只执行后面的操作。
--mixed(默认选项)
会先移动HEAD指向的分支，然后会将HEAD指向的当前快照的内容来更新暂存区。
```shell
git reset   #HEAD不移动，使用当前HEAD指向的快照填充暂存区
git reset HEAD . #和上面相同，只不过上面是简略用法
git reset (--mixed) HEAD~ .  #使用上一次提交的快照填充暂存区，不会移动HEAD指针
git reset (--mixed) commit_id file_name  
#使用指定快照中的file_name文件填充暂存区，不会移动HEAD指针
```
- [ ] 如果只是更新暂存区中的某个/某些文件，会不会暂存区中现有的同名文件冲突？

--soft
只移动HEAD指向的分支，不改变暂存区和工作区中的内容，比如，当我对这次提交不满意，或者又新增了一些新的文件，和修改，我想将这些修改和当前commit合并，此时就可以
```shell
git reset --soft HEAD~  #不带路径，将HEAD分支指向上一次提交，保持暂存区 和 工作区内容不变
git commit -m "提交"    #重新提交，会生成一个新的commit，同时舍弃原本的commit
```
也可以使用
```shell
git commit --amend
```
这个命令是将暂存区中的内容重新提交到当前commit，如果暂存区中内容没有发生变化，可以使用该命令补充提交信息。

--hard
移动HEAD指向的分支，重置暂存区 和 工作区后，会用HEAD移动后分支指向的快照填充暂存区，并将暂存区中的内容复制到工作区。

此命令会造成工作区中未提交内容的丢失，谨慎使用。

3. `git checkout`
`git checkout <branch>`和`git reset --hard [branch]`有点像，但是它会检查当前分支是否有修改，对于已追踪的文件，如果有未提交的修改，它会提示你先进行提交，否则不允许切换分支。对于未追踪的文件，由于工作区和暂存区是公共的，可以直接进行切换。而`git reset --hard`则是全面替换，会造成工作区的修改丢失。
`git checkout [branch]`本质上是改变HEAD指针的指向，从而达到切换分支的目的，`git reset`则是改变HEAD指向的分支的指向，会造成commit丢失。

此外，可以使用`git checkout`创建一个和本地远程追踪指针关联的本地指针
```shell
git checkout xxx   #xxx和远程追踪分支同名
```
如果是使用`git branch xxx`，则只会创建一个xxx分支，并不会将远程追踪分支作为其上游分支(关联)。

同`git reset`一样，`git checkout`如果后面带了路径，它不会移动HEAD指针，会使用指定的commit中的内容填充缓冲区，并覆盖工作目录
#理解
其实git checkout 带了路径，如果是当前commit，它会从缓冲区中检出文件，覆盖到工作区，如果缓冲区发生改变，并不会使用当前commit指向的快照填充缓冲区，但如果是其他commit，首先会将指定的commit指向的快照填充缓冲区，然后再覆盖工作区中的对应文件。
```shell
git checkout . #带了路径，会将当前的commit指向的快照填充缓冲区，并覆盖工作区
git checkout dev #如果dev是文件名，则是用当前commit指向的快照中的dev 填充缓冲区，并覆盖工作区，如果是分支名，则是切换分支
git checkout commit_id file_name #不改变HEAD指针，使用指定commit的file_name填充当前缓冲区，并覆盖工作区。
```
所以为什么在应用的时候，如果修改了已经追踪的文件，如果还未提交到暂存区，使用`git checkout .`会撤销修改，将其恢复到修改之前。若已经提交到了暂存区，使用`git checkout .`


---



7. 查看当前修改的内容
```shell
git diff 文件名
```

8. 版本回退
- 查看提交版本
```shell
git reflog
```
- 相对回退
```shell
git reset --hard HEAD^ #回退到上一次
git reset --hard HEAD~ #默认回退一次，后可接数字，表示回退几次提交，如
git reset --hard HEAD~3  #回退三次提交
```
注: 在windows的cmd控制台下操作git，想要回滚到上一次提交，但是输入`git reset --hard HEAD^`后就显示more?，多按几次回车后就报错如下，如何解决呢？这是因为cmd控制台中换行符默认是`^`，而不是`\` ，所以它的more？的意思是问你下一行是否需要再输入，而`^` 符号就被当做换行符而被git命令忽略掉了，解决方法有如下几种：
```shell
#加引号：
git reset --hard "HEAD^"
#加一个^
git reset --hard HEAD^^
#换成~
git reset --hard HEAD~
```
- 绝对回退
根据基限码，指定回退到哪个版本
![[Pasted image 20230805130220.png|600]]
只有使用`git commit -m "注释"`将内容提交到本地仓库后，才会创建回退节点
```shell
#使用基限码指定回退的版本
git reset --hard 147c1c2
```

9. 修改和删除
`git checkout`用于切换分支 或 恢复工作树文件

---
##### 分支命令
先切换到目标分支，可以将指定分支合并入目标分支，如假设当前是master分支，将test分支合并到master分支中如下实例
如果出现了两个分支都修改同一个文件，合并过程会有冲突(手动解决，然后提交)
```shell
git merge debug
#将debug分支合并到master分支，内容合并
```


---
##### 缓存命令
1. 隐藏修改的但是还是没有添加到缓存区的内容
```shell
git stash
```

2. 查看隐藏列表
```shell
git stash list
```

3. 删除列表
```shell
git stash drop
```

4. 恢复列表内容
```shell
git stash apply
```

---
##### 远程仓库命令
1. 克隆仓库
```shell
git clone 仓库链接
#如
git clone https://github.com/shenshelin/test.git
```

2. 添加远程仓库
```shell
git remote add origin https://github.com/AZ-ZYYQ/System.git
#可以理解为origin就代表后面的链接，即远程origin库
```

3. 将远程库获取最新版本到本地
```shell
git fetch origin master:temp #将远程的master分支中的内容下载到本地的temp分支
git pull <远程库名> <远程分支名>:<本地分支名> 
#从远程库中获取某个分支的更新，与本地指定的分支进行自动合并
#若是与本地当前分支merge，则可省略冒号后面的本地分支名
git pull origin
#若远程分支本地当前分支之间具有追踪关系，则可直接使用
```
`git pull`等同于先做了`git fetch`，再做了`git merge`。

4. 比较本地仓库与下载的temp分支
比较本地当前所在仓库与下载的temp分支
```shell
git diff temp
```

5. 同步本地分支到远程库中
```shell
git push origin <本地分支名>
```
推送后会在远程库中看到新建的本地分支，两者之间建立了追踪关系。

---
##### 链接到Github远程库
参考链接：[git和github的连接](https://blog.csdn.net/qq_29493173/article/details/113094143)
进行[git和github的连接](https://blog.csdn.net/secretstarlyp/article/details/106576882)
在github里创建一个远程仓库，将本地仓库与远程仓库进行绑定
1. 本地生成SSH公钥
```shell
ssh-keygen -t ed25519 -C "Gitee SSH Key" #gitee库
```
生成的秘钥保存在`~/.ssh`文件夹中
2. 测试
```shell
ssh -T git@gitee.com  #gitee测试
Hi USERNAME! You've successfully authenticated, but GITEE.COM does not provide shell access.
```
3. 链接到仓库
```shell
git remote add origin 后接仓库的地址
```

4. 将本地仓库推送到远程仓库中
```shell
git push -u origin master #第一次(远程仓库为空)
git push origin master #远程仓库非空
git pull --rebase origin master #若是远程库和本地库不一致，需要采用
```
把远程库中的更新合并到本地库中（pull）, 并取消掉本地库中刚刚的commit，并把它们接到更新后的版本库之中（--rebase）

---
##### gitignore文件
参考链接: [gitignore使用](https://zhuanlan.zhihu.com/p/264995020#:~:text=gitignor,%E6%8F%90%E4%BA%A4%E7%9A%84%E6%96%87%E4%BB%B6%E5%8D%B3%E5%8F%AF%E3%80%82)  [gitignore文件的配置使用](https://zhuanlan.zhihu.com/p/52885189)
在本地仓库中新建一个文件`.gitignore`（配置文件，没有文件名），在该文件中添加特定命令实现对文件的忽略，`.gitignore`只对还没有加入版本管理的文件起作用，所以要写先创建`.gitignore`文件再使用`git add .`命令
忽略一个文件夹中所有内容：`文件夹名/*`
忽略特定类型的文件：`*/.text`
待补充...
如果在创建`.gitignore`文件之前已经跟踪了一些你想要忽略的文件或文件夹(此处是指已经把文件加入了缓存区或者是本地仓库)，那么修改`.gitignore`文件并不会自动取消跟踪他们。需要使用`git rm -r --cached .`清空缓存区中的记录，创建并配置好`.gitignore`文件后，再重新`git add .`提交文件到缓存区

---
##### Git 进阶
把commit过的文件称为已经追踪文件，工作区内新建文件为未追踪文件
1. 撤销


