CREATE USER user
    PASSWORD 'pass';

CREATE DATABASE gophermart
    OWNER 'user'
    ENCODING 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8';

CREATE DATABASE accrual
    OWNER 'user'
    ENCODING 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8';