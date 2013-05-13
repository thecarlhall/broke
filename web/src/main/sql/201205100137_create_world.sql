create table if not exists `frequency` (
    `id` integer primary key auto_increment,
    `key` varchar(25) not null,
    `value` varchar(50) not null
);

insert into `frequency` (`key`, `value`)
values ('month', 'Month'),
    ('quarter', 'Quarter'),
    ('semi_annual', 'Semi-annual'),
    ('annual', 'Annual');


create table if not exists `accounts` (
    `id` integer primary key auto_increment
    `name` varchar(255) not null,
    `due_day` integer not null,
    `frequency_id` integer not null
);

create table if not exists `ledger` (
    `id` integer primary key auto_increment,
    `account_id` integer not null,
    `amount` decimal(10,2) not null,
    `date_paid` date not null
);