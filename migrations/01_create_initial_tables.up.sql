CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE EXTENSION IF NOT EXISTS btree_gist;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS emails CASCADE;

CREATE TABLE emails
(
    email_id     UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    address_from VARCHAR(250) NOT NULL CHECK ( address_from <> '' ),
    address_to   VARCHAR(250) NOT NULL CHECK ( address_to <> '' ),
    subject      VARCHAR(250) NOT NULL CHECK ( subject <> '' ),
    message      TEXT         NOT NULL CHECK ( message <> '' ),
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE emails
    ADD COLUMN document_with_idx tsvector;
UPDATE emails
set document_with_idx = to_tsvector(emails.address_to || ' ' || emails.subject || ' ' || emails.message);
CREATE INDEX document_idx ON emails USING gin (document_with_idx);

CREATE FUNCTION emails_tsvector_trigger() RETURNS trigger AS
$$
begin
    new.document_with_idx := to_tsvector(new.address_to || ' ' || new.subject || ' ' || new.message);
    return new;
end
$$ LANGUAGE plpgsql;

CREATE TRIGGER tsvectorupdate
    BEFORE INSERT OR UPDATE
    ON emails
    FOR EACH ROW
EXECUTE PROCEDURE emails_tsvector_trigger();