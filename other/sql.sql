
create database beegosushuo;
use beegosushuo;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '文章标题',
  `content` text COMMENT '正文',
  `userid`  int(11)  NOT NULL DEFAULT 0 COMMENT '作者id',
  `time` int(10) NOT NULL DEFAULT 0 COMMENT '发布时间',
  `views` int NOT NULL DEFAULT 0 COMMENT '阅读次数',
  `comments` int NOT NULL DEFAULT 0 COMMENT '评论次数',
  `classifyid` int DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `userid` (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='文章';


CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(50) NOT NULL DEFAULT '' COMMENT '手机号',
  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `time` int(10) NOT NULL DEFAULT 0 COMMENT '注册时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `account` (`account`),
  UNIQUE KEY `nickname` (`nickname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户';

CREATE TABLE IF NOT EXISTS `comment` (
  `articleid` int(11) NOT NULL,
  `usersid` int(11) NOT NULL,
  `content` text COMMENT '',
  `time` int(10) NOT NULL DEFAULT 0 COMMENT '时间',
  KEY `articleid` (`articleid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户';

CREATE TABLE IF NOT EXISTS `classify`(
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '名称',
  UNIQUE KEY(`name`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章分类';

insert into classify(name) values('default');
update classify set id = 0 where name='default';