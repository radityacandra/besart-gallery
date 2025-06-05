CREATE TABLE IF NOT EXISTS products (
  id VARCHAR PRIMARY KEY,
  name VARCHAR NOT NULL,
  image VARCHAR NOT NULL,
  description TEXT NOT NULL,
  medium VARCHAR NOT NULL,
  dimension VARCHAR NOT NULL,
  rating NUMERIC,
  original_price int8 NOT NULL,
  discounted_price int8 NULL,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR
);

CREATE TABLE IF NOT EXISTS orders (
  id VARCHAR PRIMARY KEY,
  user_id VARCHAR NOT NULL,
  status VARCHAR NOT NULL,
  receiver_full_name VARCHAR NOT NULL,
  receiver_address VARCHAR NOT NULL,
  receiver_phone_no VARCHAR NOT NULL,
  shipping_notes VARCHAR,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR
);

CREATE TABLE IF NOT EXISTS order_items (
  id VARCHAR PRIMARY KEY,
  order_id VARCHAR NOT NULL,
  product_id VARCHAR NOT NULL,
  quantity INTEGER NOT NULL,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR
);
