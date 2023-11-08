BEGIN;

-- 1. Drop trigger before because them stuck with trigger
DROP TRIGGER IF EXISTS set_updated_at_timestamp_users_table ON "users";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_oauth_table  ON "oauth";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_products_table  ON "products";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_product_images_table  ON "product_images";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_orders_table  ON "orders";


-- 2. Drop function
DROP FUNCTION IF EXISTS set_updated_at_column();

-- 3. Drop sequence function
DROP SEQUENCE IF EXISTS users_id_seq;
DROP SEQUENCE IF EXISTS products_id_seq;
DROP SEQUENCE IF EXISTS orders_id_seq;

-- 4. Drop type
DROP TYPE IF EXISTS "order_status";

-- 5. Drop table from 
DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "oauth" CASCADE;
DROP TABLE IF EXISTS "roles" CASCADE;
DROP TABLE IF EXISTS "products" CASCADE;
DROP TABLE IF EXISTS "categories" CASCADE;
DROP TABLE IF EXISTS "products_categories" CASCADE;
DROP TABLE IF EXISTS "imgaes" CASCADE;
DROP TABLE IF EXISTS "orders" CASCADE;
DROP TABLE IF EXISTS "products_orders" CASCADE;

COMMIT;