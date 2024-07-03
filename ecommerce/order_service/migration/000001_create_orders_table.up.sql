CREATE TABLE IF NOT EXISTS orders (
    id UUID DEFAULT gen_random_uuid(),
    user_id VARCHAR(255),
    product_id VARCHAR(255),
    location TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at int DEFAULT 0
);
