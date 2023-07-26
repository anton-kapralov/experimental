package kae.exp.micrograd;

import static com.google.common.collect.ImmutableList.toImmutableList;

import com.google.common.collect.ImmutableList;
import java.util.random.RandomGenerator;
import kae.exp.micrograd.engine.Value;

public final class Layer {
  private final ImmutableList<Neuron> neurons;

  public Layer(RandomGenerator rng, int sizeIn, int sizeOut) {
    var neurons = ImmutableList.<Neuron>builder();
    for (int i = 0; i < sizeOut; i++) {
      neurons.add(new Neuron(rng, sizeIn));
    }
    this.neurons = neurons.build();
  }

  public ImmutableList<Value> call(ImmutableList<Value> xs) {
    return neurons.stream().map(n -> n.call(xs)).collect(toImmutableList());
  }

}
