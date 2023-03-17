CREATE SCHEMA IF NOT EXISTS `library` ;
USE `library` ;

CREATE TABLE IF NOT EXISTS `library`.`books` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(45) NOT NULL,
  `author` VARCHAR(45) NOT NULL,
  `category` VARCHAR(45) NOT NULL,
  `year` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;
