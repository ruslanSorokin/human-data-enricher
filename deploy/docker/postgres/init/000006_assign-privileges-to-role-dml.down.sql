-- User: postgres
-- Database: human-data-enricher
--
-- psql -U postgres -d 'human-data-enricher' -a -f 000006_assign-privileges-to-role-dml.down.sql
BEGIN;

REVOKE USAGE, SELECT
ON ALL SEQUENCES
IN SCHEMA "human-data-enricher"
FROM dml_global_role;

REVOKE USAGE
ON SCHEMA "human-data-enricher"
FROM dml_global_role;

COMMIT;
