module github.com/milkcoke/bitcoin

go 1.17

replace github.com/milkcoke/bitcoin/chapter_1 => ./chapter_1

replace github.com/milkcoke/bitcoin/chapter_2 => ./chapter_2

replace github.com/milkcoke/bitcoin/chapter_3 => ./chapter_3

require (
	github.com/milkcoke/bitcoin/chapter_1 v1.0.0
	github.com/milkcoke/bitcoin/chapter_3 v1.0.0
)
