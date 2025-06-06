CREATE TABLE IF NOT EXISTS users (
    id serial4 NOT NULL,
    username varchar(50) NOT NULL,
    email varchar(100) NOT NULL,
    password_hash text NOT NULL,
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_username_key UNIQUE (username)
);

CREATE TABLE IF NOT EXISTS "groups" (
    id serial4 NOT NULL,
    user_id int4 NOT NULL,
    title varchar(100) NOT NULL,
    image varchar NULL,
    group_id int4 NULL,
    CONSTRAINT groups_pkey PRIMARY KEY (id),
    CONSTRAINT groups_users_fk FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE cards (
    id serial4 NOT NULL,
    user_id int4 NOT NULL,
    group_id int4 NULL,
    title varchar(100) NOT NULL,
    username varchar(100) NULL,
    "password" text NULL,
    website varchar(255) NULL,
    notes text NULL,
    image varchar NULL,
    is_favorite bool NULL,
    CONSTRAINT cards_pkey PRIMARY KEY (id),
    CONSTRAINT cards_group_id_fkey FOREIGN KEY (group_id) REFERENCES "groups"(id) ON DELETE SET NULL,
    CONSTRAINT cards_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);