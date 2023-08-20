# [Atlas](https://atlasgo.io/getting-started) Tutorial

The definition of atlas is written as 

> Atlas is a language-independent tool for managing and migrating database schemas using modern DevOps principles.


## atlas.hcl

The definition of hcl file below is simple template to integrate with `gorm` go orm framework.

This template is defined `Go Program mode` because of model load order. `dev` is atlas default docker basic database.

``` terraform
data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./main.go",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://mysql/8/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
```


## apply atlas with command

### schema diff command
This command compare current db schema state with desired state. And create migration sql file to keep db compatible.

``` bash
$ atlas migrate diff --env gorm
```

Above command create the sql file like below
``` sql
-- Create "products" table
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NULL,
  `price` decimal(9,2) NULL,
  `code` varchar(191) NOT NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `code` (`code`),
  INDEX `idx_products_name` (`name`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "users" table
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_users_name` (`name`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "orders" table
CREATE TABLE `orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `product_id` bigint unsigned NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_orders_product` (`product_id`),
  INDEX `fk_users_orders` (`user_id`),
  CONSTRAINT `fk_orders_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_users_orders` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
```

### apply migration 

``` bash
$ atlas migrate apply --env "gorm" --url "mysql://test:1234@127.0.0.1:3306/test"
```

### change schema

Add  user `Age` field.

``` go
type User struct {
	ID 		  uint 		`gorm:"primaryKey" json:"id"`
	Name 	  string	`gorm:"index" json:"name"`
	>> Age		  uint      `json:"age"`
	Orders []*Order		`json:"orders"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
```
And type the diff command

``` bash
$ atlas migrate diff --env gorm add_user_age
```

Below file is created.

``` sql
-- Modify "users" table
ALTER TABLE `users` ADD COLUMN `age` bigint unsigned NULL;

```

Apply change

