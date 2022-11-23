CREATE TABLE IF NOT EXISTS organizations (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    long_name text NOT NULL,
    short_name text,
    avatar_url text,
    contact_email text NOT NULL
);


CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email text NOT NULL UNIQUE,
    pw_bcrypt text NOT NULL, -- already includes salt
    full_name text,
    avatar_url text,
    organization_id uuid,
    FOREIGN KEY (organization_id)
    	REFERENCES organizations(id)
);

CREATE TYPE global_role AS ENUM ('admin');

CREATE TABLE IF NOT EXISTS global_roles (
	user_id uuid NOT NULL,
	role global_role NOT NULL,
	FOREIGN KEY (user_id)
		REFERENCES users(id),
	UNIQUE (user_id, role)
);

ALTER TABLE emergency_rooms
ADD COLUMN IF NOT EXISTS
organization_id UUID NOT NULL;


ALTER TABLE emergency_rooms
ADD CONSTRAINT emergency_rooms_organization_id_foreign_key_constraint
FOREIGN KEY (organization_id) REFERENCES organizations(id);
