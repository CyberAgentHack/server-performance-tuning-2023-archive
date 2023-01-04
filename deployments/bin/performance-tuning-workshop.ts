#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import {PerformanceTuningWorkshop} from "../lib/performance-tuning-workshop";

const app = new cdk.App();
new PerformanceTuningWorkshop(app, 'PerformanceTuningWorkshopStack');
