#!/bin/bash

echo start executing seeds

docker exec -i http-skeleton-database mysql -uapp -pth3p4ssw0rd svc_fruits_db < database/seeds/fruits.sql