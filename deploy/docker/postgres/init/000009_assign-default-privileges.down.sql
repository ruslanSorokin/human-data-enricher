-- User: migrator
-- Database: human-data-enricher
--
-- psql -U migrator -d 'human-data-enricher' -a -f 000009_assign-default-privileges.down.sql
BEGIN;

ALTER DEFAULT PRIVILEGES
IN SCHEMA "human-data-enricher"
REVOKE SELECT, INSERT, UPDATE, DELETE
ON TABLES
FROM dml_global_role;

GRANT CREATE
ON SCHEMA public
TO PUBLIC;

COMMIT;
