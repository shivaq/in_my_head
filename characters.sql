CREATE DATABASE `characterdb` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE `characterdb`.`characters` (
    -- Id
    `characterId` INT NOT NULL AUTO_INCREMENT,
    `UserID` VARCHAR(60) NOT NULL,

    -- Appearance
	`characterType` VARCHAR(60) NOT NULL,
	`circleColorName` VARCHAR(60) NOT NULL,
	`savedFileName` VARCHAR(255) NOT NULL,
	`blobImage` BLOB,
	`isPicUsed` BOOLEAN NOT NULL DEFAULT 0,

	-- Name
	`nickName` VARCHAR(255) NOT NULL,
	`firstName` VARCHAR(255),
	`middleName` VARCHAR(255),
	`lastName` VARCHAR(255),
	`originalName` VARCHAR(255),

	-- Details
	`isImaginary` BOOLEAN NOT NULL DEFAULT 0,
	`birthDate` INTEGER NOT NULL DEFAULT 1,
	`description` VARCHAR(255) ,
	`level` INTEGER NOT NULL DEFAULT 1,
	`isLiked` BOOLEAN NOT NULL DEFAULT 0,

	-- Date
	`registeredDate` INTEGER NOT NULL DEFAULT 1,
	`updatedDate` INTEGER NOT NULL DEFAULT 1,

	-- Delete status
	`onDeleteLock` BOOLEAN NOT NULL DEFAULT 0,
	`inLimbo` BOOLEAN NOT NULL DEFAULT 0,

	-- Reserved
	`reserveNumber01` INTEGER NOT NULL DEFAULT 0,
	`reserveNumber02` INTEGER NOT NULL DEFAULT 0,
	`reserveNumber03` INTEGER NOT NULL DEFAULT 0,
	`reserveNumber04` INTEGER NOT NULL DEFAULT 0,
	`reserveNumber05` INTEGER NOT NULL DEFAULT 0,
	`reserveNumber06` INTEGER NOT NULL DEFAULT 0,
	`reserveNumber07` INTEGER NOT NULL DEFAULT 0,
	`reserveNumber08` INTEGER NOT NULL DEFAULT 0,
	`reserveNumber09` INTEGER NOT NULL DEFAULT 0,
	`reserveNumber10` INTEGER NOT NULL DEFAULT 0,
	`reserveNumber11` INTEGER NOT NULL DEFAULT 0,

    PRIMARY KEY (`characterId`));

INSERT INTO `characterdb`.`characters`
(`UserID`, `characterType`, `circleColorName`, `savedFileName`, `nickName`)
VALUES ("Human A","man","yellow", "NOT SAVED", "伊邪那岐");
