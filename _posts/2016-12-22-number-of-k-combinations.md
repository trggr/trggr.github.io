---
layout: post
---

#### Number of Ways to Get K out of N items

I want to offer an another way to understand
the formula to get K out of N items.

Imagine four bubbles raising from the 
bottom of the ocean. As they are racing each other,
you can visualize the order of their arrival at the surface.
It can be {1 2 3 4} or {1 2 4 3} or any other. In all there'll
be 24 combinations.
I'd like to think of them as snakes made out of bubbles.

    1 1 1 1 1 1 |  2 2 2 2 2 2  |  3 3 3 3 3 3  |  4 4 4 4 4 4 
    2 2 3 3 4 4    1 1 3 3 4 4     1 1 2 2 4 4     1 1 2 2 3 3
    3 4 2 4 2 3    3 4 1 4 1 2     2 4 1 4 1 2     2 3 1 3 1 2
    4 3 4 2 3 2    4 3 4 1 3 1     4 2 4 1 2 1     3 2 3 1 2 1

The snakes are called the permutations, and here's how I count them:

- there is a group of snakes with 1 as a head. There's another group with 2
  as the head, another one with 3 as the head, and the last one with 4.
  There are FOUR HEAD groups altogether.
- Once the head is fixed we are left with three other bubbles to choose for the
  neck. This makes THREE NECK sub-groups within each head group.
- Now we are left with two last digits to choose for the body. Each will
  make TWO BODY groups within each neck group.
- Finally, we are left with a single digit which we stick to a single TAIL group.

    FOUR HEADS * THREE NECKS * TWO BODIES * ONE TAIL = 4 * 3 * 2 * 1 = 4! = 24

It's important to look at snakes as permutations and count them as such.

Now we want to count which snakes have their heads, necks, and bodies made exclusively
from 1, 2, and 3.
In my diagram, these are the snakes: 1, 3, 7, 9, 13, 15.
If now we look at their head, neck, and body parts, we are noticing that they
all are, in fact, again the permutations. And we already know how to count
them:

    3! = 6

So now, out of 24 snakes, only six give us what we need and we have:

    24/6 = 4!/3! = 4, or
    
    N!/K!  = permutations(N)/permutations(K)                     (1)

And this is the way to count the number of ways to pick three combinations out of four items.

Permutations divided by permutations - this so far looks good.

Only what if I know want to count the ways to get two items out of four? In other words,
will it work the same if I want to count snakes' heads and necks only?

If I'll try to use formula (1), it'll give me 12 snakes with heads and necks made out of 1 and 2.
This is too many. We know there are only four: 1, 2, 7, 8.

To adjust the calculations, what if I say that I am only interested to classify my snakes
by its first part or by its second part. The first part, the head, is made out of bubbles
that we are looking for, and the second part, the tail, will be made out of the rest.
They do it independently from each other; we know how to calculate the permuations of both:

    Number of interesting heads:  K!
    Number of whatever tails:     (N-K)!

The total number of interesting snakes is the cross of heads and tails:

     K! * (N-K)!

Thus, the formula to get K elements out of N items is:

            permutations(snake)                   N!
    ---------------------------------------  =  ---------                (2)
    permutations(head) * permutations(tail)     K! (N-K)!

In our case, it'll be:

    4!                 24
    -------------- =  ----- = 6
    (4 - 2)! * 2!]    2 * 2
 
This is a universal formula, which gives us six. 

