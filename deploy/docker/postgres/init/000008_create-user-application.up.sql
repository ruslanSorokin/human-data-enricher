-- User: postgres
-- Database: human-data-enricher
--
-- psql -U postgres -d 'human-data-enricher' -a -f 000008_create-user-application.up.sql
BEGIN;

CREATE USER "application"
WITH ENCRYPTED PASSWORD 'PASS';

GRANT dml_global_role
TO "application";

SET search_path TO 'human-data-enricher';

COMMIT;
