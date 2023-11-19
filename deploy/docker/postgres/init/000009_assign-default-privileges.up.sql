-- User: migrator
-- Database: human-data-enricher
--
-- psql -U migrator -d 'human-data-enricher' -a -f 000009_assign-default-privileges.up.sql
BEGIN;

REVOKE CREATE
ON SCHEMA public
FROM PUBLIC;

ALTER DEFAULT PRIVILEGES
IN SCHEMA "human-data-enricher"
GRANT SELECT, INSERT, UPDATE, DELETE
ON TABLES
TO dml_global_role;

COMMIT;
