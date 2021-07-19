CREATE TABLE `profiles` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
   `user_id` bigint(20) UNSIGNED NOT NULL,
   `pr` text COMMENT 'ひとこと',
   `hobby` text COMMENT '趣味',
   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`),
   UNIQUE KEY `uq_profiles_user_id` (`user_id`),
   FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
