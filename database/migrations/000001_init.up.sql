-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    phone_number VARCHAR(255) NOT NULL,
    mobile_provider VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create images table
CREATE TABLE IF NOT EXISTS images (
    id SERIAL PRIMARY KEY,
    path VARCHAR(255) NOT NULL,
    filename VARCHAR(255) NOT NULL,
    ext VARCHAR(255) NOT NULL,
    hash_uid VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create company table
CREATE TABLE IF NOT EXISTS company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    image_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_image
        FOREIGN KEY(image_id) 
        REFERENCES images(id)
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
