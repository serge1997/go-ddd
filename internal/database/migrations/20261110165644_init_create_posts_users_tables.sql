-- Create "users" table
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `id,primaryKey` bigint NULL,
  `name` longtext NULL,
  `email` longtext NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_users_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "posts" table
CREATE TABLE `posts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `id,primaryKey` bigint NULL,
  `title` longtext NULL,
  `author_id` bigint unsigned NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_users_posts` (`author_id`),
  INDEX `idx_posts_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_users_posts` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
