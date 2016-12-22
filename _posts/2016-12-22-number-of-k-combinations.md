---
layout: post
---

#### Number ways to get K items out of N items

I want to offer an another way to understand
the formula to pick K items from N items.

Let us suppose we have four bubbles, which are raising from the 
bottom of the ocean. As they are racing each other,
you can visualize the order in which they arrive at the surface. They can
arrive as {1 2 3 4} or as {1 2 4 3} or as {1 3 2 4}. If you list all combinations
of arrivals, as I did below, you'll get 24 possible ways.
Or I'd like to think of them as 24 incoming streams of bubbles.
I visualize them like this: in the first row I fix number one, and
then I loop through the remaining items 2, 3, 4. Now I fix 2, and run
through the remaining items: 3 and 4.

    1 1 1 1 1 1 |  2 2 2 2 2 2  |  3 3 3 3 3 3  |  4 4 4 4 4 4 
    2 2 3 3 4 4    1 1 3 3 4 4     1 1 2 2 4 4     1 1 2 2 3 3
    3 4 2 4 2 3    3 4 1 4 1 2     2 4 1 4 1 2     2 3 1 3 1 2
    4 3 4 2 3 2    4 3 4 1 3 1     4 2 4 1 2 1     3 2 3 1 2 1

Those are called the permutations, and here's how I count them:

- there are 4 groups in which I can fix 1, 2, 3 or 4.
- in the remaining items, there are 3 ways for me to fix one of the
  remaining items
- in the remaining of remains (2 items), there are 2 ways for me two fix 
  one of them
- which leaves me with a last item, and it's already fixed in
  its single place.


    4*(4 - 1)*(4 - 2)*(4 - 3) = 4 * 3 * 2 * 1 = 4! = 24

It's important to look at bubbles as permutations and count them as such.

Now we want to count in which streams I have 1, 2, and 3 in the
first 3 positions. Looking at my diagram, I see that they are in the streams:
1, 3, 7, 9, 13, 15. These again are the permutations, only this time they are the permutations
of {1, 2, 3}. 

We already know how to count permutations.

    3! = 6

So out of 24 streams, only 6 give us what we need and we have:

    24/6 = 4!/3! = 4, or
    
    n! / k!  = permutations(n) / permutations (k)                     (1)

There are 4 ways to pick 3 items from 4, and I have to divide the number of
permutations of n by the number of permutations of k.

Permutations divided by permutations - this so far looks good.

But what if now I want to count the ways to get 2 items out of 4? If I'll try to use formula (1),
I'll get 12 ways, and that is too many. We know there are only four streams
where the bubbles 1 and 2 are in the first two positions. Those streams are: 1, 2, 7, 8.

That's where the next visualization comes from - a snake.
Imagine a streams having a head and tail.
The head is where we have the bubbles that we are looking for. The tail is where the rest of the bubbles.
The head is permutating, and the tail is permutating.
They do it independently from each other; we know how to calculate the permuations of both:

    Head: - k!
    Tail: (n-k)!

The total number of snakes is the cross of heads and tails:

     k! * (n-k)!

Thus, the formula to get k elements from n is:

            permutations(snake)                   n!
    ---------------------------------------  =  ---------                (2)
    permutations(head) * permutations(tail)     k! (n-k)!


