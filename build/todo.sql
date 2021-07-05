-------------------------------------------------------
MYSQL_DEBUG found.libmysql started with the following: d :t :O,
/ tmp / client.trace -------------------------------------------------------
-- MySQL dump 10.13  Distrib 5.7.33, for Linux (x86_64)
--
-- Host: localhost    Database: never_todo
-- ------------------------------------------------------
-- Server version	5.7.33-debug-log
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;
/*!40101 SET NAMES utf8 */
;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */
;
/*!40103 SET TIME_ZONE='+00:00' */
;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */
;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */
;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */
;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */
;
--
-- Table structure for table `syncs`
--
DROP TABLE IF EXISTS `syncs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!40101 SET character_set_client = utf8 */
;
CREATE TABLE `syncs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `commited` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `committer` varchar(100) NOT NULL,
  `type` varchar(100) NOT NULL,
  `table` varchar(100) NOT NULL,
  `target` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Dumping data for table `syncs`
--
LOCK TABLES `syncs` WRITE;
/*!40000 ALTER TABLE `syncs` DISABLE KEYS */
;
/*!40000 ALTER TABLE `syncs` ENABLE KEYS */
;
UNLOCK TABLES;
--
-- Table structure for table `tags`
--
DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!40101 SET character_set_client = utf8 */
;
CREATE TABLE `tags` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(20) DEFAULT '',
  `desc` varchar(50) DEFAULT '',
  `color` varchar(20) NOT NULL DEFAULT '#000000',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Dumping data for table `tags`
--
LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */
;
INSERT INTO `tags`
VALUES (
    1,
    '生活',
    '生活中的小事，比如洗内裤。',
    '#FF0000',
    '2021-05-18 14:55:00',
    '2021-05-18 14:55:00',
    0
  ),
(
    2,
    '健身',
    '希望你也能像我一样成为彭于晏。',
    '#00FF00',
    '2021-05-18 14:55:00',
    '2021-05-18 14:55:00',
    0
  ),
(
    3,
    '内裤',
    '内裤是yyds！',
    '#0000FF',
    '2021-05-18 14:55:00',
    '2021-05-18 14:55:00',
    0
  ),
(
    4,
    '叽里呱啦',
    '只是凑数用的。',
    '#FFFF00',
    '2021-05-18 14:55:00',
    '2021-05-18 14:55:00',
    0
  );
/*!40000 ALTER TABLE `tags` ENABLE KEYS */
;
UNLOCK TABLES;
--
-- Table structure for table `task_tags`
--
DROP TABLE IF EXISTS `task_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!40101 SET character_set_client = utf8 */
;
CREATE TABLE `task_tags` (
  `task_id` int(11) NOT NULL,
  `tag_id` int(11) NOT NULL,
  PRIMARY KEY (`task_id`, `tag_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Dumping data for table `task_tags`
--
LOCK TABLES `task_tags` WRITE;
/*!40000 ALTER TABLE `task_tags` DISABLE KEYS */
;
INSERT INTO `task_tags`
VALUES (1, 1),
(1, 3),
(2, 2),
(3, 1),
(3, 2),
(3, 3),
(3, 4);
/*!40000 ALTER TABLE `task_tags` ENABLE KEYS */
;
UNLOCK TABLES;
--
-- Table structure for table `tasks`
--
DROP TABLE IF EXISTS `tasks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!40101 SET character_set_client = utf8 */
;
CREATE TABLE `tasks` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(100) DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `completed` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Dumping data for table `tasks`
--
LOCK TABLES `tasks` WRITE;
/*!40000 ALTER TABLE `tasks` DISABLE KEYS */
;
INSERT INTO `tasks`
VALUES (
    1,
    '记得换内裤',
    '2021-05-18 14:55:00',
    '2021-05-18 14:55:00',
    0,
    0
  ),
(
    2,
    '做广播体操',
    '2021-05-18 14:55:00',
    '2021-05-18 14:55:00',
    0,
    0
  ),
(
    3,
    '一边换内裤一边做广播体操',
    '2021-05-18 14:55:00',
    '2021-05-18 14:55:00',
    0,
    0
  ),
(
    4,
    '跳起来锤你的大脚拇指',
    '2021-05-18 14:55:00',
    '2021-05-18 14:55:00',
    0,
    1
  );
/*!40000 ALTER TABLE `tasks` ENABLE KEYS */
;
UNLOCK TABLES;
--
-- Table structure for table `users`
--
DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!40101 SET character_set_client = utf8 */
;
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `nick` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Dumping data for table `users`
--
LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */
;
INSERT INTO `users`
VALUES (
    1,
    'YuChao',
    'a2bf364d91c65964491d6ef7c0a36c46',
    'yc',
    '2021-05-18 14:55:00',
    '2021-05-18 14:55:00',
    0
  ),
(
    2,
    'ZhangHanyu',
    '3faf86140dea5edd5a06cb6715ee97a9',
    'zhy',
    '2021-05-18 14:55:00',
    '2021-05-18 14:55:00',
    0
  );
/*!40000 ALTER TABLE `users` ENABLE KEYS */
;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */
;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */
;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */
;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */
;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */
;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */
;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */
;
-- Dump completed on 2021-06-25 10:33:29