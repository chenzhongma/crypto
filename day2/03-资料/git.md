# 一、关联工程

## 1. 历史不相干的

核心命令：git pull origin master --allow-unrelated-histories  

```sh
1.本地初始化git目录
git init

2.新建文件并且写入内容
touch a.txt
echo "new data" >> a.txt

3.添加到暂存区
git add .

git commit -m "a.txt"

4.添加远程仓库
git remote add origin https://gitee.com/jianan/learnGit.git

5.本地仓库也远程仓库关联
git branch --set-upstream-to=origin/master master

6.拉取远程仓库内容到本地
这时候用git pull会提示(毕竟本地和远程仓库没啥关系指针连接不起来的缘故吧)：

fatal: refusing to merge unrelated histories

因此命令应该改为：

git pull origin master --allow-unrelated-histories  

7.将最新的内容推送到远程仓库
git push
```



## 2. 远程无代码的

```sh
3. 执行git fetch origin master

显示如下：
remote: Enumerating objects: 3, done.
remote: Counting objects: 100% (3/3), done.
remote: Total 3 (delta 0), reused 0 (delta 0)
Unpacking objects: 100% (3/3), done.
From https://gitee.com/bwcs/01_bitcoin
* branch            master     -> FETCH_HEAD
* [new branch]      master     -> origin/master
```





# 二、显示问题

## 1. 无法显示中文

```sh
当命令行出现中文无法正常显示时，执行下面的命令可以解决
git config --global core.quotepath false
```



## 2. 颜色

```sh
git 设置颜色 color
git config color.ui true
```





# 三、常用操作

## 1. 基础命令

```sh
git add file.name
git commit
git checkout file.name
git rm file.name
git mv file.name1 file.name2  相当于
    a. mv file.name1 file.name2 
    b. 删除file.name1 
    c. 添加file.name2

git branch test  
    创建test分支， 但是不切换至test分支
git checkout -b tmp
    创建tmp分支， 且切换至tmp分支
```

## 2. 推送到远程分支（push）

==当本地分支reset之后，如果想push到远程，会被拒绝，如果确定推送，可以强制push==

```sh
git push --force  //==> 慎用
```



```sh
git push origin  test:duxu/test   test 是本地分支名字， duxu/test 是远程分支名字
```

## 3. 修改最近一次提交的信息（amend）

```sh
git commit  --amend
```

## 4.查看提交历史（log）

```sh
git log --pretty=oneline
```



```sh
git log 分支哈希 -p -2
```

## 5.仓库相关（remote）

- 查看远程仓库地址 :

```sh
git remote -v 
```

- 添加远程仓库

```sh
git remote add pb git://github.com/paulboone/ticgit.git
```

## 6. 拉取代码（pull与fetch）

- fetch

  fetch只会把远程分支同步到本地，但是不会自动合并到当前分支

    ```sh
    git fetch    //将远程分支的更新同步到本地

    查看
    git branch -a

    拉倒本地  //双击tab可以补全
    git fetch origin duxu/test:duxu/test
    ```

- pull

  git pull =  git fecth + git merge

  它会把远程分支同步到本地，同时与指定的当前分支进行合并，如果不指定分支，则与当前分支合并

    ```sh
    git pull origin remote_name:指定分支名字
  
    或
    git pull origin remote_name
    
    可以关联tracking追踪
    git branch --set-upstream master origin/next
    
    git pull
    ```

- 示例

    ```sh
    $ git fetch origin master:tmp
    $ git diff tmp 
    $ git merge tmp
    ```



​	

## ==7. 打标签（tag）==

可以将发布的版本打上一些标记，从而可以快速的切换到对应的分支

- 基于当前的位置，打上标签

  ​	

  ```sh
  git tag -a v0.1 -m "my version v0.1"
  ```

- 显示所有的tag

  ```sh
  git tag
  ```

- 显示某个tag详情

  ```sh
  git show v0.1  //tag号
  ```

- 删除tag

  ```
  git tag -d v1.0
  ```

- 对某个特定的分支打tag

  ```sh
  git tag -a v1.2 9fceb02
  ```

- 提交tag

  ```sh
  git push origin --tags  //所有的tags
  或
  git push origin v1.0  //指定tag号
  ```

  ![image-20190220094542884](https://ws4.sinaimg.cn/large/006tKfTcgy1g0cnb08xgyj31ks03smyl.jpg)





## 8. 删除分支（delete）

- 远程分支

```sh
git push origin --delete remote_name
```

- 本地分支

```sh
git branch -d local_name
```



## 9. 合并分支（merge与rebase）

```sh
git rebase test1

git merge test1
```

两者都是合并两个分支， 但有所不同

git merge 在将两个分支合并的同时会显示的表明有Merge branch "xxx" into "xxx"， 有合并痕迹，不简洁。



而git rebase 是一种简洁的合并， 它的合并结果是将两个分支毫无合并痕迹的合并起来， 就像是没有合并过一样， 就像所有的操作都是在这同一个分支里完成的， 没有Merge branch "xxx" into "xxx"字样。



- 示例：

在某一个时刻基于C2创建一个分支mywork：



![img](http://www.yiibai.com/uploads/images/201707/1307/842100748_44775.png)





两个分支都做了一些修改：

![img](http://www.yiibai.com/uploads/images/201707/1307/810100749_17109.png)



现在想要将origin分支合并到mywork分支上，如果使用merge命令，相当于重新提交了一次C7，会出现：Merge branch mywork into origin字样。

![img](http://www.yiibai.com/uploads/images/201707/1307/350100750_71786.png)



如果使用reabse，就会把C5，C6的提交临时取消保存起来，然后基于C4再将mywork的提交。

（==注意，会把mywork的提交暂存起来，把基于的origin分支更新，然后把mywork的提交追加上去==）

![img](http://www.yiibai.com/uploads/images/201707/1307/845100751_76810.png)



最终效果：

![img](http://www.yiibai.com/uploads/images/201707/1307/141100752_31232.png)

## ==10.冲突处理==

- 如果merge出错

  解决冲突

  ```sh
  打开tmp.txt 文件会发现内容为
  <<<<< HEAD
  eeeeeeeee
  =========
  dddddddd
  >>>>>> ddddd into tmp.txt
  ```

  修改后：

  ```sh
  git add tmp.txt
  再次执行merge命令
  ```

  

- 如果rebase出错

  解决冲突

  ```sh
  git add tmp.txt
  git rebase --continue
  ```

  当不想合并时， 可以执行==git rebase --abort==， 并且mywork分支会回到rebase前的状态



​	

# 四、常见错误处理

## 1. git clone 错误

```sh
`server certificate verification failed. CAfile: /etc/ssl/certs/ca-certificates.crt CRLfile: none
解决方法：
export GIT_SSL_NO_VERIFY=1
#or
git config --global http.sslverify false
```

