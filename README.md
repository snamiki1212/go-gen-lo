# `go-gen-lo`

Generate lo methods for slice struct.

## Usage

```go
// TODO:
```

## Help

```shell
Generate lo methods for slice struct.

Usage:
  go-gen-lo [flags]

Flags:
  -e, --entity string   target entity name. e.g. User or *User
  -h, --help            help for go-gen-lo
  -i, --input string    input file name
  -o, --output string   output file name
  -s, --slice string    target slice name. e.g. Users
```

## [Example](./example)

- Common case ([src](./example/user.go) / [generated code 1](./example/users_list_gen.go) / [generated code 2](./example/users_ptr_gen.go))

## Support

- ✅ Support
- 🚫 Cannot Support(TODO: add issue)
- ☑️ Not supported yet
- `ー` No Need

<details>
<summary>🚀 Support List for all lo method</summary>

| samber/lo                                                                                                              | Struct        | Std | Extend |
| ---------------------------------------------------------------------------------------------------------------------- | ------------- | --- | ------ |
| [Filter](https://github.com/samber/lo?tab=readme-ov-file#filter)                                                       | `[]T`         | ✅  | ✅     |
| [Map](https://github.com/samber/lo?tab=readme-ov-file#map)                                                             | `[]T`         | ✅  | 🚫     |
| [FilterMap](https://github.com/samber/lo?tab=readme-ov-file#filtermap)                                                 | `[]T`         | ☑️  | ー     |
| [FlatMap](https://github.com/samber/lo?tab=readme-ov-file#flatmap)                                                     | `[]T`         | ☑️  | ー     |
| [Reduce](https://github.com/samber/lo?tab=readme-ov-file#reduce)                                                       | `[]T`         | 🚫  | ー     |
| [ReduceRight](https://github.com/samber/lo?tab=readme-ov-file#reduceright)                                             | `[]T`         | 🚫  | ー     |
| [ForEach](https://github.com/samber/lo?tab=readme-ov-file#foreach)                                                     | `[]T`         | ☑️  | ー     |
| [ForEachWhile](https://github.com/samber/lo?tab=readme-ov-file#foreachwhile)                                           | `[]T`         | ☑️  | ー     |
| [Times](https://github.com/samber/lo?tab=readme-ov-file#times)                                                         | `int`         | ー  | ー     |
| [Uniq](https://github.com/samber/lo?tab=readme-ov-file#uniq)                                                           | `[]T`         | ー  | ー     |
| [UniqBy](https://github.com/samber/lo?tab=readme-ov-file#uniqby)                                                       | `[]T`         | 🚫  | ☑️     |
| [GroupBy](https://github.com/samber/lo?tab=readme-ov-file#groupby)                                                     | `[]T`         | 🚫  | ☑️     |
| [Chunk](https://github.com/samber/lo?tab=readme-ov-file#chunk)                                                         | `[]T`         | ☑️  | ー     |
| [PartitionBy](https://github.com/samber/lo?tab=readme-ov-file#partitionby)                                             | `[]T`         | 🚫  | ー     |
| [Flatten](https://github.com/samber/lo?tab=readme-ov-file#flatten)                                                     | `[]T`         | ー  | ー     |
| [Interleave](https://github.com/samber/lo?tab=readme-ov-file#interleave)                                               | `[][]T`       | ー  | ー     |
| [Shuffle](https://github.com/samber/lo?tab=readme-ov-file#shuffle)                                                     | `[]T`         | ☑️  | ー     |
| [Reverse](https://github.com/samber/lo?tab=readme-ov-file#reverse)                                                     | `[]T`         | ☑️  | ー     |
| [Fill](https://github.com/samber/lo?tab=readme-ov-file#fill)                                                           | `[]T`         | ☑️  | ー     |
| [Repeat](https://github.com/samber/lo?tab=readme-ov-file#repeat)                                                       | `int`         | ー  | ー     |
| [RepeatBy](https://github.com/samber/lo?tab=readme-ov-file#repeatby)                                                   | `int`         | ー  | ー     |
| [KeyBy](https://github.com/samber/lo?tab=readme-ov-file#keyby)                                                         | `[]T`         | 🚫  | ✅     |
| [Associate / SliceToMap](https://github.com/samber/lo?tab=readme-ov-file#associate-alias-slicetomap)                   | `[]T`         | 🚫  | ー     |
| [Drop](https://github.com/samber/lo?tab=readme-ov-file#drop)                                                           | `[]T`         | ☑️  | ー     |
| [DropRight](https://github.com/samber/lo?tab=readme-ov-file#dropright)                                                 | `[]T`         | ☑️  | ー     |
| [DropWhile](https://github.com/samber/lo?tab=readme-ov-file#dropwhile)                                                 | `[]T`         | ☑️  | ー     |
| [DropRightWhile](https://github.com/samber/lo?tab=readme-ov-file#droprightwhile)                                       | `[]T`         | ☑️  | ー     |
| [DropByIndex](https://github.com/samber/lo?tab=readme-ov-file#DropByIndex)                                             | `[]T`         | ☑️  | ー     |
| [Reject](https://github.com/samber/lo?tab=readme-ov-file#reject)                                                       | `[]T`         | ☑️  | ☑️     |
| [RejectMap](https://github.com/samber/lo?tab=readme-ov-file#rejectmap)                                                 | `[]T`         | ☑️  | ー     |
| [FilterReject](https://github.com/samber/lo?tab=readme-ov-file#filterreject)                                           | `[]T`         | ☑️  | ?      |
| [Count](https://github.com/samber/lo?tab=readme-ov-file#count)                                                         | `[]T`         | ー  | ー     |
| [CountBy](https://github.com/samber/lo?tab=readme-ov-file#countby)                                                     | `[]T`         | ☑️  | ー     |
| [CountValues](https://github.com/samber/lo?tab=readme-ov-file#countvalues)                                             | `[]T`         | ー  | ー     |
| [CountValuesBy](https://github.com/samber/lo?tab=readme-ov-file#countvaluesby)                                         | `[]T`         | 🚫  | ー     |
| [Subset](https://github.com/samber/lo?tab=readme-ov-file#subset)                                                       | `[]T`         | ☑️  | ー     |
| [Slice](https://github.com/samber/lo?tab=readme-ov-file#slice)                                                         | `[]T`         | ☑️  | ー     |
| [Replace](https://github.com/samber/lo?tab=readme-ov-file#replace)                                                     | `[]T`         | ー  | ー     |
| [ReplaceAll](https://github.com/samber/lo?tab=readme-ov-file#replaceall)                                               | `[]T`         | ー  | ー     |
| [Compact](https://github.com/samber/lo?tab=readme-ov-file#compact)                                                     | `[]T`         | ー  | ー     |
| [IsSorted](https://github.com/samber/lo?tab=readme-ov-file#issorted)                                                   | `[]T`         | ー  | ☑️     |
| [IsSortedByKey](https://github.com/samber/lo?tab=readme-ov-file#issortedbykey)                                         | `[]T`         | 🚫  | ー     |
| [Splice](https://github.com/samber/lo?tab=readme-ov-file#Splice)                                                       | `[]T`         | ☑️  | ー     |
| [Keys](https://github.com/samber/lo?tab=readme-ov-file#keys)                                                           | `map[any]T`   | ☑️  | ー     |
| [UniqKeys](https://github.com/samber/lo?tab=readme-ov-file#uniqkeys)                                                   | `map[any]T`   | ー  | ー     |
| [HasKey](https://github.com/samber/lo?tab=readme-ov-file#haskey)                                                       | `map[any]T`   | ☑️  | ー     |
| [ValueOr](https://github.com/samber/lo?tab=readme-ov-file#valueor)                                                     | `map[any]T`   | ☑️  | ー     |
| [Values](https://github.com/samber/lo?tab=readme-ov-file#values)                                                       | `map[any]T`   | ☑️  | ー     |
| [UniqValues](https://github.com/samber/lo?tab=readme-ov-file#uniqvalues)                                               | `map[any]T`   | ー  | ー     |
| [PickBy](https://github.com/samber/lo?tab=readme-ov-file#pickby)                                                       | `map[any]T`   | ☑️  | ☑️     |
| [PickByKeys](https://github.com/samber/lo?tab=readme-ov-file#pickbykeys)                                               | `map[any]T`   | ☑️  | ー     |
| [PickByValues](https://github.com/samber/lo?tab=readme-ov-file#pickbyvalues)                                           | `map[any]T`   | ー  | ー     |
| [OmitBy](https://github.com/samber/lo?tab=readme-ov-file#omitby)                                                       | `map[any]T`   | ☑️  | ー     |
| [OmitByKeys](https://github.com/samber/lo?tab=readme-ov-file#omitbykeys)                                               | `map[any]T`   | ☑️  | ー     |
| [OmitByValues](https://github.com/samber/lo?tab=readme-ov-file#omitbyvalues)                                           | `map[any]T`   | ー  | ー     |
| [Entries / ToPairs](https://github.com/samber/lo?tab=readme-ov-file#entries-alias-topairs)                             | `map[any]T`   | ☑️  | ー     |
| [FromEntries / FromPairs](https://github.com/samber/lo?tab=readme-ov-file#fromentries-alias-frompairs)                 | `map[any]T`   | ー  | ー     |
| [Invert](https://github.com/samber/lo?tab=readme-ov-file#invert)                                                       | `map[any]T`   | ー  | ー     |
| [Assign (merge of maps)](https://github.com/samber/lo?tab=readme-ov-file#assign)                                       | `map[any]T`   | ー  | ー     |
| [MapKeys](https://github.com/samber/lo?tab=readme-ov-file#mapkeys)                                                     | `map[any]T`   | ー  | ー     |
| [MapValues](https://github.com/samber/lo?tab=readme-ov-file#mapvalues)                                                 | `map[any]T`   | ー  | ー     |
| [MapEntries](https://github.com/samber/lo?tab=readme-ov-file#mapentries)                                               | `map[any]T`   | ー  | ー     |
| [MapToSlice](https://github.com/samber/lo?tab=readme-ov-file#maptoslice)                                               | `map[any]T`   | ー  | ー     |
| [Range / RangeFrom / RangeWithSteps](https://github.com/samber/lo?tab=readme-ov-file#range--rangefrom--rangewithsteps) | ー            | ー  | ー     |
| [Clamp](https://github.com/samber/lo?tab=readme-ov-file#clamp)                                                         | ー            | ー  | ー     |
| [Sum](https://github.com/samber/lo?tab=readme-ov-file#sum)                                                             | ー            | ー  | ー     |
| [SumBy](https://github.com/samber/lo?tab=readme-ov-file#sumby)                                                         | ー            | ー  | ー     |
| [Mean](https://github.com/samber/lo?tab=readme-ov-file#mean)                                                           | ー            | ー  | ー     |
| [MeanBy](https://github.com/samber/lo?tab=readme-ov-file#meanby)                                                       | ー            | ー  | ー     |
| [RandomString](https://github.com/samber/lo?tab=readme-ov-file#randomstring)                                           | ー            | ー  | ー     |
| [Substring](https://github.com/samber/lo?tab=readme-ov-file#substring)                                                 | `string`      | ー  | ー     |
| [ChunkString](https://github.com/samber/lo?tab=readme-ov-file#chunkstring)                                             | `string`      | ー  | ー     |
| [RuneLength](https://github.com/samber/lo?tab=readme-ov-file#runelength)                                               | `string`      | ー  | ー     |
| [PascalCase](https://github.com/samber/lo?tab=readme-ov-file#pascalcase)                                               | `string`      | ー  | ー     |
| [CamelCase](https://github.com/samber/lo?tab=readme-ov-file#camelcase)                                                 | `string`      | ー  | ー     |
| [KebabCase](https://github.com/samber/lo?tab=readme-ov-file#kebabcase)                                                 | `string`      | ー  | ー     |
| [SnakeCase](https://github.com/samber/lo?tab=readme-ov-file#snakecase)                                                 | `string`      | ー  | ー     |
| [Words](https://github.com/samber/lo?tab=readme-ov-file#words)                                                         | `string`      | ー  | ー     |
| [Capitalize](https://github.com/samber/lo?tab=readme-ov-file#capitalize)                                               | `string`      | ー  | ー     |
| [Ellipsis](https://github.com/samber/lo?tab=readme-ov-file#ellipsis)                                                   | `string`      | ー  | ー     |
| [T2 -> T9](https://github.com/samber/lo?tab=readme-ov-file#t2---t9)                                                    | ー            | ー  | ー     |
| [Unpack2 -> Unpack9](https://github.com/samber/lo?tab=readme-ov-file#unpack2---unpack9)                                | ー            | ー  | ー     |
| [Zip2 -> Zip9](https://github.com/samber/lo?tab=readme-ov-file#zip2---zip9)                                            | ー            | ー  | ー     |
| [ZipBy2 -> ZipBy9](https://github.com/samber/lo?tab=readme-ov-file#zipby2---zipby9)                                    | ー            | ー  | ー     |
| [Unzip2 -> Unzip9](https://github.com/samber/lo?tab=readme-ov-file#unzip2---unzip9)                                    | ー            | ー  | ー     |
| [UnzipBy2 -> UnzipBy9](https://github.com/samber/lo?tab=readme-ov-file#unzipby2---unzipby9)                            | ー            | ー  | ー     |
| [ChannelDispatcher](https://github.com/samber/lo?tab=readme-ov-file#channeldispatcher)                                 | ー            | ー  | ー     |
| [SliceToChannel](https://github.com/samber/lo?tab=readme-ov-file#slicetochannel)                                       | ー            | ー  | ー     |
| [Generator](https://github.com/samber/lo?tab=readme-ov-file#generator)                                                 | ー            | ー  | ー     |
| [Buffer](https://github.com/samber/lo?tab=readme-ov-file#buffer)                                                       | ー            | ー  | ー     |
| [BufferWithTimeout](https://github.com/samber/lo?tab=readme-ov-file#bufferwithtimeout)                                 | ー            | ー  | ー     |
| [FanIn](https://github.com/samber/lo?tab=readme-ov-file#fanin)                                                         | ー            | ー  | ー     |
| [FanOut](https://github.com/samber/lo?tab=readme-ov-file#fanout)                                                       | ー            | ー  | ー     |
| [Contains](https://github.com/samber/lo?tab=readme-ov-file#contains)                                                   | `[]T`         | ー  | ー     |
| [ContainsBy](https://github.com/samber/lo?tab=readme-ov-file#containsby)                                               | `[]T`         | ☑️  | ー     |
| [Every](https://github.com/samber/lo?tab=readme-ov-file#every)                                                         | `[]T`         | ー  | ー     |
| [EveryBy](https://github.com/samber/lo?tab=readme-ov-file#everyby)                                                     | `[]T`         | ☑️  | ー     |
| [Some](https://github.com/samber/lo?tab=readme-ov-file#some)                                                           | `[]T`         | ー  | ー     |
| [SomeBy](https://github.com/samber/lo?tab=readme-ov-file#someby)                                                       | `[]T`         | ☑️  | ー     |
| [None](https://github.com/samber/lo?tab=readme-ov-file#none)                                                           | `[]T`         | ー  | ー     |
| [NoneBy](https://github.com/samber/lo?tab=readme-ov-file#noneby)                                                       | `[]T`         | ☑️  | ー     |
| [Intersect](https://github.com/samber/lo?tab=readme-ov-file#intersect)                                                 | `[]T`         | ー  | ー     |
| [Difference](https://github.com/samber/lo?tab=readme-ov-file#difference)                                               | `[]T`         | ー  | ー     |
| [Union](https://github.com/samber/lo?tab=readme-ov-file#union)                                                         | `[]T`         | ー  | ー     |
| [Without](https://github.com/samber/lo?tab=readme-ov-file#without)                                                     | `[]T`         | ー  | ー     |
| [WithoutEmpty](https://github.com/samber/lo?tab=readme-ov-file#withoutempty)                                           | `[]T`         | ー  | ー     |
| [IndexOf](https://github.com/samber/lo?tab=readme-ov-file#indexof)                                                     | `[]T`         | ☑️  | ー     |
| [LastIndexOf](https://github.com/samber/lo?tab=readme-ov-file#lastindexof)                                             | `[]T`         | ☑️  | ー     |
| [Find](https://github.com/samber/lo?tab=readme-ov-file#find)                                                           | `[]T`         | ☑️  | ☑️     |
| [FindIndexOf](https://github.com/samber/lo?tab=readme-ov-file#findindexof)                                             | `[]T`         | ☑️  | ー     |
| [FindLastIndexOf](https://github.com/samber/lo?tab=readme-ov-file#findlastindexof)                                     | `[]T`         | ☑️  | ー     |
| [FindOrElse](https://github.com/samber/lo?tab=readme-ov-file#findorelse)                                               | `[]T`         | ー  | ー     |
| [FindKey](https://github.com/samber/lo?tab=readme-ov-file#findkey)                                                     | `Map`         | ー  | ー     |
| [FindKeyBy](https://github.com/samber/lo?tab=readme-ov-file#findkeyby)                                                 | `Map`         | ☑️  | ☑️     |
| [FindUniques](https://github.com/samber/lo?tab=readme-ov-file#finduniques)                                             | `[]T`         | ー  | ー     |
| [FindUniquesBy](https://github.com/samber/lo?tab=readme-ov-file#finduniquesby)                                         | `[]T`         | ー  | ー     |
| [FindDuplicates](https://github.com/samber/lo?tab=readme-ov-file#findduplicates)                                       | `[]T`         | ー  | ー     |
| [FindDuplicatesBy](https://github.com/samber/lo?tab=readme-ov-file#findduplicatesby)                                   | `[]T`         | ー  | ー     |
| [Min](https://github.com/samber/lo?tab=readme-ov-file#min)                                                             | `[]T`         | ー  | ー     |
| [MinBy](https://github.com/samber/lo?tab=readme-ov-file#minby)                                                         | `[]T`         | ☑️  | ー     |
| [Earliest](https://github.com/samber/lo?tab=readme-ov-file#earliest)                                                   | `time`        | ー  | ー     |
| [EarliestBy](https://github.com/samber/lo?tab=readme-ov-file#earliestby)                                               | `[]T`         | ー  | ー     |
| [Max](https://github.com/samber/lo?tab=readme-ov-file#max)                                                             | `[]T`         | ー  | ー     |
| [MaxBy](https://github.com/samber/lo?tab=readme-ov-file#maxby)                                                         | `[]T`         | ☑️  | ー     |
| [Latest](https://github.com/samber/lo?tab=readme-ov-file#latest)                                                       | `[]time.Time` | ー  | ー     |
| [LatestBy](https://github.com/samber/lo?tab=readme-ov-file#latestby)                                                   | `[]T`         | ー  | ー     |
| [First](https://github.com/samber/lo?tab=readme-ov-file#first)                                                         | `[]T`         | ☑️  | ー     |
| [FirstOrEmpty](https://github.com/samber/lo?tab=readme-ov-file#FirstOrEmpty)                                           | `[]T`         | ☑️  | ー     |
| [FirstOr](https://github.com/samber/lo?tab=readme-ov-file#FirstOr)                                                     | `[]T`         | ☑️  | ー     |
| [Last](https://github.com/samber/lo?tab=readme-ov-file#last)                                                           | `[]T`         | ☑️  | ー     |
| [LastOrEmpty](https://github.com/samber/lo?tab=readme-ov-file#LastOrEmpty)                                             | `[]T`         | ☑️  | ー     |
| [LastOr](https://github.com/samber/lo?tab=readme-ov-file#LastOr)                                                       | `[]T`         | ☑️  | ー     |
| [Nth](https://github.com/samber/lo?tab=readme-ov-file#nth)                                                             | `[]T`         | ☑️  | ー     |
| [Sample](https://github.com/samber/lo?tab=readme-ov-file#sample)                                                       | `[]T`         | ☑️  | ー     |
| [Samples](https://github.com/samber/lo?tab=readme-ov-file#samples)                                                     | `[]T`         | ☑️  | ー     |
| [Ternary](https://github.com/samber/lo?tab=readme-ov-file#ternary)                                                     | `condition`   | ー  | ー     |
| [TernaryF](https://github.com/samber/lo?tab=readme-ov-file#ternaryf)                                                   | `condition`   | ー  | ー     |
| [If / ElseIf / Else](https://github.com/samber/lo?tab=readme-ov-file#if--elseif--else)                                 | `condition`   | ー  | ー     |
| [Switch / Case / Default](https://github.com/samber/lo?tab=readme-ov-file#switch--case--default)                       | `condition`   | ー  | ー     |
| [IsNil](https://github.com/samber/lo?tab=readme-ov-file#isnil)                                                         | ー            | ー  | ー     |
| [ToPtr](https://github.com/samber/lo?tab=readme-ov-file#toptr)                                                         | ー            | ー  | ー     |
| [Nil](https://github.com/samber/lo?tab=readme-ov-file#nil)                                                             | ー            | ー  | ー     |
| [EmptyableToPtr](https://github.com/samber/lo?tab=readme-ov-file#emptyabletoptr)                                       | ー            | ー  | ー     |
| [FromPtr](https://github.com/samber/lo?tab=readme-ov-file#fromptr)                                                     | ー            | ー  | ー     |
| [FromPtrOr](https://github.com/samber/lo?tab=readme-ov-file#fromptror)                                                 | ー            | ー  | ー     |
| [ToSlicePtr](https://github.com/samber/lo?tab=readme-ov-file#tosliceptr)                                               | `[]T`         | ー  | ー     |
| [FromSlicePtr](https://github.com/samber/lo?tab=readme-ov-file#fromsliceptr)                                           | `[]T`         | ー  | ー     |
| [FromSlicePtrOr](https://github.com/samber/lo?tab=readme-ov-file#fromsliceptror)                                       | `[]T`         | ー  | ー     |
| [ToAnySlice](https://github.com/samber/lo?tab=readme-ov-file#toanyslice)                                               | `[]T`         | ー  | ー     |
| [FromAnySlice](https://github.com/samber/lo?tab=readme-ov-file#fromanyslice)                                           | ー            | ー  | ー     |
| [Empty](https://github.com/samber/lo?tab=readme-ov-file#empty)                                                         | ー            | ー  | ー     |
| [IsEmpty](https://github.com/samber/lo?tab=readme-ov-file#isempty)                                                     | ー            | ー  | ー     |
| [IsNotEmpty](https://github.com/samber/lo?tab=readme-ov-file#isnotempty)                                               | ー            | ー  | ー     |
| [Coalesce](https://github.com/samber/lo?tab=readme-ov-file#coalesce)                                                   | ー            | ー  | ー     |
| [CoalesceOrEmpty](https://github.com/samber/lo?tab=readme-ov-file#coalesceorempty)                                     | ー            | ー  | ー     |
| [Partial](https://github.com/samber/lo?tab=readme-ov-file#partial)                                                     | ー            | ー  | ー     |
| [Partial2 -> Partial5](https://github.com/samber/lo?tab=readme-ov-file#partial2---partial5)                            | ー            | ー  | ー     |
| [Attempt](https://github.com/samber/lo?tab=readme-ov-file#attempt)                                                     | ー            | ー  | ー     |
| [AttemptWhile](https://github.com/samber/lo?tab=readme-ov-file#attemptwhile)                                           | ー            | ー  | ー     |
| [AttemptWithDelay](https://github.com/samber/lo?tab=readme-ov-file#attemptwithdelay)                                   | ー            | ー  | ー     |
| [AttemptWhileWithDelay](https://github.com/samber/lo?tab=readme-ov-file#attemptwhilewithdelay)                         | ー            | ー  | ー     |
| [Debounce](https://github.com/samber/lo?tab=readme-ov-file#debounce)                                                   | ー            | ー  | ー     |
| [DebounceBy](https://github.com/samber/lo?tab=readme-ov-file#debounceby)                                               | ー            | ー  | ー     |
| [Synchronize](https://github.com/samber/lo?tab=readme-ov-file#synchronize)                                             | ー            | ー  | ー     |
| [Async](https://github.com/samber/lo?tab=readme-ov-file#async)                                                         | ー            | ー  | ー     |
| [Transaction](https://github.com/samber/lo?tab=readme-ov-file#transaction)                                             | ー            | ー  | ー     |
| [WaitFor](https://github.com/samber/lo?tab=readme-ov-file#waitfor)                                                     | ー            | ー  | ー     |
| [WaitForWithContext](https://github.com/samber/lo?tab=readme-ov-file#waitforwithcontext)                               | ー            | ー  | ー     |
| [Validate](https://github.com/samber/lo?tab=readme-ov-file#validate)                                                   | ー            | ー  | ー     |
| [Must](https://github.com/samber/lo?tab=readme-ov-file#must)                                                           | ー            | ー  | ー     |
| [Try](https://github.com/samber/lo?tab=readme-ov-file#try)                                                             | ー            | ー  | ー     |
| [Try1 -> Try6](https://github.com/samber/lo?tab=readme-ov-file#try0-6)                                                 | ー            | ー  | ー     |
| [TryOr](https://github.com/samber/lo?tab=readme-ov-file#tryor)                                                         | ー            | ー  | ー     |
| [TryOr1 -> TryOr6](https://github.com/samber/lo?tab=readme-ov-file#tryor0-6)                                           | ー            | ー  | ー     |
| [TryCatch](https://github.com/samber/lo?tab=readme-ov-file#trycatch)                                                   | ー            | ー  | ー     |
| [TryWithErrorValue](https://github.com/samber/lo?tab=readme-ov-file#trywitherrorvalue)                                 | ー            | ー  | ー     |
| [TryCatchWithErrorValue](https://github.com/samber/lo?tab=readme-ov-file#trycatchwitherrorvalue)                       | ー            | ー  | ー     |
| [ErrorsAs](https://github.com/samber/lo?tab=readme-ov-file#errorsas)                                                   | ー            | ー  | ー     |

</details>

## E2E

```shell
$ go generate ./example
$ go run ./example
```

## TODO

- Support all slice lo.
- Support map struct as well as slice.

## LICENSE

[MIT](./LICENSE)
