-- User: postgres
-- Database: postgres
--
-- psql -U postgres -d postgres -a -f 000005_create-role-dml.down.sql
BEGIN;

REVOKE CONNECT
ON DATABASE "human-data-enricher"
FROM dml_global_role;

REVOKE TEMPORARY
ON DATABASE "human-data-enricher"
FROM dml_global_role;

DROP ROLE dml_global_role;

COMMIT;
