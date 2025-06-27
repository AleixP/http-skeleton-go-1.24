CREATE DATABASE IF NOT EXISTS `svc_fruits_db`;

USE `svc_fruits_db`;

CREATE USER 'app'@'%' IDENTIFIED BY 'th3p4ssw0rd';
GRANT ALL PRIVILEGES ON svc_fruits_db.* TO 'app'@'%' WITH GRANT OPTION;