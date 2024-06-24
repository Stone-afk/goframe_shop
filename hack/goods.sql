CREATE TABLE `goods` (
                       `id` INT AUTO_INCREMENT PRIMARY KEY,
                       `pic_url` VARCHAR(255) NOT NULL COMMENT '图片',
                       `name` VARCHAR(255) NOT NULL COMMENT '商品名称',
                       `price` INT NOT NULL COMMENT '价格 单位分',
                       `level1_category_id` INT NOT NULL COMMENT '1级分类id',
                       `level2_category_id` INT NOT NULL COMMENT '2级分类id',
                       `level3_category_id` INT NOT NULL COMMENT '3级分类id',
                       `brand` VARCHAR(255) NOT NULL COMMENT '品牌',
--                        `coupon_id` INT NOT NULL COMMENT '优惠券id',
                       `stock` INT NOT NULL COMMENT '库存',
                       `sale` INT NOT NULL COMMENT '销量',
                       `tags` VARCHAR(255) COMMENT '标签',
                       `detail_info` TEXT COMMENT '商品详情',
                       `created_at` DATETIME DEFAULT NULL COMMENT '',
                       `updated_at` DATETIME DEFAULT NULL COMMENT '',
                       `deleted_at` DATETIME DEFAULT NULL COMMENT '',
                       INDEX idx_level1_category_id (level1_category_id),
                       INDEX idx_level2_category_id (level2_category_id),
                       INDEX idx_level3_category_id (level3_category_id)
--                        INDEX idx_couponId (`coupon_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='商品表';