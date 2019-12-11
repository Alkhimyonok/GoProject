insert into category(description) values ('Magazines');
insert into category(description) values ('Books');
insert into category(description) values ('Books for children');
insert into category(description) values ('Cards');

insert into products(name, price, quantity, amount, category) values
('Cosmopolitan', 10, 22,(22 * 10), 1);
insert into products(name, price, quantity, amount, category) values
('Maxim', 5, 20,(20*5),1);

insert into products(name, price, quantity, amount, category) values
('Harry Potter', 20, 10,(20*10),2);
insert into products(name, price, quantity, amount, category) values
('Game of Thrones', 40, 20,(40*20),2);

insert into products(name, price, quantity, amount, category) values
('Colobok', 10, 30,(10*30),3);
insert into products(name, price, quantity, amount, category) values
('Ariel', 15, 20,(15*20),3);

insert into products(name, price, quantity, amount, category) values
('B-day', 1, 10,(1*10),4);
insert into products(name, price, quantity, amount, category) values
('New Year', 2, 20,(2*20),4);
