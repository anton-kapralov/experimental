package kae.exp.micrograd;

import com.google.common.collect.ImmutableCollection;
import com.google.common.collect.ImmutableList;
import java.util.random.RandomGenerator;
import kae.exp.micrograd.engine.Value;

public class MultilayerPerceptron {
  private final ImmutableList<Layer> layers;

  public MultilayerPerceptron(RandomGenerator rng, int[] sizes) {
    var layers = ImmutableList.<Layer>builder();
    for (int i = 0; i < sizes.length - 1; i++) {
      layers.add(new Layer(rng, sizes[i], sizes[i+1]));
    }
    this.layers = layers.build();
  }

  public ImmutableList<Value> call(ImmutableCollection<Value> xs) {
    var result = ImmutableList.copyOf(xs);
    for (Layer layer : layers) {
      result = layer.call(result);
    }
    return result;
  }

  public void clear() {
    for (Layer layer : layers) {
      layer.clear();
    }
  }

  public void descend(double step) {
   for (Layer layer : layers) {
     layer.descend(step);
   }
  }
}
