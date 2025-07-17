CREATE DATABASE IF NOT EXISTS `svc_fruits_db`;

USE `svc_fruits_db`;

CREATE USER 'app'@'%' IDENTIFIED BY 'th3p4ssw0rd';
GRANT ALL PRIVILEGES ON svc_fruits_db.* TO 'app'@'%' WITH GRANT OPTION;

CREATE TABLE IF NOT EXISTS fruits (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` varchar(255) NOT NULL,
    color ENUM('red', 'green', 'orange', 'yellow', 'other'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_name (`name`)
);

INSERT INTO fruits VALUES ('apple', 'red');