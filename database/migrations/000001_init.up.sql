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

-- Create company_socials table
CREATE TABLE IF NOT EXISTS company_socials (
    id SERIAL PRIMARY KEY,
    company_id BIGINT,
    platform VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_company
        FOREIGN KEY(company_id) 
        REFERENCES company(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- Create shops table
CREATE TABLE IF NOT EXISTS shops (
    id SERIAL PRIMARY KEY,
    image_id BIGINT,
    company_id BIGINT,
    name VARCHAR(255),
    location TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT fk_image
        FOREIGN KEY(image_id) 
        REFERENCES images(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_company
        FOREIGN KEY(company_id) 
        REFERENCES company(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- Create shop_phones table
CREATE TABLE IF NOT EXISTS shop_phones (
    id SERIAL PRIMARY KEY,
    shop_id BIGINT,
    phone_number VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_shop
        FOREIGN KEY(shop_id) 
        REFERENCES shops(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- Create shop_worktimes table
CREATE TABLE IF NOT EXISTS shop_worktimes (
    id SERIAL PRIMARY KEY,
    shop_id BIGINT,
    day_range VARCHAR(255) NOT NULL,
    opening_time VARCHAR(255) NOT NULL,
    closing_time VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_shop
        FOREIGN KEY(shop_id) 
        REFERENCES shops(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- Create coffee table
CREATE TABLE IF NOT EXISTS coffee (
    id SERIAL PRIMARY KEY,
    image_id BIGINT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_image
        FOREIGN KEY(image_id) 
        REFERENCES images(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- Create shop_coffees table
CREATE TABLE IF NOT EXISTS shop_coffees (
    id SERIAL PRIMARY KEY,
    shop_id BIGINT,
    coffee_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_shop
        FOREIGN KEY(shop_id) 
        REFERENCES shops(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_coffee
        FOREIGN KEY(coffee_id) 
        REFERENCES coffee(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- Create subscription table
CREATE TABLE IF NOT EXISTS subscription (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    coffee_limit INT NOT NULL,
    interval INT NOT NULL,
    period   INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
