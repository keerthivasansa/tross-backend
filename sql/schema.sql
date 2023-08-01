CREATE TABLE users (
    name varchar(128) not null,
    id varchar(40) not null primary key
);

create table goals (
    id int auto_increment not null,
    content varchar(256) not null,
    name varchar(64) not null,
    user_id varchar(40) not null,
    created_at datetime not null default now(),
    active bool not null default false,
    primary key (id),
    foreign key (user_id) references users(id)
);

create table relapsed_days (
    day date not null,
    goal_id int not null,
    primary key (day, goal_id),
    foreign key (goal_id) references goals(id)
)