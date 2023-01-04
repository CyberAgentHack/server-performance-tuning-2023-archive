import * as cdk from 'aws-cdk-lib';
import {CfnOutput} from 'aws-cdk-lib';
import * as apprunner from 'aws-cdk-lib/aws-apprunner';
import * as cr from 'aws-cdk-lib/custom-resources';
import * as rds from 'aws-cdk-lib/aws-rds';
import * as ec2 from "aws-cdk-lib/aws-ec2";
import {Port, SecurityGroup, SubnetType} from "aws-cdk-lib/aws-ec2";
import {Construct} from 'constructs';
import {CreateAutoScalingConfigurationCommandInput} from '@aws-sdk/client-apprunner';

export class PerformanceTuningWorkshop extends cdk.Stack {
    constructor(scope: Construct, id: string, props?: cdk.StackProps) {
        super(scope, id, props);
        /////////////////
        //// Network ////
        /////////////////
        const vpc = ec2.Vpc.fromLookup(this, 'VPC', {vpcId: this.node.tryGetContext('vpcId')})

        // AppRunner用セキュリティグループ(事前に作成済み)
        const sgAppRunner = SecurityGroup.fromLookupById(this, 'AppRunnerSecurityGroup', this.node.tryGetContext('appRunnerSecurityGroupID'))

        // Aurora用セキュリティグループ
        const sgAurora = new SecurityGroup(this, 'AuroraSecurityGroup', {
            securityGroupName: 'aurora',
            vpc: vpc,
            allowAllOutbound: true
        })
        // AppRunner→Auroraの接続を許可
        sgAurora.addIngressRule(sgAppRunner, Port.tcp(3306))

        ////////////////
        //// Aurora ////
        ////////////////
        const private_subnet = vpc.selectSubnets({subnetType: SubnetType.PRIVATE_WITH_EGRESS}).subnets
        const cluster = new rds.DatabaseCluster(this, 'Database', {
            engine: rds.DatabaseClusterEngine.auroraMysql({version: rds.AuroraMysqlEngineVersion.VER_3_02_1}),
            instanceProps: {
                instanceType: ec2.InstanceType.of(ec2.InstanceClass.R6G, ec2.InstanceSize.XLARGE2),
                vpcSubnets: {
                    subnets: private_subnet
                },
                securityGroups: [sgAurora],
                vpc
            },
            removalPolicy: cdk.RemovalPolicy.DESTROY,
        })

        new CfnOutput(this, 'dbSecretId', {value: cluster.secret?.secretName || ''})

        ///////////////////
        //// AppRunner ////
        ///////////////////
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
            networkConfiguration: {
                egressConfiguration: {
                    egressType: 'VPC',
                    vpcConnectorArn: this.node.tryGetContext('vpcConnectorArn')
                }
            },
            autoScalingConfigurationArn: autoScalingCfgArn,
            tags: [{
                key: 'Project',
                value: 'wsperf'
            }],
        })
    }
}
