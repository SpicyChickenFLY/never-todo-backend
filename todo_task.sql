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
  `create_time` datetime NOT NULL DEFAULT now(),
  `update_time` datetime NOT NULL DEFAULT now(),
  `importance` int(10) NOT NULL DEFAULT 1,
  `status` TINYINT(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tags`; 
CREATE TABLE `tags` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(20) DEFAULT '',
  `desc` varchar(50) DEFAULT '',
  `color` varchar(10) DEFAULT '#AAAAAA',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `task_tags`; 
CREATE TABLE `task_tags` (
  `task_id` int(10) NOT NULL,
  `tag_id` int(10) NOT NULL,
  PRIMARY KEY (`task_id`, `tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `nick` varchar(100) NOT NULL,
  `create_time` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ALTER TABLE `task_tags` ADD FOREIGN KEY (`task_id`) REFERENCES `tasks`(`id`);
-- ALTER TABLE `task_tags` ADD FOREIGN KEY (`tag_id`) REFERENCES `tags`(`id`);

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





