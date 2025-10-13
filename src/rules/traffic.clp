(deffacts sample
  (vehicle car)
  (speed 45)
)

(defrule test
  (vehicle ?v)
  (speed ?s)
=>
  (printout t "Detected vehicle: " ?v ", speed: " ?s crlf)
)
