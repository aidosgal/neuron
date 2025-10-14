(defrule rain-slowdown
   ?w <- (weather (condition "rain") (visibility ?v))
   ?r <- (roadSegment (id ?id))
   (test (< ?v 1))
=>
   (assert (explanation (text (str-cat "Rain on " ?id " -> vehicles slow down due to low visibility"))))
   (assert (decision 
       (action "adjust_signal") 
       (segment ?id) 
       (newCycle "extend_yellow_by_3s")
       (priority 3)))
)

(defrule fog-warning
   ?w <- (weather (condition "fog") (visibility ?v))
   (test (< ?v 0.5))
=>
   (assert (explanation (text "Fog detected -> advise caution and reduce speed")))
)
