create table `dic_tag`(
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT comment'自增ID',
    `tag` varchar(64) NOT NULL DEFAULT "" comment'文件名字',
    `status` int(11) unsigned NOT NULL DEFAULT 0 comment'0-添加 1-删除',
    `create_time` timestamp not NULL DEFAULT current_timestamp comment '创建时间',
    PRIMARY KEY (`id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;