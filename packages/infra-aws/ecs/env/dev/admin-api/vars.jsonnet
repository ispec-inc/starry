{
  cpu: "512",
  memory: "1024",
  container_def_name: "dev-pocopark-admin-api",
  image: "020604330519.dkr.ecr.ap-northeast-1.amazonaws.com/dev-admin-api",
  secrets: {
    MYSQL_POCOPARK_HOST: "/dev/aurora/main/writer_host",
    MYSQL_POCOPARK_PASSWORD: "/dev/aurora/main/pocopark_password",
    MYSQL_POCOPARK_USER: "/dev/aurora/main/pocopark_user",
    JWT_SECRET_KEY: "/dev/ecs/api/jwt_secret_key",
    OPENSEARCH_AWS_ACCESS_KEY_ID: "/dev/opensearch/access_key_id",
    OPENSEARCH_AWS_SECRET_ACCESS_KEY: "/dev/opensearch/secret_access_key",
    OPENSEARCH_ENDPOINT: "/dev/opensearch/endpoint",
    ADMIN_BASIC_AUTH_USER: "/dev/ecs/api/admin_basic_auth_user",
    ADMIN_BASIC_AUTH_PASSWORD: "/dev/ecs/api/admin_basic_auth_password",
    SECRET_KEY_BASE: "/dev/ecs/api/SECRET_KEY_BASE",
    DATADOG_API_KEY: "/dev/ecs/api/DATADOG_API_KEY",
  },
  environments: {
    PORT: "3001",
    HOST: "api.dev.pocopark.pococha.com",
    ALLOW_CORS_ORIGINS: "https://admin.dev.pocopark.pococha.com,http://localhost:3334",
    RAILS_ENV: "production",
    COGNITO_USER_POOL_ID: "ap-northeast-1_9WwiglaZW",
    RAILS_LOG_TO_STDOUT: "true",
    QA_MODE: '1',
  },
}
