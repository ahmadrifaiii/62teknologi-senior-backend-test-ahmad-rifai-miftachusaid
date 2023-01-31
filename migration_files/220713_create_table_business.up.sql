CREATE TABLE IF NOT EXISTS businesses (
	id varchar(50) PRIMARY KEY NOT NULL,
	alias varchar(100) NOT NULL,
	name varchar(50) NOT NULL,
    image_url longtext,
    is_close bool,
    url longtext,
    review_count int,
    rating decimal,
    price varchar(20),
    deleted boolean DEFAULT false
);
