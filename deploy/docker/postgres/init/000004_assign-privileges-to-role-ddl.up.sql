-- User: postgres
-- Database: human-data-enricher
--
-- psql -U postgres -d 'human-data-enricher' -a -f 000004_assign-privileges-to-role-ddl.up.sql
BEGIN;

GRANT USAGE, CREATE
ON SCHEMA "human-data-enricher"
TO ddl_global_role;

GRANT ALL PRIVILEGES
ON ALL SEQUENCES
IN SCHEMA "human-data-enricher"
TO ddl_global_role;

COMMIT;
