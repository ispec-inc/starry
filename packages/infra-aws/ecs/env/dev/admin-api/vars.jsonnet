{
  cpu: "512",
  memory: "1024",
  container_def_name: "dev-yourapp-admin-api",
  image: "020604330519.dkr.ecr.ap-northeast-1.amazonaws.com/dev-admin-api",
  secrets: {
    MYSQL_YOURAPP_HOST: "/dev/aurora/main/writer_host",
    MYSQL_YOURAPP_PASSWORD: "/dev/aurora/main/YOURAPP_password",
    MYSQL_YOURAPP_USER: "/dev/aurora/main/YOURAPP_user",
  },
  environments: {
    PORT: "3001",
    HOST: "api.dev.yourapp.com",
  },
}
