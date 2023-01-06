#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import {PerformanceTuningWorkshopStack} from "../lib/performance-tuning-workshop-stack";

const app = new cdk.App();
const githubId = app.node.tryGetContext('gh-account-id')

new PerformanceTuningWorkshopStack(app, 'PerformanceTuningWorkshopStack', {
    env: {
        account: process.env.CDK_DEFAULT_ACCOUNT,
        region: process.env.CDK_DEFAULT_REGION,
    },
    stackName: `${githubId}PerformanceTuningWorkshopStack`,
    githubId: githubId
});
