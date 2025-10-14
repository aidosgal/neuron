(defrule public-transport-priority
   (policy (publicTransportPriority TRUE))
   ?r <- (roadSegment (id ?id) (type "arterial") (avgSpeed ?s))
   (test (< ?s 25))
=>
   (assert (decision 
       (action "give_bus_priority") 
       (segment ?id) 
       (priority 1)))
   (assert (explanation (text "Public transport priority enforced")))
)

(defrule congestion-charge
   (policy (congestionCharge TRUE))
   ?r <- (roadSegment (id ?id) (vehicleCount ?vc))
   (test (> ?vc 100))
=>
   (assert (decision 
       (action "activate_congestion_charge") 
       (area "Downtown") 
       (priority 2)))
   (assert (explanation (text "Policy allows congestion charge -> activating")))
)
