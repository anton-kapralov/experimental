(defn solve [a b c]
  (let [d (- (* b b) (* 4 a c))]
    [(/ (+ (- b) (Math/sqrt d)) (* 2 a))
     (/ (- (- b) (Math/sqrt d)) (* 2 a))]))
