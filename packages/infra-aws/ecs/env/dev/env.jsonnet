{
  executionRoleArn: "{{ ecs_tfstate `output.task_execution_role_arn` }}",
  taskRoleArn: "{{ ecs_tfstate `output.task_role_arn` }}",
  bastionTaskRoleArn: "{{ ecs_tfstate `output.bastion_task_role_arn` }}",
  loadBalancer: {
    userTgArn: "{{ alb_tfstate `output.user_api_tg_arn` }}",
    adminTgArn: "{{ alb_tfstate `output.admin_api_tg_arn` }}",
  },
  vpc: {
    securityGroups: {
      api: "{{ sg_tfstate `output.security_group_id` }}",
    },
    subnets: {
      private: [
        "{{ vpc_tfstate `output.private_subnets[0]` }}",
        "{{ vpc_tfstate `output.private_subnets[1]` }}",
      ],
    }
  }
}
