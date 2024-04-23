-- DDL data definition language
-- membuat database
CREATE DATABASE db_be22;

-- memilih db yang akan digunakan/ di manage
USE db_be22;

-- membuat table
CREATE TABLE users(
	id int primary key auto_increment,
	owner_name varchar(255) not null,
	email varchar(255) not null unique,
	password varchar(255) not null,
	phone varchar(255),
	sex varchar(255),
	address text,
	bank_number varchar(255),
	store_name varchar(255)
);

-- menampilkan seluruh table di suatu db
show tables;
-- menampilkan list seluruh column di suatu table
show columns from users;
show columns from products;

-- membuat table product dengan kolom foreign key ke table user
create table products(
	id varchar(255) primary key,
	user_id int,
	product_name varchar(255),
	description text,
	price decimal,
	stock int,
	type varchar(255),
	weight int,
	constraint fk_users_products foreign key(user_id) references users(id)
);

create table profile_picture(
id int primary key auto_increment,
user_id int unique,
url text,
constraint fk_users_profilepic foreign key(user_id) references users(id)
);

create table favourites(
	id int primary key auto_increment,
	user_id int,
	product_id varchar(255),
	created_at datetime default current_timestamp,
	updated_at datetime default current_timestamp,
	deleted_at datetime,
	constraint fk_users_favourites foreign key(user_id) references users(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
	constraint fk_products_favourites foreign key(product_id) references products(id) ON UPDATE CASCADE ON DELETE CASCADE
);

create table dummy(
id int primary key auto_increment,
name varchar(255)
);

-- memodifikasi table (ALTER)
-- menambahkan kolom
ALTER table dummy 
ADD COLUMN description text;

-- mengubah tipe data column
ALTER table dummy 
MODIFY COLUMN description varchar(255);

-- menghapus kolom
ALTER table dummy 
DROP COLUMN description;

ALTER table dummy 
ADD COLUMN user_id int;

-- menambahkan foreign key
ALTER table dummy 
ADD CONSTRAINT fk_user_dummy
FOREIGN KEY(user_id) REFERENCES users(id);

-- menghapus FK
ALTER table dummy 
DROP FOREIGN KEY fk_user_dummy;

show columns from dummy;

-- menghapus table
DROP TABLE dummy;

-- menghapus database
DROP DATABASE db_be22;


ALTER TABLE db_be22.profile_picture DROP FOREIGN KEY fk_users_profilepic;
ALTER TABLE db_be22.profile_picture ADD CONSTRAINT fk_users_profilepic FOREIGN KEY (user_id) REFERENCES db_be22.users(id) ON DELETE CASCADE ON UPDATE CASCADE;

