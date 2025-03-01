#### 安装

```
sudo apt update
// 安装mariadb
sudo apt install mariadb-server mariadb-client
// 开启服务
sudo systemctl enable mariadb
//看状态
sudo systemctl status mariadb
```

#### 本机测试
```
// 登录
sudo mysql -u root -p
// 新建DATABASE
CREATE DATABASE mydata;
// 查看
SHOW DATABASES;
EXIT;

```

#### 开启远端连接

```
nano /etc/mysql/mariadb.conf.d/50-server.cnf
// 改为0.0.0.0
bind-address = 0.0.0.0
// 重启
sudo systemctl restart mariadb
```


#### 更新用户权限 从任意主机连接，可以使用'%'通配符：

```
GRANT ALL PRIVILEGES ON *.* TO 'yourusername'@'%' IDENTIFIED BY 'yourpassword' WITH GRANT OPTION;
```