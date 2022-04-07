create table transactions
(
    id          bigserial not null
        constraint transactions_pkey
            primary key,
    merchant_id bigint    not null
        constraint fk_merchants_transactions
            references merchants,
    outlet_id   bigint
        constraint fk_outlets_transactions
            references outlets,
    bill_total  numeric,
    created_at  timestamp with time zone,
    created_by  bigint,
    updated_at  timestamp with time zone,
    updated_by  bigint
);

alter table transactions
    owner to postgres;

INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (1, 1, 1, 200000, '2021-11-04 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (2, 1, 1, 150000, '2021-11-04 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (3, 1, 1, 400000, '2021-11-06 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (4, 1, 1, 50000, '2021-11-06 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (5, 1, 1, 140000, '2021-11-11 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (6, 1, 1, 80000, '2021-11-14 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (7, 1, 1, 990000, '2021-11-16 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (8, 1, 2, 400000, '2021-11-03 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (9, 1, 2, 4800000, '2021-11-04 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (10, 1, 2, 800000, '2021-11-06 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (11, 1, 2, 900000, '2021-11-08 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (12, 1, 2, 700000, '2021-11-16 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (13, 1, 2, 450000, '2021-11-21 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (14, 1, 2, 500000, '2021-11-21 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (15, 1, 2, 800000, '2021-11-23 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (16, 1, 2, 670000, '2021-11-25 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (17, 1, 2, 1020000, '2021-11-26 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (18, 2, 3, 40000, '2021-11-03 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (19, 2, 3, 65000, '2021-11-04 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (20, 2, 3, 150000, '2021-11-07 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (21, 2, 3, 45000, '2021-11-09 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (22, 2, 3, 89000, '2021-11-14 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (23, 2, 3, 90000, '2021-11-16 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);
INSERT INTO public.transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) VALUES (24, 2, 3, 40000, '2021-11-18 15:44:39.807388', 0, '2022-04-06 15:44:39.822846', 0);