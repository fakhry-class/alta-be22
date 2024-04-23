-- DML = data manipulation language
-- INSERT 

-- insert data ke table users
INSERT INTO users (owner_name, email, password, phone, sex, address, bank_number, store_name)
values("andi", "andi@mail.com", "qwerty", "08651", "L", "surabaya", "098765","maju jaya");

INSERT INTO users (owner_name, email, password, phone, sex, address, bank_number, store_name)
values("budi", "budi@mail.com", "qwerty", "08652", "L", "malang", "098764","laris jaya");
INSERT INTO users (owner_name, email, password, phone, sex, address, bank_number, store_name)
values("budi 3", "budi3@mail.com", "qwerty", "08652", "L", "malang", "098764","laris jaya");

INSERT INTO users (owner_name, email, password, store_name)
values("budi", "budisudarsono@mail.com", "qwerty", "budi sudarsono jaya");

-- error karena password tidak boleh kosong
INSERT INTO users (owner_name, email, store_name)
values("budi2", "budisudarsono2@mail.com", "budi sudarsono jaya");

-- error karena email tidak boleh sama
INSERT INTO users (owner_name, email, password, store_name)
values("budi copy", "budisudarsono@mail.com", "qwerty", "budi sudarsono jaya");

-- insert data ke table products 
INSERT INTO products (id, user_id,product_name, description,price, stock, `type`,weight)
VALUES ("P0001", 1, "mouse logitech", "mouse gaming dari logitech", 200000, 10, "elektronik", 1);

INSERT INTO products (id, user_id,product_name, description,price, stock, `type`,weight)
VALUES ("P0002", 3, "mouse fantech", "mouse gaming dari fantech", 200000, 10, "elektronik", 1);

INSERT INTO products (id, user_id,product_name, description,price, stock, `type`,weight)
VALUES ("P0003", 3, "mouse sony", "mouse gaming dari sony", 250000, 5, "elektronik", 1);

-- error karena id tidak boleh sama
INSERT INTO products (id, user_id,product_name, description,price, stock, `type`,weight)
VALUES ("P0003", 3, "mouse sony", "mouse gaming dari sony", 250000, 5, "elektronik", 1);

INSERT INTO products (id, user_id,product_name, description,price, stock, `type`,weight)
VALUES ("P0010", 3, "beras rajawali", "beras enak merk rajawali", 50000, 5, "makanan", 5);

INSERT INTO products (id, user_id,product_name, description,price, stock, `type`,weight)
VALUES ("P0009", 3, "beras rajawali 1", "beras enak merk rajawali", 50000, 5, "makanan", 5);

INSERT INTO products (id, user_id,product_name, description,price, stock, `type`,weight)
VALUES ("P0011", 6, "gula pasir rajawali", "gula enak merk rajawali", 20000, 5, "makanan", 5);

-- insert ke table profile_picture 
INSERT INTO profile_picture (user_id, url)
VALUES(1,"https://blabla.com/1");

-- tidak bisa, karena 1 user hanya boleh punya 1 foto
INSERT INTO profile_picture (user_id, url)
VALUES(1,"https://foto.com/1a");

INSERT INTO profile_picture (user_id, url)
VALUES(2,"https://foto.com/2");

INSERT INTO profile_picture (user_id, url)
VALUES(5,"https://foto.com/2");

-- SELECT 
-- membaca data, dan menampilkan seluruh kolom
SELECT * FROM users;

SELECT * FROM products;

SELECT * FROM profile_picture;

SELECT * FROM favourites;

-- menampilkan data di spesifik kolom
SELECT id, user_id, product_name from products;

-- menampilkan data tertentu dari satu table dengan where clause
SELECT * FROM products WHERE user_id = 2;
SELECT * FROM products WHERE id = "P0002";
SELECT * FROM products WHERE `type`  = "elektronik" ;
SELECT * FROM products WHERE stock  =10;

SELECT * FROM users WHERE phone is not null;
SELECT * FROM users WHERE phone is null;

-- UPDATE 
-- melakukan edit/update data
UPDATE users SET 
address ="jakarta"
WHERE id = 1;

UPDATE products SET 
price = 150000,
stock = 20
WHERE id = "P0001";

-- DELETE 
-- menghapus data
DELETE FROM products WHERE id = "P0003";

DELETE FROM users WHERE id = 2;

-- FK
-- RESTRICT 
-- kita tidak bisa menghapus data di table parent ketika masih dibutuhkan di table child
-- cara hapus: hapus dulu data di table child, baru anda bisa hapus data di parent

delete from profile_picture WHERE id = 3;
delete from profile_picture WHERE user_id  = 2;
SELECT * FROM profile_picture;
DELETE FROM users WHERE id = 2;
SELECT * from users;

