---
layout: post
---

I like Clojure and often need to explore its data structures.
When the structure is large, I'd prefer some UI to scroll
and drill down. Of course, there's `clojure.inspector`, but I
want more. 

Swing is part of JVM installation, so its available to the program
without any additional installation, but Swing is hard.
Clojure has a framework called Seesaw, which uses Swing,
it has a nice Clojure-lke syntax, but without knowing Swing well,
it's not much use to me.

So I learned some basic Swing. In the process, I also learned interops with Java.

Our goal is to make a Clojure program which shows a list functions
from `clojure.core` namespace. The program allows you to use to use F3 and F4 keys
(or Doc and Source buttons) to show a documentation and a source code for
a highlighted function. 

In the first part, we import the Java classes which we plan to use.
Technically, the import is not required but it makes the code more compact
I prefer to use `KeyListener` instead of a fully qualified 
`java.awt.event.KeyListener`. It's shorter.

    (ns tmp2.fn-browser
      (:import [javax.swing AbstractAction JFrame JLabel JColorChooser JPanel 
                            BorderFactory JList JScrollPane JButton
                            JOptionPane JTextArea KeyStroke]
               [java.util Vector]
               [java.awt BorderLayout Color Dimension FlowLayout Font TextField]
               [java.awt.event ActionListener KeyListener KeyAdapter KeyEvent] 
               [javax.swing.event ChangeListener]))


Next, we get a list of functions and store them in JList class.

    (def nsname 'clojure.core)

    (def  list1       (->> (ns-publics nsname)
                           (map key)
                           sort
                           Vector.
                           JList.))

Note the dots at the end of `Vector.` and `JList.`. The dot means a class constructor.
Thus, the last two lines are the equivalent to Java: 

    list1 = new Jlist(new Vector(...))

To handle a keyboard input we will make a function which receives a keyboard event.
The function will be called every time a users presses a key during navigation through list.
`JList` already can handle arrows, PgUp, PgDown, and probably other keys. 
All we want is to extend the standard behavior to include the keys F3 and F4, each of which
will call functions `show-doc` and `show-src`.

    (defn key-event [e]
       (let [k (.getKeyCode e)]
         (cond
           (= k KeyEvent/VK_F3)  (show-doc)
           (= k KeyEvent/VK_F4)  (show-src)
           :default nil)))

Now we attach the function to a listener, and then attach a listener to our `list1`:

    (doto list1       (. addKeyListener (proxy [KeyAdapter] [] (keyPressed [e] (key-event e))))
                      (. setSelectedIndex 1)
                      (. ensureIndexIsVisible 1))

Please note another Java interop operator - `doto`. It is useful when you want to call
several methods of the same object. The code above is the same as:

    list1.addKeyListener(new KeyAdapter() {
                           public void keyPressed(KeyEvent e) {
                                  key_event(e);
                           }
    };
    list1.setSelectedIndex(1);
    list1.ensureIndexVisible(1);

Now to the UI portion of a program. We organize the components of a program
in the following hierarchy:

    frame
        scroll-pane
            list1
                KeyAdapter
                    key-event
                        show-doc
                        show-src
        pane
            button-doc
                ActionListener
                    show-doc
            button-src
                ActionListener
                    show-src

We create a frame, which holds a scroll pane and a pane.
In turn, the scroll pane holds our list and pane contains two buttons, 
which we will use to show documentation and a source of a function.

The code to implement it is:

    (def  scroll-pane (JScrollPane. list1))

    (def  button-doc  (JButton. "Doc"))
    (doto button-doc  (. addActionListener (reify ActionListener
                          (actionPerformed [_ e] (show-doc)))))

    (def  button-src  (JButton. "Source"))
    (doto button-src  (. addActionListener (reify ActionListener
                          (actionPerformed [_ e] (show-src)))))

    (def  pane1       (JPanel. (FlowLayout.)))
    (doto pane1       (. add nstext)
                      (. add button-doc)
                      (. add button-src))

    (def  frame       (JFrame. "Interactive"))
    (doto frame       (. setSize (Dimension. 400 600))
                      (. add scroll-pane)
                      (. add BorderLayout/SOUTH pane1)
                      (. setVisible true))

Please note how I formatted the text. 
Even though the indentation is not exactly regular Clojure standard,
I find it is easier to read.
Another item which I found useful is to use two macros per component - `def` and `doto`.

First macro creates a component:

    (def frame (Frame. ...))

the second macro handles it.

    (doto frame ... )

This way each component has a name, and I can easily pick or poke at its value in REPL.
I could have combine it all into one call, but I really like this way of formating
because it breathes and is easier to read than one clob of text:

    (def frame (doto (Frame. ....)))

To finish with the layout, we request a focus on our list, so the keyboard event
would reach it.

    (doto scroll-pane (. requestFocusInWindow))

The focus request should be called AFTER the frame becomes visible.

Finally, the two functions (and a helper function) to show the doc string and the source of a function.

    (defn show-text [text]
      (JOptionPane/showMessageDialog frame
                                     (JScrollPane. (JTextArea. text 40 60))
                                     "Show"
                                     JOptionPane/INFORMATION_MESSAGE))
    (defn show-doc []
      (when-let [selected-fn (.getSelectedValue list1)]
        (show-text (-> (ns-resolve nsname selected-fn) meta :doc))))

    (defn show-src []
      (when-let [selected-fn (.getSelectedValue list1)]
        (show-text (or (clojure.repl/source-fn selected-fn) ""))))

The full text of the program is on the 
[bitbucket](https://bitbucket.org/tashepkov/tmp2/src/70bd2cfc44128308ceb47da8708b6d26c9a40587/src/tmp2/fn-browser.clj?at=default)

## Workflow

I edited the file fn-browser.clj in my editor. Simultaneously I had `lein repl` running in the parallel window.
I modified the file, and after saving it, I ran it on REPL using `(load-file "src/tmp2/fn-browser.clj")`.
The REPL displayed the frame, which I test and simultaneously update in REPL.
When I was happy with the result, I committed the code to Mercurial quite often:

    hg stat
    hg commit
    hg push    (to bitbucket)

## Conclusion

What was good:

  1. Formatting of the clojure code with extra spaces makes it easier to read;
  2. Usage of a pure Swing without any external packages simplifies life (at least for me);
  3. Ability to use my editor and file manager with REPL keeps my workflow, even though it requires
     constant copying and pasting between the editor and the REPL;
  4. Mercurial and Bitbucket rocks!

What I wish to improve:

1. Seesaw's code looks nice, but without knowing Swing very well, it's hard to understand
   sometime what is not working and what I am using wrong - whether its Seesaw or Swing itself.
2. The development environment could have be more smooth.





