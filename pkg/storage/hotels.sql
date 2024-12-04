CREATE DATABASE airbnb;

USE airbnb;

CREATE TABLE hotel (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `queue_name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '队列名称',
    `hotel_name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '酒店名',
    `star` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '星级',
    `price` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '价格',
    `price_before_taxes` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '税前价格',
    `checkin_date` BIGINT(20) NOT NULL DEFAULT '0' COMMENT '入住时间',
    `checkout_date` BIGINT(20) NOT NULL DEFAULT '0' COMMENT '退房时间',
    `guests` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '入住数',
    `create_time` BIGINT(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
    `update_time` BIGINT(20) NOT NULL DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_queue_name` (`queue_name`),
    KEY `idx_hotels_name` (`hotel_name`),
    KEY `idx_create_time` (`create_time`),
    KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'hotel';

