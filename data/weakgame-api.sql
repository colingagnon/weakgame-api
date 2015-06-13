-- autoapi -d="weakgame-api" -u="root" -h="localhost" -P="3306"
-- go run bin/main.go -d="weakgame-api" -u="root" -h="localhost" -P="3306"
-- go run bin/main.go -d="weakgame-api" -u="weakgame" -h="localhost" -P="3306"

-- CREATE DATABASE IF NOT EXISTS `weakgame-api`;
-- CREATE USER 'weakgame'@'localhost' IDENTIFIED BY 'weakpass';
-- GRANT ALL PRIVILEGES ON `weakgame-api` . * TO 'weakgame'@'localhost';
-- FLUSH PRIVILEGES;

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
DROP TABLE IF EXISTS `fights`;
DROP TABLE IF EXISTS `monsters`;
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `products`;

CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(256) NOT NULL,
  `password` varchar(2000) NOT NULL,
  `createdOn` datetime NOT NULL,
  `unicorns` int NOT NULL default 0,
  `hp` int NOT NULL default 5,
  `experienceLevel` int NOT NULL default 1,
  `experiencePoints` int NOT NULL default 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=2;

INSERT INTO `users` (`id`, `email`, `password`, `createdOn`, `unicorns`, `hp`) VALUES
    (1, 'test@test.com', sha2('test123', 256), now(), 5, 5);

CREATE TABLE IF NOT EXISTS `monsters` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(2000) default '/default.jpg',
  `monsterLevel` int NOT NULL default 1,
  `hp` int NOT NULL default 2,
  `accuracy` float(4,2) NOT NULL default 1.00,
  `dodge` float(4,2) NOT NULL default 0.00,
  `damageLow` int NOT NULL default 0,
  `damageHigh` int NOT NULL default 1,
  `experiencePoints` int NOT NULL default 1,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

INSERT INTO `monsters` (`id`, `url`, `monsterLevel`, `hp`, `accuracy`, `dodge`, `damageLow`, `damageHigh`, `experiencePoints`) VALUES

-- Level 1
  (null, 'slimeblue.jpg', 1, 2, 0.45, 0.00, 0, 1, 6),
  (null, 'npc_scaryprofessor.jpg', 1, 1000, 0.50, 0.00, 1, 1, 1000),

  (null, 'goblinlittle.jpg', 1, 3, 0.50, 0.15, 1, 1, 6),
  (null, 'animaldemonratthing.jpg', 1, 5, 0.30, 0.20, 1, 1, 6),
  
-- Level 2
  (null, 'slimegreen.jpg', 2, 5, 0.80, 0.10, 0, 1, 9),
  (null, 'goblin2010a.jpg', 2, 6, 0.60, 0.20, 1, 2, 9),
  
-- Level 3
  (null, 'undead_deathwing.jpg', 3, 5, 0.30, 0.15, 1, 2, 15),
  (null, 'slimebubble.jpg', 3, 7, 0.40, 0.10, 0, 3, 17),
  
-- Level 4
  (null, 'undead_skeletonreg.jpg', 4, 5, 0.80, 0.10, 0, 3, 30),
  (null, 'npcgnome_enchanter_by_chrismcfann-d47s5se.jpg', 4, 3, 0.60, 0.25, 0, 2, 34),

-- Level 5
  (null, 'goblinhob.jpg', 5, 8, 0.55, 0.10, 0, 6, 52),
  (null, 'undead_skeletonguardian.jpg', 5, 9, 0.60, 0.15, 1, 4, 54),

-- Level 6
  (null, 'animal_werewolf.jpg', 6, 12, 0.70, 0.25, 1, 3, 82),
  (null, 'orc_warrior_by_dimelife-d6cn5qh.jpg', 6, 23, 0.60, 0.20, 0, 4, 91),
  
-- Level 7
  (null, 'orcarcher.jpg', 7, 22, 0.90, 0.30, 0, 6, 140),
  (null, 'undead_Skeletal_Tomb_Guardian.jpg', 7, 33, 0.40, 0.20, 0, 8, 150),
  
-- Level 8
  (null, 'golem_UH_Stone_Giant_Runecarver.jpg', 8, 60, 0.35, 0.00, 0, 8, 215),
  (null, 'slime_gelcube.jpg', 8, 36, 0.40, 0.50, 2, 3, 210),

-- Level 9
  (null, 'slimeBlackooze.jpg', 9, 43, 0.70, 0.25, 0, 3, 380),
  (null, 'slimemetal.jpg', 9, 10, 0.60, 0.45, 2, 4, 420),
  
-- Level 10
  (null, 'slimeblueking.jpg', 10, 90, 0.80, 0.05, 0, 12, 500),
  (null, 'animalowlbear2.jpg', 10, 60, 0.70, 0.40, 0, 8, 550),
  
-- Level 11
  (null, 'slimebubblemetal.jpg', 11, 16, 0.80, 0.45, 4, 6, 800),
  (null, 'animaldnd___hook_horror_by.jpg', 11, 110, 0.80, 0.40, 4, 12, 1380),
  
