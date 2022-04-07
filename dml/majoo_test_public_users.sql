create table users
(
    id         bigserial not null
        constraint users_pkey
            primary key,
    name       varchar(45),
    user_name  varchar(45),
    password   varchar(225),
    created_at timestamp with time zone,
    created_by bigint,
    updated_at timestamp with time zone,
    updated_by bigint
);

alter table users
    owner to postgres;

INSERT INTO public.users (id, name, user_name, password, created_at, created_by, updated_at, updated_by) VALUES (1, 'Clara Velita Pranolo', 'claravelita', '$2a$10$dSnFxsSJnDsuYlwCrn62cOiD33OfC4P5eCUIRVRmpOq9LiPNtLbvq', '2022-04-06 15:44:39.809569', 0, '2022-04-06 15:44:39.809569', 0);
INSERT INTO public.users (id, name, user_name, password, created_at, created_by, updated_at, updated_by) VALUES (2, 'Irene', 'renebae', '$2a$10$USEl.o9CytW2KXw5jCxPr.dK5O/G7MVtsqPnvfGNdcgXv9s8FZMIq', '2022-04-06 15:44:39.809569', 0, '2022-04-06 15:44:39.809569', 0);