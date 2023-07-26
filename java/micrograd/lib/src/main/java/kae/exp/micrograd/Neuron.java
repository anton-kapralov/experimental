package kae.exp.micrograd;

import com.google.common.collect.ImmutableCollection;
import com.google.common.collect.ImmutableList;
import com.google.common.collect.Streams;
import java.util.StringJoiner;
import java.util.random.RandomGenerator;
import kae.exp.micrograd.engine.Value;

public final class Neuron {
  private final ImmutableList<Value> ws;
  private final Value b;

  public Neuron(RandomGenerator rng, int sizeIn) {
    ImmutableList.Builder<Value> ws = ImmutableList.builder();
    for (int i = 0; i < sizeIn; i++) {
      ws.add(Value.of(rng.nextDouble(-1, 1)));
    }
    this.ws = ws.build();
    b = Value.of(rng.nextDouble(-1, 1));
  }

  public Value call(ImmutableCollection<Value> xs) {
    return Streams.zip(ws.stream(), xs.stream(), Value::multiply).reduce(b, Value::add).tanh();
  }

  @Override
  public String toString() {
    return new StringJoiner(", ", Neuron.class.getSimpleName() + "[", "]")
        .add("ws=" + ws)
        .add("b=" + b)
        .toString();
  }
}
