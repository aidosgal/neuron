(defrule reroute-accidents
   ?i <- (incident (segmentId ?seg) (type "accident") (severity ?sev&:(>= ?sev 2)))
   ?r <- (roadSegment (id ?seg))
=>
   (assert (decision 
       (action "reroute") 
       (from ?seg) 
       (to "R3") 
       (reason (str-cat "accident_on_" ?seg))
       (priority 1)))
   (assert (explanation (text (str-cat "Accident on " ?seg " -> reroute recommended"))))
)

(defrule emergency-priority
   ?i <- (incident (segmentId ?seg) (type "emergency"))
=>
   (assert (decision 
       (action "clear_path") 
       (segment ?seg) 
       (priority 0)))
   (assert (explanation (text (str-cat "Emergency vehicle detected on " ?seg " -> clear path"))))
)
