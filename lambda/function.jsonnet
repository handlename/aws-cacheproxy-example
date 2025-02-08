{
  FunctionName: 'aws-cacheproxy-example',
  Environment: {
    Variables: {
    },
  },
  Handler: 'bootstarp.sh',
  MemorySize: 128,
  Role: '{{ env `ROLE_ARN` }}',
  Runtime: 'provided.al2023',
  Architectures: [
    'arm64',
  ],
  Tags: {
    Env: 'dev',
  },
  Timeout: 30,
}