-- Level 12
  (null, 'gnollBlood_War_-_046_Demonic_Gnoll_Priestess.jpg', 12, 0.70, 0.70, 0.30, 0, 12, 2200),
  (null, 'demon_114d6ea23631f49cc751aa110247e130.jpg', 12, 1000, 0.80, 0.40, 30, 50, 2300),

-- Level 13
  (null, 'demon_Spined_Devil.jpg', 13, 180, 0.80, 0.20, 3, 6, 4200),
  (null, 'demon_Shadow_demon.jpg', 13, 30, 0.50, 0.45, 0, 10, 4300),
  
-- Level 14
  (null, 'ogre.jpg', 14, 210, 0.80, 0.10, 0, 16, 6200),
  (null, 'animals_foresttreemage.jpg', 14, 80, 0.80, 0.55, 4, 16, 6300),
  
-- Level 15
  (null, 'slimemetalking.jpg', 15, 40, 0.80, 0.65, 8, 10, 12000),
  (null, 'animalfiend-kocrachon.jpg', 15, 270, 0.35, 0.40, 0, 20, 8000),
  
-- Level 16
  (null, 'demon_Dungeons_of_Dread_-_016_Shadow_Demon.jpg', 16, 200, 0.40, 0.40, 0, 30, 11000),
  (null, 'golem061b8fbf36346a0879a25bcb5f3e8601.jpg', 16, 320, 0.70, 0.00, 15, 25, 12000),

-- Level 17
  (null, 'demon_20fe6fd3d422bffc2ff696702f243d5c.jpg', 17, 300, 0.80, 0.10, 0, 20, 15000),
  (null, 'demon_01c6b6eae9ddf3efa205d2d807eb71cc.jpg', 17, 250, 0.60, 0.45, 8, 16, 16000),
  
-- Level 18
  (null, 'demon_1905d3974265c0089a56b6de3d20815e.jpg', 18, 3000, 0.70, 0.45, 8, 20, 21000),
  (null, 'slimemonstrous.jpg', 18, 350, 0.80, 0.10, 4, 12, 22000),
  
-- Level 19
  (null, 'demon_Demon4.jpg', 19, 400, 1.05, 0.20, 350, 400, 25000),
  (null, 'npc531b770cfa9d59713c9b31c107d728e8.jpg', 19, 300, 0.90, 0.50, 4, 30, 25000),

-- Level 20+
  (null, 'undead_20150327224259-WL_2.jpg', 20, 600, 1.05, 0.50, 10, 80, 30000),
  (null, 'undead_soothsayer_by_sopossum-d5504j3.jpg', 21, 800, 1.10, 0.20, 20, 80, 40000),
  (null, 'demon_Virgil_s_Shadow_Demon_form.jpg', 25, 1000, 1.15, 0.30, 30, 90, 50000),
  (null, 'dragon_portal-drake.jpg', 26, 1200, 1.15, 0.40, 40, 100, 60000),
  (null, 'dragon1e2cd71d9b351f7a000201bc34094b7b.jpg', 27, 1400, 1.15, 0.50, 50, 120, 70000);

CREATE TABLE IF NOT EXISTS `fights` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `monsterId` int(11) unsigned NOT NULL,
  `userId` int(11) unsigned NOT NULL,
  `userCurrentHp` int NOT NULL default 1,
  `monsterCurrentHp` int NOT NULL default 1,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;
  
-- Nothing to insert for this table
  
-- Products are the conversion of dollars to unicorns
CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) default 'Product',
  `tokens` int NOT NULL,
  `price` float(8,2) NOT NULL,
  `discount` int NOT NULL,
  `available` enum('n', 'y') default 'n',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=5;

INSERT INTO `products` (`id`, `name`, `tokens`, `price`, `discount`, `available`) VALUES
  (1, 'Basic Package', 10, 5, 0, 'y'),
  (2, 'More Package', 20, 10, 5, 'y'),
  (3, 'Large Package', 40, 20, 15, 'y'),
  (4, 'Super Package', 80, 40, 25, 'y');

ALTER TABLE `fights`
ADD CONSTRAINT `fights_ibfk_1` FOREIGN KEY (`userId`) REFERENCES `users` (`id`);

ALTER TABLE `fights`
ADD CONSTRAINT `fights_ibfk_2` FOREIGN KEY (`monsterId`) REFERENCES `monsters` (`id`);

DROP TRIGGER IF EXISTS `discount_checker_insert`;
DELIMITER //
CREATE TRIGGER `discount_checker_insert` BEFORE INSERT ON `products`
 FOR EACH ROW IF (NEW.discount < 0 OR NEW.discount > 25) THEN
  SIGNAL SQLSTATE '12345' SET MESSAGE_TEXT = 'The largest discount is a maximum of 25%';
END IF
//
DELIMITER ;

DROP TRIGGER IF EXISTS `discount_checker_update`;
DELIMITER //
CREATE TRIGGER `discount_checker_update` BEFORE UPDATE ON `products`
 FOR EACH ROW IF (NEW.discount < 0 OR NEW.discount > 25) THEN
  SIGNAL SQLSTATE '12345' SET MESSAGE_TEXT = 'The largest discount is a maximum of 25%';
END IF
//
DELIMITER ;