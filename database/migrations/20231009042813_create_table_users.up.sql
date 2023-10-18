CREATE TABLE User (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255),
    password VARCHAR(255),
    token VARCHAR(255),
    role ENUM('user', 'admin') NOT NULL DEFAULT 'user',
    UNIQUE KEY (username),
    UNIQUE KEY (email)
);
