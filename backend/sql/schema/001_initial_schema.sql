-- Enable PGCrypto extension
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- ============================================================================
-- CLUSTER A: AUTHENTICATION & IDENTITY (GLOBAL SCOPE)
-- ============================================================================

-- 1. Users Table (UUIDv7 Primary Key generated in Go)
CREATE TABLE users (
    id UUID PRIMARY KEY, -- Go: uuid.NewV7()
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255), -- Argon2id hash (NULL for OAuth-only users)
    full_name VARCHAR(150) NOT NULL,
    is_email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    totp_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    totp_secret_enc BYTEA, -- AES-256-GCM encrypted TOTP secret
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX idx_users_email_lower ON users (LOWER(email));

-- 2. OAuth Third-Party Identities (Google, Facebook)
CREATE TABLE user_identities (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider VARCHAR(50) NOT NULL CHECK (provider IN ('google', 'facebook')),
    provider_user_id VARCHAR(255) NOT NULL,
    provider_email VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_provider_uid UNIQUE (provider, provider_user_id)
);
CREATE INDEX idx_user_identities_user ON user_identities(user_id);

-- 3. Verification & Password Reset Tokens
CREATE TABLE auth_verification_tokens (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash VARCHAR(64) NOT NULL, -- SHA-256 of raw random token
    token_type VARCHAR(32) NOT NULL CHECK (token_type IN ('EMAIL_VERIFICATION', 'PASSWORD_RESET')),
    expires_at TIMESTAMPTZ NOT NULL,
    consumed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_auth_verification_lookup 
ON auth_verification_tokens(token_hash, token_type) 
WHERE consumed_at IS NULL;

-- 4. Refresh Token Families (Device & Session Boundaries)
CREATE TABLE refresh_token_families (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    device_name VARCHAR(255),
    ip_address VARCHAR(45),
    user_agent TEXT,
    is_revoked BOOLEAN NOT NULL DEFAULT FALSE,
    revoked_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_refresh_token_families_user ON refresh_token_families(user_id) WHERE is_revoked = FALSE;

-- 5. Refresh Tokens (Rotation & Replay Detection - RFC 6819)
CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY,
    family_id UUID NOT NULL REFERENCES refresh_token_families(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash VARCHAR(64) NOT NULL UNIQUE, -- SHA-256 of raw secret
    parent_token_id UUID REFERENCES refresh_tokens(id) ON DELETE SET NULL,
    is_consumed BOOLEAN NOT NULL DEFAULT FALSE,
    consumed_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_refresh_tokens_lookup ON refresh_tokens(token_hash);
CREATE INDEX idx_refresh_tokens_family ON refresh_tokens(family_id);

-- 6. TOTP Backup Recovery Codes
CREATE TABLE user_totp_backup_codes (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    code_hash VARCHAR(64) NOT NULL, -- SHA-256 of one-time recovery code
    used_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_backup_codes_user ON user_totp_backup_codes(user_id) WHERE used_at IS NULL;

-- ============================================================================
-- CLUSTER B: MULTI-TENANT WORKSPACES & MEMBERSHIPS
-- ============================================================================

-- 7. Workspaces (Tenant Boundary)
CREATE TABLE workspaces (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 8. Workspace Memberships (Role-Based Access Control)
CREATE TABLE workspace_members (
    workspace_id UUID NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL DEFAULT 'member' CHECK (role IN ('owner', 'admin', 'member')),
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (workspace_id, user_id)
);
CREATE INDEX idx_workspace_members_user ON workspace_members(user_id);

-- ============================================================================
-- CLUSTER C: MARKETING & CREATIVE DOMAINS (TENANT SCOPED)
-- ============================================================================

-- 9. Brand Kits (1:1 with Workspace)
CREATE TABLE brand_kits (
    id UUID PRIMARY KEY,
    workspace_id UUID NOT NULL UNIQUE REFERENCES workspaces(id) ON DELETE CASCADE,
    brand_name VARCHAR(255) NOT NULL,
    industry VARCHAR(100) NOT NULL DEFAULT 'General',
    colors JSONB NOT NULL DEFAULT '{"primary": "#000000", "secondary": "#ffffff", "accent": "#f59e0b"}'::jsonb,
    logo_url TEXT,
    brand_voice VARCHAR(100) NOT NULL DEFAULT 'Professional',
    fonts JSONB DEFAULT '{"heading": "Inter", "body": "Inter"}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_brand_kits_workspace ON brand_kits(workspace_id);

-- 10. Connected Social & Ad Accounts (Encrypted OAuth Tokens)
CREATE TABLE social_accounts (
    id UUID PRIMARY KEY,
    workspace_id UUID NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    platform VARCHAR(50) NOT NULL CHECK (platform IN ('meta', 'instagram', 'google_ads', 'whatsapp')),
    account_id VARCHAR(255) NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    access_token_enc TEXT NOT NULL,
    refresh_token_enc TEXT,
    token_expires_at TIMESTAMPTZ,
    connected_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(workspace_id, platform, account_id)
);
CREATE INDEX idx_social_accounts_workspace ON social_accounts(workspace_id);

-- 11. Generated Creatives (Festival Posters, Banners, Copy)
CREATE TABLE creatives (
    id UUID PRIMARY KEY,
    workspace_id UUID NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    prompt TEXT NOT NULL,
    aspect_ratio VARCHAR(20) NOT NULL DEFAULT '1:1' CHECK (aspect_ratio IN ('1:1', '4:5', '9:16', '16:9')),
    image_url TEXT NOT NULL,
    copy_caption TEXT,
    tags TEXT[],
    status VARCHAR(50) NOT NULL DEFAULT 'ready' CHECK (status IN ('generating', 'ready', 'failed')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_creatives_workspace ON creatives(workspace_id);

-- 12. Scheduled Social Posts (Organic Calendar)
CREATE TABLE scheduled_posts (
    id UUID PRIMARY KEY,
    workspace_id UUID NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    creative_id UUID NOT NULL REFERENCES creatives(id) ON DELETE RESTRICT,
    platforms VARCHAR(50)[] NOT NULL,
    scheduled_for TIMESTAMPTZ NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'scheduled' CHECK (status IN ('scheduled', 'publishing', 'published', 'failed')),
    published_at TIMESTAMPTZ,
    error_log TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_scheduled_posts_workspace ON scheduled_posts(workspace_id, scheduled_for);

-- 13. Paid Ad Campaigns (Meta, Google, WhatsApp)
CREATE TABLE campaigns (
    id UUID PRIMARY KEY,
    workspace_id UUID NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    channel VARCHAR(50) NOT NULL CHECK (channel IN ('meta', 'google', 'whatsapp')),
    external_campaign_id VARCHAR(255),
    campaign_name VARCHAR(255) NOT NULL,
    daily_budget NUMERIC(12, 2) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'active', 'paused', 'completed')),
    metrics JSONB DEFAULT '{"spend": 0, "impressions": 0, "clicks": 0, "leads": 0}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_campaigns_workspace ON campaigns(workspace_id);
