-- Adminer 4.8.1 MySQL 5.7.35 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `auth`;
CREATE TABLE `auth` (
  `aid` int(8) unsigned zerofill NOT NULL AUTO_INCREMENT COMMENT '编号',
  `icon` varchar(32) NOT NULL DEFAULT 'fa' COMMENT '图标',
  `title` varchar(16) NOT NULL COMMENT '标题',
  `href` varchar(64) NOT NULL COMMENT '页面地址',
  `target` varchar(16) DEFAULT '_self' COMMENT '目标',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型 0-菜单 1-页面 2-操作',
  `parent` int(8) unsigned zerofill DEFAULT NULL COMMENT '依赖',
  PRIMARY KEY (`aid`),
  KEY `parent` (`parent`),
  CONSTRAINT `auth_ibfk_1` FOREIGN KEY (`parent`) REFERENCES `auth` (`aid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `auth` (`aid`, `icon`, `title`, `href`, `target`, `type`, `parent`) VALUES
(00000013,	'fa fa-address-book-o',	'角色与权限管理',	'',	'_self',	0,	NULL),
(00000015,	'fa fa-500px',	'角色列表',	'src/page/role/index.html',	'_self',	1,	00000035),
(00000016,	'fa',	'添加角色',	'src/page/role/add.html',	'_self',	1,	00000035),
(00000017,	'fa',	'权限列表',	'src/page/auth/index.html',	'_self',	1,	00000034),
(00000018,	'fa',	'添加权限',	'src/page/auth/add.html',	'_self',	1,	00000034),
(00000020,	'fa',	'getAllRole',	'',	'_self',	2,	00000015),
(00000021,	'fa',	'addRole',	'',	'_self',	2,	00000016),
(00000022,	'fa',	'getAllAuth',	'',	'_self',	2,	00000017),
(00000023,	'fa',	'addAuth',	'',	'_self',	2,	00000018),
(00000028,	'fa',	'updateRole',	'',	'_self',	2,	00000015),
(00000029,	'fa',	'deleteRole',	'',	'_self',	2,	00000015),
(00000030,	'fa',	'authAssign',	'',	'_self',	2,	00000015),
(00000031,	'fa',	'deleteAuth',	'',	'_self',	2,	00000017),
(00000032,	'fa',	'updateAuth',	'',	'_self',	2,	00000017),
(00000034,	'fa fa',	'权限管理',	'',	'_self',	1,	00000013),
(00000035,	'fa fa',	'角色管理',	'',	'_self',	1,	00000013),
(00000038,	'fa fa',	'问题与比赛管理',	'',	'_self',	0,	NULL),
(00000039,	'fa fa',	'标签管理',	'',	'_self',	1,	00000038),
(00000040,	'fa fa',	'标签列表',	'src/page/tag/index.html',	'_self',	1,	00000039),
(00000041,	'fa fa',	'添加标签',	'src/page/tag/add.html',	'_self',	1,	00000039),
(00000042,	'fa fa',	'getAllTag',	'',	'_self',	2,	00000040),
(00000043,	'fa fa',	'addTag',	'',	'_self',	2,	00000041),
(00000044,	'fa fa',	'deleteTag',	'',	'_self',	2,	00000040),
(00000045,	'fa fa',	'updateTag',	'',	'_self',	2,	00000040),
(00000046,	'fa fa',	'OJ管理',	'',	'_self',	0,	NULL),
(00000047,	'fa fa',	'通知管理',	'',	'_self',	1,	00000046),
(00000048,	'fa fa',	'通知列表',	'src/page/notice/index.html',	'_self',	1,	00000047),
(00000049,	'fa fa',	'添加通知',	'src/page/notice/add.html',	'_self',	1,	00000047),
(00000050,	'fa fa',	'getAllNotice',	'',	'_self',	2,	00000048),
(00000051,	'fa fa',	'deleteNotice',	'',	'_self',	2,	00000048),
(00000052,	'fa fa',	'updateNotice',	'',	'_self',	2,	00000048),
(00000053,	'fa fa',	'addNotice',	'',	'_self',	2,	00000049),
(00000054,	'fa fa',	'问题管理',	'',	'_self',	1,	00000038),
(00000055,	'fa fa',	'比赛管理',	'',	'_self',	1,	00000038),
(00000056,	'fa fa',	'提交管理',	'',	'_self',	1,	00000038),
(00000057,	'fa fa',	'问题列表',	'src/page/problem/index.html',	'_self',	1,	00000054),
(00000058,	'fa fa',	'添加问题',	'src/page/problem/add.html',	'_self',	1,	00000054),
(00000060,	'fa fa',	'比赛列表',	'src/page/contest/index.html',	'_self',	1,	00000055),
(00000061,	'fa fa',	'添加比赛',	'src/page/contest/add.html',	'_self',	1,	00000055),
(00000062,	'fa fa',	'提交列表',	'src/page/submit/index.html',	'_self',	1,	00000056),
(00000064,	'fa fa',	'用户管理',	'',	'_self',	1,	00000046),
(00000065,	'fa fa',	'用户列表',	'src/page/user/index.html',	'_self',	1,	00000064),
(00000066,	'fa fa',	'提交情况',	'src/page/user/submit.html',	'_self',	1,	00000064),
(00000067,	'fa fa',	'getAllUser',	'',	'_self',	2,	00000065),
(00000068,	'fa fa',	'addUser',	'',	'_self',	2,	00000065),
(00000069,	'fa fa',	'deleteUser',	'',	'_self',	2,	00000065),
(00000070,	'fa fa',	'updateUser',	'',	'_self',	2,	00000065),
(00000071,	'fa fa',	'roleAssign',	'',	'_self',	2,	00000065),
(00000072,	'fa fa',	'getUserSubmit',	'',	'_self',	2,	00000066),
(00000073,	'fa fa',	'getAllProblem',	'',	'_self',	2,	00000057),
(00000074,	'fa fa',	'deleteProblem',	'',	'_self',	2,	00000057),
(00000075,	'fa fa',	'updateProblem',	'',	'_self',	2,	00000057),
(00000076,	'fa fa',	'uploadData',	'',	'_self',	2,	00000057),
(00000077,	'fa fa',	'addProblem',	'',	'_self',	2,	00000058),
(00000078,	'fa fa',	'getAllContest',	'',	'_self',	2,	00000060),
(00000079,	'fa fa',	'deleteContest',	'',	'_self',	2,	00000060),
(00000080,	'fa fa',	'updateContest',	'',	'_self',	2,	00000060),
(00000081,	'fa fa',	'addContest',	'',	'_self',	2,	00000061),
(00000082,	'fa fa',	'getAllSubmit',	'',	'_self',	2,	00000062),
(00000083,	'fa fa',	'rejudge',	'',	'_self',	2,	00000062),
(00000086,	'fa fa-adn',	'发送气球',	'src/page/balloon/index.html',	'_self',	1,	00000055),
(00000087,	'fa fa-adn',	'getBalloonStatus',	'',	'_self',	2,	00000086),
(00000088,	'fa fa',	'setBalloonStatus',	'',	'_self',	2,	00000086),
(00000089,	'fa fa',	'代码打印请求',	'src/page/print/index.html',	'_self',	1,	00000038),
(00000090,	'fa fa',	'getPrintRequest',	'',	'_self',	2,	00000089);

DROP TABLE IF EXISTS `contest`;
CREATE TABLE `contest` (
  `contest_id` int(11) NOT NULL AUTO_INCREMENT,
  `contest_name` varchar(255) NOT NULL,
  `begin_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `end_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `frozen` double NOT NULL DEFAULT '0.2',
  `problems` varchar(512) NOT NULL,
  `prize` varchar(512) NOT NULL COMMENT '奖励比例',
  `colors` varchar(512) NOT NULL COMMENT '题目对应颜色',
  `rule` int(11) NOT NULL DEFAULT '0' COMMENT '比赛规则',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`contest_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


SET NAMES utf8mb4;

DROP TABLE IF EXISTS `contest_problem`;
CREATE TABLE `contest_problem` (
  `contest_id` int(11) NOT NULL,
  `problem_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`contest_id`,`problem_id`),
  KEY `contest_id` (`contest_id`),
  CONSTRAINT `contest_problem_ibfk_1` FOREIGN KEY (`contest_id`) REFERENCES `contest` (`contest_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `contest_users`;
CREATE TABLE `contest_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `contest_id` int(11) NOT NULL COMMENT '比赛序号',
  `user_id` int(11) NOT NULL COMMENT '用户编号',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0-正常参赛，1-打星参赛，2-女队，3-作弊',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `discuss`;
CREATE TABLE `discuss` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `contest_id` int(11) NOT NULL COMMENT '比赛id',
  `problem_id` int(11) NOT NULL COMMENT '题目id',
  `user_id` int(11) NOT NULL COMMENT '提问者id',
  `title` text NOT NULL COMMENT '标题',
  `content` text NOT NULL COMMENT '提问内容',
  `time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提问时间',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否有管理员回复(0-没有，1-有)',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `notice`;
CREATE TABLE `notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `title` varchar(64) NOT NULL COMMENT '标题',
  `content` text NOT NULL COMMENT '内容',
  `link` varchar(128) NOT NULL DEFAULT '""' COMMENT '跳转链接',
  `begintime` datetime NOT NULL COMMENT '开始时间',
  `endtime` datetime NOT NULL COMMENT '结束时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `notification`;
CREATE TABLE `notification` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` text NOT NULL,
  `content` text NOT NULL,
  `submit_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `contest_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0-不可用 1-可用',
  `end_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `contest_id` (`contest_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `notification_ibfk_1` FOREIGN KEY (`contest_id`) REFERENCES `contest` (`contest_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `print_log`;
CREATE TABLE `print_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `submit_id` int(11) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  `print_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `problem`;
CREATE TABLE `problem` (
  `problem_id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(128) NOT NULL,
  `background` text,
  `describe` text,
  `input_format` text,
  `output_format` text,
  `hint` text,
  `public` int(11) NOT NULL DEFAULT '1',
  `source` varchar(512) DEFAULT NULL,
  `time` float NOT NULL DEFAULT '0',
  `memory` int(11) NOT NULL DEFAULT '0',
  `type` varchar(20) NOT NULL DEFAULT 'Normal',
  `tag` varchar(512) DEFAULT NULL,
  `path` varchar(128) NOT NULL DEFAULT ' ' COMMENT '数据路径',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`problem_id`),
  KEY `problem_id` (`problem_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `problem_submit_log`;
CREATE TABLE `problem_submit_log` (
  `problem_id` int(11) NOT NULL,
  `ac` int(11) NOT NULL,
  `wa` int(11) NOT NULL,
  `tle` int(11) NOT NULL,
  `mle` int(11) NOT NULL,
  `re` int(11) NOT NULL,
  `se` int(11) NOT NULL,
  `ce` int(11) NOT NULL,
  PRIMARY KEY (`problem_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `reply`;
CREATE TABLE `reply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `discuss_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `content` text NOT NULL,
  `time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `rid` int(8) unsigned zerofill NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL COMMENT '名字',
  `desc` varchar(128) NOT NULL COMMENT '描述',
  PRIMARY KEY (`rid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `role` (`rid`, `name`, `desc`) VALUES
(00000001,	'超级管理员',	'admin'),
(00000002,	'？？？',	'？？？');

DROP TABLE IF EXISTS `role_auth`;
CREATE TABLE `role_auth` (
  `rid` int(8) unsigned zerofill NOT NULL,
  `aid` int(8) unsigned zerofill NOT NULL,
  UNIQUE KEY `rid_aid` (`rid`,`aid`),
  KEY `aid` (`aid`),
  CONSTRAINT `role_auth_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `role` (`rid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `role_auth_ibfk_2` FOREIGN KEY (`aid`) REFERENCES `auth` (`aid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `role_auth` (`rid`, `aid`) VALUES
(00000001,	00000013),
(00000001,	00000015),
(00000001,	00000016),
(00000001,	00000017),
(00000001,	00000018),
(00000001,	00000020),
(00000001,	00000021),
(00000001,	00000022),
(00000001,	00000023),
(00000001,	00000028),
(00000001,	00000029),
(00000001,	00000030),
(00000001,	00000031),
(00000001,	00000032),
(00000001,	00000034),
(00000001,	00000035),
(00000001,	00000038),
(00000001,	00000039),
(00000001,	00000040),
(00000001,	00000041),
(00000001,	00000042),
(00000001,	00000043),
(00000001,	00000044),
(00000001,	00000045),
(00000001,	00000046),
(00000001,	00000047),
(00000001,	00000048),
(00000001,	00000049),
(00000001,	00000050),
(00000001,	00000051),
(00000001,	00000052),
(00000001,	00000053),
(00000001,	00000054),
(00000001,	00000055),
(00000001,	00000056),
(00000001,	00000057),
(00000001,	00000058),
(00000001,	00000060),
(00000001,	00000061),
(00000001,	00000062),
(00000001,	00000064),
(00000001,	00000065),
(00000001,	00000066),
(00000001,	00000067),
(00000001,	00000068),
(00000001,	00000069),
(00000001,	00000070),
(00000001,	00000071),
(00000001,	00000072),
(00000001,	00000073),
(00000001,	00000074),
(00000001,	00000075),
(00000001,	00000076),
(00000001,	00000077),
(00000001,	00000078),
(00000001,	00000079),
(00000001,	00000080),
(00000001,	00000081),
(00000001,	00000082),
(00000001,	00000083),
(00000001,	00000086),
(00000001,	00000087),
(00000001,	00000088),
(00000001,	00000089),
(00000001,	00000090);

DROP TABLE IF EXISTS `rotation`;
CREATE TABLE `rotation` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `title` varchar(32) NOT NULL COMMENT '标题',
  `url` varchar(128) NOT NULL COMMENT '链接',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `sample`;
CREATE TABLE `sample` (
  `sample_id` int(11) NOT NULL AUTO_INCREMENT,
  `problem_id` int(11) NOT NULL,
  `input` varchar(512) DEFAULT NULL,
  `output` varchar(512) DEFAULT NULL,
  PRIMARY KEY (`sample_id`),
  KEY `problem_id` (`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `submit`;
CREATE TABLE `submit` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `user_id` int(11) NOT NULL COMMENT '用户编号',
  `nick` varchar(64) NOT NULL COMMENT '用户昵称',
  `problem_id` int(11) NOT NULL COMMENT '题目编号',
  `contest_id` int(11) DEFAULT '0' COMMENT '比赛id',
  `source_code` text NOT NULL COMMENT '源代码',
  `language` tinyint(4) NOT NULL DEFAULT '1' COMMENT '语言',
  `status` varchar(16) NOT NULL DEFAULT 'judging' COMMENT '状态',
  `msg` mediumtext NOT NULL COMMENT 'judge message',
  `time` bigint(20) NOT NULL DEFAULT '0',
  `memory` int(11) NOT NULL,
  `submit_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提交时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `problem_id` (`problem_id`),
  KEY `submit_time` (`submit_time`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `support_language`;
CREATE TABLE `support_language` (
  `name` varchar(100) NOT NULL,
  `is_support` int(11) NOT NULL COMMENT '0不支持，1支持',
  `build_path` varchar(160) NOT NULL COMMENT '编译脚本路径',
  `with_proc` int(11) NOT NULL DEFAULT '0' COMMENT '是否挂载proc，0-不挂载,1-挂载',
  `with_rootfs` int(11) NOT NULL COMMENT '0-不需要rootfs，1-需要',
  `env_path` varchar(160) NOT NULL COMMENT 'rootfs路径',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '标签序号',
  `name` varchar(64) NOT NULL COMMENT '标签名字',
  `description` varchar(200) NOT NULL COMMENT '描述',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '标签状态(0-禁用,1-正常)',
  `color` varchar(16) NOT NULL DEFAULT '#FFFFFF' COMMENT '标签颜色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `nick` varchar(25) NOT NULL,
  `password` varchar(32) NOT NULL,
  `realname` varchar(32) NOT NULL,
  `avatar` varchar(128) NOT NULL DEFAULT '""',
  `school` varchar(64) NOT NULL,
  `major` varchar(64) NOT NULL,
  `grade` int(11) NOT NULL DEFAULT '2020' COMMENT '年级',
  `class` varchar(32) NOT NULL,
  `contact` varchar(64) NOT NULL,
  `identity` int(11) NOT NULL DEFAULT '0',
  `desc` text,
  `mail` varchar(64) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `register_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  PRIMARY KEY (`user_id`,`register_time`),
  KEY `user_id` (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

INSERT INTO `users` (`user_id`, `nick`, `password`, `realname`, `avatar`, `school`, `major`, `grade`, `class`, `contact`, `identity`, `desc`, `mail`, `status`, `register_time`) VALUES
(1,	'Coding_Panda',	'39f264a9c984299634d03e9b29718c24',	'Zhiway Zhang',	'../uploads/image/20200214/fc3d5f691e86c9f621621682c57de59b.jpg',	'武汉理工大学',	'软件工程',	2020,	'软件zy1902',	'17515099009',	1,	'',	'zhiwayzhang@outlook.com',	0,	'2021-05-27 09:48:08'),
(2,	'wqyyy',	'87d9bb400c0634691f0e3baaf1e2fd0d',	'wqyyy',	'../uploads/image/20200214/fc3d5f691e86c9f621621682c57de59b.jpg',	'1',	'1',	2020,	'1',	'11111111111',	1,	'',	'1@qq.com',	0,	'2021-05-27 09:48:55'),
(462,	'zzwzzw',	'eb5637cef0d0ba8a35a8091116d07561',	'zzz',	'../uploads/image/20200214/fc3d5f691e86c9f621621682c57de59b.jpg',	'zzz',	'zz',	2020,	'zz',	'111',	1,	'',	'1@1.com',	0,	'2021-05-27 09:51:08');

DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `user_id` int(8) unsigned zerofill NOT NULL,
  `rid` int(8) unsigned zerofill NOT NULL,
  UNIQUE KEY `uid_rid` (`user_id`,`rid`),
  KEY `rid` (`rid`),
  CONSTRAINT `user_role_ibfk_2` FOREIGN KEY (`rid`) REFERENCES `role` (`rid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `user_role` (`user_id`, `rid`) VALUES
(00000001,	00000001),
(00000002,	00000001),
(00000001,	00000002);

DROP TABLE IF EXISTS `user_submit_log`;
CREATE TABLE `user_submit_log` (
  `user_id` int(11) NOT NULL,
  `ac` int(11) NOT NULL,
  `wa` int(11) NOT NULL,
  `tle` int(11) NOT NULL,
  `mle` int(11) NOT NULL,
  `re` int(11) NOT NULL,
  `se` int(11) NOT NULL,
  `ce` int(11) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


-- 2021-09-05 14:44:50
