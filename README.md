# 个人博客系统


基于Go语言和beego框架 开发的个人博客系统

## 编译安装说明：

1 . 设置GOPATH(安装目录)
    $ export GOPATH=/go/gocode/

2 . 下载安装
  $ go get github.com/lesroad/wblog

3 . cd 到 beego_blog 目录 执行
  main.go路径示例：go/gocode/src/beego项目/wblog/main.go
  $ bee run


账号： admin  密码 :admin

1 . 设置GOPATH(安装目录)

    $ export GOPATH=/path/to/


2 . 下载安装

    $ go get github.com/Echosong/beego_blog

4 . 加入数据库

   mysql 新建db_beego数据库把根目录 db_beego.sql 导入

5 . 修改 app.conf 配置

    #MYSQL地址
    dbhost = localhost

    #MYSQL端口
    dbport = 3306

    #MYSQL用户名
    dbuser = root

    #MYSQL密码
    dbpassword =

    #MYSQL数据库名称
    dbname = db_beego

    #MYSQL表前缀
    dbprefix = tb_

 6 . 运行

    cd 到 beego_blog 目录 执行
    $ bee run

 7 . 浏览器演示

http://localhost:8099 (前台)

http://localhost:8099/admin/login (后台)

qq:342705814
