-- douyin.dy_comment definition

CREATE TABLE `dy_comment` (
                              `comment_id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论ID',
                              `user_id` bigint NOT NULL COMMENT '用户名',
                              `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '评论内容',
                              `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
                              `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `video_id` bigint NOT NULL,
                              PRIMARY KEY (`comment_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;


-- douyin.dy_favorite definition

CREATE TABLE `dy_favorite` (
                               `favorite_id` bigint NOT NULL AUTO_INCREMENT,
                               `user_id` bigint NOT NULL COMMENT '用户ID',
                               `video_id` bigint NOT NULL DEFAULT '0' COMMENT '视频ID',
                               `is_del` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
                               `created_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                               PRIMARY KEY (`favorite_id`) USING BTREE,
                               UNIQUE KEY `dy_favorite_user_id_IDX` (`user_id`,`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1767 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;


-- douyin.dy_relation definition

CREATE TABLE `dy_relation` (
                               `relation_id` bigint NOT NULL AUTO_INCREMENT COMMENT '关系ID',
                               `follower_id` bigint NOT NULL DEFAULT '0' COMMENT '粉丝ID',
                               `following_id` bigint NOT NULL DEFAULT '0' COMMENT '博主ID',
                               `is_del` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
                               PRIMARY KEY (`relation_id`) USING BTREE,
                               UNIQUE KEY `dy_relation_follower_following` (`follower_id`,`following_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1177 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='用来描述用户和粉丝之间的关系';


-- douyin.dy_user definition

CREATE TABLE `dy_user` (
                           `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                           `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
                           `follower_count` int NOT NULL DEFAULT '0' COMMENT '粉丝总数',
                           `is_follow` tinyint NOT NULL DEFAULT '0' COMMENT '是否已关注',
                           `password` char(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户密码',
                           `salt` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'uuid，密码MD5加密',
                           `follow_count` bigint NOT NULL DEFAULT '0' COMMENT '关注总数',
                           PRIMARY KEY (`user_id`) USING BTREE,
                           KEY `dy_user_name_IDX` (`name`,`salt`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;


-- douyin.dy_video definition

CREATE TABLE `dy_video` (
                            `video_id` bigint NOT NULL AUTO_INCREMENT COMMENT '视频ID',
                            `user_id` bigint DEFAULT NULL COMMENT '用户ID',
                            `play_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '播放地址',
                            `cover_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '视频封面地址',
                            `favorite_count` bigint DEFAULT '0' COMMENT '点赞量',
                            `comment_count` bigint DEFAULT '0' COMMENT '评论量',
                            `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '视频标题',
                            `create_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `update_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            PRIMARY KEY (`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;


-- douyin.dy_video_counter definition

CREATE TABLE `dy_video_counter` (
                                    `video_id` bigint NOT NULL,
                                    `like` bigint NOT NULL DEFAULT '0',
                                    `comment` bigint NOT NULL DEFAULT '0',
                                    PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;