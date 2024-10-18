-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- Set the search path for the schemas
SET search_path TO public, auth;

-- Create schemas if they don't exist
CREATE SCHEMA IF NOT EXISTS auth;
CREATE SCHEMA IF NOT EXISTS public;

-- Define a custom ENUM type for document roles
CREATE TYPE document_role AS ENUM ('viewer', 'editor', 'contributor');

-- Users table definition
CREATE TABLE IF NOT EXISTS public.users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    oauth_id VARCHAR(100) UNIQUE NOT NULL,
    image_url VARCHAR(255) DEFAULT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    CHECK (email <> '')  -- Ensure email is not empty
);

-- Sessions table definition
CREATE TABLE IF NOT EXISTS auth.sessions (
    session_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    online BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,  
    FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE
);

-- Documents table definition
CREATE TABLE IF NOT EXISTS public.documents (
    document_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    document_name VARCHAR(255) NOT NULL,
    mongo_id VARCHAR(24) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    CHECK (document_name <> '')  -- Ensure document name is not empty
);

-- Document ownership table definition
CREATE TABLE IF NOT EXISTS public.document_ownerships (
    user_id UUID NOT NULL,  
    document_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (user_id, document_id),
    FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (document_id) REFERENCES public.documents(document_id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Document contributions table definition
CREATE TABLE IF NOT EXISTS public.document_contributions (
    user_id UUID NOT NULL,
    document_id UUID NOT NULL,
    role document_role DEFAULT 'viewer', 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (user_id, document_id),
    FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (document_id) REFERENCES public.documents(document_id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Blocklist table definition
CREATE TABLE IF NOT EXISTS public.blocklist (
    blocker_id UUID NOT NULL,
    blocked_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (blocker_id, blocked_id),
    FOREIGN KEY (blocker_id) REFERENCES public.users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (blocked_id) REFERENCES public.users(user_id) ON DELETE CASCADE
);

-- Document access table definition
CREATE TABLE IF NOT EXISTS public.document_access (
    access_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    document_id UUID NOT NULL,
    public_view BOOLEAN NOT NULL DEFAULT TRUE,
    public_edit BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (document_id) REFERENCES public.documents(document_id) ON DELETE CASCADE
);

-- Create indexes for improved query performance
CREATE INDEX IF NOT EXISTS idx_document_ownerships_user ON public.document_ownerships(user_id);
CREATE INDEX IF NOT EXISTS idx_document_ownerships_doc ON public.document_ownerships(document_id);
CREATE INDEX IF NOT EXISTS idx_document_contributions_user ON public.document_contributions(user_id);
CREATE INDEX IF NOT EXISTS idx_document_contributions_doc ON public.document_contributions(document_id);
CREATE INDEX IF NOT EXISTS idx_document_access_id ON public.document_access(access_id);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON public.users(deleted_at);    
CREATE INDEX IF NOT EXISTS idx_documents_deleted_at ON public.documents(deleted_at); 
CREATE INDEX IF NOT EXISTS idx_document_ownerships_deleted_at ON public.document_ownerships(deleted_at);
CREATE INDEX IF NOT EXISTS idx_document_contributions_deleted_at ON public.document_contributions(deleted_at);
CREATE INDEX IF NOT EXISTS idx_blocklist_deleted_at ON public.blocklist(deleted_at);
CREATE INDEX IF NOT EXISTS idx_document_access_deleted_at ON public.document_access(deleted_at);
CREATE INDEX IF NOT EXISTS idx_sessions_deleted_at ON auth.sessions(deleted_at);
CREATE INDEX IF NOT EXISTS idx_sessions_token ON auth.sessions(token);
CREATE INDEX IF NOT EXISTS idx_users_user_id ON public.users(user_id)


-- Function to update the updated_at timestamp
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- Create triggers to automatically update the updated_at column
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON public.users
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_documents_updated_at
BEFORE UPDATE ON public.documents
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_document_ownerships_updated_at
BEFORE UPDATE ON public.document_ownerships
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_document_contributions_updated_at
BEFORE UPDATE ON public.document_contributions
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

-- Drop tables and schema if they exist
DROP TABLE IF EXISTS public.document_contributions;
DROP TABLE IF EXISTS public.document_ownerships;  -- Fixed the table name
DROP TABLE IF EXISTS auth.sessions;  
DROP TABLE IF EXISTS public.documents;
DROP TABLE IF EXISTS public.users;
DROP TABLE IF EXISTS public.blocklist;
DROP TABLE IF EXISTS public.document_access;
DROP SCHEMA IF EXISTS auth;

-- Drop triggers and functions if they exist
DROP TRIGGER IF EXISTS update_documents_updated_at ON public.documents;
DROP TRIGGER IF EXISTS update_users_updated_at ON public.users;
DROP TRIGGER IF EXISTS update_document_ownerships_updated_at ON public.document_ownerships;
DROP TRIGGER IF EXISTS update_document_contributions_updated_at ON public.document_contributions;
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop custom types and indexes if they exist
DROP TYPE IF EXISTS document_role;
DROP INDEX IF EXISTS idx_document_ownerships_user;
DROP INDEX IF EXISTS idx_document_ownerships_doc;
DROP INDEX IF EXISTS idx_document_contributions_user;
DROP INDEX IF EXISTS idx_document_contributions_doc;
DROP TABLE IF EXISTS public.document_access;
DROP INDEX IF EXISTS idx_document_access_id;
DROP INDEX IF EXISTS idx_users_deleted_at;
DROP INDEX IF EXISTS idx_documents_deleted_at;
DROP INDEX IF EXISTS idx_document_ownerships_deleted_at;
DROP INDEX IF EXISTS idx_document_contributions_deleted_at;
DROP INDEX IF EXISTS idx_blocklist_deleted_at;
DROP INDEX IF EXISTS idx_document_access_deleted_at;
DROP INDEX IF EXISTS idx_sessions_deleted_at;