-- CASCADE
-- ketika kita menghapus data di table parent, maka data di table child akan ikut kehapus
DELETE FROM users WHERE id = 5;

-- SOFT DELETE 
SELECT * from users;
SELECT * from products p ;
INSERT INTO favourites (user_id, product_id)
values (6, "P0002");
INSERT INTO favourites (user_id, product_id)
values (6, "P0001");

SELECT * FROM favourites ;
-- menghapus data menggunakan pendekatan soft delete
-- soft delete = data tidak benar" dihapus, melainkan hanya diberi tanda/flag bahwa data tsb dihapus
update favourites set deleted_at = "2024-04-02 21:30:28" WHERE id = 1;
SELECT * FROM favourites WHERE deleted_at is NULL ;

-- LIKE / BETWEEN
SELECT * FROM products p ;
SELECT * FROM products WHERE user_id = 3;
SELECT * FROM products WHERE `type`  = "elektroni" ;
SELECT * FROM products WHERE product_name LIKE "mouse%";
SELECT * FROM products WHERE product_name LIKE "%rajawali";
SELECT * FROM products WHERE product_name LIKE "%rajawali%";
SELECT * FROM products WHERE `type`  LIKE  "%elektroni%" ;

SELECT * FROM products WHERE price BETWEEN 0 AND 100000 ;

-- AND OR 
SELECT * FROM products ; 
SELECT * FROM products WHERE stock > 10 and `type` = 'elektronik' ;
SELECT * FROM products WHERE `type` = 'elektronik' and price BETWEEN 100000 and 200000;

SELECT * FROM products WHERE stock > 10 OR `type` = 'elektronik' ;
SELECT * FROM products WHERE stock > 10 OR `type` = 'elektronik' OR weight = 1 ;

SELECT * FROM products WHERE `type` = 'elektronik' and price BETWEEN 100000 and 200000 or weight =1;

-- ORDER BY
-- ASC : kecil ke besar
-- DESC : besar ke kecil
SELECT * FROM products ;
SELECT * FROM products order by stock ASC;
SELECT * FROM products order by stock DESC;
SELECT * FROM products WHERE stock > 10 OR `type` = 'elektronik' OR weight = 1 order by stock ASC;

-- LIMIT
SELECT * FROM products LIMIT 2;
SELECT * FROM products order by stock ASC LIMIT 1;
SELECT * FROM products WHERE stock > 10 OR `type` = 'elektronik' OR weight = 1 order by stock ASC LIMIT 1;


-- JOIN
select * FROM products;
SELECT * FROM users;
SELECT * from profile_picture;
SELECT * from favourites f ;
SELECT id, user_id, product_name, price from products;

SELECT id, owner_name, store_name from users;

-- INNER JOIN
-- menampilkan data product beserta data pemilik (owner name dan store name)
SELECT products.id, products.user_id, products.product_name, products.price, users.id, users.owner_name, users.store_name 
from products
INNER JOIN users ON products.user_id = users.id ;

-- menampilkan data user beserta url foto profilnya
-- memberikan alias ke table dan kolom
SELECT u.id, u.owner_name , u.email , pp.url as url_photo  
from users u
INNER JOIN profile_picture pp ON u.id = pp.user_id ;

-- LEFT JOIN =  seluruh data di table kiri akan dimunculkan
-- menampilkan seluruh/semua data user beserta url foto profilnya
SELECT u.id, u.owner_name , u.email , pp.url as url_photo  
from users u
LEFT JOIN profile_picture pp ON u.id = pp.user_id ;

-- RIGHT JOIN = seluruh data di table kanan akan dimunculkan
SELECT u.id, u.owner_name , u.email , pp.url as url_photo  
from users u
RIGHT JOIN profile_picture pp ON u.id = pp.user_id ;

SELECT u.id, u.owner_name , u.email , pp.url as url_photo  
from profile_picture pp
RIGHT JOIN users u ON u.id = pp.user_id ;


-- menampilkan seluruh data product beserta data pemilik dan foto pemilik
SELECT 
products.id , products.product_name , products.price , products.stock, products.user_id,
users.owner_name , profile_picture.url 
from products 
INNER JOIN users ON products.user_id  = users.id 
LEFT JOIN profile_picture ON users.id = profile_picture.user_id ;

-- menampilkan product yg difavoritkan oleh user, tampilkan nama yang memfavoritkan, tampilkan nama pemilik productnya
select products.id, products.product_name, uowner.owner_name  ,favourites.user_id , users.owner_name as name, favourites.created_at 
from products
inner join favourites ON products.id = favourites.product_id 
inner join users ON favourites.user_id = users.id
inner join users uowner ON products.user_id = uowner.id;

-- UNION = distinct
-- UNION ALL = duplicate
SELECT product_name from products
UNION
select owner_name from users;

