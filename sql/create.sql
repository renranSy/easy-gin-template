-- MySQL dump 10.13  Distrib 8.4.3, for macos14 (arm64)
--
-- Host: 127.0.0.1    Database: easy_admin
-- ------------------------------------------------------
-- Server version	8.4.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `name` varchar(255) NOT NULL DEFAULT '',
                        `code` varchar(255) NOT NULL DEFAULT '',
                        `description` varchar(255) NOT NULL DEFAULT '',
                        `sequence` int NOT NULL DEFAULT '0',
                        `type` varchar(10) NOT NULL DEFAULT '' COMMENT 'page/button',
                        `path` varchar(255) NOT NULL DEFAULT '',
                        `status` tinyint NOT NULL DEFAULT '1' COMMENT '0.disable, 1.enable',
                        `parent_id` int NOT NULL DEFAULT '0',
                        `created_at` datetime NOT NULL,
                        `updated_at` datetime NOT NULL,
                        `deleted_at` datetime DEFAULT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (10,'工作台','Workbench','工作台',99,'page','/',1,0,'2025-02-03 12:14:21','2025-02-04 13:19:49',NULL),(11,'用户','User','用户',99,'page','/user',1,0,'2025-02-03 12:15:08','2025-02-04 13:21:06',NULL),(12,'功能','Feature','功能',99,'page','/feature',1,0,'2025-02-03 12:15:51','2025-02-04 13:21:09',NULL),(13,'富文本编辑器','RichEditor','',99,'page','/feature/rich-editor',1,12,'2025-02-03 12:16:35','2025-02-03 12:16:35',NULL),(14,'资源推荐','Resource','',99,'page','/resource',1,0,'2025-02-03 12:17:08','2025-02-04 13:21:14',NULL),(15,'多级菜单','Multi','',99,'page','/multi',1,0,'2025-02-03 12:18:07','2025-02-04 13:21:18',NULL),(16,'一级菜单','MultiOne','',99,'page','/multi/one',1,15,'2025-02-03 12:18:42','2025-02-04 13:21:26',NULL),(17,'二级菜单','MultiTwo','',99,'page','/multi/two',1,15,'2025-02-03 12:19:14','2025-02-04 13:21:31',NULL),(18,'页面一','MultiTwoOne','',99,'page','/multi/two/one',1,17,'2025-02-03 12:19:55','2025-02-04 13:21:36',NULL),(19,'页面二','MultiTwoTwo','',99,'page','/multi/two/two',1,17,'2025-02-03 12:20:27','2025-02-04 13:21:40',NULL),(20,'系统管理','System','',99,'page','/system',1,0,'2025-02-03 12:20:55','2025-02-04 13:21:00',NULL),(21,'用户管理','AdminManage','',99,'page','/system/admin',1,20,'2025-02-03 12:23:51','2025-02-04 13:19:32',NULL),(22,'角色管理','RoleManage','',99,'page','/system/role',1,20,'2025-02-03 12:27:06','2025-02-04 21:42:46',NULL),(23,'菜单管理','MenuManage','',99,'page','/system/menu',1,20,'2025-02-03 12:30:22','2025-02-04 21:47:22',NULL),(24,'查询','query','',0,'button','',1,21,'2025-02-04 13:16:14','2025-02-04 13:16:14',NULL),(25,'添加','add','',0,'button','',1,21,'2025-02-04 13:17:00','2025-02-04 13:19:02',NULL),(26,'编辑','edit','',0,'button','',1,21,'2025-02-04 13:17:32','2025-02-04 13:19:13',NULL),(27,'删除','delete','',0,'button','',1,21,'2025-02-04 13:18:02','2025-02-04 13:18:02',NULL),(28,'查询','query','',0,'button','',1,22,'2025-02-04 21:43:26','2025-02-04 21:43:26',NULL),(29,'添加','add','',0,'button','',1,22,'2025-02-04 21:43:56','2025-02-04 21:43:56',NULL),(30,'编辑','edit','',0,'button','',1,22,'2025-02-04 21:44:33','2025-02-04 23:01:35',NULL),(31,'删除','delete','',0,'button','',1,22,'2025-02-04 21:44:51','2025-02-04 21:44:51',NULL),(32,'查询','query','',0,'button','',1,23,'2025-02-04 21:48:49','2025-02-04 21:48:49',NULL),(33,'添加','add','',0,'button','',1,23,'2025-02-04 21:49:12','2025-02-04 21:49:17',NULL),(34,'编辑','edit','',0,'button','',1,23,'2025-02-04 21:50:00','2025-02-05 01:18:09',NULL),(35,'删除','delete','',0,'button','',1,23,'2025-02-04 21:51:01','2025-02-04 21:51:01',NULL);
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `username` varchar(255) NOT NULL DEFAULT '',
                         `password` varchar(255) NOT NULL DEFAULT '',
                         `email` varchar(255) NOT NULL DEFAULT '',
                         `status` tinyint NOT NULL DEFAULT '1' COMMENT '0.disable, 1.enable',
                         `created_at` datetime NOT NULL,
                         `updated_at` datetime NOT NULL,
                         `deleted_at` datetime DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,'admin','admin','2480901422@qq.com',1,'2024-11-29 14:50:02','2024-12-18 20:10:27',NULL),(17,'user','111111','',1,'2025-02-05 00:10:32','2025-02-05 01:05:31',NULL);
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_role`
--

DROP TABLE IF EXISTS `admin_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_role` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `admin_id` int NOT NULL,
                              `role_id` int NOT NULL,
                              `created_at` datetime NOT NULL,
                              `updated_at` datetime NOT NULL,
                              `deleted_at` datetime DEFAULT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role`
