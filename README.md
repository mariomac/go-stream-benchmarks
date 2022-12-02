# Performance comparison of Go mapreduce-like stream processing libraries


|                 | [mariomac/gostream](https://github.com/mariomac/gostream) | [primetalk/goio](https://github.com/Primetalk/goio) | [vladimirvivien/automi](https://github.com/vladimirvivien/automi) | [koss-null/lambda](https://github.com/koss-null/lambda) |
|-----------------|-----------------------------------------------------------|-----------------------------------------------------|-------------------------------------------------------------------|-------------------------------|
| Stars           | 58                                                        | 52                                                  | 774                                                               | 114                           |
| Contributors    | 1                                                         | 1                                                   | 3                                                                 | 1                             |
| Generics        | ✅                                                         | ✅                                                   | ❌                                                                 | ✅                             |
| Parallelization | ❌                                                         | ✅                                                   | ❌                                                                 | ✅                             |



Discarded entries:
* Jucardi: requires that all the elements in the stream
  implement comparable so I can't do map or reduces e.g. on
  maps.
* Reugn: seems to be focused on other purposes (connection
  to external providers). In-memory operations were complex
  for simple benchmarks.