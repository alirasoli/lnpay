CREATE TABLE payment (
    hash TEXT NOT NULL PRIMARY KEY,
    invoice TEXT NOT NULL,
    amount INTEGER NOT NULL,
    webhook TEXT NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
    paid_at INTEGER
);

CREATE INDEX active_payment ON "payment" (
    "created_at",
    "paid_at"
);
