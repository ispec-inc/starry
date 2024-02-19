{
  cpu: '1024',
  memory: '2048',
  container_def_name: 'dev-yourapp-user-api',
  image: 'xxxxxxxxxx.dkr.ecr.ap-northeast-1.amazonaws.com/dev-api',
  secrets: {
    MYSQL_YOURAPP_HOST: '/dev/aurora/main/writer_host',
    MYSQL_YOURAPP_PASSWORD: '/dev/aurora/main/YOURAPP_password',
    MYSQL_YOURAPP_USER: '/dev/aurora/main/YOURAPP_user',
    MYSQL_YOURAPP_NAME: '/dev/aurora/main/YOURAPP_user',
  },
  environments: {
    PORT: '3000',
    HOST: 'api.dev.yourapp.com',
  },
}
