CREATE TABLE IF NOT EXISTS "queue" (
	"id" bigserial NOT NULL,
    "name" text not null,
    "priority" INT default 1,
    "status" INT default 1,
    "create_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    "update_at" timestamp default NULL,
    CONSTRAINT queue_id PRIMARY KEY (id)
);