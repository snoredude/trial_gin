# start
cd github.com\snoredude
git clone https://github.com/snoredude/trial_gin.git
cd github.com\snoredude\trial_gin
git config --local user.name "snoredude"
git config --local user.email "snoredude@aliyun.com"
go mod init github.com/snoredude/trial_gin
touch readme.md
git add .
git commit -m "feat: init project"
git branch -M main
git push -u origin main