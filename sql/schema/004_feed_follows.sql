-- +goose Up
CREATE TABLE IF NOT EXISTS feed_follows (
    id UUID PRIMARY KEY,
    feed_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    UNIQUE (feed_id, user_id),
    CONSTRAINT fk_users
        FOREIGN KEY (user_id)
        REFERENCES users(id),
    CONSTRAINT fk_feeds
        FOREIGN KEY (feed_id)
        REFERENCES feeds(id)
);  

-- +goose Down
DROP TABLE feed_follows;
