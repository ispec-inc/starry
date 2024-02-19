local env = import '../env.jsonnet';
local vars = import 'vars.jsonnet';

{
  deploymentConfiguration: {
    deploymentCircuitBreaker: {
      enable: false,
      rollback: false,
    },
    maximumPercent: 200,
    minimumHealthyPercent: 100,
  },
  deploymentController: {
    type: 'ECS',
  },
  desiredCount: 1,
  enableECSManagedTags: false,
  enableExecuteCommand: false,
  healthCheckGracePeriodSeconds: 3600,
  launchType: 'FARGATE',
  loadBalancers: [
    {
      containerName: vars.container_def_name,
      containerPort: 3001,
      targetGroupArn: env.loadBalancer.adminTgArn,
    },
  ],
  networkConfiguration: {
    awsvpcConfiguration: {
      assignPublicIp: 'DISABLED',
      securityGroups: [
        env.vpc.securityGroups.api,
      ],
      subnets: env.vpc.subnets.private,
    },
  },
  pendingCount: 0,
  platformFamily: 'Linux',
  platformVersion: 'LATEST',
  propagateTags: 'NONE',
  runningCount: 0,
  schedulingStrategy: 'REPLICA',
  tags: [
    { key: 'App', value: 'yourapp' },
    { key: 'Env', value: 'dev' },
    { key: 'Name', value: 'dev-yourapp-admin-api' },
  ],
}
