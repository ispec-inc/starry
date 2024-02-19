local env = import '../env.jsonnet';
local vars = import 'vars.jsonnet';

{
  family: 'dev-pocopark-user-api',
  cpu: vars.cpu,
  executionRoleArn: env.executionRoleArn,
  taskRoleArn: env.taskRoleArn,
  ipcMode: '',
  memory: vars.memory,
  networkMode: 'awsvpc',
  pidMode: '',
  requiresCompatibilities: [
    'FARGATE',
  ],
  tags: [
    { key: 'App', value: 'pocopark' },
    { key: 'Env', value: 'dev' },
  ],
  containerDefinitions: [
    {
      cpu: 0,
      dockerLabels: {},
      environment: [
        {
          name: k,
          value: vars.environments[k],
        }
        for k in std.objectFields(vars.environments)
      ],
      essential: true,
      image: vars.image + ':{{ must_env `TAG` }}',
      logConfiguration: {
        logDriver: 'awsfirelens',
        options: {
          Name: 'datadog',
          Host: 'http-intake.logs.datadoghq.com',
          TLS: 'on',
          dd_service: 'pocopark',
          dd_source: 'api-user-server',
          dd_tags: 'env:dev',
          provider: 'ecs',
        },
        secretOptions: [{
          name: "apikey",
          valueFrom: vars.secrets.DATADOG_API_KEY,
        }]
      },
      name: vars.container_def_name,
      portMappings: [
        {
          containerPort: 3000,
          hostPort: 3000,
          protocol: 'tcp',
        },
      ],
      readonlyRootFilesystem: false,
      secrets: [
        {
          name: k,
          valueFrom: vars.secrets[k],
        }
        for k in std.objectFields(vars.secrets)
      ],
      ulimits: [
        {
          hardLimit: 8192,
          name: 'nofile',
          softLimit: 2048,
        },
      ],
    },
    {
      name: "fluent-bit",
      image: "public.ecr.aws/aws-observability/aws-for-fluent-bit:stable",
      essential: true,
      firelensConfiguration: {
        type: "fluentbit",
        options: {
          "enable-ecs-log-metadata": "true"
        }
      },
      memoryReservation: 50,
    },
  ],
}
