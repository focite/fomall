SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_profile_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `user_type` varchar(16) NOT NULL COMMENT '认证类型',
  `passport` varchar(32) NOT NULL COMMENT '认证用户',
  `password` varchar(128) NOT NULL COMMENT '认证凭证',
  `status` tinyint(1) unsigned NOT NULL COMMENT '状态',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_user` (`user_type`,`passport`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user_auth
-- ----------------------------

-- ----------------------------
-- Table structure for user_profile
-- ----------------------------
DROP TABLE IF EXISTS `user_profile`;
CREATE TABLE `user_profile` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL COMMENT '昵称',
  `avatar` varchar(255) NOT NULL COMMENT '头像',
  `brithday` date NOT NULL COMMENT '生日',
  `email` varchar(64) NOT NULL COMMENT '电子邮箱',
  `mobile` varchar(16) NOT NULL COMMENT '手机号码',
  `status` tinyint(1) unsigned NOT NULL COMMENT '状态',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`name`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user_profile
-- ----------------------------
