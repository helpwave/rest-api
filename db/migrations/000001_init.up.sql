CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS emergency_rooms (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    location point,
    displayable_address text NOT NULL,
    is_open BOOL NOT NULL DEFAULT true,
    utilization smallint NOT NULL DEFAULT 1,
    CHECK (utilization >= 1 AND utilization <=5)
);
