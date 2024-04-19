package keyword

type KeyWord uint

const (
	ABORT KeyWord = iota
	ABS
	ABSOLUTE
	ACTION
	ADD
	ADMIN
	AFTER
	AGAINST
	ALL
	ALLOCATE
	ALTER
	ALWAYS
	ANALYZE
	AND
	ANTI
	ANY
	APPLY
	ARCHIVE
	ARE
	ARRAY
	ARRAY_AGG
	ARRAY_MAX_CARDINALITY
	AS
	ASC
	ASENSITIVE
	ASSERT
	ASYMMETRIC
	AT
	ATOMIC
	ATTACH
	AUTHORIZATION
	AUTO
	AUTOINCREMENT
	AUTO_INCREMENT
	AVG
	AVRO
	BACKWARD
	BASE64
	BEGIN
	BEGIN_FRAME
	BEGIN_PARTITION
	BETWEEN
	BIGDECIMAL
	BIGINT
	BIGNUMERIC
	BINARY
	BINDING
	BLOB
	BLOOMFILTER
	BOOL
	BOOLEAN
	BOTH
	BROWSE
	BTREE
	BY
	BYPASSRLS
	BYTEA
	BYTES
	CACHE
	CALL
	CALLED
	CARDINALITY
	CASCADE
	CASCADED
	CASE
	CAST
	CEIL
	CEILING
	CENTURY
	CHAIN
	CHANGE
	CHANNEL
	CHAR
	CHARACTER
	CHARACTERS
	CHARACTER_LENGTH
	CHARSET
	CHAR_LENGTH
	CHECK
	CLOB
	CLONE
	CLOSE
	CLUSTER
	COALESCE
	COLLATE
	COLLATION
	COLLECT
	COLLECTION
	COLUMN
	COLUMNS
	COMMENT
	COMMIT
	COMMITTED
	COMPRESSION
	COMPUTE
	CONCURRENTLY
	CONDITION
	CONFLICT
	CONNECT
	CONNECTION
	CONSTRAINT
	CONTAINS
	CONVERT
	COPY
	COPY_OPTIONS
	CORR
	CORRESPONDING
	COUNT
	COVAR_POP
	COVAR_SAMP
	CREATE
	CREATEDB
	CREATEROLE
	CREDENTIALS
	CROSS
	CSV
	CUBE
	CUME_DIST
	CURRENT
	CURRENT_CATALOG
	CURRENT_DATE
	CURRENT_DEFAULT_TRANSFORM_GROUP
	CURRENT_PATH
	CURRENT_ROLE
	CURRENT_ROW
	CURRENT_SCHEMA
	CURRENT_TIME
	CURRENT_TIMESTAMP
	CURRENT_TRANSFORM_GROUP_FOR_TYPE
	CURRENT_USER
	CURSOR
	CYCLE
	DATA
	DATABASE
	DATE
	DATETIME
	DAY
	DAYOFWEEK
	DAYOFYEAR
	DEALLOCATE
	DEC
	DECADE
	DECIMAL
	DECLARE
	DEFAULT
	DEFERRABLE
	DEFERRED
	DEFINED
	DELAYED
	DELETE
	DELIMITED
	DELIMITER
	DELTA
	DENSE_RANK
	DEREF
	DESC
	DESCRIBE
	DETACH
	DETAIL
	DETERMINISTIC
	DIRECTORY
	DISABLE
	DISCARD
	DISCONNECT
	DISTINCT
	DISTRIBUTE
	DIV
	DO
	DOUBLE
	DOW
	DOY
	DROP
	DRY
	DUPLICATE
	DYNAMIC
	EACH
	ELEMENT
	ELEMENTS
	ELSE
	EMPTY
	ENABLE
	ENCODING
	ENCRYPTION
	END
	END_EXEC
	ENDPOINT
	END_FRAME
	END_PARTITION
	ENFORCED
	ENGINE
	ENUM
	EPOCH
	EQUALS
	ERROR
	ESCAPE
	ESCAPED
	EVENT
	EVERY
	EXCEPT
	EXCEPTION
	EXCLUDE
	EXCLUSIVE
	EXEC
	EXECUTE
	EXISTS
	EXP
	EXPANSION
	EXPLAIN
	EXPLICIT
	EXPORT
	EXTENDED
	EXTENSION
	EXTERNAL
	EXTRACT
	FAIL
	FALSE
	FETCH
	FIELDS
	FILE
	FILES
	FILE_FORMAT
	FILTER
	FIRST
	FIRST_VALUE
	FLOAT
	FLOAT4
	FLOAT64
	FLOAT8
	FLOOR
	FLUSH
	FOLLOWING
	FOR
	FORCE
	FORCE_NOT_NULL
	FORCE_NULL
	FORCE_QUOTE
	FOREIGN
	FORMAT
	FORMATTED
	FORWARD
	FRAME_ROW
	FREE
	FREEZE
	FROM
	FSCK
	FULL
	FULLTEXT
	FUNCTION
	FUNCTIONS
	FUSION
	GENERAL
	GENERATE
	GENERATED
	GEOGRAPHY
	GET
	GLOBAL
	GRANT
	GRANTED
	GRAPHVIZ
	GROUP
	GROUPING
	GROUPS
	HASH
	HAVING
	HEADER
	HIGH_PRIORITY
	HISTORY
	HIVEVAR
	HOLD
	HOSTS
	HOUR
	HOURS
	IDENTITY
	IF
	IGNORE
	ILIKE
	IMMEDIATE
	IMMUTABLE
	IN
	INCLUDE
	INCLUDE_NULL_VALUES
	INCREMENT
	INDEX
	INDICATOR
	INHERIT
	INITIALLY
	INNER
	INOUT
	INPUT
	INPUTFORMAT
	INSENSITIVE
	INSERT
	INSTALL
	INT
	INT2
	INT4
	INT64
	INT8
	INTEGER
	INTERSECT
	INTERSECTION
	INTERVAL
	INTO
	IS
	ISODOW
	ISOLATION
	ISOWEEK
	ISOYEAR
	ITEMS
	JAR
	JOIN
	JSON
	JSONB
	JSONFILE
	JSON_TABLE
	JULIAN
	KEY
	KEYS
	KILL
	LAG
	LANGUAGE
	LARGE
	LAST
	LAST_VALUE
	LATERAL
	LEAD
	LEADING
	LEFT
	LEVEL
	LIKE
	LIKE_REGEX
	LIMIT
	LINES
	LISTAGG
	LN
	LOAD
	LOCAL
	LOCALTIME
	LOCALTIMESTAMP
	LOCATION
	LOCK
	LOCKED
	LOGIN
	LOGS
	LOWER
	LOW_PRIORITY
	MACRO
	MANAGEDLOCATION
	MAP
	MATCH
	MATCHED
	MATERIALIZED
	MAX
	MAXVALUE
	MEDIUMINT
	MEMBER
	MERGE
	METADATA
	METHOD
	MICROSECOND
	MICROSECONDS
	MILLENIUM
	MILLENNIUM
	MILLISECOND
	MILLISECONDS
	MIN
	MINUTE
	MINVALUE
	MOD
	MODE
	MODIFIES
	MODULE
	MONTH
	MSCK
	MULTISET
	MUTATION
	NAME
	NANOSECOND
	NANOSECONDS
	NATIONAL
	NATURAL
	NCHAR
	NCLOB
	NEW
	NEXT
	NO
	NOBYPASSRLS
	NOCREATEDB
	NOCREATEROLE
	NOINHERIT
	NOLOGIN
	NONE
	NOREPLICATION
	NORMALIZE
	NOSCAN
	NOSUPERUSER
	NOT
	NOTHING
	NOWAIT
	NO_WRITE_TO_BINLOG
	NTH_VALUE
	NTILE
	NULL
	NULLIF
	NULLS
	NUMERIC
	NVARCHAR
	OBJECT
	OCCURRENCES_REGEX
	OCTETS
	OCTET_LENGTH
	OF
	OFFSET
	OLD
	ON
	ONLY
	OPEN
	OPERATOR
	OPTIMIZE
	OPTIMIZER_COSTS
	OPTION
	OPTIONS
	OR
	ORC
	ORDER
	OUT
	OUTER
	OUTPUTFORMAT
	OVER
	OVERFLOW
	OVERLAPS
	OVERLAY
	OVERWRITE
	OWNED
	PARALLEL
	PARAMETER
	PARQUET
	PARTITION
	PARTITIONED
	PARTITIONS
	PASSWORD
	PATH
	PATTERN
	PERCENT
	PERCENTILE_CONT
	PERCENTILE_DISC
	PERCENT_RANK
	PERIOD
	PERSISTENT
	PIVOT
	PLACING
	PLANS
	PORTION
	POSITION
	POSITION_REGEX
	POWER
	PRAGMA
	PRECEDES
	PRECEDING
	PRECISION
	PREPARE
	PRESERVE
	PRIMARY
	PRIOR
	PRIVILEGES
	PROCEDURE
	PROGRAM
	PURGE
	QUALIFY
	QUARTER
	QUERY
	QUOTE
	RANGE
	RANK
	RAW
	RCFILE
	READ
	READS
	READ_ONLY
	REAL
	RECURSIVE
	REF
	REFERENCES
	REFERENCING
	REGCLASS
	REGEXP
	REGR_AVGX
	REGR_AVGY
	REGR_COUNT
	REGR_INTERCEPT
	REGR_R2
	REGR_SLOPE
	REGR_SXX
	REGR_SXY
	REGR_SYY
	RELATIVE
	RELAY
	RELEASE
	RENAME
	REORG
	REPAIR
	REPEATABLE
	REPLACE
	REPLICA
	REPLICATION
	RESET
	RESPECT
	RESTART
	RESTRICT
	RESTRICTED
	RESULT
	RESULTSET
	RETAIN
	RETURN
	RETURNING
	RETURNS
	REVOKE
	RIGHT
	RLIKE
	ROLE
	ROLLBACK
	ROLLUP
	ROOT
	ROW
	ROWID
	ROWS
	ROW_NUMBER
	RULE
	RUN
	SAFE
	SAFE_CAST
	SAVEPOINT
	SCHEMA
	SCOPE
	SCROLL
	SEARCH
	SECOND
	SECRET
	SECURITY
	SELECT
	SEMI
	SENSITIVE
	SEQUENCE
	SEQUENCEFILE
	SEQUENCES
	SERDE
	SERDEPROPERTIES
	SERIALIZABLE
	SESSION
	SESSION_USER
	SET
	SETS
	SHARE
	SHOW
	SIMILAR
	SKIP
	SLOW
	SMALLINT
	SNAPSHOT
	SOME
	SORT
	SPATIAL
	SPECIFIC
	SPECIFICTYPE
	SQL
	SQLEXCEPTION
	SQLSTATE
	SQLWARNING
	SQRT
	STABLE
	STAGE
	START
	STATIC
	STATISTICS
	STATUS
	STDDEV_POP
	STDDEV_SAMP
	STDIN
	STDOUT
	STORAGE_INTEGRATION
	STORED
	STRICT
	STRING
	STRUCT
	SUBMULTISET
	SUBSTRING
	SUBSTRING_REGEX
	SUCCEEDS
	SUM
	SUPER
	SUPERUSER
	SWAP
	SYMMETRIC
	SYNC
	SYSTEM
	SYSTEM_TIME
	SYSTEM_USER
	TABLE
	TABLES
	TABLESAMPLE
	TBLPROPERTIES
	TEMP
	TEMPORARY
	TERMINATED
	TEXT
	TEXTFILE
	THEN
	TIES
	TIME
	TIMESTAMP
	TIMESTAMPTZ
	TIMETZ
	TIMEZONE
	TIMEZONE_ABBR
	TIMEZONE_HOUR
	TIMEZONE_MINUTE
	TIMEZONE_REGION
	TINYINT
	TO
	TOP
	TRAILING
	TRANSACTION
	TRANSIENT
	TRANSLATE
	TRANSLATE_REGEX
	TRANSLATION
	TREAT
	TRIGGER
	TRIM
	TRIM_ARRAY
	TRUE
	TRUNCATE
	TRY_CAST
	TYPE
	UESCAPE
	UNBOUNDED
	UNCACHE
	UNCOMMITTED
	UNION
	UNIQUE
	UNKNOWN
	UNLOAD
	UNLOCK
	UNLOGGED
	UNNEST
	UNPIVOT
	UNSAFE
	UNSIGNED
	UNTIL
	UPDATE
	UPPER
	URL
	USAGE
	USE
	USER
	USER_RESOURCES
	USING
	UUID
	VACUUM
	VALID
	VALIDATION_MODE
	VALUE
	VALUES
	VALUE_OF
	VARBINARY
	VARCHAR
	VARIABLES
	VARYING
	VAR_POP
	VAR_SAMP
	VERBOSE
	VERSION
	VERSIONING
	VIEW
	VIRTUAL
	VOLATILE
	WEEK
	WHEN
	WHENEVER
	WHERE
	WIDTH_BUCKET
	WINDOW
	WITH
	WITHIN
	WITHOUT
	WITHOUT_ARRAY_WRAPPER
	WORK
	WRITE
	XML
	XOR
	YEAR
	ZONE
	ZORDER
)

