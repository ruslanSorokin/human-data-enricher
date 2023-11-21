-- User: postgres
-- Database: human-data-enricher
--
-- psql -U postgres -d 'human-data-enricher' -a -f 000006_assign-privileges-to-role-dml.up.sql
BEGIN;

GRANT USAGE
ON SCHEMA "human-data-enricher"
TO dml_global_role;

GRANT USAGE, SELECT
ON ALL SEQUENCES
IN SCHEMA "human-data-enricher"
TO dml_global_role;

COMMIT;
