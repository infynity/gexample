set -x

WORKDIR=$(pwd)

source $WORKDIR/ver.bash

sysctl -w net.core.somaxconn=8096

VERSION_MYSQL_UNSTABLE=`docker ps |grep mariadb:10.4`
#备份数据存储路径
store_sql="backup_mysql.sql"

is_unstable_mysql=0
#mariadb 10.4 to 10.3
if [ -n "$VERSION_MYSQL_UNSTABLE" ];then

	while ! mysqladmin ping -h 127.0.0.1 -P 10110 --silent; do
    		echo "waiting mysql start"
    		sleep 1
	done

	echo store_sql

	mysqldump -h 127.0.0.1 -P 10110 -u root --password=wwww --all-databases>$WORKDIR/$store_sql
#	cp $WORKDIR/$store_sql $WORKDIR/`date +%s`.sql
	cp $WORKDIR/$store_sql $WORKDIR/`date "+%Y-%m-%d%H:%M:%S"`.sql
  let is_unstable_mysql=1
  docker service rm dcmcRelease_mysql
  docker stack rm dcmcRelease > /dev/null 2>&1
  docker config rm $(docker config ls -f "name=dcmc" -q) > /dev/null 2>&1

  echo  "del volume of mysql"
  while ! docker volume rm dcmcRelease_mysql; do
      echo "waiting delete mysql volume"
      sleep 1
  done
else
	docker service rm dcmcRelease_msg || true
	docker service rm dcmcRelease_cms || true
	docker service rm dcmcRelease_web || true
	docker service rm dcmcRelease_notify || true

fi




docker config rm dcmc-redis
docker config create dcmc-redis $WORKDIR/stack/configs/dcmc-redis.conf
docker config rm dcmc-nginx
docker config create dcmc-nginx $WORKDIR/stack/configs/dcmc-nginx.conf
docker config rm dcmc-notify
docker config create dcmc-notify $WORKDIR/stack/configs/dcmc-notify.yaml
docker config rm dcmc-msg
docker config create dcmc-msg $WORKDIR/stack/configs/dcmc-msg.yaml
docker config rm dcmc-cms
docker config create dcmc-cms $WORKDIR/stack/configs/dcmc-cms.yaml
docker config rm dcmc-signverify
docker config create dcmc-signverify $WORKDIR/stack/configs/dcmc-sign.yaml
docker config rm dcmc-appstore-crt
docker config create dcmc-appstore-crt $WORKDIR/stack/configs/UOS_APP_Signing_CA.crt
docker config rm dcmc-developer-crt
docker config create dcmc-developer-crt $WORKDIR/stack/configs/UOS_Developer_ID_CA.crt

docker stack deploy --prune --resolve-image=never --with-registry-auth dcmcRelease -c $WORKDIR/stack/docker-compose.yml


while ! mysqladmin ping -h 127.0.0.1 -P 10110 --silent; do
    echo "waiting mysql start"
    sleep 3
done

ps_out=`docker ps |grep mysql`
if [ -n "$ps_out" ];then
    echo "docker mysql running$ps_out"
fi

## ----- 在这里添加数据库升级脚本
MYSQL_CMD='mysql -h 127.0.0.1 -P 10110 -u root --password=wwww dcmc'

if [ $is_unstable_mysql -eq 1 ];then
    #mariadb降级到10.3时    还原10.4的数据
    $MYSQL_CMD < $WORKDIR/$store_sql
fi
$MYSQL_CMD < $WORKDIR/data/sql/migrate/0.9.3.sql
## ----- 在这里添加数据库升级脚本
## ----- 产品化一期升级脚本
#$MYSQL_CMD < $WORKDIR/data/sql/migrate/product-step1-setting.sql
## ----- 工行定制项目升级脚本
#$MYSQL_CMD < $WORKDIR/data/sql/migrate/icbc.sql

$MYSQL_CMD < $WORKDIR/data/sql/migrate/1.0.2.sql
cd $WORKDIR/data/sql/migrate/
chmod +x migration$VER_ARCH
./migration$VER_ARCH

$MYSQL_CMD < $WORKDIR/data/sql/migrate/1.0.3.sql
