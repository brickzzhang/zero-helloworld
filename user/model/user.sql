CREATE TABLE `user`
(
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(255) NOT NULL UNIQUE COMMENT 'name',
  `password` varchar(255) NOT NULL COMMENT 'password',
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;