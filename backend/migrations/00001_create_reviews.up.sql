CREATE TABLE IF NOT EXISTS reviews (
    id          SERIAL      PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    rating      INT          NOT NULL CHECK (rating >= 1 AND rating <= 5),
    body        TEXT         NOT NULL,
    pros        TEXT[]       NOT NULL DEFAULT '{}',
    cons        TEXT[]       NOT NULL DEFAULT '{}',
    do_it_again BOOLEAN      NOT NULL DEFAULT TRUE,
    approved    BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
