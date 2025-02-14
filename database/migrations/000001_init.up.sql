create table images
(
    id         serial
        primary key,
    path       varchar(255) not null,
    filename   varchar(255) not null,
    ext        varchar(255) not null,
    hash_uid   varchar(255) not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp
);

alter table images
    owner to postgres;

create table users
(
    id              serial
        primary key,
    name            text,
    phone_number    text         not null
        unique,
    mobile_provider varchar(100) not null,
    refresh_token   varchar(255) not null
        unique,
    created_at      timestamp default CURRENT_TIMESTAMP,
    updated_at      timestamp default CURRENT_TIMESTAMP,
    deleted_at      timestamp,
    debit           bigint    default 0,
    credit          bigint    default 0
);

alter table users
    owner to qahvazor;

create index users_deleted_at_index
    on users (deleted_at);

create index users_refresh_token_index
    on users (refresh_token);

create index users_phone_number_index
    on users (phone_number);

create table partners
(
    id          serial
        primary key,
    logo_id     bigint,
    name        varchar(255),
    description text,
    created_at  timestamp default CURRENT_TIMESTAMP,
    updated_at  timestamp default CURRENT_TIMESTAMP,
    deleted_at  timestamp
);

alter table partners
    owner to qahvazor;

create index partners_deleted_at_index
    on partners (deleted_at);

create table partner_attributes
(
    id              serial
        primary key,
    partner_id      bigint,
    attribute_key   varchar(100) not null,
    attribute_value varchar(255) not null,
    created_at      timestamp default CURRENT_TIMESTAMP,
    updated_at      timestamp default CURRENT_TIMESTAMP,
    deleted_at      timestamp
);

alter table partner_attributes
    owner to qahvazor;

create index partner_attributes_partner_id_index
    on partner_attributes (partner_id);

create index partner_attributes_attribute_key_index
    on partner_attributes (attribute_key);

create index partner_attributes_deleted_at_index
    on partner_attributes (deleted_at);

create table shop_attributes
(
    id              serial
        primary key,
    shop_id         bigint,
    attribute_key   varchar(100) not null,
    attribute_value varchar(255) not null,
    created_at      timestamp default CURRENT_TIMESTAMP,
    updated_at      timestamp default CURRENT_TIMESTAMP,
    deleted_at      timestamp
);

alter table shop_attributes
    owner to qahvazor;

create index shop_attributes_partner_id_index
    on shop_attributes (shop_id);

create index shop_attributes_attribute_key_index
    on shop_attributes (attribute_key);

create index shop_attributes_deleted_at_index
    on shop_attributes (deleted_at);

create table shop_hours
(
    id         serial
        primary key,
    shop_id    bigint,
    week_day   varchar(50) not null,
    open_at    varchar(10) not null,
    close_at   varchar(10) not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp
);

alter table shop_hours
    owner to qahvazor;

create index shop_hours_shop_id_index
    on shop_hours (shop_id);

create index shop_hours_week_day_index
    on shop_hours (week_day);

create index shop_hours_deleted_at_index
    on shop_hours (deleted_at);

create table shop_pictures
(
    id         serial
        primary key,
    shop_id    bigint,
    image_id   bigint,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp
);

alter table shop_pictures
    owner to qahvazor;

create index shop_pictures_shop_id_index
    on shop_pictures (shop_id);

create index shop_pictures_deleted_at_index
    on shop_pictures (deleted_at);

create table drinks
(
    id         serial
        primary key,
    image_id   bigint       not null,
    name       varchar(100) not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp
);

alter table drinks
    owner to qahvazor;

create index drinks_name_index
    on drinks (name);

create index drinks_delete_at_index
    on drinks (deleted_at);

create table partner_drinks
(
    id         serial
        primary key,
    partner_id bigint not null,
    drink_id   bigint not null
);

alter table partner_drinks
    owner to qahvazor;

create index partner_drinks_partner_id_index
    on partner_drinks (partner_id);

create index partner_drinks_drink_id_index
    on partner_drinks (drink_id);

create table partner_users
(
    id              serial
        primary key,
    partner_id      bigint,
    shop_id         bigint,
    name            varchar(100),
    phone_number    varchar(255) not null,
    mobile_provider varchar(100) not null,
    refresh_token   varchar(255) not null
        unique,
    created_at      timestamp default CURRENT_TIMESTAMP,
    updated_at      timestamp default CURRENT_TIMESTAMP,
    deleted_at      timestamp
);

alter table partner_users
    owner to qahvazor;

create index partner_users_deleted_at_index
    on partner_users (deleted_at);

create index partner_users_refresh_token_index
    on partner_users (refresh_token);

