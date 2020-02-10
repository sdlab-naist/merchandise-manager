# merchandise-manager
PBL2019

merchandise-manager on sdlab


Front deployment settings
使っているサーバー：Nginx

アクセスの仕方
sshでsdlab@163.221.29.46にログイン
cdコマンドでmerchandise-manegerに移動
git pull
cd ~/merchandise-maneger/Src/Frontend/*

AddView
http://163.221.29.46/AddItem/add_item_view.html


ブランチを変えたときの設定方法
* git pull origin <branch_name>
* cd /var/www/html
* sudo cp -r ~/merchandise-manager/Src/Fronted/* .
