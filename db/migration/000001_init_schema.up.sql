CREATE TABLE "service" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" text NOT NULL,
  "description" text NOT NULL,
  "versions" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "service" ("name");