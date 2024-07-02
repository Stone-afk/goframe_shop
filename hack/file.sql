CREATE TABLE `file` (
                          `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
                          `name` VARCHAR(255) NOT NULL COMMENT '图片名称',
                          `src` VARCHAR(255) NOT NULL COMMENT '本地文件存储路径',
                          `url` VARCHAR(255) NOT NULL COMMENT 'URL地址',
                          `user_id` INT NOT NULL COMMENT '用户id',
                          `created_at` DATETIME DEFAULT NULL COMMENT '创建时间',
                          `updated_at` DATETIME DEFAULT NULL COMMENT '更新时间',
                          INDEX idx_userId (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件信息表';