# Algorithms & Data Structures

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

1. Have *fun*.
2. Learn, learn, learn. The code in this repository is solely for learning
   and self-practice only; it is not intended for production.
3. Read as much as I can stomach using varied & multiple
   [resources](#resources) so that I may fully understand the material
   by learning from myriad teachers. This can be used successfully as
   [distributed practice](http://digitalpromise.org/2015/02/07/five-learning-strategies-that-work/#distributedpractice),
   a highly effective learning technique.
4. Implement common/interesting/fundamental algorithms/data structures.
   This, of course, is [practice testing](http://digitalpromise.org/2015/02/07/five-learning-strategies-that-work/#practicetesting),
   another highly effective learning technique when the stakes are low.
5. Complete 4. in both Haskell & Go. Why? Well, a friend once called me
   a "try hard." It's true.

## TOC

* [Why Go?](#why-go)
* [Why Haskell?](#why-haskell)
* Data Structures
  * [Go](go/datastructures)
    * [SinglyLinkedList](go/datastructures/list.go)
    * [Queue](go/datastructures/queue.go)
    * [Stack](go/datastructures/stack.go)
    * [Set](go/datastructures/set.go)
  * [Haskell](haskell/datastructures)
* Algorithms
  * [Go](go/algorithms)
    * [Fisher-Yates Shuffle](go/algorithms/basics.go)
    * [In-place Reverse](go/algorithms/basics.go)
    * [Binary Search](go/algorithms/basics.go)
  * [Haskell](haskell/algorithms)
    * [Fisher-Yates Shuffle](haskell/algorithms/basics.hs)
    * [In-place Reverse](haskell/algorithms/basics.hs)
    * [Binary Search](haskell/algorithms/basics.hs)
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

## <a name="why-haskell"></a> Why Haskell?

I chose Haskell for reasons which Go advocates may disagree; however,
what's most important is what I learn and having fun.

Haskell represents interesting challenges as a pure language: not all
common algorithms/data structures have a 1:1 mapping to Haskell and when
they do the complexity is often worse due to the ubiquity of
linked lists. In some cases I ignore the inefficiency in favor of easy
implementation, in others, time/space complexity are respected with the
tradeoff of inelegant code.

Further, sometimes, I may even use mutation for an imperative approach.
_Gasp!_

> In short, Haskell is the world’s finest imperative programming
> language.
>
> -- Simon Peyton Jones

## <a name="resources"></a> Resources

No worries, I'm not using affiliate links.

* [Introduction to Algorithms, 3rd Edition](http://www.amazon.com/Introduction-Algorithms-Thomas-H-Cormen-ebook/dp/B007CNRCAO/ref=dp_kinw_strp_1)
* [Algorithms, 4th Edition](http://www.amazon.com/Algorithms-4th-Robert-Sedgewick/dp/032157351X)
* [Khan Academy: Algorithms](https://www.khanacademy.org/computing/computer-science/algorithms)
* [Purely Functional Data Structures](http://www.amazon.com/Purely-Functional-Structures-Chris-Okasaki/dp/0521663504)
* [Wikibooks Data Structures](https://en.wikibooks.org/wiki/Data_Structures)
* [Scalacaster by @vkostyukov](https://github.com/vkostyukov/scalacaster)

## <a name="thank-you"></a> Thank You

* [@deathbob](https://github.com/deathbob) for general feedback and copy editing of README.md
* [@amosley](https://github.com/amosley) for general feedback and copy editing of README.md