create index partner_users_phone_number_index
    on partner_users (phone_number);

create index partner_users_partner_id_index
    on partner_users (partner_id);

create index partner_users_shop_id_index
    on partner_users (shop_id);

create table subscriptions
(
    id         serial
        primary key,
    name       varchar(255)     not null,
    days       integer          not null,
    price      double precision not null,
    currency   varchar(50)      not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp
);

alter table subscriptions
    owner to qahvazor;

create index subscriptions_deleted_at_index
    on subscriptions (deleted_at);

create table subscription_features
(
    id              serial
        primary key,
    subscription_id bigint not null,
    feature         varchar(255),
    created_at      timestamp default CURRENT_TIMESTAMP,
    updated_at      timestamp default CURRENT_TIMESTAMP
);

alter table subscription_features
    owner to qahvazor;

create index subscription_features_subscription_id_index
    on subscription_features (subscription_id);

create table user_orders
(
    id              serial
        primary key,
    partner_id      bigint not null,
    shop_id         bigint not null,
    user_id         bigint not null,
    user_partner_id bigint not null,
    drink_id        bigint,
    created_at      timestamp default CURRENT_TIMESTAMP,
    updated_at      timestamp default CURRENT_TIMESTAMP,
    deleted_at      timestamp
);

alter table user_orders
    owner to qahvazor;

create index user_orders_partner_id_index
    on user_orders (partner_id);

create index user_orders_shop_id_index
    on user_orders (shop_id);

create index user_orders_user_id_index
    on user_orders (user_id);

create index user_orders_user_partner_id_index
    on user_orders (user_partner_id);

create index user_orders_drink_id_index
    on user_orders (drink_id);

create index user_orders_deleted_at_index
    on user_orders (deleted_at);

create table shops
(
    id            serial
        primary key,
    image_id      bigint,
    partner_id    bigint           not null,
    name          varchar(100)     not null,
    location_lat  double precision not null,
    location_long double precision not null,
    created_at    timestamp default CURRENT_TIMESTAMP,
    updated_at    timestamp default CURRENT_TIMESTAMP,
    deleted_at    timestamp
);

alter table shops
    owner to qahvazor;

create index shops_partner_id_index
    on shops (partner_id);

create index shops_deleted_at_index
    on shops (deleted_at);

create index shops_location_lat_index
    on shops (location_lat);

create index shops_location_long_index
    on shops (location_long);

create table subscription_modules
(
    id              serial
        primary key,
    subscription_id bigint not null,
    module_id       bigint not null
);

alter table subscription_modules
    owner to qahvazor;

create index subscription_modules_subscription_id_index
    on subscription_modules (subscription_id);

create index subscription_modules_module_id_index
    on subscription_modules (module_id);

create table modules
(
    id         serial
        primary key,
    name       varchar(100) not null,
    colour     varchar(50)  not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP
);

alter table modules
    owner to qahvazor;

create table shop_modules
(
    id         serial
        primary key,
    partner_id bigint not null,
    shop_id    bigint not null,
    module_id  bigint not null
);

alter table shop_modules
    owner to qahvazor;

create index shop_modules_partner_id_index
    on shop_modules (partner_id);

create index shop_modules_shop_id_index
    on shop_modules (shop_id);

create index shop_modules_module_id_index
    on shop_modules (module_id);

create table transactions
(
    id               serial
        primary key,
    phone_number     varchar             not null,
    transaction_type varchar(20)         not null,
    transaction_id   varchar(255)        not null,
    amount           integer   default 0 not null,
    created_at       timestamp default CURRENT_TIMESTAMP,
    updated_at       timestamp default CURRENT_TIMESTAMP,
    deleted_at       timestamp,
    currency         varchar(20),
    payment_type     varchar(20),
    description      text
);

alter table transactions
    owner to qahvazor;

create table payme_transactions
(
    id             serial
        primary key,
    transaction_id varchar(255)     not null
        unique,
    state          integer          not null,
    reason         integer,
    phone_number   varchar(255)     not null,
    amount         integer          not null,
    created_at     bigint default 0 not null,
    perform_at     bigint default 0,
    cancel_at      bigint default 0
);

alter table payme_transactions
    owner to qahvazor;

create table user_subscriptions
(
    id              serial
        primary key,
    user_id         integer not null,
    subscription_id integer not null,
    start_date      bigint  not null,
    end_date        bigint  not null,
    created_at      timestamp default CURRENT_TIMESTAMP,
    updated_at      timestamp default CURRENT_TIMESTAMP,
    deleted_at      timestamp
);

alter table user_subscriptions
    owner to qahvazor;

