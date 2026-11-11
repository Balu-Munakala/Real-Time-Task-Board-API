CREATE TABLE users (
id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
username TEXT NOT NULL UNIQUE,
name TEXT NOT NULL,
email TEXT NOT NULL UNIQUE,
role TEXT NOT NULL
	CHECK (role IN ('reporting_lead', 'developer', 'tester', 'business_analyst')),
password_hash TEXT NOT NULL,
created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TRIGGER users_set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();