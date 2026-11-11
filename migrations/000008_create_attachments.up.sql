CREATE TABLE attachments(
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
issue_id UUID NOT NULL REFERENCES issues(id) ON DELETE CASCADE,
comment_id UUID REFERENCES comments(id) ON DELETE CASCADE,
filename TEXT NOT NULL,
mime_type TEXT NOT NULL,
file_size BIGINT NOT NULL 
	CHECK(file_size > 0),
file_content BYTEA NOT NULL,
uploaded_by UUID NOT NULL REFERENCES users(id),
created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_attachments_issue_id ON attachments(issue_id);
CREATE INDEX idx_attachments_comment_id ON attachments(comment_id);
CREATE INDEX idx_attachments_uploaded_by ON attachments(uploaded_by);

CREATE TRIGGER attachments_set_updated_at
BEFORE UPDATE ON attachments
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();