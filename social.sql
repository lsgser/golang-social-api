DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `surname` varchar(255) NOT NULL,
    `username` varchar(255) NOT NULL UNIQUE,
    `email` varchar(255) NOT NULL UNIQUE,
    `email_verified_at` timestamp NULL DEFAULT NULL,
    `password` varchar(255) NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
	UNIQUE KEY `users_username_unique` (`username`),
  	UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `profiles`;

CREATE TABLE `profiles` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`user_id` bigint(20) unsigned NOT NULL,
	`profile_picture` varchar(255) DEFAULT NULL,
	`gender` tinyint(3) unsigned DEFAULT NULL,
	`birth_date` date DEFAULT NULL,
	`residence` varchar(255) DEFAULT NULL,
	`created_at` timestamp NULL DEFAULT NULL,
	`updated_at` timestamp NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY `profiles_user_id_unique` (`user_id`),
	KEY `profiles_user_id_index` (`user_id`),
	CONSTRAINT `profiles_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `faculties`;

CREATE TABLE `faculties` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`school_id` bigint(20) unsigned NOT NULL,
	`faculty` varchar(255) NOT NULL,
	`created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	KEY `faculties_school_id_foreign` (`school_id`),
	CONSTRAINT `faculties_school_id_foreign` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;