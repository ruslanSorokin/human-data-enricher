-- User: postgres
-- Database: human-data-enricher
--
-- psql -U postgres -d 'human-data-enricher' -a -f 000007_create-user-migrator.down.sql
BEGIN;

DROP USER "migrator";

COMMIT;
