
CREATE TABLE category (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    title varchar(255) NOT NULL,
    parent uuid NOT NULL,
    CONSTRAINT category_pkey PRIMARY KEY (id)
);