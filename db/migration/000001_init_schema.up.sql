CREATE TABLE "project" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "git_path" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "description" TEXT
);

CREATE TABLE "project_environment" (
  "id" BIGSERIAL PRIMARY KEY,
  "git_branch" VARCHAR(255) NOT NULL,
  "project_id" BIGINT NOT NULL,
  "description" TEXT,
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "template" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "path" VARCHAR(255) NOT NULL,
  "type" VARCHAR(255) NOT NULL
);  
CREATE TABLE "env_layer" (
  "id" BIGSERIAL PRIMARY KEY,
  "environment_id" BIGINT NOT NULL,
  "s3_path" VARCHAR(255) NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "process_id" BIGINT NOT NULL,
  "current_request_id" BIGINT 
  );

CREATE TABLE "process" (
  "id" BIGSERIAL PRIMARY KEY,
  "argo_id" BIGINT NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "template_id" BIGINT NOT NULL
  );

CREATE TABLE "argo_workflow" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "path" VARCHAR(255) NOT NULL,
  "description" BIGINT NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);


CREATE TABLE "request" (
  "id" BIGSERIAL PRIMARY KEY,
  "layer_id" BIGINT NOT NULL,
  "environment_id" BIGINT NOT NULL,
  "payload" JSON NOT NULL,
  "status" VARCHAR(255) NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);



CREATE INDEX ON "project_environment" ("project_id");

CREATE INDEX ON "request" ("layer_id");

ALTER TABLE "project_environment" ADD FOREIGN KEY ("project_id") REFERENCES "project" ("id");

ALTER TABLE "env_layer" ADD FOREIGN KEY ("environment_id") REFERENCES "project_environment" ("id");

ALTER TABLE "env_layer" ADD FOREIGN KEY ("process_id") REFERENCES "process" ("id");



ALTER TABLE "request" ADD FOREIGN KEY ("layer_id") REFERENCES "env_layer" ("id");

ALTER TABLE "process" ADD FOREIGN KEY ("argo_id") REFERENCES "argo_workflow" ("id");

ALTER TABLE "process" ADD FOREIGN KEY ("template_id") REFERENCES "template" ("id");
