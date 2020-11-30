/*

Source Server         : 测试环境
Source Host           : localhost:3306
Source Database       : never_todo

Target Server Type    : MYSQL

*/

DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(100) DEFAULT '',
  `createTime` datetime NOT NULL DEFAULT now(),
  `updateTime` datetime NOT NULL DEFAULT now(),
  `status` TINYINT(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tags`; 
CREATE TABLE `tags` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(10) DEFAULT '',
  `desc` varchar(20) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `task_tags`; 
CREATE TABLE `task_tags` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `task_id` int(10) NOT NULL,
  `tag_id` int(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* 以下是测试数据插入样板， 请仿照样板设计数据 */
/*
INSERT INTO `tasks`(`content`) VALUES ('task_1');
INSERT INTO `tasks`(`content`) VALUES ('task_2');
INSERT INTO `tasks`(`content`) VALUES ('task_3');

INSERT INTO `tags`(`content`, `description`) VALUES ('tag_1', 'desc_1');
INSERT INTO `tags`(`content`, `description`) VALUES ('tag_2', 'desc_2');
INSERT INTO `tags`(`content`, `description`) VALUES ('tag_3', 'desc_3');

INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (1, 2);
INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (1, 3);
INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (2, 1);
INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (2, 3);
INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (3, 1);
INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (3, 2);
*/
INSERT INTO `tasks`(`content`) VALUES ('记得换内裤');
INSERT INTO `tasks`(`content`) VALUES ('做广播体操');
INSERT INTO `tasks`(`content`) VALUES ('一边换内裤一边做广播体操');
INSERT INTO `tasks`(`content`) VALUES ('跳起来锤你的大脚拇指');

INSERT INTO `tags`(`content`, `desc`) VALUES ('生活', '生活中的小事，比如洗内裤。');
INSERT INTO `tags`(`content`, `desc`) VALUES ('健身', '希望你也能像我一样成为彭于晏。');
INSERT INTO `tags`(`content`, `desc`) VALUES ('内裤', '内裤是yyds！');
INSERT INTO `tags`(`content`, `desc`) VALUES ('叽里呱啦', '只是凑数用的。');

INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (1, 1);
INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (1, 3);
INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (2, 2);
INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (3, 1);
INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (3, 2);
INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (3, 3);
INSERT INTO `task_tags`(`task_id`, `tag_id`) VALUES (3, 4);





