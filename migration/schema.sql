CREATE TABLE organizations (
    id varchar(255) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE organization_details (
    organization_id varchar(255) NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    alias varchar(255) NOT NULL,
    organization_type tinyint(8) UNSIGNED NOT NULL,
    contact varchar(255) NOT NULL,
    description text NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_organization_details_organizations_id FOREIGN KEY (
        organization_id
    ) REFERENCES organizations (id) ON DELETE CASCADE
);

CREATE TABLE active_organizations (
    organization_id varchar(255) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_active_organizations_organizations_id FOREIGN KEY (
        organization_id
    ) REFERENCES organizations (id) ON DELETE CASCADE
);

CREATE TABLE canceled_organizations (
    organization_id varchar(255) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_canceled_organizations_organizations_id FOREIGN KEY (
        organization_id
    ) REFERENCES organizations (id) ON DELETE CASCADE
);

CREATE TABLE part_of_organizations (
    organization_id varchar(255) NOT NULL,
    part_of_organization_id varchar(255) NOT NULL,
    PRIMARY KEY (organization_id, part_of_organization_id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_part_of_organizations_organizations_id FOREIGN KEY (
        organization_id
    ) REFERENCES organizations (id) ON DELETE CASCADE,
    CONSTRAINT fk_canceled_organizations_part_of_organizations_id FOREIGN KEY (
        part_of_organization_id
    ) REFERENCES organizations (id) ON DELETE CASCADE
);
