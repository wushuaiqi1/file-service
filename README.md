```shell
# 提交到本地仓库
git add .
git commit -m "测试"
# 和远程建立关联
git fetch
git remote add origin https://gitee.com/odfive/laravel-service-a.git
# 合并远程分支并允许不相关的历史
git config pull.rebase false
git pull origin master --allow-unrelated-histories
# 重新提交到本地仓库
# 关联远程分支并推送到远程
git push --set-upstream origin master
git push origin master
```