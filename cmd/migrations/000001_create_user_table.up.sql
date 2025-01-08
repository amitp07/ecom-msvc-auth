CREATE TABLE IF NOT EXISTS msvc_users (
    id SERIAL PRIMARY KEY, 
    username varchar(50) not null,
    password text not null,
    role varchar(50) not null default 'user',
    email varchar(100) not null unique,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);