
CREATE TABLE `persons` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NULL DEFAULT NULL,
	`city` VARCHAR(255) NULL DEFAULT NULL,
	`birth_date` DATETIME NULL DEFAULT NULL,
	`weight` DECIMAL(8,4) NULL DEFAULT NULL,
	`height` DECIMAL(8,4) NULL DEFAULT NULL,
	`created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	INDEX `city` (`city`),
	INDEX `birth_date` (`birth_date`),
	INDEX `weight` (`weight`),
	INDEX `height` (`height`)
) COLLATE='utf8_general_ci' ENGINE=InnoDB;