use wallet;

CREATE TABLE IF NOT EXISTS `clients`(`id` varchar(255) NOT NULL, `name` varchar(255) NOT NULL, `email` varchar(100) NOT NULL, `created_at` datetime NOT NULL, PRIMARY KEY(`id`));
CREATE TABLE IF NOT EXISTS `accounts`(`id` varchar(255) NOT NULL, `client_id` varchar(255) NOT NULL, `balance` float NOT NULL, `created_at` datetime NOT NULL, PRIMARY KEY(`id`));
CREATE TABLE IF NOT EXISTS `transactions`(`id` varchar(255) NOT NULL, `account_id_from` varchar(255) NOT NULL, `account_id_to` varchar(255) NOT NULL, `amount` float NOT NULL, created_at datetime NOT NULL, PRIMARY KEY(`id`));

insert into clients(id, name, email, created_at) values ('e8a56770-2897-4117-bbda-e130a3c03d9a', 'Cliente 1', 'cliente1@abc.com.br', now());
insert into clients(id, name, email, created_at) values ('8416e7bd-a55e-4171-bacc-d05ffb8d76ad', 'Cliente 2', 'cliente2@abc.com.br', now());
insert into clients(id, name, email, created_at) values ('fd72832d-370f-4d62-a1c6-f7d174530107', 'Cliente 3', 'cliente3@abc.com.br', now());

insert into accounts(id, client_id, balance, created_at) values ('e885b7c4-038d-430f-9670-2d1079a22c1e', 'e8a56770-2897-4117-bbda-e130a3c03d9a', 987, now());
insert into accounts(id, client_id, balance, created_at) values ('c71ebace-7ea6-4e4f-859c-e6a2818b2890', '8416e7bd-a55e-4171-bacc-d05ffb8d76ad', 654, now());
insert into accounts(id, client_id, balance, created_at) values ('de5becee-61ae-4cc7-a03c-70d8da02bc10', 'fd72832d-370f-4d62-a1c6-f7d174530107', 321, now());