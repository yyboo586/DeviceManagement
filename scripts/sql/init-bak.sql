CREATE TABLE IF NOT EXISTS `t_product_category` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `pid` BIGINT(20) NOT NULL COMMENT '父级分类ID',
    `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
    `name` VARCHAR(50) NOT NULL COMMENT '产品分类名称',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_pid` (`pid`),
    KEY `idx_org_pid` (`org_id`, `pid`)
)ENGINE=InnoDB COMMENT='产品分类';

CREATE TABLE IF NOT EXISTS `t_product` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `category_id` BIGINT(20) NOT NULL COMMENT '产品分类ID',
    `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
    `name` VARCHAR(50) NOT NULL COMMENT '产品名称',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_org_category_id` (`org_id`,`category_id`)
)ENGINE=InnoDB COMMENT='产品';

CREATE TABLE IF NOT EXISTS `t_thing_model_template` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `org_id` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '组织ID',  
    `name` VARCHAR(50) NOT NULL COMMENT '模板名称',
    `description` TEXT COMMENT '模板描述',
    `properties` TEXT COMMENT '属性定义(JSON格式)',
    `services` TEXT COMMENT '服务定义(JSON格式)',
    `events` TEXT COMMENT '事件定义(JSON格式)',
    `is_system` TINYINT(1) DEFAULT 0 COMMENT '是否是系统内置模板',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_org_id` (`org_id`)
)ENGINE=InnoDB COMMENT='物模型模板';

CREATE TABLE IF NOT EXISTS `t_thing_model` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `product_id` BIGINT(20) NOT NULL COMMENT '产品ID',    
    `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
    `template_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '模板ID',
    `name` VARCHAR(50) NOT NULL COMMENT '物模型名称',
    `version` INT NOT NULL DEFAULT 0 COMMENT '物模型版本',
    `description` TEXT COMMENT '物模型描述',
    `properties` TEXT COMMENT '属性定义(JSON格式)',
    `services` TEXT COMMENT '服务定义(JSON格式)',
    `events` TEXT COMMENT '事件定义(JSON格式)',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_org_product_id` (`org_id`,`product_id`),
    KEY `idx_template_id` (`template_id`)
)ENGINE=InnoDB COMMENT='物模型';