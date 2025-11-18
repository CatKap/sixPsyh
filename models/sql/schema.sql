PRAGMA foreign_keys = ON;
CREATE TABLE cathegory (
  id  INTEGER PRIMARY KEY,
  name TEXT    NOT NULL
);

CREATE TABLE event (
  id   INTEGER PRIMARY KEY,
  name TEXT    NOT NULL,
  description TEXT NOT NULL,
  cathegory REFERENCES cathegory(id),
  time INTEGER DEFAULT NOT NULL 
);
