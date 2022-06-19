CREATE TABLE users (
    id bigserial PRIMARY KEY,
    username text NOT NULL,
    email text,
    password text NOT NULL
);

CREATE TABLE projects (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
	password text NOT NULL
);

CREATE TABLE users_projects (
    user_id bigint NOT NULL REFERENCES public.users (id) ON DELETE CASCADE,
    project_id bigint NOT NULL REFERENCES public.projects (id) ON DELETE CASCADE
);

CREATE TABLE tasks (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    description text,
    state text NOT NULL,
    creation_date timestamp NOT NULL,
    position bigint NOT NULL,
    project_id bigint NOT NULL REFERENCES public.projects (id) ON DELETE CASCADE
);

CREATE TABLE projects_logs (
    id bigserial PRIMARY KEY,
    date timestamp NOT NULL,
    type text NOT NULL,
    arguments text[],
    user_id bigint NOT NULL REFERENCES public.users (id) ON DELETE CASCADE,
    project_id bigint NOT NULL REFERENCES public.projects (id) ON DELETE CASCADE
);