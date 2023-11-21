-- User: postgres
-- Database: postgres
--
-- psql -U postgres -d postgres -a -f 000003_create-role-ddl.down.sql
BEGIN;

REVOKE CONNECT
ON DATABASE "human-data-enricher"
FROM ddl_global_role;

REVOKE TEMPORARY
ON DATABASE "human-data-enricher"
FROM ddl_global_role;

DROP ROLE ddl_global_role;

COMMIT;
