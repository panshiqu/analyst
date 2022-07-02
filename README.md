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

```
GET [BTC ETH MATIC RawAddress ...],[Decimals] [Amount]
Get token price

SET [BTC ETH MATIC RawAddress ...] [±Price] [Hour range with sound]
Telegram when price +higher or -lower
SET [BTC ETH MATIC RawAddress ...]
Show sets, Send same to delete

ANA [Name] [Start: 0] [End: 99999999]
Analysis buy and sell through USDC
Name: PANSHI, ZHUGE or Space(blank) for telegram username
Response format:
Time Block
Amount TokenA > Amount TokenB
Price / Average price

EXAMPLES:
GET BTC,8 0.01
SET BTC -20000 8-13,15-22
ANA PANSHI 30000000
```
