#!/bin/bash

profile=wsperf

cd deployments && npm i

# GitHubアカウントIDを入力してもらう
while true
do
  read -p "GitHubアカウントIDを入力してください: " gid
  if [ "$gid" != '' ]; then
    break
  fi
done

connection=$(aws apprunner list-connections --connection-name $gid --query 'ConnectionSummaryList[0].ConnectionArn' --profile $profile)

# GitHubとの接続がなかったら最初に接続してもらう
if [ "$connection" == '' ] || [ "$connection" == null ]; then
  connection=$(aws apprunner create-connection --profile $profile --connection-name $gid --provider-type GITHUB --query 'Connection.ConnectionArn')
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

# AppRunnerデプロイ
cdk deploy --profile $profile -c connection-arn="$connection" -c gh-account-id="$gid"
