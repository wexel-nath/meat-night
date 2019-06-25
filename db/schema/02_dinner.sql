
CREATE TABLE IF NOT EXISTS dinner (
  dinner_id SERIAL      PRIMARY KEY NOT NULL,
  date      DATE        NOT NULL,
  mateo_id  INTEGER     NOT NULL REFERENCES mateo(mateo_id),
  venue     VARCHAR(32) NOT NULL
);
