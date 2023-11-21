-- User: postgres
-- Database: human-data-enricher
--
-- psql -U postgres -d 'human-data-enricher' -a -f 000004_assign-privileges-to-role-ddl.down.sql
BEGIN;

REVOKE ALL PRIVILEGES
ON ALL SEQUENCES
IN SCHEMA "human-data-enricher"
FROM ddl_global_role;

REVOKE USAGE, CREATE
ON SCHEMA "human-data-enricher"
FROM ddl_global_role;

COMMIT;
