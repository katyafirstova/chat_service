-- +goose Up

create table chats
(
    chat_uuid uuid,
    user_uuid uuid,
    timestamp timestamp not null default now(),
    primary key (chat_uuid, user_uuid)
);

create table messages
(
    uuid       uuid primary key,
    user_uuid  uuid,
    chat_uuid  uuid,
    text       text      not null,
    created_at timestamp not null default now(),

    foreign key (chat_uuid, user_uuid) references chats (chat_uuid, user_uuid)
);

-- +goose Down
drop table messages;
drop table chats;
