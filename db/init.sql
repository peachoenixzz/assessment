CREATE TABLE IF NOT EXISTS "expenses" (
    id SERIAL PRIMARY KEY,
    title TEXT,
    amount FLOAT,
    note TEXT,
    tags TEXT[]
);


INSERT INTO "expenses" ("id", "title", "amount", "note","tags") VALUES (1, 'buy a new phone', 39000, 'buy a new phone' , '{"gadget", "shopping"}');