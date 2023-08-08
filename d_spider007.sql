-- MySQL dump 10.13  Distrib 5.7.43, for osx10.17 (x86_64)
--
-- Host: localhost    Database: d_spider007
-- ------------------------------------------------------
-- Server version	5.7.43

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `t_class`
--

DROP TABLE IF EXISTS `t_class`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` bigint(20) NOT NULL DEFAULT '0',
  `updated_at` bigint(20) NOT NULL DEFAULT '0',
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_class`
--

LOCK TABLES `t_class` WRITE;
/*!40000 ALTER TABLE `t_class` DISABLE KEYS */;
INSERT INTO `t_class` VALUES (1,1675514585902,1675514585902,'视频播放'),(2,1675514614009,1675514614009,'音乐电台'),(3,1675514621051,1675514621051,'拍摄美化'),(4,1675524314796,1675524314796,'资讯阅读'),(5,1675524326592,1675524326592,'聊天社交'),(6,1675524332187,1675524332187,'效率办公'),(7,1675524338422,1675524338422,'金融理财'),(8,1675524344265,1675524344265,'时尚购物'),(9,1675524349539,1675524349539,'便捷生活'),(10,1675524354550,1675524354550,'旅行交通'),(11,1675524359614,1675524359614,'医疗健康'),(12,1675524365086,1675524365086,'学习教育'),(13,1675524369755,1675524369755,'实用工具'),(14,1675524375095,1675524375095,'智能硬件'),(15,1675524380007,1675524380007,'游戏周边');
/*!40000 ALTER TABLE `t_class` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_log`
--

DROP TABLE IF EXISTS `t_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` bigint(20) NOT NULL DEFAULT '0',
  `updated_at` bigint(20) NOT NULL DEFAULT '0',
  `phone_num` bigint(20) NOT NULL DEFAULT '0',
  `openid` varchar(100) NOT NULL,
  `register_info` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_log`
--

LOCK TABLES `t_log` WRITE;
/*!40000 ALTER TABLE `t_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_platform`
--

DROP TABLE IF EXISTS `t_platform`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_platform` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` bigint(20) NOT NULL DEFAULT '0',
  `updated_at` bigint(20) NOT NULL DEFAULT '0',
  `class_id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(100) NOT NULL,
  `icon` varchar(100) NOT NULL,
  `developer` varchar(100) NOT NULL,
  `desc` varchar(100) NOT NULL,
  `link` varchar(100) NOT NULL,
  `tag` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_platform`
--

LOCK TABLES `t_platform` WRITE;
/*!40000 ALTER TABLE `t_platform` DISABLE KEYS */;
INSERT INTO `t_platform` VALUES (1,1675570452531,1675570452531,5,'微博','http://file.market.xiaomi.com/thumbnail/PNG/l114/AppStore/0bf2d587fa3774b7585092d9e37c54932a60295ba','微梦创科网络科技（中国）有限公司','在微博，随时随地，发现新鲜事','https://app.mi.com/details?id=com.sina.weibo','[]'),(2,1677311440955,1677311440955,9,'自如','https://file.market.xiaomi.com/thumbnail/PNG/l114/AppStore/033e2a62b29ab4bda893c00ac0fba49e4013dd428','北京自如信息科技有限公司','7大产品线，满足租住需求','https://app.mi.com/details?id=com.ziroom.ziroomcustomer','[\"网站\",\"APP\",\"微信小程序\"]'),(3,1677319299658,1677319299658,13,'百度','http://file.market.xiaomi.com/thumbnail/PNG/l114/AppStore/02821263aa1f24e35bf0fa3f40311970f1c2d85f5','百度在线网络技术（北京）有限公司','百度App，数亿用户优选的搜索和资讯客户端','https://app.mi.com/details?id=com.baidu.searchbox','[\"网站\",\"APP\",\"微信小程序\"]');
/*!40000 ALTER TABLE `t_platform` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_register`
--

DROP TABLE IF EXISTS `t_register`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_register` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` bigint(20) NOT NULL DEFAULT '0',
  `updated_at` bigint(20) NOT NULL DEFAULT '0',
  `phone_num` bigint(20) NOT NULL DEFAULT '0',
  `register_info` text NOT NULL,
  `queries` int(11) NOT NULL DEFAULT '0',
  `register_update` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_register`
--

LOCK TABLES `t_register` WRITE;
/*!40000 ALTER TABLE `t_register` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_register` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_user`
--

DROP TABLE IF EXISTS `t_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` bigint(20) NOT NULL DEFAULT '0',
  `updated_at` bigint(20) NOT NULL DEFAULT '0',
  `openid` varchar(100) NOT NULL,
  `avatar_url` text NOT NULL,
  `nick_name` varchar(100) NOT NULL,
  `credit` int(11) NOT NULL DEFAULT '0',
  `signed_at` bigint(20) NOT NULL DEFAULT '0',
  `rewarded_count` int(11) NOT NULL DEFAULT '0',
  `rewarded_at` bigint(20) NOT NULL DEFAULT '0',
  `limited` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_user`
--

LOCK TABLES `t_user` WRITE;
/*!40000 ALTER TABLE `t_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-08 22:43:45
