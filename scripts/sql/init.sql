-- 创建数据库
CREATE DATABASE IF NOT EXISTS `device_management` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `device_management`;

CREATE TABLE IF NOT EXISTS `t_device` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '设备ID',
    `name` VARCHAR(100) NOT NULL COMMENT '设备名称',
    `device_key` VARCHAR(64) NOT NULL COMMENT '设备唯一标识',
    `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
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

CREATE TABLE IF NOT EXISTS `t_device_online_log` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `device_id` BIGINT(20) NOT NULL COMMENT '设备ID',
  `device_key` VARCHAR(64) NOT NULL COMMENT '设备唯一标识',
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `event_type` TINYINT(4) NOT NULL COMMENT '事件类型(1: 上线 2: 下线)',
  `online_status` TINYINT(4) NOT NULL COMMENT '设备在线状态(1: 在线 2: 离线)',
  `ip_address` VARCHAR(45) NULL COMMENT '设备IP地址',
  `client_id` VARCHAR(64) NULL COMMENT '客户端ID',
  `reason` VARCHAR(255) NULL COMMENT '上下线原因',
  `duration` BIGINT(20) NULL DEFAULT 0 COMMENT '在线时长(秒)',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `idx_device_id` (`device_id`),
  INDEX `idx_device_key` (`device_key`),
  INDEX `idx_org_id` (`org_id`),
  INDEX `idx_event_type` (`event_type`),
  INDEX `idx_created_at` (`created_at`),
  INDEX `idx_device_event_time` (`device_id`, `event_type`, `created_at`)
) ENGINE = InnoDB COMMENT = '设备上下线日志表';

CREATE TABLE IF NOT EXISTS `t_cron_job`  (
  `id` VARCHAR(40) NOT NULL COMMENT '任务ID',
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '任务名称',
  `enabled` TINYINT(4) NULL DEFAULT 0 COMMENT '是否启用(0: 禁用 1: 启用)',  
  `remark` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '备注信息',  
  `params` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '参数',
  `invoke_target` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '调用目标(http grpc default)',
  `cron_expression` VARCHAR(255) NULL DEFAULT '' COMMENT 'cron执行表达式',
  `execute_immediately` TINYINT(4) NULL DEFAULT 0 COMMENT '创建任务后是否立即执行(0: 否 1: 是)',
  `last_execute_status` TINYINT(4) NULL DEFAULT 0 COMMENT '上次执行状态',
  `last_execute_at` DATETIME NULL DEFAULT NULL COMMENT '上次执行时间',
  `next_execute_at` DATETIME NULL DEFAULT NULL COMMENT '下次执行时间',
  `execute_count` BIGINT(20) NULL DEFAULT 0 COMMENT '执行次数',
  `success_count` BIGINT(20) NULL DEFAULT 0 COMMENT '成功次数',
  `failed_count` BIGINT(20) NULL DEFAULT 0 COMMENT '失败次数',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_org_id_name` (`org_id`, `name`),
  INDEX `idx_enabled` (`enabled`),
  INDEX `idx_next_execute_at` (`next_execute_at`)
) ENGINE = InnoDB COMMENT = '定时任务调度表';

CREATE TABLE IF NOT EXISTS `t_cron_job_log`  (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `job_id` VARCHAR(40) NOT NULL COMMENT '任务ID',
  `execute_status` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '任务执行状态',
  `result` TEXT NULL COMMENT '执行结果(JSON)',
  `start_time` DATETIME NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` DATETIME NULL DEFAULT NULL COMMENT '结束时间',
  `duration` BIGINT(20) NULL DEFAULT 0 COMMENT '执行时长(毫秒)',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `idx_org_id` (`org_id`),
  INDEX `idx_job_id` (`job_id`),
  INDEX `idx_execute_status` (`execute_status`),
  INDEX `idx_created_at` (`created_at`)
) ENGINE = InnoDB COMMENT = '定时任务执行日志表';