var stringToKeyWord = map[string]KeyWord{
	"ABORT":                            ABORT,
	"ABS":                              ABS,
	"ABSOLUTE":                         ABSOLUTE,
	"ACTION":                           ACTION,
	"ADD":                              ADD,
	"ADMIN":                            ADMIN,
	"AFTER":                            AFTER,
	"AGAINST":                          AGAINST,
	"ALL":                              ALL,
	"ALLOCATE":                         ALLOCATE,
	"ALTER":                            ALTER,
	"ALWAYS":                           ALWAYS,
	"ANALYZE":                          ANALYZE,
	"AND":                              AND,
	"ANTI":                             ANTI,
	"ANY":                              ANY,
	"APPLY":                            APPLY,
	"ARCHIVE":                          ARCHIVE,
	"ARE":                              ARE,
	"ARRAY":                            ARRAY,
	"ARRAY_AGG":                        ARRAY_AGG,
	"ARRAY_MAX_CARDINALITY":            ARRAY_MAX_CARDINALITY,
	"AS":                               AS,
	"ASC":                              ASC,
	"ASENSITIVE":                       ASENSITIVE,
	"ASSERT":                           ASSERT,
	"ASYMMETRIC":                       ASYMMETRIC,
	"AT":                               AT,
	"ATOMIC":                           ATOMIC,
	"ATTACH":                           ATTACH,
	"AUTHORIZATION":                    AUTHORIZATION,
	"AUTO":                             AUTO,
	"AUTOINCREMENT":                    AUTOINCREMENT,
	"AUTO_INCREMENT":                   AUTO_INCREMENT,
	"AVG":                              AVG,
	"AVRO":                             AVRO,
	"BACKWARD":                         BACKWARD,
	"BASE64":                           BASE64,
	"BEGIN":                            BEGIN,
	"BEGIN_FRAME":                      BEGIN_FRAME,
	"BEGIN_PARTITION":                  BEGIN_PARTITION,
	"BETWEEN":                          BETWEEN,
	"BIGDECIMAL":                       BIGDECIMAL,
	"BIGINT":                           BIGINT,
	"BIGNUMERIC":                       BIGNUMERIC,
	"BINARY":                           BINARY,
	"BINDING":                          BINDING,
	"BLOB":                             BLOB,
	"BLOOMFILTER":                      BLOOMFILTER,
	"BOOL":                             BOOL,
	"BOOLEAN":                          BOOLEAN,
	"BOTH":                             BOTH,
	"BROWSE":                           BROWSE,
	"BTREE":                            BTREE,
	"BY":                               BY,
	"BYPASSRLS":                        BYPASSRLS,
	"BYTEA":                            BYTEA,
	"BYTES":                            BYTES,
	"CACHE":                            CACHE,
	"CALL":                             CALL,
	"CALLED":                           CALLED,
	"CARDINALITY":                      CARDINALITY,
	"CASCADE":                          CASCADE,
	"CASCADED":                         CASCADED,
	"CASE":                             CASE,
	"CAST":                             CAST,
	"CEIL":                             CEIL,
	"CEILING":                          CEILING,
	"CENTURY":                          CENTURY,
	"CHAIN":                            CHAIN,
	"CHANGE":                           CHANGE,
	"CHANNEL":                          CHANNEL,
	"CHAR":                             CHAR,
	"CHARACTER":                        CHARACTER,
	"CHARACTERS":                       CHARACTERS,
	"CHARACTER_LENGTH":                 CHARACTER_LENGTH,
	"CHARSET":                          CHARSET,
	"CHAR_LENGTH":                      CHAR_LENGTH,
	"CHECK":                            CHECK,
	"CLOB":                             CLOB,
	"CLONE":                            CLONE,
	"CLOSE":                            CLOSE,
	"CLUSTER":                          CLUSTER,
	"COALESCE":                         COALESCE,
	"COLLATE":                          COLLATE,
	"COLLATION":                        COLLATION,
	"COLLECT":                          COLLECT,
	"COLLECTION":                       COLLECTION,
	"COLUMN":                           COLUMN,
	"COLUMNS":                          COLUMNS,
	"COMMENT":                          COMMENT,
	"COMMIT":                           COMMIT,
	"COMMITTED":                        COMMITTED,
	"COMPRESSION":                      COMPRESSION,
	"COMPUTE":                          COMPUTE,
	"CONCURRENTLY":                     CONCURRENTLY,
	"CONDITION":                        CONDITION,
	"CONFLICT":                         CONFLICT,
	"CONNECT":                          CONNECT,
	"CONNECTION":                       CONNECTION,
	"CONSTRAINT":                       CONSTRAINT,
	"CONTAINS":                         CONTAINS,
	"CONVERT":                          CONVERT,
	"COPY":                             COPY,
	"COPY_OPTIONS":                     COPY_OPTIONS,
	"CORR":                             CORR,
	"CORRESPONDING":                    CORRESPONDING,
	"COUNT":                            COUNT,
	"COVAR_POP":                        COVAR_POP,
	"COVAR_SAMP":                       COVAR_SAMP,
	"CREATE":                           CREATE,
	"CREATEDB":                         CREATEDB,
	"CREATEROLE":                       CREATEROLE,
	"CREDENTIALS":                      CREDENTIALS,
	"CROSS":                            CROSS,
	"CSV":                              CSV,
	"CUBE":                             CUBE,
	"CUME_DIST":                        CUME_DIST,
	"CURRENT":                          CURRENT,
	"CURRENT_CATALOG":                  CURRENT_CATALOG,
	"CURRENT_DATE":                     CURRENT_DATE,
	"CURRENT_DEFAULT_TRANSFORM_GROUP":  CURRENT_DEFAULT_TRANSFORM_GROUP,
	"CURRENT_PATH":                     CURRENT_PATH,
	"CURRENT_ROLE":                     CURRENT_ROLE,
	"CURRENT_ROW":                      CURRENT_ROW,
	"CURRENT_SCHEMA":                   CURRENT_SCHEMA,
	"CURRENT_TIME":                     CURRENT_TIME,
	"CURRENT_TIMESTAMP":                CURRENT_TIMESTAMP,
	"CURRENT_TRANSFORM_GROUP_FOR_TYPE": CURRENT_TRANSFORM_GROUP_FOR_TYPE,
	"CURRENT_USER":                     CURRENT_USER,
	"CURSOR":                           CURSOR,
	"CYCLE":                            CYCLE,
	"DATA":                             DATA,
	"DATABASE":                         DATABASE,
	"DATE":                             DATE,
	"DATETIME":                         DATETIME,
	"DAY":                              DAY,
	"DAYOFWEEK":                        DAYOFWEEK,
	"DAYOFYEAR":                        DAYOFYEAR,
	"DEALLOCATE":                       DEALLOCATE,
	"DEC":                              DEC,
	"DECADE":                           DECADE,
	"DECIMAL":                          DECIMAL,
	"DECLARE":                          DECLARE,
	"DEFAULT":                          DEFAULT,
	"DEFERRABLE":                       DEFERRABLE,
	"DEFERRED":                         DEFERRED,
	"DEFINED":                          DEFINED,
	"DELAYED":                          DELAYED,
	"DELETE":                           DELETE,
	"DELIMITED":                        DELIMITED,
	"DELIMITER":                        DELIMITER,
	"DELTA":                            DELTA,
	"DENSE_RANK":                       DENSE_RANK,
	"DEREF":                            DEREF,
	"DESC":                             DESC,
	"DESCRIBE":                         DESCRIBE,
	"DETACH":                           DETACH,
	"DETAIL":                           DETAIL,
	"DETERMINISTIC":                    DETERMINISTIC,
	"DIRECTORY":                        DIRECTORY,
	"DISABLE":                          DISABLE,
	"DISCARD":                          DISCARD,
	"DISCONNECT":                       DISCONNECT,
	"DISTINCT":                         DISTINCT,
	"DISTRIBUTE":                       DISTRIBUTE,
	"DIV":                              DIV,
	"DO":                               DO,
	"DOUBLE":                           DOUBLE,
	"DOW":                              DOW,
	"DOY":                              DOY,
	"DROP":                             DROP,
	"DRY":                              DRY,
	"DUPLICATE":                        DUPLICATE,
	"DYNAMIC":                          DYNAMIC,
	"EACH":                             EACH,
	"ELEMENT":                          ELEMENT,
	"ELEMENTS":                         ELEMENTS,
	"ELSE":                             ELSE,
	"EMPTY":                            EMPTY,
	"ENABLE":                           ENABLE,
	"ENCODING":                         ENCODING,
	"ENCRYPTION":                       ENCRYPTION,
	"END":                              END,
	"END_EXEC":                         END_EXEC,
	"ENDPOINT":                         ENDPOINT,
	"END_FRAME":                        END_FRAME,
	"END_PARTITION":                    END_PARTITION,
	"ENFORCED":                         ENFORCED,
	"ENGINE":                           ENGINE,
	"ENUM":                             ENUM,
	"EPOCH":                            EPOCH,
	"EQUALS":                           EQUALS,
	"ERROR":                            ERROR,
	"ESCAPE":                           ESCAPE,
	"ESCAPED":                          ESCAPED,
	"EVENT":                            EVENT,
	"EVERY":                            EVERY,
	"EXCEPT":                           EXCEPT,
	"EXCEPTION":                        EXCEPTION,
	"EXCLUDE":                          EXCLUDE,
	"EXCLUSIVE":                        EXCLUSIVE,
	"EXEC":                             EXEC,
	"EXECUTE":                          EXECUTE,
	"EXISTS":                           EXISTS,
	"EXP":                              EXP,
	"EXPANSION":                        EXPANSION,
	"EXPLAIN":                          EXPLAIN,
	"EXPLICIT":                         EXPLICIT,
	"EXPORT":                           EXPORT,
	"EXTENDED":                         EXTENDED,
	"EXTENSION":                        EXTENSION,
	"EXTERNAL":                         EXTERNAL,
	"EXTRACT":                          EXTRACT,
	"FAIL":                             FAIL,
	"FALSE":                            FALSE,
	"FETCH":                            FETCH,
	"FIELDS":                           FIELDS,
	"FILE":                             FILE,
	"FILES":                            FILES,
	"FILE_FORMAT":                      FILE_FORMAT,
	"FILTER":                           FILTER,
	"FIRST":                            FIRST,
	"FIRST_VALUE":                      FIRST_VALUE,
	"FLOAT":                            FLOAT,
	"FLOAT4":                           FLOAT4,
	"FLOAT64":                          FLOAT64,
	"FLOAT8":                           FLOAT8,
	"FLOOR":                            FLOOR,
	"FLUSH":                            FLUSH,
	"FOLLOWING":                        FOLLOWING,
	"FOR":                              FOR,
	"FORCE":                            FORCE,
	"FORCE_NOT_NULL":                   FORCE_NOT_NULL,
	"FORCE_NULL":                       FORCE_NULL,
	"FORCE_QUOTE":                      FORCE_QUOTE,
	"FOREIGN":                          FOREIGN,
	"FORMAT":                           FORMAT,
	"FORMATTED":                        FORMATTED,
	"FORWARD":                          FORWARD,
	"FRAME_ROW":                        FRAME_ROW,
	"FREE":                             FREE,
	"FREEZE":                           FREEZE,
	"FROM":                             FROM,
	"FSCK":                             FSCK,
	"FULL":                             FULL,
	"FULLTEXT":                         FULLTEXT,
	"FUNCTION":                         FUNCTION,
	"FUNCTIONS":                        FUNCTIONS,
	"FUSION":                           FUSION,
	"GENERAL":                          GENERAL,
	"GENERATE":                         GENERATE,
	"GENERATED":                        GENERATED,
	"GEOGRAPHY":                        GEOGRAPHY,
	"GET":                              GET,
	"GLOBAL":                           GLOBAL,
	"GRANT":                            GRANT,
	"GRANTED":                          GRANTED,
	"GRAPHVIZ":                         GRAPHVIZ,
	"GROUP":                            GROUP,
	"GROUPING":                         GROUPING,
	"GROUPS":                           GROUPS,
	"HASH":                             HASH,
	"HAVING":                           HAVING,
	"HEADER":                           HEADER,
	"HIGH_PRIORITY":                    HIGH_PRIORITY,
	"HISTORY":                          HISTORY,
	"HIVEVAR":                          HIVEVAR,
	"HOLD":                             HOLD,
	"HOSTS":                            HOSTS,
	"HOUR":                             HOUR,
	"HOURS":                            HOURS,
	"IDENTITY":                         IDENTITY,
	"IF":                               IF,
	"IGNORE":                           IGNORE,
	"ILIKE":                            ILIKE,
	"IMMEDIATE":                        IMMEDIATE,
	"IMMUTABLE":                        IMMUTABLE,
	"IN":                               IN,
	"INCLUDE":                          INCLUDE,
	"INCLUDE_NULL_VALUES":              INCLUDE_NULL_VALUES,
	"INCREMENT":                        INCREMENT,
	"INDEX":                            INDEX,
	"INDICATOR":                        INDICATOR,
	"INHERIT":                          INHERIT,
	"INITIALLY":                        INITIALLY,
	"INNER":                            INNER,
	"INOUT":                            INOUT,
	"INPUT":                            INPUT,
	"INPUTFORMAT":                      INPUTFORMAT,
	"INSENSITIVE":                      INSENSITIVE,
	"INSERT":                           INSERT,
	"INSTALL":                          INSTALL,
	"INT":                              INT,
	"INT2":                             INT2,
	"INT4":                             INT4,
	"INT64":                            INT64,
	"INT8":                             INT8,
	"INTEGER":                          INTEGER,
	"INTERSECT":                        INTERSECT,
	"INTERSECTION":                     INTERSECTION,
	"INTERVAL":                         INTERVAL,
	"INTO":                             INTO,
	"IS":                               IS,
	"ISODOW":                           ISODOW,
	"ISOLATION":                        ISOLATION,
	"ISOWEEK":                          ISOWEEK,
	"ISOYEAR":                          ISOYEAR,
	"ITEMS":                            ITEMS,
	"JAR":                              JAR,
	"JOIN":                             JOIN,
	"JSON":                             JSON,
	"JSONB":                            JSONB,
	"JSONFILE":                         JSONFILE,
	"JSON_TABLE":                       JSON_TABLE,
	"JULIAN":                           JULIAN,
	"KEY":                              KEY,
	"KEYS":                             KEYS,
	"KILL":                             KILL,
	"LAG":                              LAG,
	"LANGUAGE":                         LANGUAGE,
	"LARGE":                            LARGE,
	"LAST":                             LAST,
	"LAST_VALUE":                       LAST_VALUE,
	"LATERAL":                          LATERAL,
	"LEAD":                             LEAD,
	"LEADING":                          LEADING,
	"LEFT":                             LEFT,
	"LEVEL":                            LEVEL,
	"LIKE":                             LIKE,
	"LIKE_REGEX":                       LIKE_REGEX,
	"LIMIT":                            LIMIT,
	"LINES":                            LINES,
	"LISTAGG":                          LISTAGG,
	"LN":                               LN,
	"LOAD":                             LOAD,
	"LOCAL":                            LOCAL,
	"LOCALTIME":                        LOCALTIME,
	"LOCALTIMESTAMP":                   LOCALTIMESTAMP,
	"LOCATION":                         LOCATION,
	"LOCK":                             LOCK,
	"LOCKED":                           LOCKED,
	"LOGIN":                            LOGIN,
	"LOGS":                             LOGS,
	"LOWER":                            LOWER,
	"LOW_PRIORITY":                     LOW_PRIORITY,
	"MACRO":                            MACRO,
	"MANAGEDLOCATION":                  MANAGEDLOCATION,
	"MAP":                              MAP,
	"MATCH":                            MATCH,
	"MATCHED":                          MATCHED,
	"MATERIALIZED":                     MATERIALIZED,
	"MAX":                              MAX,
	"MAXVALUE":                         MAXVALUE,
	"MEDIUMINT":                        MEDIUMINT,
	"MEMBER":                           MEMBER,
	"MERGE":                            MERGE,
	"METADATA":                         METADATA,
	"METHOD":                           METHOD,
	"MICROSECOND":                      MICROSECOND,
	"MICROSECONDS":                     MICROSECONDS,
	"MILLENIUM":                        MILLENIUM,
	"MILLENNIUM":                       MILLENNIUM,
	"MILLISECOND":                      MILLISECOND,
	"MILLISECONDS":                     MILLISECONDS,
	"MIN":                              MIN,
	"MINUTE":                           MINUTE,
	"MINVALUE":                         MINVALUE,
	"MOD":                              MOD,
	"MODE":                             MODE,
	"MODIFIES":                         MODIFIES,
	"MODULE":                           MODULE,
	"MONTH":                            MONTH,
	"MSCK":                             MSCK,
	"MULTISET":                         MULTISET,
	"MUTATION":                         MUTATION,
	"NAME":                             NAME,
	"NANOSECOND":                       NANOSECOND,
	"NANOSECONDS":                      NANOSECONDS,
	"NATIONAL":                         NATIONAL,
	"NATURAL":                          NATURAL,
	"NCHAR":                            NCHAR,
	"NCLOB":                            NCLOB,
	"NEW":                              NEW,
	"NEXT":                             NEXT,
	"NO":                               NO,
	"NOBYPASSRLS":                      NOBYPASSRLS,
	"NOCREATEDB":                       NOCREATEDB,
	"NOCREATEROLE":                     NOCREATEROLE,
	"NOINHERIT":                        NOINHERIT,
	"NOLOGIN":                          NOLOGIN,
	"NONE":                             NONE,
	"NOREPLICATION":                    NOREPLICATION,
	"NORMALIZE":                        NORMALIZE,
	"NOSCAN":                           NOSCAN,
	"NOSUPERUSER":                      NOSUPERUSER,
	"NOT":                              NOT,
	"NOTHING":                          NOTHING,
	"NOWAIT":                           NOWAIT,
	"NO_WRITE_TO_BINLOG":               NO_WRITE_TO_BINLOG,
	"NTH_VALUE":                        NTH_VALUE,
	"NTILE":                            NTILE,
	"NULL":                             NULL,
	"NULLIF":                           NULLIF,
	"NULLS":                            NULLS,
	"NUMERIC":                          NUMERIC,
	"NVARCHAR":                         NVARCHAR,
	"OBJECT":                           OBJECT,
	"OCCURRENCES_REGEX":                OCCURRENCES_REGEX,
	"OCTETS":                           OCTETS,
	"OCTET_LENGTH":                     OCTET_LENGTH,
	"OF":                               OF,
	"OFFSET":                           OFFSET,
	"OLD":                              OLD,
	"ON":                               ON,
	"ONLY":                             ONLY,
	"OPEN":                             OPEN,
	"OPERATOR":                         OPERATOR,
	"OPTIMIZE":                         OPTIMIZE,
	"OPTIMIZER_COSTS":                  OPTIMIZER_COSTS,
	"OPTION":                           OPTION,
	"OPTIONS":                          OPTIONS,
	"OR":                               OR,
	"ORC":                              ORC,
	"ORDER":                            ORDER,
	"OUT":                              OUT,
	"OUTER":                            OUTER,
	"OUTPUTFORMAT":                     OUTPUTFORMAT,
	"OVER":                             OVER,
	"OVERFLOW":                         OVERFLOW,
	"OVERLAPS":                         OVERLAPS,
	"OVERLAY":                          OVERLAY,
	"OVERWRITE":                        OVERWRITE,
	"OWNED":                            OWNED,
	"PARALLEL":                         PARALLEL,
	"PARAMETER":                        PARAMETER,
	"PARQUET":                          PARQUET,
	"PARTITION":                        PARTITION,
	"PARTITIONED":                      PARTITIONED,
	"PARTITIONS":                       PARTITIONS,
	"PASSWORD":                         PASSWORD,
	"PATH":                             PATH,
	"PATTERN":                          PATTERN,
	"PERCENT":                          PERCENT,
	"PERCENTILE_CONT":                  PERCENTILE_CONT,
	"PERCENTILE_DISC":                  PERCENTILE_DISC,
	"PERCENT_RANK":                     PERCENT_RANK,
	"PERIOD":                           PERIOD,
	"PERSISTENT":                       PERSISTENT,
	"PIVOT":                            PIVOT,
	"PLACING":                          PLACING,
	"PLANS":                            PLANS,
	"PORTION":                          PORTION,
	"POSITION":                         POSITION,
	"POSITION_REGEX":                   POSITION_REGEX,
	"POWER":                            POWER,
	"PRAGMA":                           PRAGMA,
	"PRECEDES":                         PRECEDES,
	"PRECEDING":                        PRECEDING,
	"PRECISION":                        PRECISION,
	"PREPARE":                          PREPARE,
	"PRESERVE":                         PRESERVE,
	"PRIMARY":                          PRIMARY,
	"PRIOR":                            PRIOR,
	"PRIVILEGES":                       PRIVILEGES,
	"PROCEDURE":                        PROCEDURE,
	"PROGRAM":                          PROGRAM,
	"PURGE":                            PURGE,
	"QUALIFY":                          QUALIFY,
	"QUARTER":                          QUARTER,
	"QUERY":                            QUERY,
	"QUOTE":                            QUOTE,
	"RANGE":                            RANGE,
	"RANK":                             RANK,
	"RAW":                              RAW,
	"RCFILE":                           RCFILE,
	"READ":                             READ,
	"READS":                            READS,
	"READ_ONLY":                        READ_ONLY,
	"REAL":                             REAL,
	"RECURSIVE":                        RECURSIVE,
	"REF":                              REF,
	"REFERENCES":                       REFERENCES,
	"REFERENCING":                      REFERENCING,
	"REGCLASS":                         REGCLASS,
	"REGEXP":                           REGEXP,
	"REGR_AVGX":                        REGR_AVGX,
	"REGR_AVGY":                        REGR_AVGY,
	"REGR_COUNT":                       REGR_COUNT,
	"REGR_INTERCEPT":                   REGR_INTERCEPT,
	"REGR_R2":                          REGR_R2,
	"REGR_SLOPE":                       REGR_SLOPE,
	"REGR_SXX":                         REGR_SXX,
	"REGR_SXY":                         REGR_SXY,
	"REGR_SYY":                         REGR_SYY,
	"RELATIVE":                         RELATIVE,
	"RELAY":                            RELAY,
	"RELEASE":                          RELEASE,
	"RENAME":                           RENAME,
	"REORG":                            REORG,
	"REPAIR":                           REPAIR,
	"REPEATABLE":                       REPEATABLE,
	"REPLACE":                          REPLACE,
	"REPLICA":                          REPLICA,
	"REPLICATION":                      REPLICATION,
	"RESET":                            RESET,
	"RESPECT":                          RESPECT,
	"RESTART":                          RESTART,
	"RESTRICT":                         RESTRICT,
	"RESTRICTED":                       RESTRICTED,
	"RESULT":                           RESULT,
	"RESULTSET":                        RESULTSET,
	"RETAIN":                           RETAIN,
	"RETURN":                           RETURN,
	"RETURNING":                        RETURNING,
	"RETURNS":                          RETURNS,
	"REVOKE":                           REVOKE,
	"RIGHT":                            RIGHT,
	"RLIKE":                            RLIKE,
	"ROLE":                             ROLE,
	"ROLLBACK":                         ROLLBACK,
	"ROLLUP":                           ROLLUP,
	"ROOT":                             ROOT,
	"ROW":                              ROW,
	"ROWID":                            ROWID,
	"ROWS":                             ROWS,
	"ROW_NUMBER":                       ROW_NUMBER,
	"RULE":                             RULE,
	"RUN":                              RUN,
	"SAFE":                             SAFE,
	"SAFE_CAST":                        SAFE_CAST,
	"SAVEPOINT":                        SAVEPOINT,
	"SCHEMA":                           SCHEMA,
	"SCOPE":                            SCOPE,
	"SCROLL":                           SCROLL,
	"SEARCH":                           SEARCH,
	"SECOND":                           SECOND,
	"SECRET":                           SECRET,
	"SECURITY":                         SECURITY,
	"SELECT":                           SELECT,
	"SEMI":                             SEMI,
	"SENSITIVE":                        SENSITIVE,
	"SEQUENCE":                         SEQUENCE,
	"SEQUENCEFILE":                     SEQUENCEFILE,
	"SEQUENCES":                        SEQUENCES,
	"SERDE":                            SERDE,
	"SERDEPROPERTIES":                  SERDEPROPERTIES,
	"SERIALIZABLE":                     SERIALIZABLE,
	"SESSION":                          SESSION,
	"SESSION_USER":                     SESSION_USER,
	"SET":                              SET,
	"SETS":                             SETS,
	"SHARE":                            SHARE,
	"SHOW":                             SHOW,
	"SIMILAR":                          SIMILAR,
	"SKIP":                             SKIP,
	"SLOW":                             SLOW,
	"SMALLINT":                         SMALLINT,
	"SNAPSHOT":                         SNAPSHOT,
	"SOME":                             SOME,
	"SORT":                             SORT,
	"SPATIAL":                          SPATIAL,
	"SPECIFIC":                         SPECIFIC,
	"SPECIFICTYPE":                     SPECIFICTYPE,
	"SQL":                              SQL,
	"SQLEXCEPTION":                     SQLEXCEPTION,
	"SQLSTATE":                         SQLSTATE,
	"SQLWARNING":                       SQLWARNING,
	"SQRT":                             SQRT,
	"STABLE":                           STABLE,
	"STAGE":                            STAGE,
	"START":                            START,
	"STATIC":                           STATIC,
	"STATISTICS":                       STATISTICS,
	"STATUS":                           STATUS,
	"STDDEV_POP":                       STDDEV_POP,
	"STDDEV_SAMP":                      STDDEV_SAMP,
	"STDIN":                            STDIN,
	"STDOUT":                           STDOUT,
	"STORAGE_INTEGRATION":              STORAGE_INTEGRATION,
	"STORED":                           STORED,
	"STRICT":                           STRICT,
	"STRING":                           STRING,
	"STRUCT":                           STRUCT,
	"SUBMULTISET":                      SUBMULTISET,
	"SUBSTRING":                        SUBSTRING,
	"SUBSTRING_REGEX":                  SUBSTRING_REGEX,
	"SUCCEEDS":                         SUCCEEDS,
	"SUM":                              SUM,
	"SUPER":                            SUPER,
	"SUPERUSER":                        SUPERUSER,
	"SWAP":                             SWAP,
	"SYMMETRIC":                        SYMMETRIC,
	"SYNC":                             SYNC,
	"SYSTEM":                           SYSTEM,
	"SYSTEM_TIME":                      SYSTEM_TIME,
	"SYSTEM_USER":                      SYSTEM_USER,
	"TABLE":                            TABLE,
	"TABLES":                           TABLES,
	"TABLESAMPLE":                      TABLESAMPLE,
	"TBLPROPERTIES":                    TBLPROPERTIES,
	"TEMP":                             TEMP,
	"TEMPORARY":                        TEMPORARY,
	"TERMINATED":                       TERMINATED,
	"TEXT":                             TEXT,
	"TEXTFILE":                         TEXTFILE,
	"THEN":                             THEN,
	"TIES":                             TIES,
	"TIME":                             TIME,
	"TIMESTAMP":                        TIMESTAMP,
	"TIMESTAMPTZ":                      TIMESTAMPTZ,
	"TIMETZ":                           TIMETZ,
	"TIMEZONE":                         TIMEZONE,
	"TIMEZONE_ABBR":                    TIMEZONE_ABBR,
	"TIMEZONE_HOUR":                    TIMEZONE_HOUR,
	"TIMEZONE_MINUTE":                  TIMEZONE_MINUTE,
	"TIMEZONE_REGION":                  TIMEZONE_REGION,
	"TINYINT":                          TINYINT,
	"TO":                               TO,
	"TOP":                              TOP,
	"TRAILING":                         TRAILING,
	"TRANSACTION":                      TRANSACTION,
	"TRANSIENT":                        TRANSIENT,
	"TRANSLATE":                        TRANSLATE,
	"TRANSLATE_REGEX":                  TRANSLATE_REGEX,
	"TRANSLATION":                      TRANSLATION,
	"TREAT":                            TREAT,
	"TRIGGER":                          TRIGGER,
	"TRIM":                             TRIM,
	"TRIM_ARRAY":                       TRIM_ARRAY,
	"TRUE":                             TRUE,
	"TRUNCATE":                         TRUNCATE,
	"TRY_CAST":                         TRY_CAST,
	"TYPE":                             TYPE,
	"UESCAPE":                          UESCAPE,
	"UNBOUNDED":                        UNBOUNDED,
	"UNCACHE":                          UNCACHE,
	"UNCOMMITTED":                      UNCOMMITTED,
	"UNION":                            UNION,
	"UNIQUE":                           UNIQUE,
	"UNKNOWN":                          UNKNOWN,
	"UNLOAD":                           UNLOAD,
	"UNLOCK":                           UNLOCK,
	"UNLOGGED":                         UNLOGGED,
	"UNNEST":                           UNNEST,
	"UNPIVOT":                          UNPIVOT,
	"UNSAFE":                           UNSAFE,
	"UNSIGNED":                         UNSIGNED,
	"UNTIL":                            UNTIL,
	"UPDATE":                           UPDATE,
	"UPPER":                            UPPER,
	"URL":                              URL,
	"USAGE":                            USAGE,
	"USE":                              USE,
	"USER":                             USER,
	"USER_RESOURCES":                   USER_RESOURCES,
	"USING":                            USING,
	"UUID":                             UUID,
	"VACUUM":                           VACUUM,
	"VALID":                            VALID,
	"VALIDATION_MODE":                  VALIDATION_MODE,
	"VALUE":                            VALUE,
	"VALUES":                           VALUES,
	"VALUE_OF":                         VALUE_OF,
	"VARBINARY":                        VARBINARY,
	"VARCHAR":                          VARCHAR,
	"VARIABLES":                        VARIABLES,
	"VARYING":                          VARYING,
	"VAR_POP":                          VAR_POP,
	"VAR_SAMP":                         VAR_SAMP,
	"VERBOSE":                          VERBOSE,
	"VERSION":                          VERSION,
	"VERSIONING":                       VERSIONING,
	"VIEW":                             VIEW,
	"VIRTUAL":                          VIRTUAL,
	"VOLATILE":                         VOLATILE,
	"WEEK":                             WEEK,
	"WHEN":                             WHEN,
	"WHENEVER":                         WHENEVER,
	"WHERE":                            WHERE,
	"WIDTH_BUCKET":                     WIDTH_BUCKET,
	"WINDOW":                           WINDOW,
	"WITH":                             WITH,
	"WITHIN":                           WITHIN,
	"WITHOUT":                          WITHOUT,
	"WITHOUT_ARRAY_WRAPPER":            WITHOUT_ARRAY_WRAPPER,
	"WORK":                             WORK,
	"WRITE":                            WRITE,
	"XML":                              XML,
	"XOR":                              XOR,
	"YEAR":                             YEAR,
	"ZONE":                             ZONE,
	"ZORDER":                           ZORDER,
}
var keyWordToString = map[KeyWord]string{
	ABORT:                            "ABORT",
	ABS:                              "ABS",
	ABSOLUTE:                         "ABSOLUTE",
	ACTION:                           "ACTION",
	ADD:                              "ADD",
	ADMIN:                            "ADMIN",
	AFTER:                            "AFTER",
	AGAINST:                          "AGAINST",
	ALL:                              "ALL",
	ALLOCATE:                         "ALLOCATE",
	ALTER:                            "ALTER",
	ALWAYS:                           "ALWAYS",
	ANALYZE:                          "ANALYZE",
	AND:                              "AND",
	ANTI:                             "ANTI",
	ANY:                              "ANY",
	APPLY:                            "APPLY",
	ARCHIVE:                          "ARCHIVE",
	ARE:                              "ARE",
	ARRAY:                            "ARRAY",
	ARRAY_AGG:                        "ARRAY_AGG",
	ARRAY_MAX_CARDINALITY:            "ARRAY_MAX_CARDINALITY",
	AS:                               "AS",
	ASC:                              "ASC",
	ASENSITIVE:                       "ASENSITIVE",
	ASSERT:                           "ASSERT",
	ASYMMETRIC:                       "ASYMMETRIC",
	AT:                               "AT",
	ATOMIC:                           "ATOMIC",
	ATTACH:                           "ATTACH",
	AUTHORIZATION:                    "AUTHORIZATION",
	AUTO:                             "AUTO",
	AUTOINCREMENT:                    "AUTOINCREMENT",
	AUTO_INCREMENT:                   "AUTO_INCREMENT",
	AVG:                              "AVG",
	AVRO:                             "AVRO",
	BACKWARD:                         "BACKWARD",
	BASE64:                           "BASE64",
	BEGIN:                            "BEGIN",
	BEGIN_FRAME:                      "BEGIN_FRAME",
	BEGIN_PARTITION:                  "BEGIN_PARTITION",
	BETWEEN:                          "BETWEEN",
	BIGDECIMAL:                       "BIGDECIMAL",
	BIGINT:                           "BIGINT",
	BIGNUMERIC:                       "BIGNUMERIC",
	BINARY:                           "BINARY",
	BINDING:                          "BINDING",
	BLOB:                             "BLOB",
	BLOOMFILTER:                      "BLOOMFILTER",
	BOOL:                             "BOOL",
	BOOLEAN:                          "BOOLEAN",
	BOTH:                             "BOTH",
	BROWSE:                           "BROWSE",
	BTREE:                            "BTREE",
	BY:                               "BY",
	BYPASSRLS:                        "BYPASSRLS",
	BYTEA:                            "BYTEA",
	BYTES:                            "BYTES",
	CACHE:                            "CACHE",
	CALL:                             "CALL",
	CALLED:                           "CALLED",
	CARDINALITY:                      "CARDINALITY",
	CASCADE:                          "CASCADE",
	CASCADED:                         "CASCADED",
	CASE:                             "CASE",
	CAST:                             "CAST",
	CEIL:                             "CEIL",
	CEILING:                          "CEILING",
	CENTURY:                          "CENTURY",
	CHAIN:                            "CHAIN",
	CHANGE:                           "CHANGE",
	CHANNEL:                          "CHANNEL",
	CHAR:                             "CHAR",
	CHARACTER:                        "CHARACTER",
	CHARACTERS:                       "CHARACTERS",
	CHARACTER_LENGTH:                 "CHARACTER_LENGTH",
	CHARSET:                          "CHARSET",
	CHAR_LENGTH:                      "CHAR_LENGTH",
	CHECK:                            "CHECK",
	CLOB:                             "CLOB",
	CLONE:                            "CLONE",
	CLOSE:                            "CLOSE",
	CLUSTER:                          "CLUSTER",
	COALESCE:                         "COALESCE",
	COLLATE:                          "COLLATE",
	COLLATION:                        "COLLATION",
	COLLECT:                          "COLLECT",
	COLLECTION:                       "COLLECTION",
	COLUMN:                           "COLUMN",
	COLUMNS:                          "COLUMNS",
	COMMENT:                          "COMMENT",
	COMMIT:                           "COMMIT",
	COMMITTED:                        "COMMITTED",
	COMPRESSION:                      "COMPRESSION",
	COMPUTE:                          "COMPUTE",
	CONCURRENTLY:                     "CONCURRENTLY",
	CONDITION:                        "CONDITION",
	CONFLICT:                         "CONFLICT",
	CONNECT:                          "CONNECT",
	CONNECTION:                       "CONNECTION",
	CONSTRAINT:                       "CONSTRAINT",
	CONTAINS:                         "CONTAINS",
	CONVERT:                          "CONVERT",
	COPY:                             "COPY",
	COPY_OPTIONS:                     "COPY_OPTIONS",
	CORR:                             "CORR",
	CORRESPONDING:                    "CORRESPONDING",
	COUNT:                            "COUNT",
	COVAR_POP:                        "COVAR_POP",
	COVAR_SAMP:                       "COVAR_SAMP",
	CREATE:                           "CREATE",
	CREATEDB:                         "CREATEDB",
	CREATEROLE:                       "CREATEROLE",
	CREDENTIALS:                      "CREDENTIALS",
	CROSS:                            "CROSS",
	CSV:                              "CSV",
	CUBE:                             "CUBE",
	CUME_DIST:                        "CUME_DIST",
	CURRENT:                          "CURRENT",
	CURRENT_CATALOG:                  "CURRENT_CATALOG",
	CURRENT_DATE:                     "CURRENT_DATE",
	CURRENT_DEFAULT_TRANSFORM_GROUP:  "CURRENT_DEFAULT_TRANSFORM_GROUP",
	CURRENT_PATH:                     "CURRENT_PATH",
	CURRENT_ROLE:                     "CURRENT_ROLE",
	CURRENT_ROW:                      "CURRENT_ROW",
	CURRENT_SCHEMA:                   "CURRENT_SCHEMA",
	CURRENT_TIME:                     "CURRENT_TIME",
	CURRENT_TIMESTAMP:                "CURRENT_TIMESTAMP",
	CURRENT_TRANSFORM_GROUP_FOR_TYPE: "CURRENT_TRANSFORM_GROUP_FOR_TYPE",
	CURRENT_USER:                     "CURRENT_USER",
	CURSOR:                           "CURSOR",
	CYCLE:                            "CYCLE",
	DATA:                             "DATA",
	DATABASE:                         "DATABASE",
	DATE:                             "DATE",
	DATETIME:                         "DATETIME",
	DAY:                              "DAY",
	DAYOFWEEK:                        "DAYOFWEEK",
	DAYOFYEAR:                        "DAYOFYEAR",
	DEALLOCATE:                       "DEALLOCATE",
	DEC:                              "DEC",
	DECADE:                           "DECADE",
	DECIMAL:                          "DECIMAL",
	DECLARE:                          "DECLARE",
	DEFAULT:                          "DEFAULT",
	DEFERRABLE:                       "DEFERRABLE",
	DEFERRED:                         "DEFERRED",
	DEFINED:                          "DEFINED",
	DELAYED:                          "DELAYED",
	DELETE:                           "DELETE",
	DELIMITED:                        "DELIMITED",
	DELIMITER:                        "DELIMITER",
	DELTA:                            "DELTA",
	DENSE_RANK:                       "DENSE_RANK",
	DEREF:                            "DEREF",
	DESC:                             "DESC",
	DESCRIBE:                         "DESCRIBE",
	DETACH:                           "DETACH",
	DETAIL:                           "DETAIL",
	DETERMINISTIC:                    "DETERMINISTIC",
	DIRECTORY:                        "DIRECTORY",
	DISABLE:                          "DISABLE",
	DISCARD:                          "DISCARD",
	DISCONNECT:                       "DISCONNECT",
	DISTINCT:                         "DISTINCT",
	DISTRIBUTE:                       "DISTRIBUTE",
	DIV:                              "DIV",
	DO:                               "DO",
	DOUBLE:                           "DOUBLE",
	DOW:                              "DOW",
	DOY:                              "DOY",
	DROP:                             "DROP",
	DRY:                              "DRY",
	DUPLICATE:                        "DUPLICATE",
	DYNAMIC:                          "DYNAMIC",
	EACH:                             "EACH",
	ELEMENT:                          "ELEMENT",
	ELEMENTS:                         "ELEMENTS",
	ELSE:                             "ELSE",
	EMPTY:                            "EMPTY",
	ENABLE:                           "ENABLE",
	ENCODING:                         "ENCODING",
	ENCRYPTION:                       "ENCRYPTION",
	END:                              "END",
	END_EXEC:                         "END_EXEC",
	ENDPOINT:                         "ENDPOINT",
	END_FRAME:                        "END_FRAME",
	END_PARTITION:                    "END_PARTITION",
	ENFORCED:                         "ENFORCED",
	ENGINE:                           "ENGINE",
	ENUM:                             "ENUM",
	EPOCH:                            "EPOCH",
	EQUALS:                           "EQUALS",
	ERROR:                            "ERROR",
	ESCAPE:                           "ESCAPE",
	ESCAPED:                          "ESCAPED",
	EVENT:                            "EVENT",
	EVERY:                            "EVERY",
	EXCEPT:                           "EXCEPT",
	EXCEPTION:                        "EXCEPTION",
	EXCLUDE:                          "EXCLUDE",
	EXCLUSIVE:                        "EXCLUSIVE",
	EXEC:                             "EXEC",
	EXECUTE:                          "EXECUTE",
	EXISTS:                           "EXISTS",
	EXP:                              "EXP",
	EXPANSION:                        "EXPANSION",
	EXPLAIN:                          "EXPLAIN",
	EXPLICIT:                         "EXPLICIT",
	EXPORT:                           "EXPORT",
	EXTENDED:                         "EXTENDED",
	EXTENSION:                        "EXTENSION",
	EXTERNAL:                         "EXTERNAL",
	EXTRACT:                          "EXTRACT",
	FAIL:                             "FAIL",
	FALSE:                            "FALSE",
	FETCH:                            "FETCH",
	FIELDS:                           "FIELDS",
	FILE:                             "FILE",
	FILES:                            "FILES",
	FILE_FORMAT:                      "FILE_FORMAT",
	FILTER:                           "FILTER",
	FIRST:                            "FIRST",
	FIRST_VALUE:                      "FIRST_VALUE",
	FLOAT:                            "FLOAT",
	FLOAT4:                           "FLOAT4",
	FLOAT64:                          "FLOAT64",
	FLOAT8:                           "FLOAT8",
	FLOOR:                            "FLOOR",
	FLUSH:                            "FLUSH",
	FOLLOWING:                        "FOLLOWING",
	FOR:                              "FOR",
	FORCE:                            "FORCE",
	FORCE_NOT_NULL:                   "FORCE_NOT_NULL",
	FORCE_NULL:                       "FORCE_NULL",
	FORCE_QUOTE:                      "FORCE_QUOTE",
	FOREIGN:                          "FOREIGN",
	FORMAT:                           "FORMAT",
	FORMATTED:                        "FORMATTED",
	FORWARD:                          "FORWARD",
	FRAME_ROW:                        "FRAME_ROW",
	FREE:                             "FREE",
	FREEZE:                           "FREEZE",
	FROM:                             "FROM",
	FSCK:                             "FSCK",
	FULL:                             "FULL",
	FULLTEXT:                         "FULLTEXT",
	FUNCTION:                         "FUNCTION",
	FUNCTIONS:                        "FUNCTIONS",
	FUSION:                           "FUSION",
	GENERAL:                          "GENERAL",
	GENERATE:                         "GENERATE",
	GENERATED:                        "GENERATED",
	GEOGRAPHY:                        "GEOGRAPHY",
	GET:                              "GET",
	GLOBAL:                           "GLOBAL",
	GRANT:                            "GRANT",
	GRANTED:                          "GRANTED",
	GRAPHVIZ:                         "GRAPHVIZ",
	GROUP:                            "GROUP",
	GROUPING:                         "GROUPING",
	GROUPS:                           "GROUPS",
	HASH:                             "HASH",
	HAVING:                           "HAVING",
	HEADER:                           "HEADER",
	HIGH_PRIORITY:                    "HIGH_PRIORITY",
	HISTORY:                          "HISTORY",
	HIVEVAR:                          "HIVEVAR",
	HOLD:                             "HOLD",
	HOSTS:                            "HOSTS",
	HOUR:                             "HOUR",
	HOURS:                            "HOURS",
	IDENTITY:                         "IDENTITY",
	IF:                               "IF",
	IGNORE:                           "IGNORE",
	ILIKE:                            "ILIKE",
	IMMEDIATE:                        "IMMEDIATE",
	IMMUTABLE:                        "IMMUTABLE",
	IN:                               "IN",
	INCLUDE:                          "INCLUDE",
	INCLUDE_NULL_VALUES:              "INCLUDE_NULL_VALUES",
	INCREMENT:                        "INCREMENT",
	INDEX:                            "INDEX",
	INDICATOR:                        "INDICATOR",
	INHERIT:                          "INHERIT",
	INITIALLY:                        "INITIALLY",
	INNER:                            "INNER",
	INOUT:                            "INOUT",
	INPUT:                            "INPUT",
	INPUTFORMAT:                      "INPUTFORMAT",
	INSENSITIVE:                      "INSENSITIVE",
	INSERT:                           "INSERT",
	INSTALL:                          "INSTALL",
	INT:                              "INT",
	INT2:                             "INT2",
	INT4:                             "INT4",
	INT64:                            "INT64",
	INT8:                             "INT8",
	INTEGER:                          "INTEGER",
	INTERSECT:                        "INTERSECT",
	INTERSECTION:                     "INTERSECTION",
	INTERVAL:                         "INTERVAL",
	INTO:                             "INTO",
	IS:                               "IS",
	ISODOW:                           "ISODOW",
	ISOLATION:                        "ISOLATION",
	ISOWEEK:                          "ISOWEEK",
	ISOYEAR:                          "ISOYEAR",
	ITEMS:                            "ITEMS",
	JAR:                              "JAR",
	JOIN:                             "JOIN",
	JSON:                             "JSON",
	JSONB:                            "JSONB",
	JSONFILE:                         "JSONFILE",
	JSON_TABLE:                       "JSON_TABLE",
	JULIAN:                           "JULIAN",
	KEY:                              "KEY",
	KEYS:                             "KEYS",
	KILL:                             "KILL",
	LAG:                              "LAG",
	LANGUAGE:                         "LANGUAGE",
	LARGE:                            "LARGE",
	LAST:                             "LAST",
	LAST_VALUE:                       "LAST_VALUE",
	LATERAL:                          "LATERAL",
	LEAD:                             "LEAD",
	LEADING:                          "LEADING",
	LEFT:                             "LEFT",
	LEVEL:                            "LEVEL",
	LIKE:                             "LIKE",
	LIKE_REGEX:                       "LIKE_REGEX",
	LIMIT:                            "LIMIT",
	LINES:                            "LINES",
	LISTAGG:                          "LISTAGG",
	LN:                               "LN",
	LOAD:                             "LOAD",
	LOCAL:                            "LOCAL",
	LOCALTIME:                        "LOCALTIME",
	LOCALTIMESTAMP:                   "LOCALTIMESTAMP",
	LOCATION:                         "LOCATION",
	LOCK:                             "LOCK",
	LOCKED:                           "LOCKED",
	LOGIN:                            "LOGIN",
	LOGS:                             "LOGS",
	LOWER:                            "LOWER",
	LOW_PRIORITY:                     "LOW_PRIORITY",
	MACRO:                            "MACRO",
	MANAGEDLOCATION:                  "MANAGEDLOCATION",
	MAP:                              "MAP",
	MATCH:                            "MATCH",
	MATCHED:                          "MATCHED",
	MATERIALIZED:                     "MATERIALIZED",
	MAX:                              "MAX",
	MAXVALUE:                         "MAXVALUE",
	MEDIUMINT:                        "MEDIUMINT",
	MEMBER:                           "MEMBER",
	MERGE:                            "MERGE",
	METADATA:                         "METADATA",
	METHOD:                           "METHOD",
	MICROSECOND:                      "MICROSECOND",
	MICROSECONDS:                     "MICROSECONDS",
	MILLENIUM:                        "MILLENIUM",
	MILLENNIUM:                       "MILLENNIUM",
	MILLISECOND:                      "MILLISECOND",
	MILLISECONDS:                     "MILLISECONDS",
	MIN:                              "MIN",
	MINUTE:                           "MINUTE",
	MINVALUE:                         "MINVALUE",
	MOD:                              "MOD",
	MODE:                             "MODE",
	MODIFIES:                         "MODIFIES",
	MODULE:                           "MODULE",
	MONTH:                            "MONTH",
	MSCK:                             "MSCK",
	MULTISET:                         "MULTISET",
	MUTATION:                         "MUTATION",
	NAME:                             "NAME",
	NANOSECOND:                       "NANOSECOND",
	NANOSECONDS:                      "NANOSECONDS",
	NATIONAL:                         "NATIONAL",
	NATURAL:                          "NATURAL",
	NCHAR:                            "NCHAR",
	NCLOB:                            "NCLOB",
	NEW:                              "NEW",
	NEXT:                             "NEXT",
	NO:                               "NO",
	NOBYPASSRLS:                      "NOBYPASSRLS",
	NOCREATEDB:                       "NOCREATEDB",
	NOCREATEROLE:                     "NOCREATEROLE",
	NOINHERIT:                        "NOINHERIT",
	NOLOGIN:                          "NOLOGIN",
	NONE:                             "NONE",
	NOREPLICATION:                    "NOREPLICATION",
	NORMALIZE:                        "NORMALIZE",
	NOSCAN:                           "NOSCAN",
	NOSUPERUSER:                      "NOSUPERUSER",
	NOT:                              "NOT",
	NOTHING:                          "NOTHING",
	NOWAIT:                           "NOWAIT",
	NO_WRITE_TO_BINLOG:               "NO_WRITE_TO_BINLOG",
	NTH_VALUE:                        "NTH_VALUE",
	NTILE:                            "NTILE",
	NULL:                             "NULL",
	NULLIF:                           "NULLIF",
	NULLS:                            "NULLS",
	NUMERIC:                          "NUMERIC",
	NVARCHAR:                         "NVARCHAR",
	OBJECT:                           "OBJECT",
	OCCURRENCES_REGEX:                "OCCURRENCES_REGEX",
	OCTETS:                           "OCTETS",
	OCTET_LENGTH:                     "OCTET_LENGTH",
	OF:                               "OF",
	OFFSET:                           "OFFSET",
	OLD:                              "OLD",
	ON:                               "ON",
	ONLY:                             "ONLY",
	OPEN:                             "OPEN",
	OPERATOR:                         "OPERATOR",
	OPTIMIZE:                         "OPTIMIZE",
	OPTIMIZER_COSTS:                  "OPTIMIZER_COSTS",
	OPTION:                           "OPTION",
	OPTIONS:                          "OPTIONS",
	OR:                               "OR",
	ORC:                              "ORC",
	ORDER:                            "ORDER",
	OUT:                              "OUT",
	OUTER:                            "OUTER",
	OUTPUTFORMAT:                     "OUTPUTFORMAT",
	OVER:                             "OVER",
	OVERFLOW:                         "OVERFLOW",
	OVERLAPS:                         "OVERLAPS",
	OVERLAY:                          "OVERLAY",
	OVERWRITE:                        "OVERWRITE",
	OWNED:                            "OWNED",
	PARALLEL:                         "PARALLEL",
	PARAMETER:                        "PARAMETER",
	PARQUET:                          "PARQUET",
	PARTITION:                        "PARTITION",
	PARTITIONED:                      "PARTITIONED",
	PARTITIONS:                       "PARTITIONS",
	PASSWORD:                         "PASSWORD",
	PATH:                             "PATH",
	PATTERN:                          "PATTERN",
	PERCENT:                          "PERCENT",
	PERCENTILE_CONT:                  "PERCENTILE_CONT",
	PERCENTILE_DISC:                  "PERCENTILE_DISC",
	PERCENT_RANK:                     "PERCENT_RANK",
	PERIOD:                           "PERIOD",
	PERSISTENT:                       "PERSISTENT",
	PIVOT:                            "PIVOT",
	PLACING:                          "PLACING",
	PLANS:                            "PLANS",
	PORTION:                          "PORTION",
	POSITION:                         "POSITION",
	POSITION_REGEX:                   "POSITION_REGEX",
	POWER:                            "POWER",
	PRAGMA:                           "PRAGMA",
	PRECEDES:                         "PRECEDES",
	PRECEDING:                        "PRECEDING",
	PRECISION:                        "PRECISION",
	PREPARE:                          "PREPARE",
	PRESERVE:                         "PRESERVE",
	PRIMARY:                          "PRIMARY",
	PRIOR:                            "PRIOR",
	PRIVILEGES:                       "PRIVILEGES",
	PROCEDURE:                        "PROCEDURE",
	PROGRAM:                          "PROGRAM",
	PURGE:                            "PURGE",
	QUALIFY:                          "QUALIFY",
	QUARTER:                          "QUARTER",
	QUERY:                            "QUERY",
	QUOTE:                            "QUOTE",
	RANGE:                            "RANGE",
	RANK:                             "RANK",
	RAW:                              "RAW",
	RCFILE:                           "RCFILE",
	READ:                             "READ",
	READS:                            "READS",
	READ_ONLY:                        "READ_ONLY",
	REAL:                             "REAL",
	RECURSIVE:                        "RECURSIVE",
	REF:                              "REF",
	REFERENCES:                       "REFERENCES",
	REFERENCING:                      "REFERENCING",
	REGCLASS:                         "REGCLASS",
	REGEXP:                           "REGEXP",
	REGR_AVGX:                        "REGR_AVGX",
	REGR_AVGY:                        "REGR_AVGY",
	REGR_COUNT:                       "REGR_COUNT",
	REGR_INTERCEPT:                   "REGR_INTERCEPT",
	REGR_R2:                          "REGR_R2",
	REGR_SLOPE:                       "REGR_SLOPE",
	REGR_SXX:                         "REGR_SXX",
	REGR_SXY:                         "REGR_SXY",
	REGR_SYY:                         "REGR_SYY",
	RELATIVE:                         "RELATIVE",
	RELAY:                            "RELAY",
	RELEASE:                          "RELEASE",
	RENAME:                           "RENAME",
	REORG:                            "REORG",
	REPAIR:                           "REPAIR",
	REPEATABLE:                       "REPEATABLE",
	REPLACE:                          "REPLACE",
	REPLICA:                          "REPLICA",
	REPLICATION:                      "REPLICATION",
	RESET:                            "RESET",
	RESPECT:                          "RESPECT",
	RESTART:                          "RESTART",
	RESTRICT:                         "RESTRICT",
	RESTRICTED:                       "RESTRICTED",
	RESULT:                           "RESULT",
	RESULTSET:                        "RESULTSET",
	RETAIN:                           "RETAIN",
	RETURN:                           "RETURN",
	RETURNING:                        "RETURNING",
	RETURNS:                          "RETURNS",
	REVOKE:                           "REVOKE",
	RIGHT:                            "RIGHT",
	RLIKE:                            "RLIKE",
	ROLE:                             "ROLE",
	ROLLBACK:                         "ROLLBACK",
	ROLLUP:                           "ROLLUP",
	ROOT:                             "ROOT",
	ROW:                              "ROW",
	ROWID:                            "ROWID",
	ROWS:                             "ROWS",
	ROW_NUMBER:                       "ROW_NUMBER",
	RULE:                             "RULE",
	RUN:                              "RUN",
	SAFE:                             "SAFE",
	SAFE_CAST:                        "SAFE_CAST",
	SAVEPOINT:                        "SAVEPOINT",
	SCHEMA:                           "SCHEMA",
	SCOPE:                            "SCOPE",
	SCROLL:                           "SCROLL",
	SEARCH:                           "SEARCH",
	SECOND:                           "SECOND",
	SECRET:                           "SECRET",
	SECURITY:                         "SECURITY",
	SELECT:                           "SELECT",
	SEMI:                             "SEMI",
	SENSITIVE:                        "SENSITIVE",
	SEQUENCE:                         "SEQUENCE",
	SEQUENCEFILE:                     "SEQUENCEFILE",
	SEQUENCES:                        "SEQUENCES",
	SERDE:                            "SERDE",
	SERDEPROPERTIES:                  "SERDEPROPERTIES",
	SERIALIZABLE:                     "SERIALIZABLE",
	SESSION:                          "SESSION",
	SESSION_USER:                     "SESSION_USER",
	SET:                              "SET",
	SETS:                             "SETS",
	SHARE:                            "SHARE",
	SHOW:                             "SHOW",
	SIMILAR:                          "SIMILAR",
	SKIP:                             "SKIP",
	SLOW:                             "SLOW",
	SMALLINT:                         "SMALLINT",
	SNAPSHOT:                         "SNAPSHOT",
	SOME:                             "SOME",
	SORT:                             "SORT",
	SPATIAL:                          "SPATIAL",
	SPECIFIC:                         "SPECIFIC",
	SPECIFICTYPE:                     "SPECIFICTYPE",
	SQL:                              "SQL",
	SQLEXCEPTION:                     "SQLEXCEPTION",
	SQLSTATE:                         "SQLSTATE",
	SQLWARNING:                       "SQLWARNING",
	SQRT:                             "SQRT",
	STABLE:                           "STABLE",
	STAGE:                            "STAGE",
	START:                            "START",
	STATIC:                           "STATIC",
	STATISTICS:                       "STATISTICS",
	STATUS:                           "STATUS",
	STDDEV_POP:                       "STDDEV_POP",
	STDDEV_SAMP:                      "STDDEV_SAMP",
	STDIN:                            "STDIN",
	STDOUT:                           "STDOUT",
	STORAGE_INTEGRATION:              "STORAGE_INTEGRATION",
	STORED:                           "STORED",
	STRICT:                           "STRICT",
	STRING:                           "STRING",
	STRUCT:                           "STRUCT",
	SUBMULTISET:                      "SUBMULTISET",
	SUBSTRING:                        "SUBSTRING",
	SUBSTRING_REGEX:                  "SUBSTRING_REGEX",
	SUCCEEDS:                         "SUCCEEDS",
	SUM:                              "SUM",
	SUPER:                            "SUPER",
	SUPERUSER:                        "SUPERUSER",
	SWAP:                             "SWAP",
	SYMMETRIC:                        "SYMMETRIC",
	SYNC:                             "SYNC",
	SYSTEM:                           "SYSTEM",
	SYSTEM_TIME:                      "SYSTEM_TIME",
	SYSTEM_USER:                      "SYSTEM_USER",
	TABLE:                            "TABLE",
	TABLES:                           "TABLES",
	TABLESAMPLE:                      "TABLESAMPLE",
	TBLPROPERTIES:                    "TBLPROPERTIES",
	TEMP:                             "TEMP",
	TEMPORARY:                        "TEMPORARY",
	TERMINATED:                       "TERMINATED",
	TEXT:                             "TEXT",
	TEXTFILE:                         "TEXTFILE",
	THEN:                             "THEN",
	TIES:                             "TIES",
	TIME:                             "TIME",
	TIMESTAMP:                        "TIMESTAMP",
	TIMESTAMPTZ:                      "TIMESTAMPTZ",
	TIMETZ:                           "TIMETZ",
	TIMEZONE:                         "TIMEZONE",
	TIMEZONE_ABBR:                    "TIMEZONE_ABBR",
	TIMEZONE_HOUR:                    "TIMEZONE_HOUR",
	TIMEZONE_MINUTE:                  "TIMEZONE_MINUTE",
	TIMEZONE_REGION:                  "TIMEZONE_REGION",
	TINYINT:                          "TINYINT",
	TO:                               "TO",
	TOP:                              "TOP",
	TRAILING:                         "TRAILING",
	TRANSACTION:                      "TRANSACTION",
	TRANSIENT:                        "TRANSIENT",
	TRANSLATE:                        "TRANSLATE",
	TRANSLATE_REGEX:                  "TRANSLATE_REGEX",
	TRANSLATION:                      "TRANSLATION",
	TREAT:                            "TREAT",
	TRIGGER:                          "TRIGGER",
	TRIM:                             "TRIM",
	TRIM_ARRAY:                       "TRIM_ARRAY",
	TRUE:                             "TRUE",
	TRUNCATE:                         "TRUNCATE",
	TRY_CAST:                         "TRY_CAST",
	TYPE:                             "TYPE",
	UESCAPE:                          "UESCAPE",
	UNBOUNDED:                        "UNBOUNDED",
	UNCACHE:                          "UNCACHE",
	UNCOMMITTED:                      "UNCOMMITTED",
	UNION:                            "UNION",
	UNIQUE:                           "UNIQUE",
	UNKNOWN:                          "UNKNOWN",
	UNLOAD:                           "UNLOAD",
	UNLOCK:                           "UNLOCK",
	UNLOGGED:                         "UNLOGGED",
	UNNEST:                           "UNNEST",
	UNPIVOT:                          "UNPIVOT",
	UNSAFE:                           "UNSAFE",
	UNSIGNED:                         "UNSIGNED",
	UNTIL:                            "UNTIL",
	UPDATE:                           "UPDATE",
	UPPER:                            "UPPER",
	URL:                              "URL",
	USAGE:                            "USAGE",
	USE:                              "USE",
	USER:                             "USER",
	USER_RESOURCES:                   "USER_RESOURCES",
	USING:                            "USING",
	UUID:                             "UUID",
	VACUUM:                           "VACUUM",
	VALID:                            "VALID",
	VALIDATION_MODE:                  "VALIDATION_MODE",
	VALUE:                            "VALUE",
	VALUES:                           "VALUES",
	VALUE_OF:                         "VALUE_OF",
	VARBINARY:                        "VARBINARY",
	VARCHAR:                          "VARCHAR",
	VARIABLES:                        "VARIABLES",
	VARYING:                          "VARYING",
	VAR_POP:                          "VAR_POP",
	VAR_SAMP:                         "VAR_SAMP",
	VERBOSE:                          "VERBOSE",
	VERSION:                          "VERSION",
	VERSIONING:                       "VERSIONING",
	VIEW:                             "VIEW",
	VIRTUAL:                          "VIRTUAL",
	VOLATILE:                         "VOLATILE",
	WEEK:                             "WEEK",
	WHEN:                             "WHEN",
	WHENEVER:                         "WHENEVER",
	WHERE:                            "WHERE",
	WIDTH_BUCKET:                     "WIDTH_BUCKET",
	WINDOW:                           "WINDOW",
	WITH:                             "WITH",
	WITHIN:                           "WITHIN",
	WITHOUT:                          "WITHOUT",
	WITHOUT_ARRAY_WRAPPER:            "WITHOUT_ARRAY_WRAPPER",
	WORK:                             "WORK",
	WRITE:                            "WRITE",
	XML:                              "XML",
	XOR:                              "XOR",
	YEAR:                             "YEAR",
	ZONE:                             "ZONE",
	ZORDER:                           "ZORDER",
}
