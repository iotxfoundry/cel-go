# CEL Standard Library Functions

This document lists the standard functions available in the CEL standard library.
Each function is represented with its name, ID, expression signature, deprecation status, and examples.
The table below provides a comprehensive overview of these functions.

- Standard Functions

| name | id | expr | example |
|------|----|------| ------- |
|`timestamp`|timestamp_to_timestamp|(google.protobuf.Timestamp) -> google.protobuf.Timestamp|timestamp(timestamp('2023-01-01T00:00:00Z')) // timestamp('2023-01-01T00:00:00Z')|
|`timestamp`|int64_to_timestamp|(int) -> google.protobuf.Timestamp|timestamp(1) // timestamp('1970-01-01T00:00:01Z')|
|`timestamp`|string_to_timestamp|(string) -> google.protobuf.Timestamp|timestamp('2025-01-01T12:34:56Z') // timestamp('2025-01-01T12:34:56Z')|
|`bitwise_clear`|bitwise_clear_bytes|(bytes,bytes) -> bytes|b"\xff".bitwise_clear(b"\x0f") // b"\xf0"|
|`bitwise_clear`|int_bitwise_clear_int_int|(int,int) -> int|7.bitwise_clear(1) // 6<br>-1.bitwise_clear(0) // -1|
|`bitwise_clear`|uint_bitwise_clear_uint_uint|(uint,uint) -> uint|7u.bitwise_clear(1u) // 6u<br>255u.bitwise_clear(0u) // 255u|
|`base64.encode`|base64_encode_bytes|(bytes) -> string||
|`div`|arith_div_int64_int64|(int,int) -> int||
|`div`|arith_div_int64_uint64|(int,uint) -> int||
|`div`|arith_div_int64_double|(int,double) -> double||
|`div`|arith_div_uint64_uint64|(uint,uint) -> uint||
|`div`|arith_div_uint64_int64|(uint,int) -> int||
|`div`|arith_div_uint64_double|(uint,double) -> double||
|`div`|arith_div_double_double|(double,double) -> double||
|`div`|arith_div_double_int64|(double,int) -> double||
|`div`|arith_div_double_uint64|(double,uint) -> double||
|`math.sign`|math_sign_double|(double) -> double||
|`math.sign`|math_sign_int|(int) -> int||
|`math.sign`|math_sign_uint|(uint) -> uint||
|`optional.none`|optional_none|() -> optional_type(\<V\>)|optional.none()|
|`dyn`|to_dyn|(A) -> dyn|dyn(1) // 1|
|`math.trunc`|math_trunc_double|(double) -> double||
|`optional.unwrap`|optional_unwrap|(list) -> list(\<V\>)|optional.unwrap([optional.of(1), optional.none()]) // [1]|
|`bitwise_shr`|bitwise_shr_int|(bytes,int) -> bytes|b"\xf0".bitwise_shr(4) // b"\x0f"<br>b"\xff\xff".bitwise_shr(8) // b"\x00\xff"|
|`bitwise_shr`|int_bitwise_shr_int_int|(int,int) -> int|8.bitwise_shr(2) // 2<br>-8.bitwise_shr(2) // -2|
|`bitwise_shr`|uint_bitwise_shr_uint_int|(uint,int) -> uint|8u.bitwise_shr(2) // 2u<br>2u.bitwise_shr(-2) // 8u|
|`_!=_`|not_equals|(A,A) -> bool|1 != 2     // true<br>"a" != "a" // false<br>3.0 != 3.1 // true|
|`_[_]`|index_list|(list,int) -> \<A\>|[1, 2, 3][1] // 2|
|`_[_]`|index_map|(map,A) -> \<B\>|{'key': 'value'}['key'] // 'value'<br>{'key': 'value'}['missing'] // error|
|`_[_]`|optional_list_index_int|(optional_type,int) -> optional_type(\<V\>)||
|`_[_]`|optional_map_index_value|(optional_type,K) -> optional_type(\<V\>)||
|`math.bitShiftLeft`|math_bitShiftLeft_int_int|(int,int) -> int||
|`math.bitShiftLeft`|math_bitShiftLeft_uint_int|(uint,int) -> uint||
|`cel.@block`|cel_block_list|(list,T) -> \<T\>||
|`charAt`|string_char_at_int|(string,int) -> string||
|`getDayOfYear`|timestamp_to_day_of_year|(google.protobuf.Timestamp) -> int|timestamp('2023-01-02T00:00:00Z').getDayOfYear() // 1|
|`getDayOfYear`|timestamp_to_day_of_year_with_tz|(google.protobuf.Timestamp,string) -> int|timestamp('2023-01-01T05:00:00Z').getDayOfYear('America/Los_Angeles') // 364|
|`math.isNaN`|math_isNaN_double|(double) -> bool||
|`bool`|bool_to_bool|(bool) -> bool|bool(true) // true|
|`bool`|string_to_bool|(string) -> bool|bool('true') // true<br>bool('false') // false|
|`unwrapOpt`|optional_unwrapOpt|(list) -> list(\<V\>)|[optional.of(1), optional.none()].unwrapOpt() // [1]|
|`sets.intersects`|list_sets_intersects_list|(list,list) -> bool||
|`getDayOfMonth`|timestamp_to_day_of_month|(google.protobuf.Timestamp) -> int|timestamp('2023-07-14T10:30:45.123Z').getDayOfMonth() // 13|
|`getDayOfMonth`|timestamp_to_day_of_month_with_tz|(google.protobuf.Timestamp,string) -> int|timestamp('2023-07-01T05:00:00Z').getDayOfMonth('America/Los_Angeles') // 29|
|`uint`|uint64_to_uint64|(uint) -> uint|uint(123u) // 123u|
|`uint`|double_to_uint64|(double) -> uint|uint(123.45) // 123u|
|`uint`|int64_to_uint64|(int) -> uint|uint(123) // 123u|
|`uint`|string_to_uint64|(string) -> uint|uint('123') // 123u|
|`math.bitShiftRight`|math_bitShiftRight_int_int|(int,int) -> int||
|`math.bitShiftRight`|math_bitShiftRight_uint_int|(uint,int) -> uint||
|`matches`|matches|(string,string) -> bool|matches('123-456', '^[0-9]+(-[0-9]+)?$') // true<br>matches('hello', '^h.*o$') // true|
|`matches`|matches_string|(string,string) -> bool|'123-456'.matches('^[0-9]+(-[0-9]+)?$') // true<br>'hello'.matches('^h.*o$') // true|
|`swap`|swap_int_int|(bytes,int,int) -> bytes|b"\x01\x02\x03\x04".swap(1, 3) // b"\x01\x04\x03\x02"<br>b"\x01\x02\x03\x04".swap(2, 2) // b"\x01\x02\x03\x04"<br>b"\x01\x02\x03\x04".swap(0, 5) // error: index '5' out of range in bytes size '4'<br>b"\x01\x02\x03\x04".swap(2, 1) // b"\x01\x02\x03\x04"|
|`toi`|toi_int|(bytes,int) -> int|b"\x00\x00\x00\x01".toi(8) // 1<br>b"\x00\x00\x01\x02".toi(16) // 258<br>b"\x00\x01\x02\x03".toi(32) // 16909060<br>b"\x01\x02\x03\x04".toi(64) // 72623859790382856|
|`getMilliseconds`|timestamp_to_milliseconds|(google.protobuf.Timestamp) -> int|timestamp('2023-07-14T10:30:45.123Z').getMilliseconds() // 123|
|`getMilliseconds`|timestamp_to_milliseconds_with_tz|(google.protobuf.Timestamp,string) -> int|timestamp('2023-07-14T10:30:45.123Z').getMilliseconds('America/Los_Angeles') // 123|
|`getMilliseconds`|duration_to_milliseconds|(google.protobuf.Duration) -> int||
|`math.bitOr`|math_bitOr_int_int|(int,int) -> int||
|`math.bitOr`|math_bitOr_uint_uint|(uint,uint) -> uint||
|`math.ceil`|math_ceil_double|(double) -> double||
|`startsWith`|starts_with_string|(string,string) -> bool|'hello world'.startsWith('hello') // true<br>'hello world'.startsWith('world') // false|
|`lastIndexOf`|string_last_index_of_string|(string,string) -> int||
|`lastIndexOf`|string_last_index_of_string_int|(string,string,int) -> int||
|`regex.replace`|regex_replace_string_string_string|(string,string,string) -> string||
|`regex.replace`|regex_replace_string_string_string_int|(string,string,string,int) -> string||
|`math.@min`|math_@min_double|(double) -> double||
|`math.@min`|math_@min_int|(int) -> int||
|`math.@min`|math_@min_uint|(uint) -> uint||
|`math.@min`|math_@min_double_double|(double,double) -> double||
|`math.@min`|math_@min_int_int|(int,int) -> int||
|`math.@min`|math_@min_uint_uint|(uint,uint) -> uint||
|`math.@min`|math_@min_int_uint|(int,uint) -> dyn||
|`math.@min`|math_@min_int_double|(int,double) -> dyn||
|`math.@min`|math_@min_double_int|(double,int) -> dyn||
|`math.@min`|math_@min_double_uint|(double,uint) -> dyn||
|`math.@min`|math_@min_uint_int|(uint,int) -> dyn||
|`math.@min`|math_@min_uint_double|(uint,double) -> dyn||
|`math.@min`|math_@min_list_double|(list) -> double||
|`math.@min`|math_@min_list_int|(list) -> int||
|`math.@min`|math_@min_list_uint|(list) -> uint||
|`math_randf`|math_randf_none|() -> double|math.randf() // [0.0, 1.0) float64|
|`math_randf`|math_randf_int|(int) -> double|math.randf(32) // [0.0, 1.0) float32<br>math.randf(64) // [0.0, 1.0) float64|
|`math_randf`|math_randf_uint|(uint) -> double|math.randf(32) // [0.0, 1.0) float32<br>math.randf(64) // [0.0, 1.0) float64|
|`math_randf`|math_randf_double|(double) -> double|math.randf(32) // [0.0, 1.0) float32<br>math.randf(64) // [0.0, 1.0) float64|
|`delete`|delete_int_int|(bytes,int,int) -> bytes|b"\x01\x02\x03\x04".delete(1, 3) // b"\x01\x04"<br>b"\x01\x02\x03\x04".delete(2, 2) // b"\x01\x02\x03\x04"<br>b"\x01\x02\x03\x04".delete(0, 5) // b""<br>b"\x01\x02\x03\x04".delete(2, 1) // b"\x01\x02\x04"|
|`delete`|delete_int|(bytes,int) -> bytes|b"\x01\x02\x03\x04".delete(1) // b"\x01\x03\x04"<br>b"\x01\x02\x03\x04".delete(2) // b"\x01\x02\x04"<br>b"\x01\x02\x03\x04".delete(0) // b"\x02\x03\x04"<br>b"\x01\x02\x03\x04".delete(3) // b"\x01\x02\x03"|
|`getHours`|timestamp_to_hours|(google.protobuf.Timestamp) -> int|timestamp('2023-07-14T10:30:45.123Z').getHours() // 10|
|`getHours`|timestamp_to_hours_with_tz|(google.protobuf.Timestamp,string) -> int|timestamp('2023-07-14T10:30:45.123Z').getHours('America/Los_Angeles') // 2|
|`getHours`|duration_to_hours|(google.protobuf.Duration) -> int|duration('3723s').getHours() // 1|
|`tof`|tof_int|(bytes,int) -> double|b"\x40\x49\x0f\xdb".tof(32) // 3.141592653589793<br>b"\x40\x09\x21\xfb\x54\x44\x2d\x18".tof(64) // 3.141592653589793|
|`_-_`|subtract_double|(double,double) -> double|10.5 - 2.0 // 8.5|
|`_-_`|subtract_duration_duration|(google.protobuf.Duration,google.protobuf.Duration) -> google.protobuf.Duration|duration('1m') - duration('1s') // duration('59s')|
|`_-_`|subtract_int64|(int,int) -> int|5 - 3 // 2|
|`_-_`|subtract_timestamp_duration|(google.protobuf.Timestamp,google.protobuf.Duration) -> google.protobuf.Timestamp|timestamp('2023-01-10T12:00:00Z')<br>  - duration('12h') // timestamp('2023-01-10T00:00:00Z')|
|`_-_`|subtract_timestamp_timestamp|(google.protobuf.Timestamp,google.protobuf.Timestamp) -> google.protobuf.Duration|timestamp('2023-01-10T12:00:00Z')<br>  - timestamp('2023-01-10T00:00:00Z') // duration('12h')|
|`_-_`|subtract_uint64|(uint,uint) -> uint|// the subtraction result must be positive, otherwise an overflow<br>// error is generated.<br>42u - 3u // 39u|
|`math.sqrt`|math_sqrt_double|(double) -> double||
|`math.sqrt`|math_sqrt_int|(int) -> double||
|`math.sqrt`|math_sqrt_uint|(uint) -> double||
|`hasValue`|optional_hasValue|(optional_type) -> bool|optional.of({1: 2}).hasValue() // true|
|`mul`|arith_mul_int64_int64|(int,int) -> int||
|`mul`|arith_mul_int64_uint64|(int,uint) -> int||
|`mul`|arith_mul_int64_double|(int,double) -> double||
|`mul`|arith_mul_uint64_uint64|(uint,uint) -> uint||
|`mul`|arith_mul_uint64_int64|(uint,int) -> int||
|`mul`|arith_mul_uint64_double|(uint,double) -> double||
|`mul`|arith_mul_double_double|(double,double) -> double||
|`mul`|arith_mul_double_int64|(double,int) -> double||
|`mul`|arith_mul_double_uint64|(double,uint) -> double||
|`_/_`|divide_double|(double,double) -> double|7.0 / 2.0 // 3.5|
|`_/_`|divide_int64|(int,int) -> int|10 / 2 // 5|
|`_/_`|divide_uint64|(uint,uint) -> uint|42u / 2u // 21u|
|`_\|\|_`|logical_or|(bool,bool) -> bool|true \|\| false // true<br>false \|\| false // false<br>error \|\| true // true<br>error \|\| error // true|
|`regex.extract`|regex_extract_string_string|(string,string) -> optional_type(string)||
|`endsWith`|ends_with_string|(string,string) -> bool|'hello world'.endsWith('world') // true<br>'hello world'.endsWith('hello') // false|
|`math.round`|math_round_double|(double) -> double||
|`math.abs`|math_abs_double|(double) -> double||
|`math.abs`|math_abs_int|(int) -> int||
|`math.abs`|math_abs_uint|(uint) -> uint||
|`add`|arith_add_int64_int64|(int,int) -> int||
|`add`|arith_add_int64_uint64|(int,uint) -> int||
|`add`|arith_add_int64_double|(int,double) -> double||
|`add`|arith_add_uint64_uint64|(uint,uint) -> uint||
|`add`|arith_add_uint64_int64|(uint,int) -> int||
|`add`|arith_add_uint64_double|(uint,double) -> double||
|`add`|arith_add_double_double|(double,double) -> double||
|`add`|arith_add_double_int64|(double,int) -> double||
|`add`|arith_add_double_uint64|(double,uint) -> double||
|`math.floor`|math_floor_double|(double) -> double||
|`size`|size_bytes|(bytes) -> int|size(b'123') // 3|
|`size`|bytes_size|(bytes) -> int|b'123'.size() // 3|
|`size`|size_list|(list) -> int|size([1, 2, 3]) // 3|
|`size`|list_size|(list) -> int|[1, 2, 3].size() // 3|
|`size`|size_map|(map) -> int|size({'a': 1, 'b': 2}) // 2|
|`size`|map_size|(map) -> int|{'a': 1, 'b': 2}.size() // 2|
|`size`|size_string|(string) -> int|size('hello') // 5|
|`size`|string_size|(string) -> int|'hello'.size() // 5|
|`sets.contains`|list_sets_contains_list|(list,list) -> bool||
|`indexOf`|string_index_of_string|(string,string) -> int||
|`indexOf`|string_index_of_string_int|(string,string,int) -> int||
|`getMonth`|timestamp_to_month|(google.protobuf.Timestamp) -> int|timestamp('2023-07-14T10:30:45.123Z').getMonth() // 6|
|`getMonth`|timestamp_to_month_with_tz|(google.protobuf.Timestamp,string) -> int|timestamp('2023-01-01T05:30:00Z').getMonth('America/Los_Angeles') // 11|
|`bitwise_popcnt`|bitwise_popcnt_bytes|(bytes) -> int|b"\xf0".bitwise_popcnt() // 4<br>b"\xff\x00".bitwise_popcnt() // 8|
|`_==_`|equals|(A,A) -> bool|1 == 1 // true<br>'hello' == 'world' // false<br>bytes('hello') == b'hello' // true<br>duration('1h') == duration('60m') // true<br>dyn(3.0) == 3 // true|
|`getSeconds`|timestamp_to_seconds|(google.protobuf.Timestamp) -> int|timestamp('2023-07-14T10:30:45.123Z').getSeconds() // 45|
|`getSeconds`|timestamp_to_seconds_tz|(google.protobuf.Timestamp,string) -> int|timestamp('2023-07-14T10:30:45.123Z').getSeconds('America/Los_Angeles') // 45|
|`getSeconds`|duration_to_seconds|(google.protobuf.Duration) -> int|duration('3723.456s').getSeconds() // 3723|
|`_\>=_`|greater_equals_bool|(bool,bool) -> bool|true \>= false // true|
|`_\>=_`|greater_equals_int64|(int,int) -> bool|3 \>= -2 // true|
|`_\>=_`|greater_equals_int64_double|(int,double) -> bool|2 \>= 1.1 // true<br>1 \>= 1.0 // true|
|`_\>=_`|greater_equals_int64_uint64|(int,uint) -> bool|3 \>= 2u // true|
|`_\>=_`|greater_equals_uint64|(uint,uint) -> bool|2u \>= 1u // true|
|`_\>=_`|greater_equals_uint64_double|(uint,double) -> bool|2u \>= 1.9 // true|
|`_\>=_`|greater_equals_uint64_int64|(uint,int) -> bool|23u \>= 1 // true<br>1u \>= 1 // true|
|`_\>=_`|greater_equals_double|(double,double) -> bool|2.4 \>= 2.0 // true|
|`_\>=_`|greater_equals_double_int64|(double,int) -> bool|3.1 \>= 3 // true|
|`_\>=_`|greater_equals_double_uint64|(double,uint) -> bool|2.3 \>= 2u // true|
|`_\>=_`|greater_equals_string|(string,string) -> bool|'b' \>= 'a' // true|
|`_\>=_`|greater_equals_bytes|(bytes,bytes) -> bool|b'world' \>= b'hello' // true|
|`_\>=_`|greater_equals_timestamp|(google.protobuf.Timestamp,google.protobuf.Timestamp) -> bool|timestamp('2001-01-01T02:03:04Z') \>= timestamp('2001-01-01T02:03:04Z') // true|
|`_\>=_`|greater_equals_duration|(google.protobuf.Duration,google.protobuf.Duration) -> bool|duration('60s') \>= duration('1m') // true|
|`_?._`|select_optional_field|(dyn,string) -> optional_type(\<V\>)|msg.?field // optional.of(field) if non-empty, otherwise optional.none()<br>msg.?field.?nested_field // optional.of(nested_field) if both field and nested_field are non-empty.|
|`double`|double_to_double|(double) -> double|double(1.23) // 1.23|
|`double`|int64_to_double|(int) -> double|double(123) // 123.0|
|`double`|string_to_double|(string) -> double|double('1.23') // 1.23|
|`double`|uint64_to_double|(uint) -> double|double(123u) // 123.0|
|`lists.range`|lists_range|(int) -> list(int)||
|`bitwise_not`|bitwise_not_bytes|(bytes) -> bytes|b"\xf0".bitwise_not() // b"\x0f"<br>b"\x00".bitwise_not() // b"\xff"|
|`bitwise_not`|int_bitwise_not_int|(int) -> int|0.bitwise_not() // -1<br>-1.bitwise_not() // 0|
|`bitwise_not`|uint_bitwise_not_uint|(uint) -> uint|0u.bitwise_not() // 18446744073709551615u<br>1u.bitwise_not() // 18446744073709551614u|
|`format`|string_format|(string,list) -> string||
|`_\<=_`|less_equals_bool|(bool,bool) -> bool|false \<= true // true|
|`_\<=_`|less_equals_int64|(int,int) -> bool|-2 \<= 3 // true|
|`_\<=_`|less_equals_int64_double|(int,double) -> bool|1 \<= 1.1 // true|
|`_\<=_`|less_equals_int64_uint64|(int,uint) -> bool|1 \<= 2u // true<br>-1 \<= 0u // true|
|`_\<=_`|less_equals_uint64|(uint,uint) -> bool|1u \<= 2u // true|
|`_\<=_`|less_equals_uint64_double|(uint,double) -> bool|1u \<= 1.0 // true<br>1u \<= 1.1 // true|
|`_\<=_`|less_equals_uint64_int64|(uint,int) -> bool|1u \<= 23 // true|
|`_\<=_`|less_equals_double|(double,double) -> bool|2.0 \<= 2.4 // true|
|`_\<=_`|less_equals_double_int64|(double,int) -> bool|2.1 \<= 3 // true|
|`_\<=_`|less_equals_double_uint64|(double,uint) -> bool|2.0 \<= 2u // true<br>-1.0 \<= 1u // true|
|`_\<=_`|less_equals_string|(string,string) -> bool|'a' \<= 'b' // true<br>'a' \<= 'a' // true<br>'cat' \<= 'cab' // false|
|`_\<=_`|less_equals_bytes|(bytes,bytes) -> bool|b'hello' \<= b'world' // true|
|`_\<=_`|less_equals_timestamp|(google.protobuf.Timestamp,google.protobuf.Timestamp) -> bool|timestamp('2001-01-01T02:03:04Z') \<= timestamp('2002-02-02T02:03:04Z') // true|
|`_\<=_`|less_equals_duration|(google.protobuf.Duration,google.protobuf.Duration) -> bool|duration('1ms') \<= duration('1s') // true|
|`getFullYear`|timestamp_to_year|(google.protobuf.Timestamp) -> int|timestamp('2023-07-14T10:30:45.123Z').getFullYear() // 2023|
|`getFullYear`|timestamp_to_year_with_tz|(google.protobuf.Timestamp,string) -> int|timestamp('2023-01-01T05:30:00Z').getFullYear('-08:00') // 2022|
|`_+_`|add_bytes|(bytes,bytes) -> bytes|b'hi' + bytes('ya') // b'hiya'|
|`_+_`|add_double|(double,double) -> double|3.14 + 1.59 // 4.73|
|`_+_`|add_duration_duration|(google.protobuf.Duration,google.protobuf.Duration) -> google.protobuf.Duration|duration('1m') + duration('1s') // duration('1m1s')|
|`_+_`|add_duration_timestamp|(google.protobuf.Duration,google.protobuf.Timestamp) -> google.protobuf.Timestamp|duration('24h') + timestamp('2023-01-01T00:00:00Z') // timestamp('2023-01-02T00:00:00Z')|
|`_+_`|add_timestamp_duration|(google.protobuf.Timestamp,google.protobuf.Duration) -> google.protobuf.Timestamp|timestamp('2023-01-01T00:00:00Z') + duration('24h1m2s') // timestamp('2023-01-02T00:01:02Z')|
|`_+_`|add_int64|(int,int) -> int|1 + 2 // 3|
|`_+_`|add_list|(list,list) -> list(\<A\>)|[1] + [2, 3] // [1, 2, 3]|
|`_+_`|add_string|(string,string) -> string|"Hello, " + "world!" // "Hello, world!"|
|`_+_`|add_uint64|(uint,uint) -> uint|22u + 33u // 55u|
|`strings.quote`|strings_quote|(string) -> string||
|`duration`|duration_to_duration|(google.protobuf.Duration) -> google.protobuf.Duration|duration(duration('1s')) // duration('1s')|
|`duration`|string_to_duration|(string) -> google.protobuf.Duration|duration('1h2m3s') // duration('3723s')|
|`reverse`|string_reverse|(string) -> string||
|`reverse`|list_reverse|(list) -> list(\<T\>)||
|`join`|list_join|(list) -> string||
|`join`|list_join_string|(list,string) -> string||
|`first`|list_first|(list) -> optional_type(\<V\>)|[].first() // optional.none()<br>[1, 2, 3].first() ? optional.of(1)|
|`bitwise_index`|bitwise_index_int|(bytes,int) -> bytes|b"\x0f".bitwise_index(3) // b"\x01"|
|`bitwise_index`|int_bitwise_index_int_int|(int,int) -> bytes|10.bitwise_index(1) // b"\x01"<br>10.bitwise_index(0) // b"\x00"|
|`bitwise_index`|uint_bitwise_index_uint_int|(uint,int) -> bytes|10u.bitwise_index(1) // b"\x01"<br>10u.bitwise_index(0) // b"\x00"|
|`_\<_`|less_bool|(bool,bool) -> bool|false \< true // true|
|`_\<_`|less_int64|(int,int) -> bool|-2 \< 3 // true<br>1 \< 0 // false|
|`_\<_`|less_int64_double|(int,double) -> bool|1 \< 1.1 // true|
|`_\<_`|less_int64_uint64|(int,uint) -> bool|1 \< 2u // true|
|`_\<_`|less_uint64|(uint,uint) -> bool|1u \< 2u // true|
|`_\<_`|less_uint64_double|(uint,double) -> bool|1u \< 0.9 // false|
|`_\<_`|less_uint64_int64|(uint,int) -> bool|1u \< 23 // true<br>1u \< -1 // false|
|`_\<_`|less_double|(double,double) -> bool|2.0 \< 2.4 // true|
|`_\<_`|less_double_int64|(double,int) -> bool|2.1 \< 3 // true|
|`_\<_`|less_double_uint64|(double,uint) -> bool|2.3 \< 2u // false<br>-1.0 \< 1u // true|
|`_\<_`|less_string|(string,string) -> bool|'a' \< 'b' // true<br>'cat' \< 'cab' // false|
|`_\<_`|less_bytes|(bytes,bytes) -> bool|b'hello' \< b'world' // true|
|`_\<_`|less_timestamp|(google.protobuf.Timestamp,google.protobuf.Timestamp) -> bool|timestamp('2001-01-01T02:03:04Z') \< timestamp('2002-02-02T02:03:04Z') // true|
|`_\<_`|less_duration|(google.protobuf.Duration,google.protobuf.Duration) -> bool|duration('1ms') \< duration('1s') // true|
|`_*_`|multiply_double|(double,double) -> double|3.5 * 40.0 // 140.0|
|`_*_`|multiply_int64|(int,int) -> int|-2 * 6 // -12|
|`_*_`|multiply_uint64|(uint,uint) -> uint|13u * 3u // 39u|
|`_\>_`|greater_bool|(bool,bool) -> bool|true \> false // true|
|`_\>_`|greater_int64|(int,int) -> bool|3 \> -2 // true|
|`_\>_`|greater_int64_double|(int,double) -> bool|2 \> 1.1 // true|
|`_\>_`|greater_int64_uint64|(int,uint) -> bool|3 \> 2u // true|
|`_\>_`|greater_uint64|(uint,uint) -> bool|2u \> 1u // true|
|`_\>_`|greater_uint64_double|(uint,double) -> bool|2u \> 1.9 // true|
|`_\>_`|greater_uint64_int64|(uint,int) -> bool|23u \> 1 // true<br>0u \> -1 // true|
|`_\>_`|greater_double|(double,double) -> bool|2.4 \> 2.0 // true|
|`_\>_`|greater_double_int64|(double,int) -> bool|3.1 \> 3 // true<br>3.0 \> 3 // false|
|`_\>_`|greater_double_uint64|(double,uint) -> bool|2.3 \> 2u // true|
|`_\>_`|greater_string|(string,string) -> bool|'b' \> 'a' // true|
|`_\>_`|greater_bytes|(bytes,bytes) -> bool|b'world' \> b'hello' // true|
|`_\>_`|greater_timestamp|(google.protobuf.Timestamp,google.protobuf.Timestamp) -> bool|timestamp('2002-02-02T02:03:04Z') \> timestamp('2001-01-01T02:03:04Z') // true|
|`_\>_`|greater_duration|(google.protobuf.Duration,google.protobuf.Duration) -> bool|duration('1ms') \> duration('1us') // true|
|`bitwise_shl`|bitwise_shl_int|(bytes,int) -> bytes|b"\xf0".bitwise_shl(4) // b"\x00"<br>b"\x00\xff".bitwise_shl(8) // b"\xff\x00"|
|`bitwise_shl`|int_bitwise_shl_int_int|(int,int) -> int|2.bitwise_shl(2) // 8<br>8.bitwise_shl(-2) // 2|
|`bitwise_shl`|uint_bitwise_shl_uint_int|(uint,int) -> uint|2u.bitwise_shl(2) // 8u<br>8u.bitwise_shl(-2) // 2u|
|`replace`|string_replace_string_string|(string,string,string) -> string||
|`replace`|string_replace_string_string_int|(string,string,string,int) -> string||
|`toui`|toui_int|(bytes,int) -> uint|b"\x00\x00\x00\x01".tou(8) // 1<br>b"\x00\x00\x01\x02".tou(16) // 258<br>b"\x00\x01\x02\x03".tou(32) // 16909060<br>b"\x01\x02\x03\x04".tou(64) // 72623859790382856|
|`optional.of`|optional_of|(V) -> optional_type(\<V\>)|optional.of(1) // optional(1)|
|`@in`|in_list|(A,list) -> bool|2 in [1, 2, 3] // true<br>"a" in ["b", "c"] // false|
|`@in`|in_map|(A,map) -> bool|'key1' in {'key1': 'value1', 'key2': 'value2'} // true<br>3 in {1: "one", 2: "two"} // false|
|`math.bitAnd`|math_bitAnd_int_int|(int,int) -> int||
|`math.bitAnd`|math_bitAnd_uint_uint|(uint,uint) -> uint||
|`_%_`|modulo_int64|(int,int) -> int|3 % 2 // 1|
|`_%_`|modulo_uint64|(uint,uint) -> uint|6u % 3u // 0u|
|`@not_strictly_false`|not_strictly_false|(bool) -> bool||
|`type`|type|(A) -> type(\<A\>)|type(1) // int<br>type('hello') // string<br>type(int) // type<br>type(type) // type|
|`math_randui`|math_randui_none|() -> uint|math.randui() // [0, MaxUint64]|
|`math_randui`|math_randui_int|(int) -> uint|math.randui(32) // [0, MaxUint32]<br>math.randui(64) // [0, MaxUint64]|
|`math_randui`|math_randui_uint|(uint) -> uint|math.randui(32) // [0, MaxUint32]<br>math.randui(64) // [0, MaxUint64]|
|`math_randui`|math_randui_double|(double) -> uint|math.randui(32) // [0, MaxUint32]<br>math.randui(64) // [0, MaxUint64]|
|`flatten`|list_flatten|(list) -> list(\<T\>)||
|`flatten`|list_flatten_int|(list,int) -> list(dyn)||
|`int`|int64_to_int64|(int) -> int|int(123) // 123|
|`int`|double_to_int64|(double) -> int|int(123.45) // 123|
|`int`|duration_to_int64|(google.protobuf.Duration) -> int|int(duration('1s')) // 1000000000|
|`int`|string_to_int64|(string) -> int|int('123') // 123<br>int('-456') // -456|
|`int`|timestamp_to_int64|(google.protobuf.Timestamp) -> int|int(timestamp('1970-01-01T00:00:01Z')) // 1|
|`int`|uint64_to_int64|(uint) -> int|int(123u) // 123|
|`math.@max`|math_@max_double|(double) -> double||
|`math.@max`|math_@max_int|(int) -> int||
|`math.@max`|math_@max_uint|(uint) -> uint||
|`math.@max`|math_@max_double_double|(double,double) -> double||
|`math.@max`|math_@max_int_int|(int,int) -> int||
|`math.@max`|math_@max_uint_uint|(uint,uint) -> uint||
|`math.@max`|math_@max_int_uint|(int,uint) -> dyn||
|`math.@max`|math_@max_int_double|(int,double) -> dyn||
|`math.@max`|math_@max_double_int|(double,int) -> dyn||
|`math.@max`|math_@max_double_uint|(double,uint) -> dyn||
|`math.@max`|math_@max_uint_int|(uint,int) -> dyn||
|`math.@max`|math_@max_uint_double|(uint,double) -> dyn||
|`math.@max`|math_@max_list_double|(list) -> double||
|`math.@max`|math_@max_list_int|(list) -> int||
|`math.@max`|math_@max_list_uint|(list) -> uint||
|`lowerAscii`|string_lower_ascii|(string) -> string||
|`substring`|string_substring_int|(string,int) -> string||
|`substring`|string_substring_int_int|(string,int,int) -> string||
|`bytes`|bytes_to_bytes|(bytes) -> bytes|bytes(b'abc') // b'abc'|
|`bytes`|string_to_bytes|(string) -> bytes|bytes('hello') // b'hello'|
|`getDate`|timestamp_to_day_of_month_1_based|(google.protobuf.Timestamp) -> int|timestamp('2023-07-14T10:30:45.123Z').getDate() // 14|
|`getDate`|timestamp_to_day_of_month_1_based_with_tz|(google.protobuf.Timestamp,string) -> int|timestamp('2023-07-01T05:00:00Z').getDate('America/Los_Angeles') // 30|
|`-_`|negate_double|(double) -> double|-(3.14) // -3.14|
|`-_`|negate_int64|(int) -> int|-(5) // -5|
|`sub`|arith_sub_int64_int64|(int,int) -> int||
|`sub`|arith_sub_int64_uint64|(int,uint) -> int||
|`sub`|arith_sub_int64_double|(int,double) -> double||
|`sub`|arith_sub_uint64_uint64|(uint,uint) -> uint||
|`sub`|arith_sub_uint64_int64|(uint,int) -> int||
|`sub`|arith_sub_uint64_double|(uint,double) -> double||
|`sub`|arith_sub_double_double|(double,double) -> double||
|`sub`|arith_sub_double_int64|(double,int) -> double||
|`sub`|arith_sub_double_uint64|(double,uint) -> double||
|`mod`|arith_mod_int64_int64|(int,int) -> int||
|`mod`|arith_mod_int64_uint64|(int,uint) -> int||
|`mod`|arith_mod_uint64_uint64|(uint,uint) -> uint||
|`mod`|arith_mod_uint64_int64|(uint,int) -> int||
|`bitwise_xor`|bitwise_xor_bytes|(bytes,bytes) -> bytes|b"\x0f".bitwise_xor(b"\xf0") // b"\xff"|
|`bitwise_xor`|int_bitwise_xor_int_int|(int,int) -> int|5.bitwise_xor(3) // 6<br>-1.bitwise_xor(-1) // 0|
|`bitwise_xor`|uint_bitwise_xor_uint_uint|(uint,uint) -> uint|5u.bitwise_xor(3u) // 6u<br>255u.bitwise_xor(255u) // 0u|
|`index`|index_int|(bytes,int) -> bytes|b"\x01\x02\x03\x04".index(1) // b"\x02"<br>b"\x01\x02\x03\x04".index(2) // b"\x03"<br>b"\x01\x02\x03\x04".index(0) // b"\x01"<br>b"\x01\x02\x03\x04".index(3) // b"\x04"|
|`slice`|slice_int_int|(bytes,int,int) -> bytes|b"\x01\x02\x03\x04".slice(1, 3) // b"\x02\x03"<br>b"\x01\x02\x03\x04".slice(2, 2) // b""<br>b"\x01\x02\x03\x04".slice(0, 5) // b"\x01\x02\x03\x04"<br>b"\x01\x02\x03\x04".slice(2, 1) // b""|
|`slice`|list_slice|(list,int,int) -> list(\<T\>)||
|`base64.decode`|base64_decode_string|(string) -> bytes||
|`to_bytes`|int_to_bytes_int|(int,int) -> bytes|42.to_bytes(8) // b"\x2a"<br>42.to_bytes(16) // b"\x00\x2a"<br>42.to_bytes(32) // b"\x00\x00\x00\x2a"<br>42.to_bytes(64) // b"\x00\x00\x00\x00\x00\x00\x00\x2a"<br>42.to_bytes(128) // error: base '128' out of int8, int16, int32, int64 size|
|`to_bytes`|uint_to_bytes_int|(uint,int) -> bytes|42.to_bytes(8) // b"\x2a"<br>42.to_bytes(16) // b"\x00\x2a"<br>42.to_bytes(32) // b"\x00\x00\x00\x2a"<br>42.to_bytes(64) // b"\x00\x00\x00\x00\x00\x00\x00\x2a"<br>42.to_bytes(128) // error: base '128' out of uint8, uint16, uint32, uint64 size|
|`to_bytes`|float64_to_bytes_int|(double,int) -> bytes|3.14.to_bytes(32) // b"\x40\x48\xf5\xc3"<br>3.14.to_bytes(64) // b"\x40\x09\x21\xfb\x54\x44\x2d\x18"<br>3.14.to_bytes(16) // error: base '16' out of float32, float64 size|
|`math.isInf`|math_isInf_double|(double) -> bool||
|`contains`|contains_string|(string,string) -> bool|'hello world'.contains('o w') // true<br>'hello world'.contains('goodbye') // false|
|`split`|string_split_string|(string,string) -> list(string)||
|`split`|string_split_string_int|(string,string,int) -> list(string)||
|`math.isFinite`|math_isFinite_double|(double) -> bool||
|`bitwise_or`|bitwise_or_bytes|(bytes,bytes) -> bytes|b"\x0f".bitwise_or(b"\xf0") // b"\xff"|
|`bitwise_or`|int_bitwise_or_int_int|(int,int) -> int|5.bitwise_or(3) // 7<br>0.bitwise_or(1) // 1|
|`bitwise_or`|uint_bitwise_or_uint_uint|(uint,uint) -> uint|5u.bitwise_or(3u) // 7u<br>0u.bitwise_or(1u) // 1u|
|`_[?_]`|list_optindex_optional_int|(list,int) -> optional_type(\<V\>)|[1, 2, 3][?x] // element value if x is in the list size, else optional.none()|
|`_[?_]`|optional_list_optindex_optional_int|(optional_type,int) -> optional_type(\<V\>)||
|`_[?_]`|map_optindex_optional_value|(map,K) -> optional_type(\<V\>)|map_value[?key] // value at the key if present, else optional.none()<br>// map key-value if index is a valid map key, else optional.none()<br>{0: 2, 2: 4, 6: 8}[?index]|
|`_[?_]`|optional_map_optindex_optional_value|(optional_type,K) -> optional_type(\<V\>)||
|`or`|optional_or_optional|(optional_type,optional_type) -> optional_type(\<V\>)|optional.none().or(optional.of(1)) // optional.of(1)<br>// either a value from the first list, a value from the second, or optional.none()<br>[1, 2, 3][?x].or([3, 4, 5][?y])|
|`cel.@mapInsert`|@mapInsert_map_key_value|(map,K,V) -> map(\<K\>, \<V\>)||
|`cel.@mapInsert`|@mapInsert_map_map|(map,map) -> map(\<K\>, \<V\>)||
|`distinct`|list_distinct|(list) -> list(\<T\>)||
|`_&&_`|logical_and|(bool,bool) -> bool|true && true   // true<br>true && false  // false<br>error && true  // error<br>error && false // false|
|`trim`|string_trim|(string) -> string||
|`!_`|logical_not|(bool) -> bool|!true // false<br>!false // true<br>!error // error|
|`getMinutes`|timestamp_to_minutes|(google.protobuf.Timestamp) -> int|timestamp('2023-07-14T10:30:45.123Z').getMinutes() // 30|
|`getMinutes`|timestamp_to_minutes_with_tz|(google.protobuf.Timestamp,string) -> int|timestamp('2023-07-14T10:30:45.123Z').getMinutes('America/Los_Angeles') // 30|
|`getMinutes`|duration_to_minutes|(google.protobuf.Duration) -> int|duration('3723s').getMinutes() // 62|
|`math_randi`|math_randi_none|() -> int|math.randi() // [0, MaxInt64]|
|`math_randi`|math_randi_int|(int) -> int|math.randi(32) // [0, MaxInt32]<br>math.randi(64) // [0, MaxInt64]|
|`math_randi`|math_randi_int_int|(int,int) -> int|math.randi(32, 64) // [0, 64)<br>math.randi(64, 128) // [0, 128)|
|`math_randi`|math_randi_int_uint|(int,uint) -> int|math.randi(32, 64) // [0, 64)<br>math.randi(64, 128) // [0, 128)|
|`math_randi`|math_randi_int_double|(int,double) -> int|math.randi(32, 64) // [0, 64)<br>math.randi(64, 128) // [0, 128)|
|`math_randi`|math_randi_uint|(uint) -> int|math.randi(32) // [0, MaxInt32]<br>math.randi(64) // [0, MaxInt64]|
|`math_randi`|math_randi_uint_int|(uint,int) -> int|math.randi(32, 64) // [0, 64)<br>math.randi(64, 128) // [0, 128)|
|`math_randi`|math_randi_uint_uint|(uint,uint) -> int|math.randi(32, 64) // [0, 64)<br>math.randi(64, 128) // [0, 128)|
|`math_randi`|math_randi_uint_double|(uint,double) -> int|math.randi(32, 64) // [0, 64)<br>math.randi(64, 128) // [0, 128)|
|`math_randi`|math_randi_double|(double) -> int|math.randi(32) // [0, MaxInt32]<br>math.randi(64) // [0, MaxInt64]|
|`math_randi`|math_randi_double_int|(double,int) -> int|math.randi(32, 64) // [0, 64)<br>math.randi(64, 128) // [0, 128)|
|`math_randi`|math_randi_double_uint|(double,uint) -> int|math.randi(32, 64) // [0, 64)<br>math.randi(64, 128) // [0, 128)|
|`math_randi`|math_randi_double_double|(double,double) -> int|math.randi(32, 64) // [0, 64)<br>math.randi(64, 128) // [0, 128)|
|`upperAscii`|string_upper_ascii|(string) -> string||
|`@sortByAssociatedKeys`|list_int_sortByAssociatedKeys|(list,list) -> list(\<T\>)||
|`@sortByAssociatedKeys`|list_uint_sortByAssociatedKeys|(list,list) -> list(\<T\>)||
|`@sortByAssociatedKeys`|list_double_sortByAssociatedKeys|(list,list) -> list(\<T\>)||
|`@sortByAssociatedKeys`|list_bool_sortByAssociatedKeys|(list,list) -> list(\<T\>)||
|`@sortByAssociatedKeys`|list_google.protobuf.Duration_sortByAssociatedKeys|(list,list) -> list(\<T\>)||
|`@sortByAssociatedKeys`|list_google.protobuf.Timestamp_sortByAssociatedKeys|(list,list) -> list(\<T\>)||
|`@sortByAssociatedKeys`|list_string_sortByAssociatedKeys|(list,list) -> list(\<T\>)||
|`@sortByAssociatedKeys`|list_bytes_sortByAssociatedKeys|(list,list) -> list(\<T\>)||
|`getDayOfWeek`|timestamp_to_day_of_week|(google.protobuf.Timestamp) -> int|timestamp('2023-07-14T10:30:45.123Z').getDayOfWeek() // 5|
|`getDayOfWeek`|timestamp_to_day_of_week_with_tz|(google.protobuf.Timestamp,string) -> int|timestamp('2023-07-16T05:00:00Z').getDayOfWeek('America/Los_Angeles') // 6|
|`orValue`|optional_orValue_value|(optional_type,V) -> \<V\>|// pick the value for the given key if the key exists, otherwise return 'you'<br>{'hello': 'world', 'goodbye': 'cruel world'}[?greeting].orValue('you')|
|`last`|list_last|(list) -> optional_type(\<V\>)|[].last() // optional.none()<br>[1, 2, 3].last() ? optional.of(3)|
|`_?_:_`|conditional|(bool,A,A) -> \<A\>|'hello'.contains('lo') ? 'hi' : 'bye' // 'hi'<br>32 % 3 == 0 ? 'divisible' : 'not divisible' // 'not divisible'|
|`math.bitXor`|math_bitXor_int_int|(int,int) -> int||
|`math.bitXor`|math_bitXor_uint_uint|(uint,uint) -> uint||
|`optional.ofNonZeroValue`|optional_ofNonZeroValue|(V) -> optional_type(\<V\>)|optional.ofNonZeroValue(null) // optional.none()<br>optional.ofNonZeroValue("") // optional.none()<br>optional.ofNonZeroValue("hello") // optional.of('hello')|
|`sets.equivalent`|list_sets_equivalent_list|(list,list) -> bool||
|`bitwise_and`|bitwise_and_bytes|(bytes,bytes) -> bytes|b"\xff".bitwise_and(b"\x0f") // b"\x0f"|
|`bitwise_and`|int_bitwise_and_int_int|(int,int) -> int|5.bitwise_and(3) // 1<br>-1.bitwise_and(7) // 7|
|`bitwise_and`|uint_bitwise_and_uint_uint|(uint,uint) -> uint|5u.bitwise_and(3u) // 1u<br>255u.bitwise_and(15u) // 15u|
|`value`|optional_value|(optional_type) -> \<V\>|optional.of(1).value() // 1<br>optional.none().value() // error|
|`sort`|list_int_sort|(list) -> list(int)||
|`sort`|list_uint_sort|(list) -> list(uint)||
|`sort`|list_double_sort|(list) -> list(double)||
|`sort`|list_bool_sort|(list) -> list(bool)||
|`sort`|list_google.protobuf.Duration_sort|(list) -> list(google.protobuf.Duration)||
|`sort`|list_google.protobuf.Timestamp_sort|(list) -> list(google.protobuf.Timestamp)||
|`sort`|list_string_sort|(list) -> list(string)||
|`sort`|list_bytes_sort|(list) -> list(bytes)||
|`regex.extractAll`|regex_extractAll_string_string|(string,string) -> list(string)||
|`math.bitNot`|math_bitNot_int_int|(int) -> int||
|`math.bitNot`|math_bitNot_uint_uint|(uint) -> uint||
|`string`|string_to_string|(string) -> string|string('hello') // 'hello'|
|`string`|bool_to_string|(bool) -> string|string(true) // 'true'|
|`string`|bytes_to_string|(bytes) -> string|string(b'hello') // 'hello'|
|`string`|double_to_string|(double) -> string|string(-1.23e4) // '-12300'|
|`string`|duration_to_string|(google.protobuf.Duration) -> string|string(duration('1h30m')) // '5400s'|
|`string`|int64_to_string|(int) -> string|string(-123) // '-123'|
|`string`|timestamp_to_string|(google.protobuf.Timestamp) -> string|string(timestamp('1970-01-01T00:00:00Z')) // '1970-01-01T00:00:00Z'|
|`string`|uint64_to_string|(uint) -> string|string(123u) // '123'|
- Standard Macros

