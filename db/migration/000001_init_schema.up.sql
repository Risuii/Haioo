CREATE TABLE `Haioo`.`Cart` (
    `ID` INT NOT NULL AUTO_INCREMENT,
    `nama` VARCHAR(255) NULL,
    `kodeProduk` VARCHAR(255) NOT NULL,
    `kuantitas` INT NULL,
    `created_at` DATE NULL DEFAULT (now()),
    `update_at` DATE NULL DEFAULT (now()),
    PRIMARY KEY (`ID`)
)