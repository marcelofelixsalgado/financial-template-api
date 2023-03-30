CREATE DATABASE IF NOT EXISTS financial_TEMPLATE_db;
USE financial_TEMPLATE_db;

DROP TABLE IF EXISTS TEMPLATES;

CREATE TABLE TEMPLATES(
    id varchar(255) primary key,
    tenant_id varchar(255) not null,
    name varchar(255) not null,
    created_at datetime not null,
    updated_at datetime default '0001-01-01 00:00:00'
);