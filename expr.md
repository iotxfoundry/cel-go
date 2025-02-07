# Expr

- Standard Functions

| name                     | id                                        | expr                                                                              |
| ------------------------ | ----------------------------------------- | --------------------------------------------------------------------------------- |
| `_?_:_`                  | conditional                               | (bool,A,A) -> A                                                                   |
| `_&&_`                   | logical_and                               | (bool,bool) -> bool                                                               |
| `_\|\|_`                 | logical_or                                | (bool,bool) -> bool                                                               |
| `!_`                     | logical_not                               | (bool) -> bool                                                                    |
| `@not_strictly_false`    | not_strictly_false                        | (bool) -> bool                                                                    |
| `__not_strictly_false__` | **not_strictly_false**                    | (bool) -> bool                                                                    |
| `_==_`                   | equals                                    | (A,A) -> bool                                                                     |
| `_!=_`                   | not_equals                                | (A,A) -> bool                                                                     |
| `_+_`                    | add_bytes                                 | (bytes,bytes) -> bytes                                                            |
| `_+_`                    | add_double                                | (double,double) -> double                                                         |
| `_+_`                    | add_duration_duration                     | (google.protobuf.Duration,google.protobuf.Duration) -> google.protobuf.Duration   |
| `_+_`                    | add_duration_timestamp                    | (google.protobuf.Duration,google.protobuf.Timestamp) -> google.protobuf.Timestamp |
| `_+_`                    | add_timestamp_duration                    | (google.protobuf.Timestamp,google.protobuf.Duration) -> google.protobuf.Timestamp |
| `_+_`                    | add_int64                                 | (int,int) -> int                                                                  |
| `_+_`                    | add_list                                  | (list,list) -> list(A)                                                            |
| `_+_`                    | add_string                                | (string,string) -> string                                                         |
| `_+_`                    | add_uint64                                | (uint,uint) -> uint                                                               |
| `_/_`                    | divide_double                             | (double,double) -> double                                                         |
| `_/_`                    | divide_int64                              | (int,int) -> int                                                                  |
| `_/_`                    | divide_uint64                             | (uint,uint) -> uint                                                               |
| `_%_`                    | modulo_int64                              | (int,int) -> int                                                                  |
| `_%_`                    | modulo_uint64                             | (uint,uint) -> uint                                                               |
| `_*_`                    | multiply_double                           | (double,double) -> double                                                         |
| `_*_`                    | multiply_int64                            | (int,int) -> int                                                                  |
| `_*_`                    | multiply_uint64                           | (uint,uint) -> uint                                                               |
| `-_`                     | negate_double                             | (double) -> double                                                                |
| `-_`                     | negate_int64                              | (int) -> int                                                                      |
| `_-_`                    | subtract_double                           | (double,double) -> double                                                         |
| `_-_`                    | subtract_duration_duration                | (google.protobuf.Duration,google.protobuf.Duration) -> google.protobuf.Duration   |
| `_-_`                    | subtract_int64                            | (int,int) -> int                                                                  |
| `_-_`                    | subtract_timestamp_duration               | (google.protobuf.Timestamp,google.protobuf.Duration) -> google.protobuf.Timestamp |
| `_-_`                    | subtract_timestamp_timestamp              | (google.protobuf.Timestamp,google.protobuf.Timestamp) -> google.protobuf.Duration |
| `_-_`                    | subtract_uint64                           | (uint,uint) -> uint                                                               |
| `_<_`                    | less_bool                                 | (bool,bool) -> bool                                                               |
| `_<_`                    | less_int64                                | (int,int) -> bool                                                                 |
| `_<_`                    | less_int64_double                         | (int,double) -> bool                                                              |
| `_<_`                    | less_int64_uint64                         | (int,uint) -> bool                                                                |
| `_<_`                    | less_uint64                               | (uint,uint) -> bool                                                               |
| `_<_`                    | less_uint64_double                        | (uint,double) -> bool                                                             |
| `_<_`                    | less_uint64_int64                         | (uint,int) -> bool                                                                |
| `_<_`                    | less_double                               | (double,double) -> bool                                                           |
| `_<_`                    | less_double_int64                         | (double,int) -> bool                                                              |
| `_<_`                    | less_double_uint64                        | (double,uint) -> bool                                                             |
| `_<_`                    | less_string                               | (string,string) -> bool                                                           |
| `_<_`                    | less_bytes                                | (bytes,bytes) -> bool                                                             |
| `_<_`                    | less_timestamp                            | (google.protobuf.Timestamp,google.protobuf.Timestamp) -> bool                     |
| `_<_`                    | less_duration                             | (google.protobuf.Duration,google.protobuf.Duration) -> bool                       |
| `_<=_`                   | less_equals_bool                          | (bool,bool) -> bool                                                               |
| `_<=_`                   | less_equals_int64                         | (int,int) -> bool                                                                 |
| `_<=_`                   | less_equals_int64_double                  | (int,double) -> bool                                                              |
| `_<=_`                   | less_equals_int64_uint64                  | (int,uint) -> bool                                                                |
| `_<=_`                   | less_equals_uint64                        | (uint,uint) -> bool                                                               |
| `_<=_`                   | less_equals_uint64_double                 | (uint,double) -> bool                                                             |
| `_<=_`                   | less_equals_uint64_int64                  | (uint,int) -> bool                                                                |
| `_<=_`                   | less_equals_double                        | (double,double) -> bool                                                           |
| `_<=_`                   | less_equals_double_int64                  | (double,int) -> bool                                                              |
| `_<=_`                   | less_equals_double_uint64                 | (double,uint) -> bool                                                             |
| `_<=_`                   | less_equals_string                        | (string,string) -> bool                                                           |
| `_<=_`                   | less_equals_bytes                         | (bytes,bytes) -> bool                                                             |
| `_<=_`                   | less_equals_timestamp                     | (google.protobuf.Timestamp,google.protobuf.Timestamp) -> bool                     |
| `_<=_`                   | less_equals_duration                      | (google.protobuf.Duration,google.protobuf.Duration) -> bool                       |
| `_>_`                    | greater_bool                              | (bool,bool) -> bool                                                               |
| `_>_`                    | greater_int64                             | (int,int) -> bool                                                                 |
| `_>_`                    | greater_int64_double                      | (int,double) -> bool                                                              |
| `_>_`                    | greater_int64_uint64                      | (int,uint) -> bool                                                                |
| `_>_`                    | greater_uint64                            | (uint,uint) -> bool                                                               |
| `_>_`                    | greater_uint64_double                     | (uint,double) -> bool                                                             |
| `_>_`                    | greater_uint64_int64                      | (uint,int) -> bool                                                                |
| `_>_`                    | greater_double                            | (double,double) -> bool                                                           |
| `_>_`                    | greater_double_int64                      | (double,int) -> bool                                                              |
| `_>_`                    | greater_double_uint64                     | (double,uint) -> bool                                                             |
| `_>_`                    | greater_string                            | (string,string) -> bool                                                           |
| `_>_`                    | greater_bytes                             | (bytes,bytes) -> bool                                                             |
| `_>_`                    | greater_timestamp                         | (google.protobuf.Timestamp,google.protobuf.Timestamp) -> bool                     |
| `_>_`                    | greater_duration                          | (google.protobuf.Duration,google.protobuf.Duration) -> bool                       |
| `_>=_`                   | greater_equals_bool                       | (bool,bool) -> bool                                                               |
| `_>=_`                   | greater_equals_int64                      | (int,int) -> bool                                                                 |
| `_>=_`                   | greater_equals_int64_double               | (int,double) -> bool                                                              |
| `_>=_`                   | greater_equals_int64_uint64               | (int,uint) -> bool                                                                |
| `_>=_`                   | greater_equals_uint64                     | (uint,uint) -> bool                                                               |
| `_>=_`                   | greater_equals_uint64_double              | (uint,double) -> bool                                                             |
| `_>=_`                   | greater_equals_uint64_int64               | (uint,int) -> bool                                                                |
| `_>=_`                   | greater_equals_double                     | (double,double) -> bool                                                           |
| `_>=_`                   | greater_equals_double_int64               | (double,int) -> bool                                                              |
| `_>=_`                   | greater_equals_double_uint64              | (double,uint) -> bool                                                             |
| `_>=_`                   | greater_equals_string                     | (string,string) -> bool                                                           |
| `_>=_`                   | greater_equals_bytes                      | (bytes,bytes) -> bool                                                             |
| `_>=_`                   | greater_equals_timestamp                  | (google.protobuf.Timestamp,google.protobuf.Timestamp) -> bool                     |
| `_>=_`                   | greater_equals_duration                   | (google.protobuf.Duration,google.protobuf.Duration) -> bool                       |
| `_[_]`                   | index_list                                | (list,int) -> A                                                                   |
| `_[_]`                   | index_map                                 | (map,A) -> B                                                                      |
| `@in`                    | in_list                                   | (A,list) -> bool                                                                  |
| `@in`                    | in_map                                    | (A,map) -> bool                                                                   |
| `_in_`                   | in_list                                   | (A,list) -> bool                                                                  |
| `_in_`                   | in_map                                    | (A,map) -> bool                                                                   |
| `in`                     | in_list                                   | (A,list) -> bool                                                                  |
| `in`                     | in_map                                    | (A,map) -> bool                                                                   |
| `size`                   | size_bytes                                | (bytes) -> int                                                                    |
| `size`                   | bytes_size                                | (bytes) -> int                                                                    |
| `size`                   | size_list                                 | (list) -> int                                                                     |
| `size`                   | list_size                                 | (list) -> int                                                                     |
| `size`                   | size_map                                  | (map) -> int                                                                      |
| `size`                   | map_size                                  | (map) -> int                                                                      |
| `size`                   | size_string                               | (string) -> int                                                                   |
| `size`                   | string_size                               | (string) -> int                                                                   |
| `type`                   | type                                      | (A) -> type(A)                                                                    |
| `bool`                   | bool_to_bool                              | (bool) -> bool                                                                    |
| `bool`                   | string_to_bool                            | (string) -> bool                                                                  |
| `bytes`                  | bytes_to_bytes                            | (bytes) -> bytes                                                                  |
| `bytes`                  | string_to_bytes                           | (string) -> bytes                                                                 |
| `double`                 | double_to_double                          | (double) -> double                                                                |
| `double`                 | int64_to_double                           | (int) -> double                                                                   |
| `double`                 | string_to_double                          | (string) -> double                                                                |
| `double`                 | uint64_to_double                          | (uint) -> double                                                                  |
| `duration`               | duration_to_duration                      | (google.protobuf.Duration) -> google.protobuf.Duration                            |
| `duration`               | int64_to_duration                         | (int) -> google.protobuf.Duration                                                 |
| `duration`               | string_to_duration                        | (string) -> google.protobuf.Duration                                              |
| `dyn`                    | to_dyn                                    | (A) -> dyn                                                                        |
| `int`                    | int64_to_int64                            | (int) -> int                                                                      |
| `int`                    | double_to_int64                           | (double) -> int                                                                   |
| `int`                    | duration_to_int64                         | (google.protobuf.Duration) -> int                                                 |
| `int`                    | string_to_int64                           | (string) -> int                                                                   |
| `int`                    | timestamp_to_int64                        | (google.protobuf.Timestamp) -> int                                                |
| `int`                    | uint64_to_int64                           | (uint) -> int                                                                     |
| `string`                 | string_to_string                          | (string) -> string                                                                |
| `string`                 | bool_to_string                            | (bool) -> string                                                                  |
| `string`                 | bytes_to_string                           | (bytes) -> string                                                                 |
| `string`                 | double_to_string                          | (double) -> string                                                                |
| `string`                 | duration_to_string                        | (google.protobuf.Duration) -> string                                              |
| `string`                 | int64_to_string                           | (int) -> string                                                                   |
| `string`                 | timestamp_to_string                       | (google.protobuf.Timestamp) -> string                                             |
| `string`                 | uint64_to_string                          | (uint) -> string                                                                  |
| `timestamp`              | timestamp_to_timestamp                    | (google.protobuf.Timestamp) -> google.protobuf.Timestamp                          |
| `timestamp`              | int64_to_timestamp                        | (int) -> google.protobuf.Timestamp                                                |
| `timestamp`              | string_to_timestamp                       | (string) -> google.protobuf.Timestamp                                             |
| `uint`                   | uint64_to_uint64                          | (uint) -> uint                                                                    |
| `uint`                   | double_to_uint64                          | (double) -> uint                                                                  |
| `uint`                   | int64_to_uint64                           | (int) -> uint                                                                     |
| `uint`                   | string_to_uint64                          | (string) -> uint                                                                  |
| `contains`               | contains_string                           | (string,string) -> bool                                                           |
| `endsWith`               | ends_with_string                          | (string,string) -> bool                                                           |
| `startsWith`             | starts_with_string                        | (string,string) -> bool                                                           |
| `matches`                | matches                                   | (string,string) -> bool                                                           |
| `matches`                | matches_string                            | (string,string) -> bool                                                           |
| `getFullYear`            | timestamp_to_year                         | (google.protobuf.Timestamp) -> int                                                |
| `getFullYear`            | timestamp_to_year_with_tz                 | (google.protobuf.Timestamp,string) -> int                                         |
| `getMonth`               | timestamp_to_month                        | (google.protobuf.Timestamp) -> int                                                |
| `getMonth`               | timestamp_to_month_with_tz                | (google.protobuf.Timestamp,string) -> int                                         |
| `getDayOfYear`           | timestamp_to_day_of_year                  | (google.protobuf.Timestamp) -> int                                                |
| `getDayOfYear`           | timestamp_to_day_of_year_with_tz          | (google.protobuf.Timestamp,string) -> int                                         |
| `getDayOfMonth`          | timestamp_to_day_of_month                 | (google.protobuf.Timestamp) -> int                                                |
| `getDayOfMonth`          | timestamp_to_day_of_month_with_tz         | (google.protobuf.Timestamp,string) -> int                                         |
| `getDate`                | timestamp_to_day_of_month_1_based         | (google.protobuf.Timestamp) -> int                                                |
| `getDate`                | timestamp_to_day_of_month_1_based_with_tz | (google.protobuf.Timestamp,string) -> int                                         |
| `getDayOfWeek`           | timestamp_to_day_of_week                  | (google.protobuf.Timestamp) -> int                                                |
| `getDayOfWeek`           | timestamp_to_day_of_week_with_tz          | (google.protobuf.Timestamp,string) -> int                                         |
| `getHours`               | timestamp_to_hours                        | (google.protobuf.Timestamp) -> int                                                |
| `getHours`               | timestamp_to_hours_with_tz                | (google.protobuf.Timestamp,string) -> int                                         |
| `getHours`               | duration_to_hours                         | (google.protobuf.Duration) -> int                                                 |
| `getMinutes`             | timestamp_to_minutes                      | (google.protobuf.Timestamp) -> int                                                |
| `getMinutes`             | timestamp_to_minutes_with_tz              | (google.protobuf.Timestamp,string) -> int                                         |
| `getMinutes`             | duration_to_minutes                       | (google.protobuf.Duration) -> int                                                 |
| `getSeconds`             | timestamp_to_seconds                      | (google.protobuf.Timestamp) -> int                                                |
| `getSeconds`             | timestamp_to_seconds_tz                   | (google.protobuf.Timestamp,string) -> int                                         |
| `getSeconds`             | duration_to_seconds                       | (google.protobuf.Duration) -> int                                                 |
| `getMilliseconds`        | timestamp_to_milliseconds                 | (google.protobuf.Timestamp) -> int                                                |
| `getMilliseconds`        | timestamp_to_milliseconds_with_tz         | (google.protobuf.Timestamp,string) -> int                                         |
| `getMilliseconds`        | duration_to_milliseconds                  | (google.protobuf.Duration) -> int                                                 |

- Standard Macros

| name         | id         | expr                      | example                                                        |
| ------------ | ---------- | ------------------------- | -------------------------------------------------------------- |
| `has`        | has        | has(e.f) -> bool          | `has({'a':1}.a) -> true`                                       |
| `all`        | all        | e.all(x, p) -> bool       | `["a", "b"].all(x, x in ["a", "b"]) -> true`                   |
| `exists`     | exists     | e.exists(x, p) -> bool    | `[1, 2, 3, 4, 5u, 1.0].exists(e, type(e) == uint) -> true`     |
| `exists_one` | exists_one | e.exists_one(x,p) -> bool | `[1, 2, 3, 4, 5u, 1.0].exists_one(e, type(e) == uint) -> true` |
| `map`        | map        | e.map(x, t) -> e          | `[1, 2, 3].map(x, x * 2) -> [2 4 6]`                           |
| `map`        | map        | e.map(x, p, t) -> e       | `[-1, 0, 1, 2, 3].map(x, x > 0, x * 2) -> [2 4 6]`             |
| `filter`     | filter     | e.filter(x, p) -> e       | `[-1, 0, 1, 2, 3].filter(x, x > 0) -> [1 2 3]`                 |
