# GoFrame Template For SingleRepo

Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.

### DAO 生成器

```
gf gen dao -l "mysql:root:root@tcp(127.0.0.1:3306)/db" -p "app"

gf gen service -s "app/service/logic" -d "app/service"
```


### 收单

思想：传统联单，如三联单：买家，卖家，供货商

数字化之后：按照不同角色划分不同纬度的数据库，数据状态通过内部通知更新其他数据库的数据。


### 模块

Nginx 是一个高性能的 HTTP 和反向代理服务器，代码完全用 C 实现，基于它的高性能以及诸多优点，我们可以把它设置为 gomall 的前置服务器，实现负载均衡或 HTTPS 前置服务器等。

```
upstream goauth {
    server 127.0.0.1:8001;
}
upstream goconsole {
    server 127.0.0.1:8002;
}
upstream gomobile {
    server 127.0.0.1:8003;
}
upstream gopartner {
    server 127.0.0.1:8004;
}
upstream goportal {
    server 127.0.0.1:8005;
}
upstream goseller {
    server 127.0.0.1:8006;
}
upstream goshop {
    server 127.0.0.1:8007;
}
upstream gostore {
    server 127.0.0.1:8008;
}
upstream gosupplier {
    server 127.0.0.1:8009;
}
upstream gouser {
    server 127.0.0.1:8010;
}
upstream goweb {
    server 127.0.0.1:8000;
}
upstream websocket {
    ip_hash;
    server 127.0.0.1:9000;
}

server {
    listen 80; 
    server_name demo.gomall.dev;

    location /auth {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass http://goauth;
    }

    location /console {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass http://goconsole;
    }

    location /mobile {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass http://gomobile;
    }

    location /partner {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass http://gopartner;
    }

    location /portal {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass http://goportal;
    }

    location /seller {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass http://goseller;
    }

    location /shop {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass http://goshop;
    }

    location /store {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass http://gostore;
    }

    location /supplier {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass http://gosupplier;
    }

    location /user {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass http://gouser;
    }

    location / {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass http://goweb;
    }
    
    location /ws {
        proxy_http_version 1.1;
        proxy_set_header Upgrade websocket;
        proxy_set_header Connection "Upgrade";
        
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header Host $http_host;
    
        proxy_read_timeout 60s ;
        
        proxy_pass http://websocket;
    }
}
```