CREATE TABLE "project" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "git_path" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "description" TEXT
);

CREATE TABLE "project_environment" (
  "id" BIGSERIAL PRIMARY KEY,
  "git_branch" VARCHAR NOT NULL,
  "project_id" BIGINT NOT NULL,
  "description" TEXT,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "template" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "path" VARCHAR NOT NULL
);

CREATE TABLE "env_layer" (
  "id" BIGSERIAL PRIMARY KEY,
  "environment_id" BIGINT NOT NULL,
  "s3_path" VARCHAR NOT NULL,
  "template_id" BIGINT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "process" (
  "id" BIGSERIAL PRIMARY KEY,
  "name_of_argo_process" VARCHAR NOT NULL,
  "currentstep" TEXT NOT NULL,
  "layer_id" BIGINT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "notifications" (
  "id" BIGSERIAL PRIMARY KEY,
  "message" TEXT NOT NULL,
  "process_id" BIGINT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "request" (
  "id" BIGSERIAL PRIMARY KEY,
  "layer_id" BIGINT NOT NULL,
  "environment_id" BIGINT NOT NULL,
  "payload" JSON NOT NULL,
  "status" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE INDEX ON "project_environment" ("project_id");

CREATE INDEX ON "env_layer" ("environment_id");

CREATE INDEX ON "process" ("layer_id");

CREATE INDEX ON "notifications" ("process_id");

CREATE INDEX ON "request" ("environment_id");

ALTER TABLE "project_environment" ADD FOREIGN KEY ("project_id") REFERENCES "project" ("id");

ALTER TABLE "env_layer" ADD FOREIGN KEY ("environment_id") REFERENCES "project_environment" ("id");

ALTER TABLE "env_layer" ADD FOREIGN KEY ("template_id") REFERENCES "template" ("id");

ALTER TABLE "process" ADD FOREIGN KEY ("layer_id") REFERENCES "env_layer" ("id");

ALTER TABLE "notifications" ADD FOREIGN KEY ("process_id") REFERENCES "process" ("id");

ALTER TABLE "request" ADD FOREIGN KEY ("layer_id") REFERENCES "env_layer" ("id");

ALTER TABLE "request" ADD FOREIGN KEY ("environment_id") REFERENCES "project_environment" ("id");