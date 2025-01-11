-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    phone_number VARCHAR(255) NOT NULL UNIQUE,
    mobile_provider VARCHAR(100) NOT NULL,
    refresh_token VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

create index users_deleted_at_index
    on users (deleted_at);

create index users_phone_number_index
    on users (phone_number);

create index users_refresh_token_index
    on users (refresh_token);

-- Create images table
CREATE TABLE IF NOT EXISTS images (
    id SERIAL PRIMARY KEY,
    path VARCHAR(255) NOT NULL,
    filename VARCHAR(255) NOT NULL,
    ext VARCHAR(255) NOT NULL,
    hash_uid VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

create index images_deleted_at_index
    on images (deleted_at);

-- Create partners table
CREATE TABLE IF NOT EXISTS partners (
    id SERIAL PRIMARY KEY,
    logo_id BIGINT,
    name VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

create index partners_deleted_at_index
    on partners (deleted_at);

-- Create partner_attributes table
CREATE TABLE IF NOT EXISTS partner_attributes (
    id SERIAL PRIMARY KEY,
    partner_id BIGINT,
    attribute_key VARCHAR(100) NOT NULL,
    attribute_value VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

create index partner_attributes_partner_id_index
    on partner_attributes (partner_id);

create index partner_attributes_attribute_key_index
    on partner_attributes (attribute_key);

create index partner_attributes_deleted_at_index
    on partner_attributes (deleted_at);

-- Create shops table
CREATE TABLE IF NOT EXISTS shops (
    id SERIAL PRIMARY KEY,
    partner_id BIGINT,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

create index shops_partner_id_index
    on shops (partner_id);

create index shops_deleted_at_index
    on shops (deleted_at);

-- Create shop_attributes table
CREATE TABLE IF NOT EXISTS shop_attributes (
    id SERIAL PRIMARY KEY,
    shop_id BIGINT,
    attribute_key VARCHAR(100) NOT NULL,
    attribute_value VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

create index shop_attributes_partner_id_index
    on shop_attributes (shop_id);

create index shop_attributes_attribute_key_index
    on shop_attributes (attribute_key);

create index shop_attributes_deleted_at_index
    on shop_attributes (deleted_at);

-- Create shop_hours table
CREATE TABLE IF NOT EXISTS shop_hours (
    id SERIAL PRIMARY KEY,
    shop_id BIGINT,
    week_day VARCHAR(50) NOT NULL,
    open_at VARCHAR(10) NOT NULL,
    close_at VARCHAR(10) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

create index shop_hours_shop_id_index
    on shop_hours (shop_id);

create index shop_hours_week_day_index
    on shop_hours (week_day);

create index shop_hours_deleted_at_index
    on shop_hours (deleted_at);

-- Create shop_pictures table
CREATE TABLE IF NOT EXISTS shop_pictures (
    id SERIAL PRIMARY KEY,
    shop_id BIGINT,
    image_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

create index shop_pictures_shop_id_index
    on shop_pictures (shop_id);

create index shop_pictures_deleted_at_index
    on shop_pictures (deleted_at);

-- Create drinks table
CREATE TABLE IF NOT EXISTS drinks (
    id SERIAL PRIMARY KEY,
    image_id BIGINT NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

create index drinks_name_index
    on drinks (name);

create index drinks_delete_at_index
    on drinks (deleted_at);

-- Create partner_drinks table
CREATE TABLE IF NOT EXISTS partner_drinks (
    id SERIAL PRIMARY KEY,
    partner_id BIGINT NOT NULL,
    drink_id BIGINT NOT NULL
);

create index partner_drinks_partner_id_index
    on partner_drinks (partner_id);

create index partner_drinks_drink_id_index
    on partner_drinks (drink_id);

-- Create subscriptions table
CREATE TABLE IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    days INT NOT NULL,
    price   INT NOT NULL,
    currency VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

create index subscriptions_deleted_at_index
    on subscriptions (deleted_at);

-- Create subscription_shops table
CREATE TABLE IF NOT EXISTS subscription_shops (
    id SERIAL PRIMARY KEY,
    partner_id BIGINT NOT NULL,
    shop_id BIGINT NOT NULL,
    subscription_id BIGINT NOT NULL
);

create index subscription_shops_partner_id_index
    on subscription_shops (partner_id);

create index subscription_shops_shop_id_index
    on subscription_shops (shop_id);

create index subscription_shops_subscription_id_index
    on subscription_shops (subscription_id);

-- Create partner_users table
CREATE TABLE IF NOT EXISTS partner_users (
    id SERIAL PRIMARY KEY,
    partner_id BIGINT,
    shop_id BIGINT,
    name VARCHAR(100),
    phone_number VARCHAR(255) NOT NULL,
    mobile_provider VARCHAR(100) NOT NULL,
    refresh_token VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

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