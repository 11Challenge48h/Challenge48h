#!/bin/bash

cd /home/ubuntu/project/

#Build app
npm run build

#Copy distributing files to proper directory
sudo cp ./dist/* /var/www/html/

#Set correct rights on destination folder
sudo chown www-data:www-data -R /var/www/html/