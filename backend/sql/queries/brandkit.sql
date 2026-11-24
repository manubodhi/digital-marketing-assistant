-- name: UpsertBrandKit :one
INSERT INTO brand_kits (
    id, workspace_id, brand_name, industry, colors, logo_url, brand_voice, fonts, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, NOW()
) ON CONFLICT (workspace_id) DO UPDATE SET
    brand_name = EXCLUDED.brand_name,
    industry = EXCLUDED.industry,
    colors = EXCLUDED.colors,
    logo_url = COALESCE(EXCLUDED.logo_url, brand_kits.logo_url),
    brand_voice = EXCLUDED.brand_voice,
    fonts = EXCLUDED.fonts,
    updated_at = NOW()
RETURNING *;

-- name: GetBrandKitByWorkspaceID :one
SELECT * FROM brand_kits
WHERE workspace_id = $1 LIMIT 1;
