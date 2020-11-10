/*
 Navicat Premium Data Transfer

 Source Server         : Mac
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : localhost:3306
 Source Schema         : repo_demo

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 10/11/2020 21:12:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for demo
-- ----------------------------
DROP TABLE IF EXISTS `demo`;
CREATE TABLE `demo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of demo
-- ----------------------------
BEGIN;
INSERT INTO `demo` VALUES (1, '测试数据i.', '2020-11-02 11:32:06', '2020-11-02 11:32:10');
INSERT INTO `demo` VALUES (2, '测试数据ii', '2020-11-03 17:05:05', '2020-11-03 17:05:08');
COMMIT;

-- ----------------------------
-- Table structure for demo_dt
-- ----------------------------
DROP TABLE IF EXISTS `demo_dt`;
CREATE TABLE `demo_dt` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `demo_id` int(11) DEFAULT NULL,
  `dt_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of demo_dt
-- ----------------------------
BEGIN;
INSERT INTO `demo_dt` VALUES (1, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for demo_to_many
-- ----------------------------
DROP TABLE IF EXISTS `demo_to_many`;
CREATE TABLE `demo_to_many` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `demo_id` int(11) DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of demo_to_many
-- ----------------------------
BEGIN;
INSERT INTO `demo_to_many` VALUES (1, 1, '哈哈哈1', '2020-11-02 15:33:32', '2020-11-03 22:05:13');
INSERT INTO `demo_to_many` VALUES (2, 1, '哈哈哈2', '2020-11-03 22:05:04', '2020-11-03 22:05:07');
INSERT INTO `demo_to_many` VALUES (3, 2, '呵呵呵1', '2020-11-04 11:31:53', '2020-11-04 11:31:57');
INSERT INTO `demo_to_many` VALUES (4, 2, '呵呵呵2', '2020-11-04 11:32:11', '2020-11-04 11:32:14');
COMMIT;

-- ----------------------------
-- Table structure for dt
-- ----------------------------
DROP TABLE IF EXISTS `dt`;
CREATE TABLE `dt` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of dt
-- ----------------------------
BEGIN;
INSERT INTO `dt` VALUES (1, '测试多对多关联表');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
