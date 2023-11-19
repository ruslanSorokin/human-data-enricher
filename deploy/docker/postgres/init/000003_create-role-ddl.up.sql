-- User: postgres
-- Database: postgres
--
-- psql -U postgres -d postgres -a -f 000003_create-role-ddl.up.sql
BEGIN;

CREATE ROLE ddl_global_role
WITH ENCRYPTED PASSWORD 'PASS';

GRANT CONNECT
ON DATABASE "human-data-enricher"
TO ddl_global_role;

GRANT TEMPORARY
ON DATABASE "human-data-enricher"
TO ddl_global_role;


COMMIT;
