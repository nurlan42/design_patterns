CREATE TABLE IF NOT EXISTS todo(
     id uuid PRIMARY KEY,
     public_id uuid,
     title VARCHAR(255),
     description VARCHAR(255),
     due_date timestamp,
     completed BOOLEAN default false
);