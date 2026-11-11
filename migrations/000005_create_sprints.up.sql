CREATE TABLE sprints (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
name TEXT NOT NULL,
goal TEXT,
start_date DATE,
end_date DATE, 
status TEXT NOT NULL DEFAULT 'planned'
	CHECK (status IN ('planned', 'active', 'completed')),
started_at TIMESTAMPTZ,
ended_at TIMESTAMPTZ,
created_by UUID NOT NULL REFERENCES users(id),
created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

CONSTRAINT sprint_dates_valid 
	CHECK (
		start_date IS NULL
		OR end_date IS NULL
		OR end_date > start_date
	)
);

CREATE INDEX idx_sprints_project_id ON sprints(project_id);
CREATE INDEX idx_sprints_status ON sprints(status);

CREATE UNIQUE INDEX one_active_sprint_per_project ON sprints(project_id) WHERE status = 'active';

CREATE TRIGGER sprints_set_updated_at
BEFORE UPDATE ON sprints
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();