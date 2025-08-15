CREATE TABLE programs (
    int_id SERIAL PRIMARY KEY,
    id UUID DEFAULT gen_random_uuid (),
    type VARCHAR(20) NOT NULL CHECK (type IN ('абонемент', 'сертификат')),
    name VARCHAR(255) NOT NULL,
    image TEXT,
    fixed_price INT NOT NULL,
    total_services_cost INT NOT NULL,
    discount_percent INT NOT NULL,
    valid_until DATE NOT NULL,
    terms TEXT,
    created_at numeric not null default extract(epoch from now()),
    updated_at numeric not null default extract(epoch from now()),
    active BOOLEAN DEFAULT TRUE
);

CREATE TABLE services (
    service_id UUID,
    name TEXT NOT NULL,
    tarif INT NOT NULL,
    duration INT NOT NULL
);

CREATE TABLE program_services (
    program_id UUID NOT NULL REFERENCES programs (id) ON DELETE CASCADE,
    service_id UUID NOT NULL REFERENCES services (service_id) ON DELETE CASCADE,
    PRIMARY KEY (program_id, service_id)
);
