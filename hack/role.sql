-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
                             `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
                             `created_at` datetime DEFAULT NULL,
                             `updated_at` datetime DEFAULT NULL,
                             `deleted_at` datetime DEFAULT NULL,
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `unique_index` (`name`) USING BTREE COMMENT '角色昵称唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of role_info
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, '运营1', '测试', '2022-09-25 10:35:52', '2022-12-24 10:51:24', NULL);
INSERT INTO `role` VALUES (3, '运营', '', '2022-12-21 10:43:33', '2022-12-21 10:43:33', NULL);
COMMIT;

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission` (
                                        `id` int NOT NULL AUTO_INCREMENT,
                                        `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
                                        `permission_id` int NOT NULL COMMENT '权限id',
                                        `created_at` datetime DEFAULT NULL,
                                        `updated_at` datetime DEFAULT NULL,
                                        PRIMARY KEY (`id`),
                                        UNIQUE KEY `unique_index` (`role_id`,`permission_id`) USING BTREE COMMENT '唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
