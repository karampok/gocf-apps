NAME = dbApp
ARGS = ""

include ../Makefile.include

cfpush: build
	cf push $(NAME) $(ARGS) -c $(OUTPUTDIR)/$(NAME) -b https://github.com/cloudfoundry/binary-buildpack.git --no-start
	cf set-env $(NAME) DB_URL "$(DB_URL)"
	cf start $(NAME)

localpush: build
	docker run -v ${PWD}/bin:/opt/bin  --env-file ./cf.env -p 4000:4000  --link mariadb:mariadb  -it cloudfoundry/cflinuxfs2 /opt/bin/$(NAME)

db-start:
	@echo  "$(OK_COLOR)==> Starting the mariadb $(NO_COLOR)"
	docker run -d --name mariadb --env-file ./mariadb.env  -p 3306:3306/tcp mariadb  2>/dev/null || echo "MariaDB is already running (make db-stop to start from scratch)"

db-stop:
	@echo  "$(OK_COLOR)==> Stoping the mariadb  $(NO_COLOR)"
	docker rm -f  mariadb || exit 0

db-client:
	@echo  "$(OK_COLOR)==> Start a client $(NO_COLOR)"
	docker run -it --env-file ./mariadb.env  --link mariadb:mysql --rm mariadb sh -c 'exec mysql -h"$$MYSQL_PORT_3306_TCP_ADDR" -P"$$MYSQL_PORT_3306_TCP_PORT" -u"$$MYSQL_USER" -p"$$MYSQL_PASSWORD"'

db-client-root:
	@echo  "$(OK_COLOR)==> Start a client $(NO_COLOR)"
	docker run -it --env-file ./mariadb.env  --link mariadb:mysql --rm mariadb sh -c 'exec mysql -h"$$MYSQL_PORT_3306_TCP_ADDR" -P"$$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$$MYSQL_ROOT_PASSWORD"'


