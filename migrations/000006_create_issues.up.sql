CREATE TABLE issues(
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
sprint_id UUID REFERENCES sprints(id) ON DELETE SET NULL,
parent_id UUID REFERENCES issues(id) ON DELETE CASCADE,
issue_type TEXT NOT NULL
	CHECK(issue_type IN('story', 'bug', 'feature', 'subtask')),
title TEXT NOT NULL,
description TEXT,
status TEXT NOT NULL DEFAULT 'todo'
	CHECK(status IN('todo', 'in_progress', 'qa', 'done')),
priority TEXT 
	CHECK(priority IN('low', 'medium', 'high', 'urgent')),
creator_id UUID NOT NULL REFERENCES users(id),
reporter_id UUID NOT NULL REFERENCES users(id),
assignee_id UUID REFERENCES users(id),
due_date DATE,
created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

CONSTRAINT subtask_must_have_parent
	CHECK(
		issue_type <> 'subtask'
		OR parent_id IS NOT NULL
	),
CONSTRAINT parent_cannot_be_subtask
	CHECK(
		parent_id IS NULL
		OR issue_type <> 'subtask'
	)
);

CREATE INDEX idx_issues_project_id ON issues(project_id);
CREATE INDEX idx_issues_sprint_id ON issues(sprint_id);
CREATE INDEX idx_issues_parent_id ON issues(parent_id);
CREATE INDEX idx_issues_assignee_id ON issues(assignee_id);
CREATE INDEX idx_issues_creator_id ON issues(creator_id);
CREATE INDEX idx_issues_reporter_id ON issues(reporter_id);
CREATE INDEX idx_issues_status ON issues(status);

CREATE TRIGGER issues_set_updated_at
BEFORE UPDATE ON issues
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();