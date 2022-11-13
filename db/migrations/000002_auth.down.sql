ALTER TABLE emergency_rooms
DROP CONSTRAINT IF EXISTS emergency_rooms_owned_by_foreign_key_constraint;

ALTER TABLE emergency_rooms
DROP COLUMN IF EXISTS owned_by;

DROP TABLE IF EXISTS organizations_have_users;
DROP TABLE IF EXISTS organizations;
DROP TABLE IF EXISTS users;
