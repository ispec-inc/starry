local env = import '../env.jsonnet';
local vars = import 'vars.jsonnet';

{
  family: 'dev-yourapp-admin-api',
  cpu: vars.cpu,
  executionRoleArn: env.executionRoleArn,
  taskRoleArn: env.taskRoleArn,
  ipcMode: '',
  memory: vars.memory,
  networkMode: 'awsvpc',
  pidMode: '',
  requiresCompatibilities: ['FARGATE'],
  tags: [
    { key: 'App', value: 'yourapp' },
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
          dd_service: 'yourapp',
          dd_source: 'api-admin-server',
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
          containerPort: 3001,
          hostPort: 3001,
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
      logConfiguration: {
        logDriver: 'awslogs',
        options: {
          'awslogs-group': '/aws/ecs/cluster/dev/pcpk-api/ecs',
          'awslogs-region': 'ap-northeast-1',
          'awslogs-stream-prefix': 'fluent-bit-',
        },
      },
      memoryReservation: 50,
    },
  ],
}
