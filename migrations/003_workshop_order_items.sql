create table
  workshop_order_items (
    id serial primary key,
    code varchar,
    quantity integer,
    price numeric,
    description text,
    orderId integer references workshop_orders
  );

---- create above / drop below ----

drop table
  workshop_order_items;
