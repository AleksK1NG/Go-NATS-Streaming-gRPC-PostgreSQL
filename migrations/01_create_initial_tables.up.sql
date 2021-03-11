CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE EXTENSION IF NOT EXISTS btree_gist;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS emails CASCADE;
DROP TABLE IF EXISTS addresses CASCADE;


CREATE TABLE emails
(
    email_id    UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    addressFrom VARCHAR(250) NOT NULL CHECK ( addressFrom <> '' ),
    addressTo   VARCHAR(250) NOT NULL CHECK ( addressTo <> '' ),
    subject     VARCHAR(250) NOT NULL CHECK ( subject <> '' ),
    message     TEXT         NOT NULL CHECK ( message <> '' ),
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE addresses
(
    address_id UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    email      VARCHAR(250) UNIQUE NOT NULL CHECK ( email <> '' ),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE emails
    ADD COLUMN document_with_idx tsvector;
UPDATE emails
set document_with_idx = to_tsvector(emails.addressto || ' ' || emails.subject || ' ' || emails.message);
CREATE INDEX document_idx ON emails USING gin (document_with_idx);

CREATE FUNCTION emails_tsvector_trigger() RETURNS trigger AS
$$
begin
    new.document_with_idx := to_tsvector(new.addressto || ' ' || new.subject || ' ' || new.message);
    return new;
end
$$ LANGUAGE plpgsql;

CREATE TRIGGER tsvectorupdate BEFORE INSERT OR UPDATE
    ON emails FOR EACH ROW EXECUTE PROCEDURE emails_tsvector_trigger();