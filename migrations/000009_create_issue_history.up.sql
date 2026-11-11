CREATE TABLE issue_history(
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
issue_id UUID NOT NULL REFERENCES issues(id),
actor_id UUID NOT NULL REFERENCES users(id),
action TEXT NOT NULL
	CHECK(action IN('created', 'updated', 'status_changed', 'assigned', 'unassigned', 'sprint_changed', 'comment_added', 'comment_updated', 'comment_deleted', 'attachment_added', 'attachment_deleted')),
field TEXT,
old_value TEXT,
new_value TEXT,
created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

CONSTRAINT old_new_values_different
CHECK (
    old_value IS NULL
    OR new_value IS NULL
    OR old_value <> new_value
)
);

CREATE INDEX idx_issue_history_issue_id ON issue_history(issue_id);
CREATE INDEX idx_issue_history_actor_id ON issue_history(actor_id);
CREATE INDEX idx_issue_history_created_at ON issue_history(created_at DESC);