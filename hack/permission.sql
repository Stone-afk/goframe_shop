-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
                                   `id` int NOT NULL AUTO_INCREMENT,
                                   `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限名称',
                                   `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路径',
                                   `created_at` datetime DEFAULT NULL,
                                   `updated_at` datetime DEFAULT NULL,
                                   `deleted_at` datetime DEFAULT NULL,
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `unique_name` (`name`) COMMENT '名称唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of permission
-- ----------------------------
BEGIN;
INSERT INTO `permission` VALUES (1, '文章1', 'admin.article.index', '2022-09-25 15:03:01', '2022-09-25 15:03:43', NULL);
INSERT INTO `permission` VALUES (2, '测试2', 'admin.test.index', NULL, NULL, NULL);
INSERT INTO `permission` VALUES (5, '商品3', 'admin/goods', '2022-12-26 19:51:44', '2022-12-26 19:52:29', NULL);
INSERT INTO `permission` VALUES (6, '商品2', 'admin/goods', '2022-12-26 19:52:01', '2022-12-26 19:52:01', NULL);
COMMIT;