SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema melisprint
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema melisprint
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `melisprint` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `melisprint` ;

-- -----------------------------------------------------
-- Table `melisprint`.`buyers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`buyers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `card_number_id` TEXT NOT NULL,
  `first_name` TEXT NOT NULL,
  `last_name` TEXT NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 3
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `melisprint`.`warehouses`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`warehouses` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `address` TEXT NOT NULL,
  `telephone` TEXT NOT NULL,
  `warehouse_code` TEXT NOT NULL,
  `minimum_capacity` INT NOT NULL,
  `minimum_temperature` INT NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 3
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `melisprint`.`employees`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`employees` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `card_number_id` TEXT NOT NULL,
  `first_name` TEXT NOT NULL,
  `last_name` TEXT NOT NULL,
  `warehouses_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_employees_warehouses1_idx` (`warehouses_id` ASC) VISIBLE,
  CONSTRAINT `fk_employees_warehouses1`
    FOREIGN KEY (`warehouses_id`)
    REFERENCES `melisprint`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
AUTO_INCREMENT = 4
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `melisprint`.`sellers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`sellers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `cid` INT NOT NULL,
  `company_name` TEXT NOT NULL,
  `address` TEXT NOT NULL,
  `telephone` VARCHAR(15) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 4
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `melisprint`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`products` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` TEXT NOT NULL,
  `expiration_rate` FLOAT NOT NULL,
  `freezing_rate` FLOAT NOT NULL,
  `height` FLOAT NOT NULL,
  `lenght` FLOAT NOT NULL,
  `netweight` FLOAT NOT NULL,
  `product_code` TEXT NOT NULL,
  `recommended_freezing_temperature` FLOAT NOT NULL,
  `width` FLOAT NOT NULL,
  `id_product_type` INT NOT NULL,
  `sellers_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_products_sellers1_idx` (`sellers_id` ASC) VISIBLE,
  CONSTRAINT `fk_products_sellers1`
    FOREIGN KEY (`sellers_id`)
    REFERENCES `melisprint`.`sellers` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
AUTO_INCREMENT = 8
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `melisprint`.`sections`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`sections` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `section_number` INT NOT NULL,
  `current_temperature` INT NOT NULL,
  `minimum_temperature` INT NOT NULL,
  `current_capacity` INT NOT NULL,
  `minimum_capacity` INT NOT NULL,
  `maximum_capacity` INT NOT NULL,
  `id_product_type` INT NOT NULL,
  `warehouses_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_sections_warehouses1_idx` (`warehouses_id` ASC) VISIBLE,
  CONSTRAINT `fk_sections_warehouses1`
    FOREIGN KEY (`warehouses_id`)
    REFERENCES `melisprint`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
AUTO_INCREMENT = 4
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `melisprint`.`localities`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`localities` (
  `id` INT NOT NULL,
  `locality_name` VARCHAR(45) NOT NULL,
  `province_name` VARCHAR(45) NOT NULL,
  `country_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `melisprint`.`carries`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`carries` (
  `id` INT NOT NULL,
  `cid` INT NOT NULL,
  `company_name` VARCHAR(45) NOT NULL,
  `address` VARCHAR(45) NOT NULL,
  `telephone` VARCHAR(45) NOT NULL,
  `localities_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_carries_localities1_idx` (`localities_id` ASC) VISIBLE,
  CONSTRAINT `fk_carries_localities1`
    FOREIGN KEY (`localities_id`)
    REFERENCES `melisprint`.`localities` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `melisprint`.`product_batches`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`product_batches` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `batch_number` INT NOT NULL,
  `current_quantity` INT NOT NULL,
  `current_temperature` DECIMAL(4,1) NOT NULL,
  `due_date` DATETIME NOT NULL,
  `initial_quantity` INT NOT NULL,
  `manufacturing_date` DATE NOT NULL,
  `manufacturing_hour` TIME NOT NULL,
  `minumum_temperature` DECIMAL(4,1) NOT NULL,
  `products_id` INT NOT NULL,
  `sections_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_product_Batches_products1_idx` (`products_id` ASC) VISIBLE,
  INDEX `fk_product_batches_sections1_idx` (`sections_id` ASC) VISIBLE,
  CONSTRAINT `fk_product_Batches_products1`
    FOREIGN KEY (`products_id`)
    REFERENCES `melisprint`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_product_batches_sections1`
    FOREIGN KEY (`sections_id`)
    REFERENCES `melisprint`.`sections` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `melisprint`.`product_records`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`product_records` (
  `id` INT NOT NULL,
  `last_update_date` DATE NOT NULL,
  `purchase_price` FLOAT NOT NULL,
  `sale_price` FLOAT NOT NULL,
  `products_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_product_records_products1_idx` (`products_id` ASC) VISIBLE,
  CONSTRAINT `fk_product_records_products1`
    FOREIGN KEY (`products_id`)
    REFERENCES `melisprint`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `melisprint`.`inbound_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`inbound_orders` (
  `id` INT NOT NULL,
  `order_date` DATE NOT NULL,
  `order_number` INT NOT NULL,
  `employees_id` INT NOT NULL,
  `warehouses_id` INT NOT NULL,
  `product_Batches_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_inbound_orders_employees1_idx` (`employees_id` ASC) VISIBLE,
  INDEX `fk_inbound_orders_warehouses1_idx` (`warehouses_id` ASC) VISIBLE,
  INDEX `fk_inbound_orders_product_Batches1_idx` (`product_Batches_id` ASC) VISIBLE,
  CONSTRAINT `fk_inbound_orders_employees1`
    FOREIGN KEY (`employees_id`)
    REFERENCES `melisprint`.`employees` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_inbound_orders_warehouses1`
    FOREIGN KEY (`warehouses_id`)
    REFERENCES `melisprint`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_inbound_orders_product_Batches1`
    FOREIGN KEY (`product_Batches_id`)
    REFERENCES `melisprint`.`product_batches` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `melisprint`.`order_status`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`order_status` (
  `id` INT NOT NULL,
  `description` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `melisprint`.`purchase_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `melisprint`.`purchase_orders` (
  `id` INT NOT NULL,
  `order_number` INT NOT NULL,
  `order_date` DATE NOT NULL,
  `tracking_code` VARCHAR(45) NOT NULL,
  `product_records_id` INT NOT NULL,
  `buyers_id` INT NOT NULL,
  `order_status_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_purchase_orders_product_records1_idx` (`product_records_id` ASC) VISIBLE,
  INDEX `fk_purchase_orders_buyers1_idx` (`buyers_id` ASC) VISIBLE,
  INDEX `fk_purchase_orders_order_status1_idx` (`order_status_id` ASC) VISIBLE,
  CONSTRAINT `fk_purchase_orders_product_records1`
    FOREIGN KEY (`product_records_id`)
    REFERENCES `melisprint`.`product_records` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_purchase_orders_buyers1`
    FOREIGN KEY (`buyers_id`)
    REFERENCES `melisprint`.`buyers` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_purchase_orders_order_status1`
    FOREIGN KEY (`order_status_id`)
    REFERENCES `melisprint`.`order_status` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;