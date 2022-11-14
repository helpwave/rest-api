ALTER TABLE emergency_rooms
DROP CONSTRAINT IF EXISTS emergency_rooms_organization_id_foreign_key_constraint;

ALTER TABLE emergency_rooms
DROP COLUMN IF EXISTS organization_id;

DROP TABLE IF EXISTS organizations_have_users;
DROP TABLE IF EXISTS organizations;
DROP TABLE IF EXISTS global_roles;
DROP TYPE IF EXISTS global_role;
DROP TABLE IF EXISTS users;
