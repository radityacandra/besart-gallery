CREATE TABLE IF NOT EXISTS users (
  id VARCHAR PRIMARY KEY,
  username VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  created_by VARCHAR NOT NULL,
  updated_at int8,
  updated_by VARCHAR
);

CREATE TABLE IF NOT EXISTS applied_positions (
  id VARCHAR PRIMARY KEY,
  user_id VARCHAR NOT NULL,
  job_id VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  created_by VARCHAR NOT NULL,
  updated_at int8,
  updated_by VARCHAR
);
