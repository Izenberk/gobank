CREATE TABLE accounts (
  id BIGSERIAL PRIMARY KEY,
  fullname TEXT,
  balance BIGINT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);