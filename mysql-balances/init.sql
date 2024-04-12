use balances;

CREATE TABLE IF NOT EXISTS `balances`(`id` int NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `account_id` varchar(255) NOT NULL, balance float, PRIMARY KEY(`id`));

insert into balances(name, account_id, balance) values ('Cliente 1', 'e885b7c4-038d-430f-9670-2d1079a22c1e', 987);
insert into balances(name, account_id, balance) values ('Cliente 2', 'c71ebace-7ea6-4e4f-859c-e6a2818b2890', 654);
insert into balances(name, account_id, balance) values ('Cliente 3', 'de5becee-61ae-4cc7-a03c-70d8da02bc10', 321);