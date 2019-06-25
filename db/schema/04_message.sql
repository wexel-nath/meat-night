
CREATE TABLE IF NOT EXISTS message (
  message_id        SERIAL      PRIMARY KEY NOT NULL,
  message_event     VARCHAR(20) NOT NULL,
  message_timestamp TIMESTAMP   NOT NULL DEFAULT NOW(),
  mateo_id          INTEGER     NOT NULL REFERENCES mateo(mateo_id)
);
