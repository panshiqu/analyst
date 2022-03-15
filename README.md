```
/home/ubuntu
│
└───analyst
│   │   analyst
│   │   monit.sh
│   │   config.json
│
└───router
│   │   index.ts
│   │   package.json
│   │   tsconfig.json
│
└───conf.d
│   │   line.conf
│   │   tg.conf
│
└───line
    │   index.html

/tmp
│   alerts.json
│   prices.json
```

```
sudo timedatectl set-timezone Asia/Singapore
date

sudo apt update
sudo apt install nginx
sudo apt install certbot
sudo apt install npm
cd router
npm install

sudo certbot certonly --webroot
sudo vi /etc/nginx/nginx.conf
sudo systemctl reload nginx

nohup npx ts-node index.ts > output.log 2>&1 &
./monit.sh start analyst
```
