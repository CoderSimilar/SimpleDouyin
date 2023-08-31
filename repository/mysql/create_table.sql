create database if not exists douyin;
use douyin;

/*
drop table if exists `user`;
create table `users`(
        `id` bigint(20) unsigned not null AUTO_INCREMENT,
        `username` varchar(32) collate utf8mb4_general_ci not null unique , -- collate排序
        `password` varchar(32) collate utf8mb4_general_ci not null,
        `follow_count` bigint DEFAULT NULL,
        `follower_count` bigint DEFAULT NULL,
        `is_follow` tinyint(1) DEFAULT NULL,
        `avatar` varchar(255) DEFAULT NULL,
        `background_image` varchar(255) DEFAULT NULL,
        `signature` varchar(255) DEFAULT NULL,
        `total_favorited` int DEFAULT NULL,
        `work_count` int DEFAULT NULL,
        `favorite_count` int DEFAULT NULL,
        `create_time` timestamp not null default current_timestamp,
        `update_time` timestamp not null default current_timestamp on update current_timestamp  COMMENT '用户信息更新时间',
        primary key (`id`),
        unique key `idx_username` (`username`) using btree ,
        unique key `idx_user_id` (`userid`) using btree
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `follow_count` bigint(20) DEFAULT '0',
  `follower_count` bigint(20) DEFAULT '0',
  `is_follow` tinyint(1) DEFAULT '1',
  `avatar` varchar(191) DEFAULT NULL,
  `background_image` longtext,
  `signature` longtext,
  `total_favorited` bigint(20) DEFAULT '0',
  `work_count` bigint(20) DEFAULT NULL,
  `favorite_count` bigint(20) DEFAULT NULL,
  `token` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_name` (`username`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8"

CREATE TABLE `videos` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `author_id` bigint NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `play_url` varchar(255) DEFAULT NULL,
  `cover_url` varchar(255) DEFAULT NULL,
  `favorite_count` bigint DEFAULT NULL,
  `comment_count` bigint DEFAULT NULL,
  `is_favorite` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_videos_deleted_at` (`deleted_at`)
  CONSTRAINT `fk_videos_author_id_users_id` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

 CREATE TABLE `comments` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `comment_id` bigint DEFAULT NULL,
  `author_id` bigint DEFAULT NULL,
  `video_id` bigint DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_comments_deleted_at` (`deleted_at`)
  CONSTRAINT `fk_comments_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci


CREATE TABLE `user_video_relations` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `video_id` bigint(20) DEFAULT NULL,
  `is_favorite` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8

 */

