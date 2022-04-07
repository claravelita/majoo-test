create table outlets
(
    id          bigserial   not null
        constraint outlets_pkey
            primary key,
    merchant_id bigint      not null
        constraint fk_merchants_outlets
            references merchants,
    outlet_name varchar(40) not null,
    created_at  timestamp with time zone,
    created_by  bigint,
    updated_at  timestamp with time zone,
    updated_by  bigint
);

alter table outlets
    owner to postgres;

INSERT INTO public.outlets (id, merchant_id, outlet_name, created_at, created_by, updated_at, updated_by) VALUES (1, 1, 'Merchalbum Bekasi', '2022-04-06 15:44:39.817341', 0, '2022-04-06 15:44:39.817341', 0);
INSERT INTO public.outlets (id, merchant_id, outlet_name, created_at, created_by, updated_at, updated_by) VALUES (2, 1, 'Merchalbum Jakarta', '2022-04-06 15:44:39.817341', 0, '2022-04-06 15:44:39.817341', 0);
INSERT INTO public.outlets (id, merchant_id, outlet_name, created_at, created_by, updated_at, updated_by) VALUES (3, 2, 'Goodies Tangerang', '2022-04-06 15:44:39.817341', 0, '2022-04-06 15:44:39.817341', 0);