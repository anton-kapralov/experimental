package kae.exp.micrograd;

import com.google.common.collect.ImmutableList;
import com.google.common.collect.Streams;
import kae.exp.micrograd.engine.Value;

public interface LossFunction {

  Value calculate(ImmutableList<Value> target, ImmutableList<Value> predicted);

  public static LossFunction meanSquareError() {
    return (target, predicted) -> {
      int n = target.size();
      if (n == 0 || n != predicted.size()) {
        throw new IllegalArgumentException(
            "Both target and predicted must be non-empty and of equal size");
      }
      return Streams.zip(
              target.stream(),
              predicted.stream(),
              (t, p) -> {
                Value v = t.subtract(p);
                return v.multiply(v);
              })
          .reduce(Value::add)
          .map(v -> v.divide(Value.of(n)))
          .orElseThrow();
    };
  }
}
