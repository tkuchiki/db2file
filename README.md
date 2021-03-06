# db2file

## Usage

```console
$ ./db2file --help
usage: db2file --dbname=DBNAME --query=QUERY --dump=DUMP [<flags>]

Database dump to file

Flags:
  --help                     Show context-sensitive help (also try --help-long and --help-man).
  --dbuser="root"            Database user
  --dbpass=DBPASS            Database password
  --dbhost="localhost"       Database host
  --dbport=3306              Database port
  --dbsock=DBSOCK            Database socket
  --dbname=DBNAME            Database name
  --query=QUERY              SQL
  --dump=DUMP                Dump file from database column
  --filename=FILENAME        Filename column
  --filename-template=FILENAME-TEMPLATE
                             Filename Go text/template syntax
  --mimetype=MIMETYPE        Mimetype column
  --auto                     Autodetect file extension
  --out-dir=$TMPDIR/db2file  Output directory
  --overwrite                Overwrite file same filename
  --version                  Show application version.
```


```console
mysql> desc image;
+-------+---------------------+------+-----+---------+----------------+
| Field | Type                | Null | Key | Default | Extra          |
+-------+---------------------+------+-----+---------+----------------+
| id    | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment | -- e.g.) 1
| name  | varchar(191)        | YES  |     | NULL    |                | -- e.g.) default.png
| data  | longblob            | YES  |     | NULL    |                |
| mime  | varchar(128)        | YES  |     | NULL    |                | -- e.g.) image/jpeg
+-------+---------------------+------+-----+---------+----------------+
4 rows in set (0.01 sec)

$ ./db2file --dbname isubata --query "SELECT * from image" --out-dir ./tmp --dump data --filename name
2018/07/17 23:52:23 [dump] tmp/default.png
2018/07/17 23:52:23 [dump] tmp/1ce0c4ff504f19f267e877a9e244d60ac0bf1a41.png
2018/07/17 23:52:23 [dump] tmp/846f4f0bde2a2103c71936091e82bc1354f11b3a.png
2018/07/17 23:52:23 [dump] tmp/8628ef0f034d734729e3a735362e6008b30bb72b.png
2018/07/17 23:52:23 [dump] tmp/d9efb5732e0ee53618bd10d2ddc5a6b33edc4751.png
2018/07/17 23:52:23 [dump] tmp/851e9e15e1d1fff39c2d182881926d107154c44a.png
...

$ ./db2file --dbname isubata --query "SELECT * from image" --out-dir ./tmp --dump data --filename-template "{{ .id }}.jpg"
2018/07/18 15:40:08 [dump] tmp/1.jpg
2018/07/18 15:40:08 [dump] tmp/2.jpg
2018/07/18 15:40:08 [dump] tmp/3.jpg
...

$ ./db2file --dbname isubata --query "SELECT * from image" --out-dir ./tmp --dump data --mimetype mime --filename-template "{{ .id }}"
2018/07/18 23:07:42 [dump] tmp/1.jpg
2018/07/18 23:07:42 [dump] tmp/2.jpg
2018/07/18 23:07:42 [dump] tmp/3.jpg
...

$ ./db2file --dbname isubata --query "SELECT * from image" --out-dir ./tmp --dump data --auto --filename-template "{{ .id }}"
2018/07/19 11:30:56 [dump] tmp/1.png
2018/07/19 11:30:56 [dump] tmp/2.png
2018/07/19 11:30:56 [dump] tmp/15.jpg
...
```
