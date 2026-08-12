CREATE TABLE IF NOT EXISTS `yweet` (
  `id` varchar(255) NOT NULL,
  `user_id` varchar(255) NOT NULL UNIQUE,
  `content` text NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);
