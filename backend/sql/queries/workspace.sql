-- name: CreateWorkspace :one
INSERT INTO workspaces (
    id, name, slug
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetWorkspaceByID :one
SELECT * FROM workspaces
WHERE id = $1 LIMIT 1;

-- name: GetWorkspaceBySlug :one
SELECT * FROM workspaces
WHERE slug = $1 LIMIT 1;

-- name: AddWorkspaceMember :one
INSERT INTO workspace_members (
    workspace_id, user_id, role
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: ListUserWorkspaces :many
SELECT w.id, w.name, w.slug, wm.role, w.created_at
FROM workspaces w
JOIN workspace_members wm ON w.id = wm.workspace_id
WHERE wm.user_id = $1
ORDER BY w.created_at DESC;

-- name: GetWorkspaceMembership :one
SELECT * FROM workspace_members
WHERE workspace_id = $1 AND user_id = $2 LIMIT 1;
