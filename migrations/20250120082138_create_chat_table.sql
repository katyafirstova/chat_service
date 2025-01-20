-- +goose Up

create table chats
(
    uuid      uuid primary key,
    user_uuid uuid unique,
    timestamp timestamp not null default now()
);

create table messages
(
    uuid       uuid primary key,
    user_uuid  uuid unique,
    chat_uuid  uuid unique,
    text       text      not null,
    created_at timestamp not null default now(),

    foreign key (chat_uuid) references chats (uuid)
);

-- +goose Down

