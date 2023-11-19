-- User: postgres
-- Database: human-data-enricher
--
-- psql -U postgres -d 'human-data-enricher' -a -f 000007_create-user-migrator.up.sql
BEGIN;

CREATE USER "migrator"
WITH ENCRYPTED PASSWORD 'PASS';

GRANT ddl_global_role
TO "migrator";

COMMIT;
