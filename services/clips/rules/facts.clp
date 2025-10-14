(deftemplate roadSegment
   (slot id)
   (slot vehicleCount)
   (slot avgSpeed)
   (slot capacity)
   (slot type))

(deftemplate incident
   (slot segmentId)
   (slot type)
   (slot severity))

(deftemplate weather
   (slot condition) ;; rain, fog, snow, clear
   (slot visibility))

(deftemplate policy
   (slot congestionCharge)
   (slot publicTransportPriority)
   (slot maxAllowedSpeed))

(deftemplate decision
   (slot action)
   (slot segment)
   (slot from)
   (slot to)
   (slot reason)
   (slot newCycle)
   (slot area)
   (slot priority))

(deftemplate explanation
   (slot text))
