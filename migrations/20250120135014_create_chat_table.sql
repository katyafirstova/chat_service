-- +goose Up

create table chats
(
    chat_uuid uuid,
    user_uuid uuid,
    timestamp timestamp not null default now(),
    primary key (chat_uuid, user_uuid)
);

-- +goose Down
drop table chats;
