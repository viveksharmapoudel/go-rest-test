FROM mysql:8.0.22
#copying the my.cnf file to the custom.cnf file for arrangement of the db
COPY ./db/my.cnf /etc/mysql/conf.d/custom.cnf