#!/bin/bash

cd /home/ubuntu/project/

#Delete old stuff
rm -rf ./*

#Clone repo
git clone git@github.com:11Challenge48h/Challenge48h.git .

#Install dependencies
npm install

#Build app
npm run build

#Copy distributing files to proper directory
sudo cp -r ./dist/* /var/www/html/

#Set correct rights on destination folder
sudo chown www-data:www-data -R /var/www/html/