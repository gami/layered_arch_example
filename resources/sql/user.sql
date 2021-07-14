CREATE TABLE `users` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
   `name` varchar(255) NOT NULl,
   `email` varchar(255),
   `status` ENUM('active', 'invited', 'inactive'),
   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
