create table tb_user (
`user_id` int not null auto_increment,
`user_uuid` varchar(36) not null,
`name` varchar(255),
`email` varchar(255) not null,
`phone` varchar(11) not null,
`cpf` varchar(11) not null,
`password` varchar(255) not null,
`password_salt` varchar(255) not null,
`created_at` datetime not null default current_timestamp,
`updated_at` datetime not null default current_timestamp on update current_timestamp,
`deleted_at_unix` int NOT NULL DEFAULT '0',
PRIMARY KEY (`user_id`),
UNIQUE KEY `tb_user_uuid_idx` (`user_uuid`),
UNIQUE KEY `tb_user_email_uq` (`email`,`deleted_at_unix`),
UNIQUE KEY `tb_user_cpf_uq` (`cpf`,`deleted_at_unix`),
KEY `tb_user_email_idx` (`email`),
KEY `tb_user_cpf_idx` (`cpf`),
KEY `tb_user_phone_idx` (`phone`)
);

create table tb_card (
`card_id` int not null auto_increment,
`user_id` int not null,
`card_uuid` varchar(36) not null,
`card_name` varchar(255),
`max_limit` float not null,
`bill_date` date not null,
`current_limit` float not null,
`balance` float not null,
PRIMARY KEY (`card_id`),
CONSTRAINT `fk_user_card_id` FOREIGN KEY (`user_id`) REFERENCES `tb_user` (`user_id`)
);

create table tb_expenses (
`expense_id` int not null auto_increment,
`card_id` int not null,
`bill_id` int not null,
`expense_uuid` varchar(36) not null,
`expense_value` float not null,
`expense_type` varchar(255) not null,
`payment_type` enum('debit','credit') not null,
`purchase_date` date not null,
PRIMARY KEY (`expense_id`),
constraint `fk_expense_card_id` FOREIGN KEY (`card_id`) REFERENCES `tb_card` (`card_id`),
constraint `fk_expense_bill_id` foreign key (`bill_id`) references `tb_bill` (`bill_id`)
);

create table tb_bill (
`bill_id` int not null auto_increment,
`card_id` int not null,
`bill_value` float,
`bill_date` date not null,
PRIMARY KEY (`bill_id`),
constraint `fk_bill_card_id` FOREIGN KEY (`card_id`) REFERENCES `tb_card` (`card_id`)
);