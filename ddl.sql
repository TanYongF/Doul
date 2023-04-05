-- douyin.dy_relation definition

CREATE TABLE `dy_relation` (
                               `relation_id` bigint NOT NULL AUTO_INCREMENT COMMENT '关系ID',
                               `follower_id` bigint NOT NULL DEFAULT '0' COMMENT '粉丝ID',
                               `following_id` bigint NOT NULL DEFAULT '0' COMMENT '博主ID',
                               `is_del` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
                               PRIMARY KEY (`relation_id`) USING BTREE,
                               UNIQUE KEY `dy_relation_follower_following` (`follower_id`,`following_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='用来描述用户和粉丝之间的关系';