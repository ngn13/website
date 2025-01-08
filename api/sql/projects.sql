CREATE TABLE IF NOT EXISTS projects(
  name    TEXT NOT NULL UNIQUE,
  desc    TEXT NOT NULL,
  url     TEXT NOT NULL,
  license TEXT
);
