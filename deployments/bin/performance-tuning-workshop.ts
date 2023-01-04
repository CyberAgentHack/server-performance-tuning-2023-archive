#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import {AppRunnerStack} from "../lib/app-runner-stack";

const app = new cdk.App();
new AppRunnerStack(app, 'PerformanceTuningWorkshopStack');