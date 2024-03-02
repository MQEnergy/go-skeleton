-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v6` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `v7` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES (NULL, 'p', '1', '/backend/auth/*', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (NULL, 'p', '1', '/backend/attachment/*', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (NULL, 'p', '1', '/backend/role/*', 'GET,POST', NULL, NULL, NULL, '', '');
INSERT INTO `casbin_rule` VALUES (NULL, 'p', '1', '/backend/menu/*', 'GET,POST', NULL, NULL, NULL, '', '');
INSERT INTO `casbin_rule` VALUES (NULL, 'p', '1', '/backend/permission/*', 'GET,POST', NULL, NULL, NULL, '', '');
INSERT INTO `casbin_rule` VALUES (NULL, 'p', '1', '/backend/admin/*', 'GET,POST', NULL, NULL, NULL, '', '');
COMMIT;

