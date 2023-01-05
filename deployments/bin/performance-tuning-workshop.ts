#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import {PerformanceTuningWorkshop} from "../lib/performance-tuning-workshop";

const app = new cdk.App();
const githubId = app.node.tryGetContext('gh-account-id')

new PerformanceTuningWorkshop(app, 'PerformanceTuningWorkshopStack', {
    env: {
        account: process.env.CDK_DEFAULT_ACCOUNT,
        region: process.env.CDK_DEFAULT_REGION,
    },
    stackName: `${githubId}PerformanceTuningWorkshopStack`,
    githubId: githubId
});
