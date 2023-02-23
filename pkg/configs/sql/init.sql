create table comments
(
    id          bigint auto_increment
        primary key,
    video_id    bigint          null,
    user_id     bigint unsigned null,
    content     longtext        null,
    create_date longtext        null,
    delate_at   tinyint(1)      null,
    create_time bigint          null
);

create index idx_comments_create_time
    on comments (create_time);

create index idx_comments_delate_at
    on comments (delate_at);

create index vd
    on comments (video_id);

create table favorites
(
    id          bigint auto_increment
        primary key,
    video_id    bigint     null,
    user_id     bigint     null,
    action_type int        null,
    delate_at   tinyint(1) null
);

create index idx_favorites_delate_at
    on favorites (delate_at);

create index idx_member
    on favorites (video_id, user_id);

create index ua
    on favorites (action_type);

create index vu
    on favorites (video_id, user_id);

create table user
(
    id               bigint auto_increment
        primary key,
    name             longtext   null,
    password         longtext   null,
    follow_count     bigint     null,
    follower_count   bigint     null,
    avatar           longtext   null,
    signature        longtext   null,
    background_image longtext   null,
    total_favorited  longtext   null,
    is_follow        tinyint(1) null,
    work_count       bigint     null,
    favorite_count   bigint     null
);

create table videos
(
    id             bigint auto_increment
        primary key,
    author_id      bigint          null,
    play_url       longtext        null,
    cover_url      longtext        null,
    favorite_count bigint unsigned null,
    comment_count  bigint unsigned null,
    title          longtext        null,
    created_at     bigint          null
);

create index idx_videos_author_id
    on videos (author_id);

create index idx_videos_created_at
    on videos (created_at);

