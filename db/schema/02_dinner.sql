
CREATE TABLE IF NOT EXISTS dinner (
  id      SERIAL      PRIMARY KEY NOT NULL,
  date    DATE        NOT NULL DEFAULT NOW(),
  host_id INTEGER     NOT NULL REFERENCES mateo(id),
  venue   VARCHAR(32) NOT NULL
);
