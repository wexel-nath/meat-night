
CREATE TABLE IF NOT EXISTS mateo (
  mateo_id   SERIAL      PRIMARY KEY NOT NULL,
  first_name VARCHAR(12) NOT NULL,
  last_name  VARCHAR(12) NOT NULL,
  email      VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS message (
  message_id        SERIAL      PRIMARY KEY NOT NULL,
  message_event     VARCHAR(20) NOT NULL,
  message_timestamp TIMESTAMP   NOT NULL DEFAULT NOW(),
  mateo_id          INTEGER     NOT NULL REFERENCES mateo(mateo_id)
);
