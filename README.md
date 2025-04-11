# userManageSystem

该系统用到的数据库是MySQL8.0，需要用到数据库中两个表结构，建表语句奉上：

use usermanagesystem;

CREATE TABLE IF NOT EXISTS users (
id INT PRIMARY KEY AUTO_INCREMENT,
username VARCHAR(255) UNIQUE NOT NULL,
password VARCHAR(255) NOT NULL,
role ENUM('管理员','普通用户') NOT NULL DEFAULT '普通用户',
email VARCHAR(255) NULL,
avatar VARCHAR(255),
status TINYINT(1) NOT NULL DEFAULT 1,
last_login DATETIME,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
is_logged_in BOOLEAN DEFAULT FALSE
);

CREATE TABLE index_visits (
id INT AUTO_INCREMENT PRIMARY KEY,
visit_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

如果想借鉴并使用此系统可以下载压缩包或者通过git获取https的https://github.com/isPrisoner/userManageSystem.git

或者ssh的git@github.com:isPrisoner/userManageSystem.git

后两种方式要求已经安装过git才能成功

