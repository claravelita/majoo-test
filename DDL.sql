create table if not exists users
(
	id bigserial not null
		constraint users_pkey
			primary key,
	name varchar(45),
	user_name varchar(45),
	password varchar(225),
	created_at timestamp with time zone,
	created_by bigint,
	updated_at timestamp with time zone,
	updated_by bigint
);

alter table users owner to postgres;

create table if not exists merchants
(
	id bigserial not null
		constraint merchants_pkey
			primary key,
	user_id bigint not null
		constraint fk_users_merchants
			references users,
	merchant_name varchar(40) not null,
	created_at timestamp with time zone,
	created_by bigint,
	updated_at timestamp with time zone,
	updated_by bigint
);

alter table merchants owner to postgres;

create table if not exists outlets
(
	id bigserial not null
		constraint outlets_pkey
			primary key,
	merchant_id bigint not null
		constraint fk_merchants_outlets
			references merchants,
	outlet_name varchar(40) not null,
	created_at timestamp with time zone,
	created_by bigint,
	updated_at timestamp with time zone,
	updated_by bigint
);

alter table outlets owner to postgres;

create table if not exists transactions
(
	id bigserial not null
		constraint transactions_pkey
			primary key,
	merchant_id bigint not null
		constraint fk_merchants_transactions
			references merchants,
	outlet_id bigint
		constraint fk_outlets_transactions
			references outlets,
	bill_total numeric,
	created_at timestamp with time zone,
	created_by bigint,
	updated_at timestamp with time zone,
	updated_by bigint
);

alter table transactions owner to postgres;

