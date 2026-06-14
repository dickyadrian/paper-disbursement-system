CREATE TYPE disbursement_status AS ENUM ('pending', 'completed', 'failed');

CREATE TABLE disbursements (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    amount BIGINT NOT NULL CHECK (amount > 0),
    status disbursement_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_disbursements_user_id ON disbursements(user_id);
CREATE INDEX idx_disbursements_status ON disbursements(status);
