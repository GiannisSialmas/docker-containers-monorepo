FROM postgres:10-alpine AS populate

ENV POSTGRES_PASSWORD=chinook
ENV POSTGRES_DB=chinook
# Changing the PGDATA so that it is not created in the /var/lib/postgresql/data dir which is a volume
# because when commiting an image, data in volumes are scraped
ENV PGDATA=/var/lib/postgresql/populated-database

COPY ./Chinook_PostgreSql.sql /docker-entrypoint-initdb.d/

# Remove the last line from the entrypoint script so that the database is not started ater initial population is done.
RUN sed -i '$ d' /usr/local/bin/docker-entrypoint.sh
# Manually execute the startup script in order to populate the database
RUN bash /docker-entrypoint.sh postgres -c 'wal_level=minimal' -c 'max_wal_senders=0' -c 'min_wal_size=5MB' -c 'max_wal_size=20MB'
# Adding auth md5 for all in order to access the database in any way
RUN sed -i 's/trust/md5/g' /var/lib/postgresql/populated-database/pg_hba.conf
# Tweak the configuration to minimize the docker image
RUN sed -i 's/#wal_level = replica/wal_level = minimal/g' /var/lib/postgresql/populated-database/postgresql.conf
RUN sed -i 's/#max_wal_senders = 10/max_wal_senders = 0/g' /var/lib/postgresql/populated-database/postgresql.conf
RUN sed -i 's/#min_wal_size = 80MB/min_wal_size = 5MB/g' /var/lib/postgresql/populated-database/postgresql.conf
RUN sed -i 's/#max_wal_size = 1GB/max_wal_size = 20MB/g' /var/lib/postgresql/populated-database/postgresql.conf
# RUN pg_archivecleanup /var/lib/postgresql/populated-database/pg_wal $(ls /var/lib/postgresql/populated-database/pg_wal | sort -r | tail -n +5 | head -n 1)

RUN chown -R root:root /var/lib/postgresql/populated-database


# Begin from scratch and take only the postgres data dir to minimize the image and remove clutter created by previous steps
FROM postgres:10-alpine
ENV POSTGRES_PASSWORD=chinook
ENV PGDATA=/var/lib/postgresql/populated-database
COPY --chown=postgres --from=populate /var/lib/postgresql/populated-database /var/lib/postgresql/populated-database