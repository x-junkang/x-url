CREATE TABLE `tiny_url` (
  `id` int NOT NULL AUTO_INCREMENT,
  `base_url` varchar(128) NOT NULL DEFAULT '',
  `suffix_url` varchar(256) NOT NULL DEFAULT '',
  `full_url` varchar(384) NOT NULL DEFAULT '',
  `shot_code` varchar(36) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `shot_code` (`shot_code`),
  KEY `full_url` (`full_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;