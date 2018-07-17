# db2file

## Usage

```console
$ ./db2file --help
usage: db2file --dbname=DBNAME --query=QUERY --dump=DUMP --filename=FILENAME [<flags>]

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
  --filename=FILENAME        filename column
  --out-dir=$TMPDIR/db2file  Output directory
  --overwrite                Overwrite file same filename
  --version                  Show application version.

```


```console
mysql> desc image;
+-------+---------------------+------+-----+---------+----------------+
| Field | Type                | Null | Key | Default | Extra          |
+-------+---------------------+------+-----+---------+----------------+
| id    | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
| name  | varchar(191)        | YES  |     | NULL    |                |
| data  | longblob            | YES  |     | NULL    |                |
+-------+---------------------+------+-----+---------+----------------+
3 rows in set (0.01 sec)

$ ./db2file --dbname isubata --query "SELECT * from image" --out-dir ./tmp --dump data --filename name
2018/07/17 23:52:23 [dump] tmp/default.png
2018/07/17 23:52:23 [dump] tmp/1ce0c4ff504f19f267e877a9e244d60ac0bf1a41.png
2018/07/17 23:52:23 [dump] tmp/846f4f0bde2a2103c71936091e82bc1354f11b3a.png
2018/07/17 23:52:23 [dump] tmp/8628ef0f034d734729e3a735362e6008b30bb72b.png
2018/07/17 23:52:23 [dump] tmp/d9efb5732e0ee53618bd10d2ddc5a6b33edc4751.png
2018/07/17 23:52:23 [dump] tmp/851e9e15e1d1fff39c2d182881926d107154c44a.png
...
```