| name |  expr | example |
|------|-------|---------|
|`has`|`has:1:false`|// true if the 'address' field exists in the 'user' message<br>has(user.address)<br>// test whether the 'key_name' is set on the map which defines it<br>has({'key_name': 'value'}.key_name) // true<br>// test whether the 'id' field is set to a non-default value on the Expr{} message literal<br>has(Expr{}.id) // false|
|`all`|`all:2:true`|[1, 2, 3].all(x, x \> 0) // true<br>[1, 2, 0].all(x, x \> 0) // false<br>['apple', 'banana', 'cherry'].all(fruit, fruit.size() \> 3) // true<br>[3.14, 2.71, 1.61].all(num, num \< 3.0) // false<br>{'a': 1, 'b': 2, 'c': 3}.all(key, key != 'b') // false<br>// an empty list or map as the range will result in a trivially true result<br>[].all(x, x \> 0) // true|
|`exists`|`exists:2:true`|[1, 2, 3].exists(i, i % 2 != 0) // true<br>[0, -1, 5].exists(num, num \< 0) // true<br>{'x': 'foo', 'y': 'bar'}.exists(key, key.startsWith('z')) // false<br>// an empty list or map as the range will result in a trivially false result<br>[].exists(i, i \> 0) // false<br>// test whether a key name equalling 'iss' exists in the map and the<br>// value contains the substring 'cel.dev'<br>// tokens = {'sub': 'me', 'iss': 'https://issuer.cel.dev'}<br>tokens.exists(k, k == 'iss' && tokens[k].contains('cel.dev'))|
|`exists_one`|`exists_one:2:true`|[1, 2, 2].exists_one(i, i \< 2) // true<br>{'a': 'hello', 'aa': 'hellohello'}.exists_one(k, k.startsWith('a')) // false<br>[1, 2, 3, 4].exists_one(num, num % 2 == 0) // false<br>// ensure exactly one key in the map ends in @acme.co<br>{'wiley@acme.co': 'coyote', 'aa@milne.co': 'bear'}.exists_one(k, k.endsWith('@acme.co')) // true|
|`map`|`map:2:true`|[1, 2, 3].map(x, x * 2) // [2, 4, 6]<br>[5, 10, 15].map(x, x / 5) // [1, 2, 3]<br>['apple', 'banana'].map(fruit, fruit.upperAscii()) // ['APPLE', 'BANANA']<br>// Combine all map key-value pairs into a list<br>{'hi': 'you', 'howzit': 'bruv'}.map(k,<br>    k + ":" + {'hi': 'you', 'howzit': 'bruv'}[k]) // ['hi:you', 'howzit:bruv']|
|`map`|`map:3:true`|// multiply only numbers divisible two, by 2<br>[1, 2, 3, 4].map(num, num % 2 == 0, num * 2) // [4, 8]|
|`filter`|`filter:2:true`|[1, 2, 3].filter(x, x \> 1) // [2, 3]<br>['cat', 'dog', 'bird', 'fish'].filter(pet, pet.size() == 3) // ['cat', 'dog']<br>[{'a': 10, 'b': 5, 'c': 20}].map(m, m.filter(key, m[key] \> 10)) // [['c']]<br>// filter a list to select only emails with the @cel.dev suffix<br>['alice@buf.io', 'tristan@cel.dev'].filter(v, v.endsWith('@cel.dev')) // ['tristan@cel.dev']<br>// filter a map into a list, selecting only the values for keys that start with 'http-auth'<br>{'http-auth-agent': 'secret', 'user-agent': 'mozilla'}.filter(k,<br>     k.startsWith('http-auth')) // ['secret']|
|`optMap`|`optMap:2:true`|// sub with the prefix 'dev.cel' or optional.none()<br>request.auth.tokens.?sub.optMap(id, 'dev.cel.' + id)<br>optional.none().optMap(i, i * 2) // optional.none()|
|`optFlatMap`|`optFlatMap:2:true`|// m = {'key': {}}<br>m.?key.optFlatMap(k, k.?subkey) // optional.none()<br>// m = {'key': {'subkey': 'value'}}<br>m.?key.optFlatMap(k, k.?subkey) // optional.of('value')|
|`randf`|`randf:*:true`||
|`randi`|`randi:*:true`||
|`randui`|`randui:*:true`||
|`getExt`|`getExt:2:true`||
|`hasExt`|`hasExt:2:true`||
|`least`|`least:*:true`||
|`greatest`|`greatest:*:true`||
|`sortBy`|`sortBy:2:true`||
|`all`|`all:3:true`||
|`exists`|`exists:3:true`||
|`existsOne`|`existsOne:3:true`||
|`exists_one`|`exists_one:3:true`||
|`transformList`|`transformList:3:true`||
|`transformList`|`transformList:4:true`||
|`transformMap`|`transformMap:3:true`||
|`transformMap`|`transformMap:4:true`||
|`transformMapEntry`|`transformMapEntry:3:true`||
|`transformMapEntry`|`transformMapEntry:4:true`||
|`bind`|`bind:3:true`||
