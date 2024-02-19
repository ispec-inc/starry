local env = import '../env.jsonnet';

{
  family: "dev-bastion",
  cpu: "256",
  memory: "512",
  executionRoleArn: env.executionRoleArn,
  taskRoleArn: env.bastionTaskRoleArn,
  networkMode: "awsvpc",
  requiresCompatibilities: ["FARGATE"],
  containerDefinitions: [
    {
      name: "bastion",
      image: "yumafuu/sleepy",
      command: ["3000"],
      essential: true,
    }
  ],
}
