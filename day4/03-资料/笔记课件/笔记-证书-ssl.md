# 今日内容

```sh
1. 自己本地生成数字证书  ==> openssl  => done


2. 介绍SSL/TLS  ==> done


3. 使用go搭建一个简单的SSL请求的server
两套代码：
A. 单向认证
- 我们客户端单向认证服务器，服务器不认证客户端


B. 双向认证
- 客户端认证服务器
服务器把证书发给客户端，客户端校验（登录淘宝）

- 服务器认证客户端
客户端把自己的证书发送给服务器，由服务器进行校验（金融领域的通信，网银盾）
```



# 四、生成自签名证书

下列两种方式生成的证书都是pem格式的，可以导入到计算机。

- 私钥文件
- 数字证书（包含公钥）



## 1. 方式一

1. 创建一个目录如Mytest, 进入该目录, 在该目录下打开命令行窗口

2. 启动openssl

   ```shell
   openssl    # 执行该命令即可
   ```

3. 使用openssl工具生成一个RSA私钥, 注意：生成私钥，需要提供一个至少4位的密码。

   ```shell
   genrsa -des3 -out server.key 2048  // 2048私钥的位数，可以不指定，默认值：？？？
   	- des3: 使用3des对私钥进行加密，//使用req参数的可以不指定这个参数，加下面
   ```

4. 生成CSR（证书签名请求）

   会引导我们填写申请方的信息：国家，省份，城市，部门…, 格式是pem格式

   ```shell
   req -new -key server.key -out server.csr
   
   
   #查看请求
   req -in server.csr -text
   ```

5. 删除私钥中的密码, 第一步给私钥文件设置密码是必须要做的, 如果不想要可以删掉

   ```shell
   rsa -in server.key -out server.key
   	-out 参数后的文件名可以随意起
   ```

6. 生成自签名证书

   ```shell
   x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
   
   
   #生成的证书是pem进行base64编码的
   #查看方式
   ```



> 在Windows下安装，Openssl-Win64.exe
>
> 
>
> 进入到：C:\Program Files\OpenSSL-Win64\bin\openssl.exe
>
> 右键单击->管理员运行 -> OPenSSL >
>
> 如果不是管理员打开: Permission Denied —> 权限不够
>
> 执行 : genrsa -des3 -out server.key 2048





==自签名证书，自己颁发给自己，自己验证自己。==



## 2. 方式二

不需要生成csr，直接生成证书，没有指定Subject相关的数据，所以还会引导输入

```sh
openssl req -x509 -newkey rsa:4096 -keyout server2.key -out cert.crt -days 365 -nodes
```

>  -nodes 不设置密码



解析证书：

```sh
 openssl x509 -in cert.pem  -text
```



# 常见的证书格式

## 1. pem格式

我们使用openssl生成的都是pem格式的

解析过程

```sh
openssl x509 -in cert.crt -text
```



- Privacy Enhanced Mail(信封)

- 查看内容，以"-----BEGIN..."开头，以"-----END..."结尾。

- 查看PEM格式证书的信息：

  ```sh
  `Apache和*NIX服务器偏向于使用这种编码格式。
  openssl x509 -in certificate.pem -text -noout
  ```

  

## 2. der格式

我们使用Windows导出的可以使der

对于der格式的，解析方式如下：

```sh
openssl x509 -in itcastcrt.cer -text -inform der  // 额外的参数 -inform der
```



- Distinguished Encoding Rules

- 打开看是二进制格式，不可读。

- Java和Windows服务器偏向于使用这种编码格式。

- 查看DER格式证书的信息

  ```sh
  `der是格式，与证书的后缀名没有直接关系
  openssl x509 -in certificate.der -inform der -text -noout  `请试试-pubkey参数
  ```



## 3. windows导数格式选择

![image-20190309113548996](https://ws1.sinaimg.cn/large/006tKfTcgy1g0we0twfz7j315s0u04qp.jpg)





## 2. PKI的组成要素



PKI的组成要素主要有以下三个：

- 用户 --- 使用PKI的人
- 认证机构 --- 颁发证书的人
- 仓库 --- 保存证书的数据库

![](https://ws4.sinaimg.cn/large/006tNc79ly1fyzcj2aivdj30fh0do0tg.jpg)



# HTTP, HTTPS, SSL/TLS

- HTTPS = HTTP + SSL/TLS

- 早期的版本SSL （3.0之后叫TLS）

- 现在：TLS

- 1.0 TLS = 3.0 SSL

- 1.1  TLS = 3.1 SSL, 目前版本TLS1.2



# 关系图示

![](https://ws1.sinaimg.cn/large/006tKfTcgy1g0weuhfpt5j31ia0p2dnu.jpg)



# SSL通信图示

![](https://ws3.sinaimg.cn/large/006tNc79ly1fyzcjfegbmj30jd0bwq42.jpg)



# SSL协议细节（拓展）



![tls协议分析](https://ws3.sinaimg.cn/large/006tKfTcgy1g0wet8z8yvj31sy0u01ky.jpg)





