
CREATE TABLE IF NOT EXISTS guest(
  guest_id  SERIAL  PRIMARY KEY NOT NULL,
  dinner_id INTEGER NOT NULL REFERENCES dinner(dinner_id),
  mateo_id  INTEGER NOT NULL REFERENCES mateo(mateo_id)
);
