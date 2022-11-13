
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email text NOT NULL UNIQUE,
    pw_bcrypt text NOT NULL, -- already includes salt
    full_name text,
    avatar_url text,
    is_admin BOOL DEFAULT FALSE
);


CREATE TABLE IF NOT EXISTS organizations (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    long_name text NOT NULL,
    short_name text,
    avatar_url text,
    contact_email text NOT NULL
);

CREATE TABLE IF NOT EXISTS organizations_have_users (
	organization_id UUID NOT NULL,
	user_id UUID NOT NULL,
	FOREIGN KEY (organization_id)
		REFERENCES organizations(id),
	FOREIGN KEY (user_id)
		REFERENCES users(id)
);

ALTER TABLE emergency_rooms
ADD COLUMN IF NOT EXISTS
owned_by UUID NOT NULL;


ALTER TABLE emergency_rooms
ADD CONSTRAINT emergency_rooms_owned_by_foreign_key_constraint
FOREIGN KEY (owned_by) REFERENCES organizations(id);
