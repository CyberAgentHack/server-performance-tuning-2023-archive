#!/bin/bash

profile=wsperf

cd deployments && cdk destroy --profile $profile

connection=$(aws apprunner list-connections --query 'ConnectionSummaryList[0].ConnectionArn')
# エラー出るのでダブルクォーテーションを削除
connection=$(echo $connection | sed 's/"//g')

aws apprunner delete-connection --profile $profile --connection-arn "$connection"