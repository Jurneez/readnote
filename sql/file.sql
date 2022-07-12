create table `file`(
	  `id` int(11) unsigned NOT NULL AUTO_INCREMENT comment'自增ID',
    `order` int(11) unsigned NOT NULL DEFAULT 1 comment'排序',
    `name` varchar(64) NOT NULL DEFAULT "" comment'文件名字',
    `ftype` varchar(10) not null DEFAULT "" comment'文件类型，pdf,txt等',
    `content` varchar(1024) not null DEFAULT "" comment'文件内容',
    `author` varchar(128) not null DEFAULT "" comment'作者',
    `remark` varchar(256) not null DEFAULT "" comment'备注',
    `create_time` timestamp not NULL DEFAULT current_timestamp comment '创建时间',
    `modify_time` timestamp not null DEFAULT current_timestamp on update current_timestamp comment'更新时间',
    PRIMARY KEY (`id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;