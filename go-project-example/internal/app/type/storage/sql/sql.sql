CREATE TABLE `relay_op_log`
(
    `id`         int         NOT NULL AUTO_INCREMENT,
    `op_type`    varchar(32) NOT NULL DEFAULT '' COMMENT '操作类型，bigdata/...',
    `req_type`   varchar(8)  NOT NULL DEFAULT '' COMMENT '请求类型：create/delete',
    `target`     varchar(32) NOT NULL DEFAULT '' COMMENT '处理目标',
    `req_data`   text COMMENT '请求数据',
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
    `updated_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 comment 'relay_op_log';