create table merchants
(
    id            bigserial   not null
        constraint merchants_pkey
            primary key,
    user_id       bigint      not null
        constraint fk_users_merchants
            references users,
    merchant_name varchar(40) not null,
    created_at    timestamp with time zone,
    created_by    bigint,
    updated_at    timestamp with time zone,
    updated_by    bigint
);

alter table merchants
    owner to postgres;

INSERT INTO public.merchants (id, user_id, merchant_name, created_at, created_by, updated_at, updated_by) VALUES (1, 1, 'Merchalbum', '2022-04-06 15:44:39.813184', 0, '2022-04-06 15:44:39.813184', 0);
INSERT INTO public.merchants (id, user_id, merchant_name, created_at, created_by, updated_at, updated_by) VALUES (2, 2, 'Goodies', '2022-04-06 15:44:39.813184', 0, '2022-04-06 15:44:39.813184', 0);