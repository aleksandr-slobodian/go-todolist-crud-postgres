CREATE TABLE todos (
    id bigserial PRIMARY KEY,
    item varchar(100) NOT NULL,
    completed boolean NOT NULL DEFAULT FALSE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);