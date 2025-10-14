(defrule detect-high-congestion
   ?r <- (roadSegment (id ?id) (vehicleCount ?vc) (avgSpeed ?s) (capacity ?cap))
   (test (and (> ?vc 100) (< ?s 20)))
=>
   (assert (decision 
       (action "adjust_signal") 
       (segment ?id) 
       (newCycle "extend_green_by_15s")
       (priority 1)))
   (assert (explanation (text (str-cat ?id " congested: vehicleCount " ?vc ", avgSpeed " ?s " -> HIGH congestion"))))
)

(defrule detect-medium-congestion
   ?r <- (roadSegment (id ?id) (vehicleCount ?vc) (avgSpeed ?s))
   (test (and (> ?vc 50) (<= ?vc 100) (< ?s 30)))
=>
   (assert (decision 
       (action "adjust_signal") 
       (segment ?id) 
       (newCycle "extend_green_by_5s")
       (priority 2)))
   (assert (explanation (text (str-cat ?id " moderate congestion detected"))))
)

