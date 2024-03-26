-- +goose Up
CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    title VARCHAR(90) NOT NULL,
    url VARCHAR(90) UNIQUE NOT NULL,
    description TEXT,
    published_at TIMESTAMPTZ,
    feed_id UUID NOT NULL,
    CONSTRAINT fk_feeds
        FOREIGN KEY(feed_id) 
        REFERENCES feeds(id)
);  

-- +goose Down
DROP TABLE posts;