--

LOCK TABLES `admin_role` WRITE;
/*!40000 ALTER TABLE `admin_role` DISABLE KEYS */;
INSERT INTO `admin_role` VALUES (3,1,1,'2024-12-18 20:10:27','2024-12-18 20:10:27',NULL),(8,14,2,'2024-12-20 23:36:24','2024-12-20 23:36:24',NULL),(10,15,2,'2024-12-20 23:36:33','2024-12-20 23:36:33',NULL),(11,16,1,'2024-12-22 15:41:39','2024-12-22 15:41:39',NULL),(12,17,3,'2025-02-05 01:05:31','2025-02-05 01:05:31',NULL);
/*!40000 ALTER TABLE `admin_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu_resource`
--

DROP TABLE IF EXISTS `menu_resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu_resource` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `menu_id` int NOT NULL,
                                 `method` varchar(255) NOT NULL DEFAULT '',
                                 `path` varchar(255) NOT NULL DEFAULT '',
                                 `created_at` datetime NOT NULL,
                                 `updated_at` datetime NOT NULL,
                                 `deleted_at` datetime DEFAULT NULL,
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=332 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu_resource`
--

LOCK TABLES `menu_resource` WRITE;
/*!40000 ALTER TABLE `menu_resource` DISABLE KEYS */;
INSERT INTO `menu_resource` VALUES (64,13,'GET','','2025-02-03 12:16:35','2025-02-03 12:16:35',NULL),(307,24,'GET','/admin','2025-02-04 13:16:14','2025-02-04 13:16:14',NULL),(310,27,'DELETE','/admin/:id','2025-02-04 13:18:02','2025-02-04 13:18:02',NULL),(311,25,'POST','/admin','2025-02-04 13:19:02','2025-02-04 13:19:02',NULL),(312,26,'PUT','/admin/:id','2025-02-04 13:19:13','2025-02-04 13:19:13',NULL),(313,28,'GET','/role','2025-02-04 21:43:26','2025-02-04 21:43:26',NULL),(314,29,'POST','/role','2025-02-04 21:43:56','2025-02-04 21:43:56',NULL),(316,31,'DELETE','/role/:id','2025-02-04 21:44:51','2025-02-04 21:44:51',NULL),(318,32,'GET','/menu/:id','2025-02-04 21:48:49','2025-02-04 21:48:49',NULL),(319,32,'GET','/menu','2025-02-04 21:48:49','2025-02-04 21:48:49',NULL),(321,33,'POST','/menu','2025-02-04 21:49:17','2025-02-04 21:49:17',NULL),(323,35,'DELETE','/menu/:id','2025-02-04 21:51:01','2025-02-04 21:51:01',NULL),(327,30,'PUT','/role/:id','2025-02-04 23:01:35','2025-02-04 23:01:35',NULL),(328,30,'GET','/role/menu/:id','2025-02-04 23:01:35','2025-02-04 23:01:35',NULL),(329,30,'PUT','/role/menu/:id','2025-02-04 23:01:35','2025-02-04 23:01:35',NULL),(330,30,'GET','/menu/:id','2025-02-04 23:01:35','2025-02-04 23:01:35',NULL),(331,34,'PUT','/menu/:id','2025-02-05 01:18:09','2025-02-05 01:18:09',NULL);
/*!40000 ALTER TABLE `menu_resource` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `name` varchar(255) NOT NULL DEFAULT '',
                        `description` varchar(255) NOT NULL DEFAULT '',
                        `sequence` int NOT NULL DEFAULT '0',
                        `status` tinyint NOT NULL DEFAULT '1' COMMENT '0.disable, 1.enable',
                        `created_at` datetime NOT NULL,
                        `updated_at` datetime NOT NULL,
                        `deleted_at` datetime DEFAULT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'超级管理员','最高权限',999,1,'2024-12-06 00:00:16','2024-12-22 15:45:00',NULL),(3,'用户','',0,1,'2025-02-05 00:10:44','2025-02-05 00:10:44',NULL);
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_menu`
--

DROP TABLE IF EXISTS `role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_menu` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `role_id` int NOT NULL,
                             `menu_id` int NOT NULL,
                             `created_at` datetime NOT NULL,
                             `updated_at` datetime NOT NULL,
                             `deleted_at` datetime DEFAULT NULL,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1987 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_menu`
--

LOCK TABLES `role_menu` WRITE;
/*!40000 ALTER TABLE `role_menu` DISABLE KEYS */;
INSERT INTO `role_menu` VALUES (1817,1,10,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1818,1,11,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1819,1,13,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1820,1,12,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1821,1,14,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1822,1,16,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1823,1,18,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1824,1,19,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1825,1,17,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1826,1,15,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1827,1,24,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1828,1,25,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1829,1,26,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1830,1,27,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1831,1,21,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1832,1,28,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1833,1,29,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1834,1,30,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1835,1,31,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1836,1,22,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1837,1,32,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1838,1,33,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1839,1,34,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1840,1,35,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1841,1,23,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1842,1,20,'2025-02-05 00:27:35','2025-02-05 00:27:35',NULL),(1970,3,10,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1971,3,11,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1972,3,13,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1973,3,12,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1974,3,14,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1975,3,16,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1976,3,18,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1977,3,19,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1978,3,17,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1979,3,15,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1980,3,24,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1981,3,28,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1982,3,32,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1983,3,21,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1984,3,22,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1985,3,23,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL),(1986,3,20,'2025-02-05 01:23:13','2025-02-05 01:23:13',NULL);
/*!40000 ALTER TABLE `role_menu` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-05 17:31:55
