CREATE TABLE `goods_options` (
                                  `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
                                  `goods_id` INT NOT NULL COMMENT '商品id',
                                  `pic_url` VARCHAR(255) NOT NULL COMMENT '图片',
                                  `name` VARCHAR(255) NOT NULL COMMENT '商品名称',
                                  `price` INT NOT NULL COMMENT '价格 单位分',
                                  `stock` INT NOT NULL COMMENT '库存',
                                  `created_at` DATETIME DEFAULT NULL COMMENT '创建时间',
                                  `updated_at` DATETIME DEFAULT NULL COMMENT '更新时间',
                                  `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
                                  INDEX idx_goodsId (goods_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品规格信息表';