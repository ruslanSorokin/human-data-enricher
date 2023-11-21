BEGIN;

CREATE TABLE
  persons (
    "id" uuid PRIMARY KEY,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "deleted_at" timestamptz,
    "name" varchar NOT NULL,
    "surname" varchar NOT NULL,
    "middle_name" varchar,
    "nationality" varchar NOT NULL,
    "gender" varchar NOT NULL,
    "age" smallint NOT NULL
  );


COMMIT;
