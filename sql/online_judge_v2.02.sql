-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： 127.0.0.1:3306
-- 生成日期： 2021-02-02 05:11:59
-- 服务器版本： 5.7.26
-- PHP 版本： 7.3.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `online_judge_dev`
--
CREATE DATABASE IF NOT EXISTS `online_judge_dev` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `online_judge_dev`;

-- --------------------------------------------------------

--
-- 表的结构 `auth`
--

DROP TABLE IF EXISTS `auth`;
CREATE TABLE IF NOT EXISTS `auth` (
  `aid` int(8) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '编号',
  `icon` varchar(32) NOT NULL DEFAULT 'fa' COMMENT '图标',
  `title` varchar(16) NOT NULL COMMENT '标题',
  `href` varchar(64) NOT NULL COMMENT '页面地址',
  `target` varchar(16) DEFAULT '_self' COMMENT '目标',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型 0-菜单 1-页面 2-操作',
  `parent` int(8) UNSIGNED ZEROFILL DEFAULT NULL COMMENT '依赖',
  PRIMARY KEY (`aid`),
  KEY `parent` (`parent`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `contest`
--

DROP TABLE IF EXISTS `contest`;
CREATE TABLE IF NOT EXISTS `contest` (
  `contest_id` int(11) NOT NULL AUTO_INCREMENT,
  `contest_name` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
  `begin_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `end_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `frozen` double NOT NULL DEFAULT '0.2',
  `problems` varchar(512) NOT NULL,
  `prize` varchar(512) NOT NULL COMMENT '奖励比例',
  `colors` varchar(512) CHARACTER SET utf8mb4 NOT NULL COMMENT '题目对应颜色',
  `rule` int(2) NOT NULL DEFAULT '0' COMMENT '比赛规则',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`contest_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `contest_users`
--

DROP TABLE IF EXISTS `contest_users`;
CREATE TABLE IF NOT EXISTS `contest_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `contest_id` int(11) NOT NULL COMMENT '比赛序号',
  `user_id` int(11) NOT NULL COMMENT '用户编号',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0-正常参赛，1-打星参赛，2-女队，3-作弊',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `discuss`
--

DROP TABLE IF EXISTS `discuss`;
CREATE TABLE IF NOT EXISTS `discuss` (
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

-- --------------------------------------------------------

--
-- 表的结构 `notice`
--

DROP TABLE IF EXISTS `notice`;
CREATE TABLE IF NOT EXISTS `notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `title` varchar(64) NOT NULL COMMENT '标题',
  `content` text NOT NULL COMMENT '内容',
  `link` varchar(128) NOT NULL DEFAULT '""' COMMENT '跳转链接',
  `begintime` datetime NOT NULL COMMENT '开始时间',
  `endtime` datetime NOT NULL COMMENT '结束时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `notification`
--

DROP TABLE IF EXISTS `notification`;
CREATE TABLE IF NOT EXISTS `notification` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` text NOT NULL,
  `content` text NOT NULL,
  `submit_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `contest_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0-不可用 1-可用',
  PRIMARY KEY (`id`),
  KEY `contest_id` (`contest_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `problem`
--

DROP TABLE IF EXISTS `problem`;
CREATE TABLE IF NOT EXISTS `problem` (
  `problem_id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(128) NOT NULL,
  `background` text,
  `describe` text,
  `input_format` text,
  `output_format` text,
  `hint` text,
  `public` int(1) NOT NULL DEFAULT '1',
  `source` varchar(512) DEFAULT NULL,
  `time` float NOT NULL DEFAULT '0',
  `memory` int(20) NOT NULL DEFAULT '0',
  `type` varchar(20) NOT NULL DEFAULT 'Normal',
  `tag` varchar(512) DEFAULT NULL,
  `path` varchar(128) NOT NULL DEFAULT ' ' COMMENT '数据路径',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`problem_id`),
  KEY `problem_id` (`problem_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `problem_submit_log`
--

DROP TABLE IF EXISTS `problem_submit_log`;
CREATE TABLE IF NOT EXISTS `problem_submit_log` (
  `problem_id` int(11) NOT NULL,
  `ac` int(8) NOT NULL,
  `wa` int(8) NOT NULL,
  `tle` int(8) NOT NULL,
  `mle` int(8) NOT NULL,
  `re` int(8) NOT NULL,
  `se` int(8) NOT NULL,
  `ce` int(8) NOT NULL,
  PRIMARY KEY (`problem_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `reply`
--

DROP TABLE IF EXISTS `reply`;
CREATE TABLE IF NOT EXISTS `reply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `discuss_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `content` text NOT NULL,
  `time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `role`
--

DROP TABLE IF EXISTS `role`;
CREATE TABLE IF NOT EXISTS `role` (
  `rid` int(8) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL COMMENT '名字',
  `desc` varchar(128) NOT NULL COMMENT '描述',
  PRIMARY KEY (`rid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `role_auth`
--

DROP TABLE IF EXISTS `role_auth`;
CREATE TABLE IF NOT EXISTS `role_auth` (
  `rid` int(8) UNSIGNED ZEROFILL NOT NULL,
  `aid` int(8) UNSIGNED ZEROFILL NOT NULL,
  UNIQUE KEY `rid_aid` (`rid`,`aid`),
  KEY `aid` (`aid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `rotation`
--

DROP TABLE IF EXISTS `rotation`;
CREATE TABLE IF NOT EXISTS `rotation` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `title` varchar(32) NOT NULL COMMENT '标题',
  `url` varchar(128) NOT NULL COMMENT '链接',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `sample`
--

DROP TABLE IF EXISTS `sample`;
CREATE TABLE IF NOT EXISTS `sample` (
  `sample_id` int(11) NOT NULL AUTO_INCREMENT,
  `problem_id` int(11) NOT NULL,
  `input` varchar(512) DEFAULT NULL,
  `output` varchar(512) DEFAULT NULL,
  PRIMARY KEY (`sample_id`),
  KEY `problem_id` (`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `submit`
--

DROP TABLE IF EXISTS `submit`;
CREATE TABLE IF NOT EXISTS `submit` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `user_id` int(11) NOT NULL COMMENT '用户编号',
  `nick` varchar(64) CHARACTER SET utf8mb4 NOT NULL COMMENT '用户昵称',
  `problem_id` int(11) NOT NULL COMMENT '题目编号',
  `contest_id` int(11) DEFAULT '0' COMMENT '比赛id',
  `source_code` text NOT NULL COMMENT '源代码',
  `language` tinyint(4) NOT NULL DEFAULT '1' COMMENT '语言',
  `status` varchar(16) NOT NULL COMMENT '状态',
  `msg` mediumtext NOT NULL COMMENT 'judge message',
  `time` bigint(12) NOT NULL,
  `memory` int(11) NOT NULL,
  `submit_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提交时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `problem_id` (`problem_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8
PARTITION BY HASH (id)
(
PARTITION p0 ENGINE=MyISAM,
PARTITION p1 ENGINE=MyISAM,
PARTITION p2 ENGINE=MyISAM,
PARTITION p3 ENGINE=MyISAM,
PARTITION p4 ENGINE=MyISAM,
PARTITION p5 ENGINE=MyISAM,
PARTITION p6 ENGINE=MyISAM,
PARTITION p7 ENGINE=MyISAM,
PARTITION p8 ENGINE=MyISAM,
PARTITION p9 ENGINE=MyISAM,
PARTITION p10 ENGINE=MyISAM,
PARTITION p11 ENGINE=MyISAM,
PARTITION p12 ENGINE=MyISAM,
PARTITION p13 ENGINE=MyISAM,
PARTITION p14 ENGINE=MyISAM,
PARTITION p15 ENGINE=MyISAM,
PARTITION p16 ENGINE=MyISAM,
PARTITION p17 ENGINE=MyISAM,
PARTITION p18 ENGINE=MyISAM,
PARTITION p19 ENGINE=MyISAM,
PARTITION p20 ENGINE=MyISAM,
PARTITION p21 ENGINE=MyISAM,
PARTITION p22 ENGINE=MyISAM,
PARTITION p23 ENGINE=MyISAM,
PARTITION p24 ENGINE=MyISAM,
PARTITION p25 ENGINE=MyISAM,
PARTITION p26 ENGINE=MyISAM,
PARTITION p27 ENGINE=MyISAM,
PARTITION p28 ENGINE=MyISAM,
PARTITION p29 ENGINE=MyISAM,
PARTITION p30 ENGINE=MyISAM,
PARTITION p31 ENGINE=MyISAM
);

-- --------------------------------------------------------

--
-- 表的结构 `support_language`
--

DROP TABLE IF EXISTS `support_language`;
CREATE TABLE IF NOT EXISTS `support_language` (
  `name` varchar(100) NOT NULL,
  `is_support` int(1) NOT NULL COMMENT '0不支持，1支持',
  `build_path` varchar(160) NOT NULL COMMENT '编译脚本路径',
  `with_proc` int(11) NOT NULL DEFAULT '0' COMMENT '是否挂载proc，0-不挂载,1-挂载',
  `with_rootfs` int(11) NOT NULL COMMENT '0-不需要rootfs，1-需要',
  `env_path` varchar(160) NOT NULL COMMENT 'rootfs路径',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `tag`
--

DROP TABLE IF EXISTS `tag`;
CREATE TABLE IF NOT EXISTS `tag` (
  `id` int(4) NOT NULL AUTO_INCREMENT COMMENT '标签序号',
  `name` varchar(64) NOT NULL COMMENT '标签名字',
  `description` varchar(200) NOT NULL COMMENT '描述',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '标签状态(0-禁用,1-正常)',
  `color` varchar(16) NOT NULL DEFAULT '#FFFFFF' COMMENT '标签颜色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `nick` varchar(25) NOT NULL,
  `password` varchar(32) NOT NULL,
  `realname` varchar(32) NOT NULL,
  `avatar` varchar(128) NOT NULL DEFAULT '""',
  `school` varchar(64) NOT NULL,
  `major` varchar(64) NOT NULL,
  `grade` int(4) NOT NULL DEFAULT '2020' COMMENT '年级',
  `class` varchar(32) NOT NULL,
  `contact` varchar(64) NOT NULL,
  `identity` int(2) NOT NULL DEFAULT '0',
  `desc` text,
  `mail` varchar(64) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `register_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  PRIMARY KEY (`user_id`,`register_time`),
  KEY `user_id` (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user_role`
--

DROP TABLE IF EXISTS `user_role`;
CREATE TABLE IF NOT EXISTS `user_role` (
  `uid` int(8) UNSIGNED ZEROFILL NOT NULL,
  `rid` int(8) UNSIGNED ZEROFILL NOT NULL,
  UNIQUE KEY `uid_rid` (`uid`,`rid`),
  KEY `rid` (`rid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user_submit_log`
--

DROP TABLE IF EXISTS `user_submit_log`;
CREATE TABLE IF NOT EXISTS `user_submit_log` (
  `user_id` int(11) NOT NULL,
  `ac` int(8) NOT NULL,
  `wa` int(8) NOT NULL,
  `tle` int(8) NOT NULL,
  `mle` int(8) NOT NULL,
  `re` int(8) NOT NULL,
  `se` int(8) NOT NULL,
  `ce` int(8) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4;

--
-- 限制导出的表
--

--
-- 限制表 `auth`
--
ALTER TABLE `auth`
  ADD CONSTRAINT `auth_ibfk_1` FOREIGN KEY (`parent`) REFERENCES `auth` (`aid`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- 限制表 `notification`
--
ALTER TABLE `notification`
  ADD CONSTRAINT `notification_ibfk_1` FOREIGN KEY (`contest_id`) REFERENCES `contest` (`contest_id`) ON DELETE CASCADE;

--
-- 限制表 `role_auth`
--
ALTER TABLE `role_auth`
  ADD CONSTRAINT `role_auth_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `role` (`rid`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `role_auth_ibfk_2` FOREIGN KEY (`aid`) REFERENCES `auth` (`aid`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- 限制表 `user_role`
--
ALTER TABLE `user_role`
  ADD CONSTRAINT `user_role_ibfk_2` FOREIGN KEY (`rid`) REFERENCES `role` (`rid`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
