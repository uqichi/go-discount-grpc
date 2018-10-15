-- -----------------------------------------------------
-- Table `discounts`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `discounts` (
  `id` CHAR(36) NOT NULL,
  `tenant_id` CHAR(36) NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `type` TINYINT UNSIGNED NOT NULL,
  `max_available_count` SMALLINT UNSIGNED NOT NULL,
  `start_time` DATETIME NOT NULL,
  `end_time` DATETIME NOT NULL,
  `repeat_period_type` TINYINT UNSIGNED NOT NULL,
  `repeat_period_value` SMALLINT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  `updated_at` TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_discounts_tenants1_idx` (`tenant_id` ASC));


-- -----------------------------------------------------
-- Table `discount_conditions`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `discount_conditions` (
  `id` CHAR(36) NOT NULL,
  `discount_id` CHAR(36) NOT NULL,
  `category_id` CHAR(36) NOT NULL,
  `min_total_purchase_quantity` SMALLINT UNSIGNED NOT NULL,
  `min_total_purchase_amount` MEDIUMINT UNSIGNED NOT NULL,
  `is_member` TINYINT UNSIGNED NOT NULL,
  `is_first_purchase` TINYINT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  `updated_at` TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_discount_conditions_discounts1_idx` (`discount_id` ASC),
  INDEX `fk_discount_conditions_categories1_idx` (`category_id` ASC));


-- -----------------------------------------------------
-- Table `discount_condition_products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `discount_condition_products` (
  `id` CHAR(36) NOT NULL,
  `condition_id` CHAR(36) NOT NULL,
  `product_id` CHAR(36) NOT NULL,
  `min_purchase_quantity` SMALLINT UNSIGNED NOT NULL,
  `min_purchase_amount` MEDIUMINT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  `updated_at` TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_discount_condition_products_discount_conditions1_idx` (`condition_id` ASC),
  INDEX `fk_discount_condition_products_products1_idx` (`product_id` ASC));


-- -----------------------------------------------------
-- Table `discount_rewards`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `discount_rewards` (
  `id` CHAR(36) NOT NULL,
  `discount_id` CHAR(36) NOT NULL,
  `total_fixed_amount` MEDIUMINT UNSIGNED NOT NULL,
  `total_percent_point` FLOAT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  `updated_at` TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_discount_rewards_discounts1_idx` (`discount_id` ASC));


-- -----------------------------------------------------
-- Table `discount_reward_products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `discount_reward_products` (
  `id` CHAR(36) NOT NULL,
  `reward_id` CHAR(36) NOT NULL,
  `product_id` CHAR(36) NOT NULL,
  `fixed_amount` MEDIUMINT UNSIGNED NOT NULL,
  `percent_point` FLOAT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  `updated_at` TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_discount_reward_products_discount_rewards1_idx` (`reward_id` ASC),
  INDEX `fk_discount_reward_products_products1_idx` (`product_id` ASC));

