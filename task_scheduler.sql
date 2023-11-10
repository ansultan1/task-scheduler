CREATE DATABASE task_scheduler;

use task_scheduler;
CREATE TABLE your_table_name (
                                 id SERIAL PRIMARY KEY,
                                 name VARCHAR(255),
                                 command VARCHAR(255),
                                 scheduled_time TIMESTAMP,
                                 recurring BOOLEAN,
                                 time_zone VARCHAR(255)
);