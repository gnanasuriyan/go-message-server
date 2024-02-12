-- +goose Up

CREATE TABLE IF NOT EXISTS `users` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `active` BOOLEAN NOT NULL DEFAULT TRUE,
    `shadow_active` BOOLEAN GENERATED ALWAYS AS (IF(`active`, true, NULL)) STORED,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     UNIQUE `active_unique_user`(`username`, `shadow_active`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS `messages` (
    `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `fk_user` int(11) NOT NULL,
    `content` text NOT NULL,
    `active` BOOLEAN NOT NULL DEFAULT TRUE,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    FOREIGN KEY (`fk_user`) REFERENCES `users`(`id`) ON DELETE NO ACTION
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;
