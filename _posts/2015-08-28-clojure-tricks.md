---
layout: post
---

# Clojure Tricks

### Methods in Java
To get the list of the methods of a Java object use `reflect`

    (use 'clojure.reflect 'clojure.pprint)
    (-> (all-ns) first reflect pprint)

    (def x (Object.))
    (->> x reflect :members (map (juxt :name :parameter-types)) pprint)

returns

    ([getClass []]
     [wait []]
     [finalize []]
     [java.lang.Object []]
     [equals [java.lang.Object]]
     [notifyAll []]
     [hashCode []]
     [toString []]
     [registerNatives []]
     [clone []]
     [notify []]
     [wait [long int]]
     [wait [long]])

### Repl and Run
if you want to run you program AND stay in REPL, add

  :repl-options {:init-ns tmp2.fn-browser}

to your project.clj. It'll keep 
