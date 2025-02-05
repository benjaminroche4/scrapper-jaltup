-- -----------------------------------------------------
-- Schema jaltup
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `jaltup` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `jaltup` ;

-- -----------------------------------------------------
-- Table `jaltup`.`category`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jaltup`.`category` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `public_id` VARCHAR(20) NOT NULL,
  `name` VARCHAR(120) NOT NULL,
  `slug` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `UNIQ_64C19C1B5B48B91` (`public_id` ASC) VISIBLE)
ENGINE = InnoDB
AUTO_INCREMENT = 1298
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `jaltup`.`company`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jaltup`.`company` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `public_id` VARCHAR(20) NOT NULL,
  `name` VARCHAR(120) NOT NULL,
  `siret` VARCHAR(14) NULL DEFAULT NULL,
  `contact_email` VARCHAR(120) NULL DEFAULT NULL,
  `phone_number` VARCHAR(20) NULL DEFAULT NULL,
  `website_url` VARCHAR(255) NULL DEFAULT NULL,
  `logo` VARCHAR(255) NULL DEFAULT NULL,
  `created_at` DATETIME NOT NULL,
  `slug` VARCHAR(255) NOT NULL,
  `verified` TINYINT(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `UNIQ_4FBF094FB5B48B91` (`public_id` ASC) VISIBLE)
ENGINE = InnoDB
AUTO_INCREMENT = 7762
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `jaltup`.`contact`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jaltup`.`contact` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(60) NOT NULL,
  `last_name` VARCHAR(60) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `phone_number` VARCHAR(25) NULL DEFAULT NULL,
  `contact_type` VARCHAR(255) NOT NULL,
  `help_type` VARCHAR(255) NOT NULL,
  `message` LONGTEXT NULL DEFAULT NULL,
  `created_at` DATETIME NOT NULL COMMENT '(DC2Type:datetime_immutable)',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `jaltup`.`newsletter`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jaltup`.`newsletter` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NOT NULL COMMENT '(DC2Type:datetime_immutable)',
  `subscribe` TINYINT(1) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `jaltup`.`offer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jaltup`.`offer` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `company_id` INT NULL DEFAULT NULL,
  `public_id` VARCHAR(20) NOT NULL,
  `title` VARCHAR(120) NOT NULL,
  `place` JSON NOT NULL,
  `job` JSON NOT NULL,
  `url` VARCHAR(255) NULL DEFAULT NULL,
  `tag` JSON NULL DEFAULT NULL,
  `status` VARCHAR(40) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `end_date` DATE NOT NULL,
  `end_premium` DATETIME NULL DEFAULT NULL COMMENT '(DC2Type:datetime_immutable)',
  `slug` VARCHAR(255) NOT NULL,
  `premium` TINYINT(1) NOT NULL,
  `external_id` VARCHAR(255) NULL DEFAULT NULL,
  `service_name` VARCHAR(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `UNIQ_29D6873EB5B48B91` (`public_id` ASC) VISIBLE,
  INDEX `IDX_29D6873E979B1AD6` (`company_id` ASC) VISIBLE,
  CONSTRAINT `FK_29D6873E979B1AD6`
    FOREIGN KEY (`company_id`)
    REFERENCES `jaltup`.`company` (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 35633
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `jaltup`.`offer_category`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jaltup`.`offer_category` (
  `offer_id` INT NOT NULL,
  `category_id` INT NOT NULL,
  PRIMARY KEY (`offer_id`, `category_id`),
  INDEX `IDX_7F31A9A353C674EE` (`offer_id` ASC) VISIBLE,
  INDEX `IDX_7F31A9A312469DE2` (`category_id` ASC) VISIBLE,
  CONSTRAINT `FK_7F31A9A312469DE2`
    FOREIGN KEY (`category_id`)
    REFERENCES `jaltup`.`category` (`id`)
    ON DELETE CASCADE,
  CONSTRAINT `FK_7F31A9A353C674EE`
    FOREIGN KEY (`offer_id`)
    REFERENCES `jaltup`.`offer` (`id`)
    ON DELETE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `jaltup`.`refresh_tokens`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jaltup`.`refresh_tokens` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `refresh_token` VARCHAR(128) NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  `valid` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `UNIQ_9BACE7E1C74F2195` (`refresh_token` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `jaltup`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `jaltup`.`user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(180) NOT NULL,
  `roles` JSON NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `public_id` VARCHAR(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `UNIQ_IDENTIFIER_EMAIL` (`email` ASC) VISIBLE,
  UNIQUE INDEX `UNIQ_8D93D649B5B48B91` (`public_id` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;
