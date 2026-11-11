CREATE TABLE project_members (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
role TEXT NOT NULL DEFAULT 'member'
	CHECK (role IN ('admin', 'member', 'viewer')),
joined_at TIMESTAMPTZ NOT NULL DEFAULT now(),
created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

CONSTRAINT project_member_project_user_unique
UNIQUE (project_id, user_id)
);

CREATE INDEX idx_project_members_project_id ON project_members(project_id);
CREATE INDEX idx_project_members_users_id ON project_members(user_id);

CREATE TRIGGER project_members_set_updated_at
BEFORE UPDATE ON project_members
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();