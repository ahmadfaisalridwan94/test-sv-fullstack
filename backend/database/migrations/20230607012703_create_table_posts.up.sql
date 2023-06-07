CREATE TABLE `posts` (
  `Id` INT unsigned NOT NULL AUTO_INCREMENT,
  `Title` VARCHAR(200) NOT NULL,
  `Content` TEXT NOT NULL,
  `Category` VARCHAR(100) NOT NULL,
  `CreatedAt` TIMESTAMP NOT NULL DEFAULT current_timestamp(),
  `UpdatedAt` TIMESTAMP NOT NULL DEFAULT current_timestamp(),
  `Status` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Title` (`Title`)
) ENGINE=InnoDB