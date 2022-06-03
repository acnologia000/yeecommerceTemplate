drop table if exists items ; 

create table items (
    item_name varchar(200),
    item_description varchar(9000),
    price int,
    item_priority int,
    totalClicks int, 
    added_on timestamp default current_timestamp
)


-- random timestamp generation logic 


-- update items 
-- set added_on = (select timestamp '1970-01-10 20:00:00' +
--        random() * (timestamp '2022-01-20 20:00:00' -
--                    timestamp '1970-01-10 10:00:00'))

-- v2 that needs to multiple times but works as expected 

-- update items 
-- set added_on = (select timestamp '1970-01-10 20:00:00' +
--        random() * (timestamp '2022-01-20 20:00:00' -
--                    timestamp '1970-01-10 10:00:00')) 

-- where price between random()*10000 + 1 and random()*10000 +1;



-- to search for between time stamps collected in number format
-- select * from items

-- where added_on between TO_TIMESTAMP('2014-03-31', 'YYYY-MM-DD') and TO_TIMESTAMP('2018-03-31', 'YYYY-MM-DD');