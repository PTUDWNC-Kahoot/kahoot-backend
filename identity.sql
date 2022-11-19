CREATE TABLE `users` (
  `id` int AUTO_INCREMENT NOT NULL,
  `username` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP,  
  PRIMARY KEY (`id`)
);
insert into users(username,password)values("user1","user12345678");
insert into users(username,password) values("user2","user12345678");
insert into users(username,password) values("user3"," user12345678");