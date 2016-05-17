# Algorithms & Data Structures [![BuildStatus](https://travis-ci.org/jpfuentes2/algorithms-and-datastructures.svg?branch=master)](https://travis-ci.org/jpfuentes2/algorithms-and-datastructures) [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/jpfuentes2/algorithms-and-datastructures) [![Coverage Status](https://coveralls.io/repos/github/jpfuentes2/algorithms-and-datastructures/badge.svg?branch=master)](https://coveralls.io/github/jpfuentes2/algorithms-and-datastructures?branch=master)

I never took a Computer Science course in university since my major was
Philosophy and my goal was to be a lawyer. I dropped out after two years.
Phew, glad I dodged that bullet!

During my travails as a self-taught programmer I have sometimes felt
unprepared to solve certain problems such as implementing a DB or in
whiteboard interviews. Lately, my curiosity is captivated by distributed
systems which frequently use algorithms and data structures foreign to me.

I want to fill the void in my education as a professional, because choking
during whiteboard interviews sucks :wink:. Really, I had no choice given my
obsession with learning everything I can to improve as a professional
programmer.

## Goals

I created this repository to implement common algorithms/data structures
with the following goals:

0. Have *fun*.
1. LEARN. The code in this repository is solely for learning
   and self-practice only; it is not intended for production.
2. Read as much as I can stomach<sup>0</sup> using varied & multiple
   [resources](#resources) so that I may fully understand the material
   by learning from myriad teachers. This can be used successfully as
   [distributed practice](http://digitalpromise.org/2015/02/07/five-learning-strategies-that-work/#distributedpractice),
   a highly effective learning technique.

   > <sup>0</sup> Some books are to be tasted, others to be swallowed, and some few to be chewed and digested: that is, some books are to be read only in parts, others to be read, but not curiously, and some few to be read wholly, and with diligence and attention.
   >
   > -- Francis Bacon
3. Implement common/interesting/fundamental algorithms/data structures.
   This, of course, is [practice testing](http://digitalpromise.org/2015/02/07/five-learning-strategies-that-work/#practicetesting),
   another highly effective learning technique when the stakes are low.
4. >= 90% test coverage.
5. Provide abundant documentation, links, and resources so
   others can learn from this work.

## TOC

* [Why Go?](#why-go)
* Data Structures: :snowflake: = immutable, :lock: = thread-safe
  * [SinglyLinkedList](datastructures/list.go)
  * [Queue :lock:](datastructures/queue.go)
  * [Stack :lock:](datastructures/stack.go)
  * [Set](datastructures/set.go)
  * [Binary Search Tree](datastructures/bstree/bstree.go)
* Algorithms: :floppy_disk: = in-place, :arrows_clockwise: = iterative, :recycle: = recursive
  * Sorting
    * [Selection :floppy_disk:](algorithms/selection_sort.go)
    * [Insertion :floppy_disk:](algorithms/insertion_sort.go)
  * [Fisher-Yates Shuffle :floppy_disk:](algorithms/basics.go)
  * [Reverse :arrows_clockwise: :floppy_disk:](algorithms/basics.go)
  * [Binary Search :arrows_clockwise: :recycle: :floppy_disk:](algorithms/binary_search.go)
  * [IsPalindrome :arrows_clockwise: :recycle:](algorithms/basics.go)
  * [Fibonacci :arrows_clockwise: :recycle:](algorithms/fibonacci.go)
* [Resources](#resources)
* [Thank you](#thank-you)

## <a name="why-go"></a> Why Go?

I chose Go for reasons which FP advocates may disagree; however, what's most
important is what I learn and having fun.

In some cases I eschew polymorphism by using primitive built-in types
because there's often no advantage to using `interface{}` for these
implementations. My reasoning: none of the code is meant for production and using `int`
or `string` makes testing easy and concise.

> Data dominates. If you've chosen the right data structures and
> organized things well, the algorithms will almost always be
> self-evident. Data structures, not algorithms, are central to
> programming.
>
> -- Rob Pike

## <a name="resources"></a> Resources

### Books/Blogs/etc

* [Introduction to Algorithms, 3rd Edition (CLRS)](http://www.amazon.com/Introduction-Algorithms-Thomas-H-Cormen-ebook/dp/B007CNRCAO/ref=dp_kinw_strp_1)
* [Algorithms, 4th Edition, Robert Sedgewick, et al.](http://www.amazon.com/Algorithms-4th-Robert-Sedgewick/dp/032157351X)
* [Khan Academy: Algorithms](https://www.khanacademy.org/computing/computer-science/algorithms)
* [Wikibooks Data Structures](https://en.wikibooks.org/wiki/Data_Structures)

### Videos

* [Algorithms Part I, Robert Sedgewick](https://www.youtube.com/watch?v=YIFWCpquoS8&list=PLUX6FBiUa2g4YWs6HkkCpXL6ru02i7y3Q)
* [Algorithms Part II, Robert Sedgewick](https://www.youtube.com/watch?v=0B745ZPxdBE&list=PLqD_OdMOd_6YixsHkd9f4sNdof4IhIima)

## <a name="thank-you"></a> Thank You

* [@deathbob](https://github.com/deathbob) for general feedback and copy editing of README.md
* [@amosley](https://github.com/amosley) for general feedback and copy editing of README.md
