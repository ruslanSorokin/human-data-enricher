-- User: postgres
-- Database: postgres
--
-- psql -U postgres -d postgres -a -f 000005_create-role-dml.up.sql
BEGIN;

CREATE ROLE dml_global_role
WITH ENCRYPTED PASSWORD 'PASS';

GRANT CONNECT
ON DATABASE "human-data-enricher"
TO dml_global_role;

GRANT TEMPORARY
ON DATABASE "human-data-enricher"
TO dml_global_role;

COMMIT;