-- AGGREGATION
-- MIN
select min(stock) as min_stock from products;
select min(created_at) as min_created from favourites;

-- MAX
select max(stock) as max_stock from products;

-- SUM / jumlah total nilai
SELECT * from products p ;
SELECT SUM(stock) as jumlah_stock from products; 
SELECT SUM(price) as jumlah_harga from products; 

-- AVG / average / rata-rata
SELECT AVG(stock) as rata_stock from products; 
SELECT AVG(price) as rata_harga from products; 

-- COUNT / banyaknya data/ jumlah data
-- banyaknya product
SELECT COUNT(id) from products; 
SELECT COUNT(id) from users u ; 

SELECT * FROM products p ;
-- tampilkan data user beserta jumlah product yg dimilikinya
SELECT users.id, users.owner_name , users.email , COUNT(products.id) as jumlah_product
from users
INNER JOIN products ON users.id = products.user_id 
GROUP BY products.user_id ;

SELECT users.id, users.owner_name , users.email , profile_picture.url , COUNT(products.id) as jumlah_product
from users
INNER JOIN products ON users.id = products.user_id 
LEFT JOIN profile_picture ON users.id = profile_picture.user_id  
GROUP BY products.user_id ;

-- order of execution
SELECT users.id, users.owner_name , users.email , profile_picture.url , COUNT(products.id) as jumlah_product
from users
INNER JOIN products ON users.id = products.user_id 
LEFT JOIN profile_picture ON users.id = profile_picture.user_id  
WHERE users.id  = 3
GROUP BY products.user_id 
;

SELECT users.id, users.owner_name , users.email , profile_picture.url , COUNT(products.id) as jumlah_product
from users
INNER JOIN products ON users.id = products.user_id 
LEFT JOIN profile_picture ON users.id = profile_picture.user_id  
GROUP BY products.user_id
ORDER  BY COUNT(products.id) ASC
LIMIT 2
;

-- HAVING
-- tampilkan data user beserta jumlah product yg dimilikinya, tampilkan hanya user yang memiliki product lebih dari 2
SELECT users.id, users.owner_name , users.email , COUNT(products.id) as jumlah_product
from users
INNER JOIN products ON users.id = products.user_id 
GROUP BY products.user_id
HAVING COUNT(products.id) > 2 ;

-- SUBQUERY
-- 1 tampilkan data user yang punya product

-- 1a tampilkan semua user id yang mempunyai product
SELECT * from products;
SELECT user_id from products;
SELECT DISTINCT user_id from products;
SELECT user_id from products GROUP BY user_id ;

-- 1b tampilkan data user
SELECT * from users;

-- 1c tampilkan data user yang punya product
SELECT * from users WHERE id IN (SELECT user_id from products GROUP BY user_id);
-- SELECT * from users WHERE id IN (1,3,6)

-- update users set status = "banned" where id in (select user_id from reports group by user_id having ...)

-- tampilkan data user yang tidak punya product
SELECT * from users WHERE id NOT IN (SELECT user_id from products GROUP BY user_id);

-- tampilkan data user yang punya product pakai pendekatan join
select DISTINCT users.id, users.owner_name , users.email 
from users 
inner join products ON users.id = products.user_id ;

-- FUNCTION
SELECT COUNT(*) FROM products WHERE user_id = 9;

-- buat function di file terpisah
DELIMITER $$
CREATE FUNCTION sf_count_product_peruser(user_id_p int) RETURNS INT DETERMINISTIC
BEGIN
DECLARE total INT;
SELECT COUNT(*) INTO total FROM products WHERE user_id = user_id_p;
RETURN total;
END$$
DELIMITER ;

-- menjalankan function
SELECT sf_count_product_peruser(3);

-- menampilkan list function beserta status func
show function status where db='db_be22';

-- menghapus function
DROP function if exists `db_be22`.`sf_count_product_peruser`;

-- TRIGGER
-- buat dulu triggernya di file terpisah, lalu ketika sudah berhasil, kalian bisa jalankan statement yg mentrigger
DELIMITER $$
CREATE TRIGGER delete_all_data_user
BEFORE DELETE ON users FOR EACH ROW
BEGIN
-- declare variables
DECLARE v_user_id INT;
SET v_user_id=OLD.id;
-- trigger code
DELETE FROM profile_picture WHERE user_id=v_user_id;
DELETE FROM favourites WHERE user_id=v_user_id;
DELETE FROM products WHERE user_id=v_user_id;
END$$
DELIMITER ;

-- jalankan statement delete, maka secara otomatis trigger akan berjalan
DELETE from users where id = 6;

SELECT * FROM users u ;
SELECT * FROM products p ;

-- hapus trigger
DROP trigger if exists `db_be22`.`delete_all_data_user`;


SELECT id, owner_name, email, password, phone, sex, address, bank_number, store_name FROM users;
