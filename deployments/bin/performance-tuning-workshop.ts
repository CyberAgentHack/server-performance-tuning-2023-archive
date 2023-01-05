#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import {PerformanceTuningWorkshop} from "../lib/performance-tuning-workshop";

const app = new cdk.App();
const githubId = app.node.tryGetContext('gh-account-id')
if (githubId == undefined) {
    throw new Error(`GitHub ID is Empty`)
}

new PerformanceTuningWorkshop(app, `${githubId}PerformanceTuningWorkshopStack`, {
    env: {
        account: process.env.CDK_DEFAULT_ACCOUNT,
        region: process.env.CDK_DEFAULT_REGION
    }
});
