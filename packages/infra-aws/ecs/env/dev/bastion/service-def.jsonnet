{
  launchType: "FARGATE",
  enableExecuteCommand: true,
  networkConfiguration: {
    awsvpcConfiguration: {
      assignPublicIp: "DISABLED",
      securityGroups: [
        "{{ sg_tfstate `output.security_group_id` }}",
      ],
      subnets: [
        "{{ vpc_tfstate `output.private_subnets[0]` }}",
        "{{ vpc_tfstate `output.private_subnets[1]` }}",
      ],
    }
  },
}
