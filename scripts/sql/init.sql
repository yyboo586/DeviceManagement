-- 创建数据库
CREATE DATABASE IF NOT EXISTS `device_management` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `device_management`;

CREATE TABLE IF NOT EXISTS `t_product_category` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
    `name` VARCHAR(50) NOT NULL COMMENT '产品分类名称',
    `desc` VARCHAR(255) NULL COMMENT '产品分类描述',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_org_id_name` (`org_id`, `name`)
)ENGINE=InnoDB COMMENT='产品分类';

CREATE TABLE IF NOT EXISTS `t_product` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
    `category_id` BIGINT(20) NOT NULL COMMENT '产品分类ID',
    `name` VARCHAR(50) NOT NULL COMMENT '产品名称',
    `desc` VARCHAR(255) NULL COMMENT '产品描述',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_org_id` (`org_id`),
    KEY `idx_category_id` (`category_id`)
)ENGINE=InnoDB COMMENT='产品';

CREATE TABLE IF NOT EXISTS `t_device` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '设备ID',
    `product_id` BIGINT(20) NOT NULL COMMENT '产品ID',
    `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
    `creator_id` VARCHAR(40) NOT NULL COMMENT '创建者ID',
    `device_key` VARCHAR(64) NOT NULL COMMENT '设备唯一标识',
    `name` VARCHAR(100) NOT NULL COMMENT '设备名称',
    `enabled` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '设备状态',
    `online_status` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '设备在线状态',
    `location` VARCHAR(255) NULL COMMENT '设备位置',
    `description` TEXT NULL COMMENT '设备描述',
    `last_online_time` DATETIME NULL COMMENT '最后在线时间',
    `last_offline_time` DATETIME NULL COMMENT '最后离线时间',    
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_org_device_key` (`org_id`, `device_key`),
    KEY `idx_org_enabled` (`org_id`, `enabled`),
    KEY `idx_org_online_status` (`org_id`, `online_status`)
) ENGINE=InnoDB COMMENT='设备表';

CREATE TABLE IF NOT EXISTS `t_device_permission` ( 
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `user_id` VARCHAR(40) NOT NULL COMMENT '用户ID',
  `device_id` BIGINT(20) NOT NULL COMMENT '设备ID',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_org_user_device` (`org_id`, `user_id`, `device_id`),
  KEY `idx_user_device` (`user_id`, `device_id`)
) ENGINE = InnoDB COMMENT = '设备权限表';

CREATE TABLE IF NOT EXISTS `t_device_config`  (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '配置ID',
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `type` TINYINT(4) NOT NULL COMMENT '配置类型',
  `key` VARCHAR(40) NOT NULL COMMENT '配置键',
  `value` VARCHAR(40) NOT NULL COMMENT '配置值',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_org_key`(`org_id`, `key`)
) ENGINE = InnoDB COMMENT = '设备配置表';

CREATE TABLE IF NOT EXISTS `t_device_log` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `device_id` BIGINT(20) NOT NULL COMMENT '设备ID',
  `device_name` VARCHAR(100) NOT NULL COMMENT '设备名称',
  `device_key` VARCHAR(64) NOT NULL COMMENT '设备唯一标识',
  `type` TINYINT(4) NOT NULL COMMENT '日志类型(1:上线 2:下线 3:报警)',
  `content` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '日志内容',
  `timestamp` BIGINT(20) NOT NULL COMMENT '时间戳',
  `created_at` BIGINT(20) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `idx_org_device_type_created_at` (`org_id`, `device_id`, `type`, `created_at`),
  INDEX `idx_org_device_created_at` (`org_id`, `device_id`, `created_at`),  
  INDEX `idx_org_type_created_at` (`org_id`, `type`, `created_at`),
  INDEX `idx_org_created_at` (`org_id`, `created_at`)
) ENGINE = InnoDB COMMENT = '设备日志表';

CREATE TABLE IF NOT EXISTS `t_cron_job_template` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '模板ID',
  `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '模板名称',
  `invoke_type` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '调用类型(http grpc default)',
  `config` TEXT COMMENT '配置',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
)ENGINE = InnoDB COMMENT = '任务模板表';

CREATE TABLE IF NOT EXISTS `t_cron_job`  (
  `id` VARCHAR(40) NOT NULL COMMENT '任务ID',
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '任务名称',
  `enabled` TINYINT(4) NULL DEFAULT 0 COMMENT '是否启用(0: 禁用 1: 启用)',  
  `remark` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '备注信息', 
  `params` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '参数(租户可自定义)',
  `invoke_type` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '调用目标(http grpc default)',
  `cron_expression` VARCHAR(255) NULL DEFAULT '' COMMENT 'cron执行表达式',
  `description` VARCHAR(500) NOT NULL COMMENT '描述',
  `last_execute_status` TINYINT(4) NULL DEFAULT 0 COMMENT '上次执行状态',
  `last_execute_at` DATETIME NULL DEFAULT NULL COMMENT '上次执行时间',
  `execute_count` BIGINT(20) NULL DEFAULT 0 COMMENT '执行次数',
  `success_count` BIGINT(20) NULL DEFAULT 0 COMMENT '成功次数',
  `failed_count` BIGINT(20) NULL DEFAULT 0 COMMENT '失败次数',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_org_id_name` (`org_id`, `name`),
  INDEX `idx_enabled` (`enabled`)
) ENGINE = InnoDB COMMENT = '定时任务调度表';

CREATE TABLE IF NOT EXISTS `t_cron_job_log`  (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `job_id` VARCHAR(40) NOT NULL COMMENT '任务ID',
  `job_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '任务名称',
  `execute_status` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '任务执行状态',
  `result` TEXT NULL COMMENT '执行结果(JSON)',
  `start_time` DATETIME NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` DATETIME NULL DEFAULT NULL COMMENT '结束时间',
  `duration` BIGINT(20) NULL DEFAULT 0 COMMENT '执行时长(毫秒)',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `idx_org_id` (`org_id`, `job_id`, `execute_status`, `created_at`)
) ENGINE = InnoDB COMMENT = '定时任务执行日志表';