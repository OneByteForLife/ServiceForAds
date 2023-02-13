/*
    Базовая структура для проверки работоспособности сервиса
*/

CREATE TABLE advertisements (
    id SERIAL PRIMARY KEY,
    product_name text not null,
    product_description varchar(1000) not null,
    product_main_picture varchar(1000) not null,
    product_more_pictures varchar(1000) array,
    date_create timestamp not null,
    price decimal(8, 2) not null
);

-- INSERT INTO advertisements (product_name, product_description, product_main_picture, product_more_pictures, date_create, price)
-- VALUES(
--     'Iphone 5s 64GB', -- Название товара
--     'Iphone 5s 64GB в идеальном состоянии. Документы в наличии. Не ремонтировался.', -- Описание товара 200 <= 1000
--     'https://iphone5s/image.png', -- Основная фотография товара
--     '{
--         "https://iphone5s/image0.png", 
--         "https://iphone5s/image1.png", 
--         "https://iphone5s/image2.png", 
--         "https://iphone5s/image3.png", 
--         "https://iphone5s/image4.png"
--     }', -- Дополнительные фотографии товара до 5 штук
--     '2023-01-27 10:22:34', -- Дата создания
--     15000 -- Цена товара
-- ), 
-- (
--     'Ноутбук Acer', -- Название товара
--     'Новый ноутбук имеется гарантия 12 месяцев. Использовался 3 дня. Продаю потому что я барыга перекуп', -- Описание товара 200 <= 1000
--     'https://acer/notepad/image.png', -- Основная фотография товара
--     '{
--         "https://acer/notepad/image0.png", 
--         "https://acer/notepad/image1.png", 
--         "https://acer/notepad/image2.png", 
--         "https://acer/notepad/image3.png", 
--         "https://acer/notepad/image4.png"
--     }', -- Дополнительные фотографии товара до 5 штук
--     '2023-01-28 10:33:04', -- Дата создания
--     30000 -- Цена товара
-- ),
-- (
--     'Ноутбук Lenovo', -- Название товара
--     'Новый ноутбук имеется гарантия. Использовался 3 месяца. Продаю потому что надоел!', -- Описание товара 200 <= 1000
--     'https://acer/notepad/image.png', -- Основная фотография товара
--     '{
--         "https://acer/notepad/image0.png", 
--         "https://acer/notepad/image1.png", 
--         "https://acer/notepad/image2.png", 
--         "https://acer/notepad/image3.png", 
--         "https://acer/notepad/image4.png"
--     }', -- Дополнительные фотографии товара до 5 штук
--     '2023-02-12 10:33:04', -- Дата создания
--     30000 -- Цена товара
-- )