# Performance comparison of Go mapreduce-like stream processing libraries


|                     | Mariomac | Youthlin | Automi | Lambda |
|---------------------|----------|----------|-------|--------|
| Stars               | 58       | 70       | 774   | 114    |
| Contributors        | 1        | 1        | 3     | 1      |
| Generics            | ✅        | ✅        |  ❌      | ✅      |
| Parallel            | ❌        | ✅        |    ❌   | ✅      |


Discarded entries:
* Jucardi: requires that all the elements in the stream
  implement comparable so I can't do map or reduces e.g. on
  maps.
* Reugn: seems to be focused on other purposes (connection
  to external providers). In-memory operations were complex
  for simple benchmarks.