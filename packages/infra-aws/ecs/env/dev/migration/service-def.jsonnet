local env = import '../env.jsonnet';
{
  launchType: 'FARGATE',
  networkConfiguration: {
    awsvpcConfiguration: {
      assignPublicIp: 'ENABLED',
      securityGroups: [
        env.vpc.securityGroups.api,
      ],
      subnets: env.vpc.subnets.private,
    },
  },
}
