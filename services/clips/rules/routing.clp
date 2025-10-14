(defrule prioritize-decisions
   ?d1 <- (decision (action ?a1) (segment ?s) (priority ?p1))
   ?d2 <- (decision (action ?a2) (segment ?s) (priority ?p2&:(< ?p2 ?p1)))
=>
   (retract ?d2)
   (assert (explanation (text (str-cat "Removed lower priority decision " ?a2 " on segment " ?s))))
)
