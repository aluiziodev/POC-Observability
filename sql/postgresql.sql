CREATE TABLE orders (
    id          TEXT PRIMARY KEY,
    item        TEXT NOT NULL,
    quantity    INTEGER NOT NULL CHECK (quantity > 0),
    status      TEXT NOT NULL CHECK (status IN ('pending', 'confirmed', 'shipped', 'delivered', 'cancelled')),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_orders_status ON orders (status);