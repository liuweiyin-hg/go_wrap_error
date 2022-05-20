# Readme

## core 目录
初始化 iris application
初始化 database

## user_app 目录
创建和查询 user
请求先进入 service.go
然后进入 dao.go 查数据库

## main
两个路径配到 main.go 里面<br>
在这个请求里边 /get_user/{id:uint}, 试验了两个函数 errors.As 和 errors.Is, 用于判断自定义 error<br>
<br>
![image](https://github.com/liuweiyin-hg/go_wrap_error/blob/master/WX20220521-011546%402xb.jpg)