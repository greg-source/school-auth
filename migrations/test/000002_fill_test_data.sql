-- +migrate Up
INSERT INTO `auth` VALUES (1,'www-dfq92-sqfwf'),(2,'ffff-2918-xcas');
INSERT INTO `user` VALUES (1,'test'),(2,'admin'),(3,'guest');
INSERT INTO `user_data` VALUES (1,'гімназія №179 міста Києва'),(2,'ліцей №227'),(3,'Медична гімназія №33 міста Києва');
INSERT INTO `user_profile` VALUES (1,'Александр','Школьный','+38050123455','ул. Сибирская 2','Киев'),(2,'Дмитрий','Арбузов','+38065133223','ул. Белая 4','Харьков'),(3,'Василий','Шпак','+38055221166','ул. Северная 5','Житомир');