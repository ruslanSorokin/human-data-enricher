-- User: postgres
-- Database: human-data-enricher
--
-- psql -U postgres -d 'human-data-enricher' -a -f 000002_create-schema.down.sql
BEGIN;

DROP SCHEMA "human-data-enricher" CASCADE;

COMMIT;
