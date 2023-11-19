-- User: postgres
-- Database: human-data-enricher
--
-- psql -U postgres -d 'human-data-enricher' -a -f 000008_create-user-application.down.sql
BEGIN;

DROP USER "application";

COMMIT;
