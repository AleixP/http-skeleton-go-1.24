CREATE TABLE IF NOT EXISTS fruits (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` varchar(255) NOT NULL,
    color ENUM('red', 'green', 'orange', 'yellow', 'other'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_name (`name`)
);

INSERT INTO fruits VALUES ('apple', 'red');