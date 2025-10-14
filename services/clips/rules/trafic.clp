(deffacts input
  (area "Downtown")
  (sector)
  (avg_speed 18)
  (public_transport_usage 0.25)
  (air_quality_index 120))

(defrule recommend-public-transport
  (avg_speed ?s&:(< ?s 20))
  (public_transport_usage ?u&:(< ?u 0.3))
=>
  (assert (recommendation "Increase public transport frequency")))
