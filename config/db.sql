CREATE TABLE `word` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `src_content` varchar(256) NOT NULL DEFAULT '' COMMENT '需要翻译单词',
  `dst_content` varchar(256) NOT NULL DEFAULT '' COMMENT '翻译结果',
  `dst_attr` varchar(512) NOT NULL DEFAULT '' COMMENT '词性',
  `dst_explain` text  NULL  COMMENT '解释',
  `dst_example` text  NULL  COMMENT '例子',
  `media_id` varchar(256) NOT NULL DEFAULT '' COMMENT '素材id',
  `created_at` timestamp NOT NULL DEFAULT '2020-01-01 12:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='单词表';
