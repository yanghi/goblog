create table if not exists gb_user(
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(64) NOT NULL COMMENT '用户名',
	password VARCHAR(64) NOT NULL COMMENT '密码',
	time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(),
	PRIMARY KEY(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';