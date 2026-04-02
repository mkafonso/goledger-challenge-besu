-- Create or update storage value (singleton row)
-- Parameters: value
-- Returns: Updated storage row
-- name: SetStorage :one
INSERT INTO storage (id, value, updated_at)
VALUES (1, $1, NOW())
ON CONFLICT (id)
DO UPDATE SET
    value = EXCLUDED.value,
    updated_at = NOW()
RETURNING *;


-- Get current storage value
-- Parameters: none
-- Returns: Single storage row
-- name: GetStorage :one
SELECT *
FROM storage
WHERE id = 1
LIMIT 1;
