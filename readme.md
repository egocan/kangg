## kangg 看过够CMS
本项目在[go-blog](https://github.com/1920853199/go-blog/)的基础上完善，后端基于[beego](https://github.com/astaxie/beego/)，后台前端：使用[X-admin](http://x.xuebingsi.com/)、[layui](https://github.com/sentsin/layui/)、MD编辑器使用[Editor.md](https://github.com/pandao/editor.md)；主页前端使用：bootstrap、jquery、[slick](https://github.com/kenwheeler/slick)等组件完成。[https://kanguogou.com](https://kanguogou.com)

### 主要功能

![后台](/images/kangg.png)

| 主要功能 | 功能描述                                               | 对应文件                                      |
| :------: | :----------------------------------------------------- | --------------------------------------------- |
|  一句话  | 主页顶部滚动展示，类似说说                             | aword相关文件，eg:aword.html、aword-edit.html |
| 用户管理 | 系统管理员                                             | user相关文件                                  |
| 站点设置 | 站点基本信息、SEO、banner、广告等                      | setting、ad、banner等                         |
| 导航管理 | 支持按分类设置/list.html?cate=ebook                    | menu相关                                      |
| 文章管理 | 文章、评论管理；文章MD编辑支持html标签、评论默认不展示 | article、review相关                           |
| 分类管理 | 分类新增别名，别名用于导航设置与前台检索               | cate相关                                      |
| 留言管理 | 留言默认不展示                                         | message相关                                   |
|  公众号  | 拉取公众号粉丝                                         | wechat相关                                    |
| 友情链接 | 友链管理                                               | link相关                                      |

新增：搜索功能、完善分页功能等

### Install 

1. 把[kangg](https://github.com/egocan/kangg)项目拉到本地

```
https://github.com/egocan/kangg.git
```

2. 新建数据库，导入数据库文件，数据库文件/database/kangg.sql

3. 修改项目配置信息

```
#conf/app.conf

appname = kangg
httpport = 8088
runmode = dev
EnableAdmin = false
sessionon = true
url = 127.0.0.1:8088
view = kangg		//主页模板

limit = 10
title = kanguogou
autograph = 如今的我，谈不上幸福，也谈不上不幸。
rootStaticExt = .txt,.html,.xml  //static根目录下允许直接被访问的文件后缀，方便robots、sitemap使用

[db]
dbType = mysql   // 可选 'mysql'，'postgres'
dbUser = root
dbPass = root
dbHost = 127.0.0.1
dbPort = 3306
dbName = kangg

[redis]
rHost = 127.0.0.1
rPort = 6379

[wechat]
AppID = xxxxxxx
AppSecret = xxxxxxx
Token = xxxxxxx
EncodingAESKey = xxxxxxx


```

4. 在kangg 根目录下执行go run main.go ，访问 http://127.0.0.1:8088 即可

   ```sh
   ./kangg            //linux 系统可使用已编译的文件运行，请保持kangg同目录下包含：kangg、static、conf、views文件与目录；常驻运行可结合nohup
   ```

5. 后台默认账号密码

   ```sh
   后台：http://127.0.0.1:8088/admin
   账号：user
   密码：123456
   ```

6. nginx代理示例 [参考](https://www.cnblogs.com/shanchen/p/11725799.html)

```
user www-data;
worker_processes auto;
pid /run/nginx.pid;
 
 
events {
  worker_connections 768;
  # multi_accept on;
}
 
 
http {
 
 
  ##
  # Basic Settings
  ##
 
 
  sendfile on;
  tcp_nopush on;
  tcp_nodelay on;
  keepalive_timeout 65;
  types_hash_max_size 2048;
  # server_tokens off;
 
 
  # server_names_hash_bucket_size 64;
  # server_name_in_redirect off;
 
 
  include /etc/nginx/mime.types;
  default_type application/octet-stream;
 
 
  ##
  # SSL Settings
  ##
 
 
  ssl_protocols TLSv1 TLSv1.1 TLSv1.2; # Dropping SSLv3, ref: POODLE
  ssl_prefer_server_ciphers on;
 
 
  ##
  # Logging Settings
  ##
 
 
  access_log /var/log/nginx/access.log;
  error_log /var/log/nginx/error.log;
 
 
  ##
  # Gzip Settings
  ##
 
 
  gzip on;
  gzip_disable "msie6";
 
 
  # gzip_vary on;
  # gzip_proxied any;
  # gzip_comp_level 6;
  # gzip_buffers 16 8k;
  # gzip_http_version 1.1;
  # gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;
 
 
  ##
  # Virtual Host Configs
  ##
 
 
  include /etc/nginx/conf.d/*.conf;
  include /etc/nginx/sites-enabled/*;
 
 
  server {
           listen       80;
          server_name  www.XXXXXX.com; #你的域名
 
 
          # 配置https
          ssl          on;
          ssl_certificate 1_XXXXXX.com_bundle.crt;   #你申请的nginx中的ssl证书
          ssl_certificate_key 2_XXXXXX.com.key;     #你申请的nginx中的key
          ssl_session_timeout 5m;
          ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
          ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;
          ssl_prefer_server_ciphers on;
 
 
 
 
        charset utf-8;
        access_log  /home/ubuntu/beego/access.log; #该目录建议是你的部署文件目录
 
 
        location /(css|js|fonts|img)/ {
            access_log off;
            expires 1d;
 
 
            root "/home/ubuntu/beego/static";#该文件是你beego项目中的static静态文件路径
          try_files $uri @backend;
        }
 
 
        location / {
            try_files /_not_exists_ @backend;
        }
 
 
        location @backend {
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header Host            $http_host;
 
 
            proxy_pass http://127.0.0.1:8088;  #代理指向的beego项目地址
        }
  }
}

```

### [主页截图](https://kanguogou.com)

![主页](/images/kanguogou.jpg)

