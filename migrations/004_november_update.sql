alter table workshop_order_items
drop column code,
drop column quantity;

alter table workshop_orders
drop column address,
drop column dni,
drop column ruc,
drop column color,
add column car_year int default 0,
add column discount float default 0;

alter table workshop_orders
rename column name to propietary;

---- create above / drop below ----

alter table workshop_orders
rename column propietary to name;
