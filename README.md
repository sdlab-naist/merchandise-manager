# merchandise-manager
PBL2019

merchandise-manager on sdlab

AddItem URL
http://163.221.29.46:13131/addItemHTML

Front deployment settings
使っているサーバー：Nginx

アクセスの仕方
*  `ssh sdlab@163.221.29.46`
*  `cd merchandise-maneger`  
*  `git pull`  
*  `cd ~/merchandise-maneger/Src/Frontend/` 

ブランチ切り替え
*  `git pull origin <branch_name>`  
*  `cd /var/www/html`  
*  `sudo cp -r ~/merchandise-manager/Src/Fronted/* .`  

デプロイ
* `sudo docker ps -a`
* find container id has name:great_stonebraker
* `sudo docker attach <containar id>`
* `C-c`
* `git pull`
* `cd /Src/Backend`
* `go run main_auth.go`
* `C-pq`
