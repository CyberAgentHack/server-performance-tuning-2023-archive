import * as cdk from 'aws-cdk-lib';
import {Construct} from 'constructs';
import * as apprunner from 'aws-cdk-lib/aws-apprunner';
import * as cr from 'aws-cdk-lib/custom-resources'
import {CreateAutoScalingConfigurationCommandInput} from '@aws-sdk/client-apprunner'

export class AppRunnerStack extends cdk.Stack {
    constructor(scope: Construct, id: string, props?: cdk.StackProps) {
        super(scope, id, props);

        // TODO: 削除時にautoScalingConfigurationが残ってしまうので、カスタムリソースのLambdaで削除するか、CLIで削除する
        // NOTE: 学生にオートスケーリングを設定してもらうため、デフォルトでオートスケーリングを無効にする
        const autoScalingConfiguration: CreateAutoScalingConfigurationCommandInput = {
            AutoScalingConfigurationName: 'auto-scaling-cfg-no-scale',
            MaxConcurrency: 100,
            MinSize: 1,
            MaxSize: 1,
        }

        const createAutoScalingConfiguration = new cr.AwsCustomResource(this, 'AutoScalingConfiguration', {
            onCreate: {
                service: 'AppRunner',
                action: 'createAutoScalingConfiguration',
                parameters: autoScalingConfiguration,
                physicalResourceId: cr.PhysicalResourceId.fromResponse('AutoScalingConfiguration.AutoScalingConfigurationArn')
            },
            policy: cr.AwsCustomResourcePolicy.fromSdkCalls({
                resources: cr.AwsCustomResourcePolicy.ANY_RESOURCE
            })
        })

        const autoScalingCfgArn = createAutoScalingConfiguration.getResponseField(
            'AutoScalingConfiguration.AutoScalingConfigurationArn'
        )

        new apprunner.CfnService(this, 'AppRunner', {
            serviceName: "wsperf-app-runner",
            sourceConfiguration: {
                authenticationConfiguration: {
                    connectionArn: this.node.tryGetContext('connection-arn')
                },
                codeRepository: {
                    codeConfiguration: {
                        configurationSource: 'REPOSITORY'
                    },
                    sourceCodeVersion: {
                        type: 'BRANCH',
                        value: 'main'
                    },
                    repositoryUrl: `https://github.com/${this.node.tryGetContext('gh-account-id')}/server-performance-tuning-2023`
                },
                autoDeploymentsEnabled: true
            },
            autoScalingConfigurationArn: autoScalingCfgArn,
            tags: [{
                key: 'Project',
                value: 'wsperf'
            }],
        })
    }
}