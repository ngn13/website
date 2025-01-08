CREATE TABLE IF NOT EXISTS services(
  name       TEXT NOT NULL UNIQUE,
  desc       TEXT NOT NULL,
  check_time INTEGER NOT NULL,
  check_res  INTEGER NOT NULL,
  check_url  TEXT NOT NULL,
  clear      TEXT,
  onion      TEXT,
  i2p        TEXT
);
