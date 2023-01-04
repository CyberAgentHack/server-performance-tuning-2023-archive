#!/bin/bash

profile=wsperf

cd deployments && npm i

connection=$(aws apprunner list-connections --query 'ConnectionSummaryList[0].ConnectionArn')

# GitHubとの接続がなかったら最初に接続してもらう
if [ "$connection" == '' ] || [ "$connection" == null ]; then
  connection=$(aws apprunner create-connection --profile $profile --connection-name wsperf-connection --provider-type GITHUB --query 'Connection.ConnectionArn')
  printf "App Runnerのコンソール(https://ap-northeast-1.console.aws.amazon.com/apprunner/home?region=ap-northeast-1#/connections)\nにアクセスして『ハンドシェイクを完了』ボタンを押し、自分のGitHubアカウントを接続してください。\n"

  # ハンドシェイクを完了してもらう
  while true
  do
    read -p "接続完了しましたか?(y/n): " yn
    case "$yn" in [yY]*) break ;; *) ;; esac
  done
fi

# エラー出るのでダブルクォーテーションを削除
connection=$(echo $connection | sed 's/"//g')

# GitHubアカウントIDを入力してもらう
while true
do
  read -p "GitHubアカウントIDを入力してください: " gid
  if [ "$gid" != '' ]; then
    break
  fi
done


# AppRunnerデプロイ
cdk deploy --profile $profile -c connection-arn="$connection" -c gh-account-id="$gid"