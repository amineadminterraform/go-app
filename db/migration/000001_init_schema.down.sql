-- Drop Foreign Keys in Reverse Order

-- request table foreign keys
ALTER TABLE "request" DROP CONSTRAINT IF EXISTS "request_environment_id_fkey";
ALTER TABLE "request" DROP CONSTRAINT IF EXISTS "request_layer_id_fkey";

-- notifications table foreign key
ALTER TABLE "notifications" DROP CONSTRAINT IF EXISTS "notifications_process_id_fkey";

-- process table foreign key
ALTER TABLE "process" DROP CONSTRAINT IF EXISTS "process_layer_id_fkey";

-- env_layer table foreign keys
ALTER TABLE "env_layer" DROP CONSTRAINT IF EXISTS "env_layer_template_id_fkey";
ALTER TABLE "env_layer" DROP CONSTRAINT IF EXISTS "env_layer_environment_id_fkey";

-- project_environment table foreign key
ALTER TABLE "project_environment" DROP CONSTRAINT IF EXISTS "project_environment_project_id_fkey";


-- Drop Indexes in Reverse Order

-- request table index
DROP INDEX IF EXISTS "request_environment_id_idx";

-- notifications table index
DROP INDEX IF EXISTS "notifications_process_id_idx";

-- process table index
DROP INDEX IF EXISTS "process_layer_id_idx";

-- env_layer table index
DROP INDEX IF EXISTS "env_layer_environment_id_idx";

-- project_environment table index
DROP INDEX IF EXISTS "project_environment_project_id_idx";


DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS request;
DROP TABLE IF EXISTS process;
DROP TABLE IF EXISTS env_layer;
DROP TABLE IF EXISTS project_environment;
DROP TABLE IF EXISTS template;
DROP TABLE IF EXISTS project;
