set -e
# 仮想環境の作成のために毎回実行する

dep ensure
exec "$@"
