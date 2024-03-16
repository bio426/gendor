create table
  workshop_orders (
    id serial primary key,
    name varchar,
    address varchar,
    dni varchar,
    ruc varchar,
    brand varchar,
    model varchar,
    color varchar,
    plate varchar,
    mileage int,
    observation text,
    created_at timestamp default now(),
    userId integer references users
  );

---- create above / drop below ----

drop table
  workshop_orders;
