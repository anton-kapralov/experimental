package kae.exp.micrograd.engine;

import static com.google.common.collect.ImmutableList.toImmutableList;

import com.google.common.collect.ImmutableList;
import java.util.Arrays;
import java.util.Collection;
import java.util.StringJoiner;
import java.util.function.Consumer;

public abstract class Value {

  double data;
  final ImmutableList<Value> children;
  private final Operation operation;
  private double grad;

  Value(double data, Collection<Value> children, Operation operation) {
    this.data = data;
    this.children = ImmutableList.copyOf(children);
    this.operation = operation;
  }

  public double asDouble() {
    return data;
  }

  public String asExpression() {
    StringBuilder sb = new StringBuilder();
    appendTo(sb);
    return sb.toString();
  }

  void appendTo(StringBuilder sb) {
    if (children.isEmpty()) {
      sb.append(data);
      return;
    }

    if (children.size() == 1) {
      sb.append(operation);
    }
    sb.append('(');
    int i = 0;
    for (Value child : children) {
      if (i > 0) {
        sb.append(' ');
        sb.append(operation);
        sb.append(' ');
      }
      child.appendTo(sb);
      i++;
    }

    sb.append(')');
  }

  public double grad() {
    return grad;
  }

  public void backward() {
    backward(null);
  }

  private void backward(Value parent) {
    grad += parent == null ? 1 : parent.derivativeWithRespectTo(this);
    for (Value child : children) {
      child.backward(this);
    }
  }

  private double derivativeWithRespectTo(Value child) {
    return grad * localDerivativeWithRespectTo(child);
  }

  abstract double localDerivativeWithRespectTo(Value x);

  private Value negate() {
    return this.multiply(Value.of(-1));
  }

  public Value add(Value another) {
    return new Sum(ImmutableList.of(this, another));
  }

  public Value subtract(Value p) {
    return add(p.negate());
  }

  public Value multiply(Value another) {
    return new Product(ImmutableList.of(this, another));
  }

  public Value multiply(Value another1, Value another2) {
    return new Product(ImmutableList.of(this, another1, another2));
  }

  public Value divide(Value another) {
    return multiply(another.pow(-1));
  }

  public Value tanh() {
    return new Tanh(this);
  }

  public Value pow(double e) {
    return new Power(this, e);
  }

  public Value exp() {
    return new Exponent(this);
  }

  public void descend(double step) {
    data += -step * grad;
  }

  public void traverseTopologically(Consumer<Value> consumer) {
    for (Value child : children) {
      child.traverseTopologically(consumer);
    }
    consumer.accept(this);
  }

  @Override
  public String toString() {
    return new StringJoiner(", ", Value.class.getSimpleName() + "[", "]")
        .add("data=" + data)
        .add(operation.toString())
        .add("grad=" + grad)
        .toString();
  }

  public static Value of(double d) {
    return new Scalar(d);
  }

  public static ImmutableList<Value> listOf(double... ds) {
    return Arrays.stream(ds).mapToObj(Value::of).collect(toImmutableList());
  }

  public void clear() {
    grad = 0;
  }
}

class Scalar extends Value {

  Scalar(double data) {
    super(data, ImmutableList.of(), Operation.NONE);
  }

  @Override
  double localDerivativeWithRespectTo(Value x) {
    if (this == x) {
      return 1;
    }
    return 0;
  }
}

class Sum extends Value {
  Sum(Collection<Value> addends) {
    super(addends.stream().mapToDouble(Value::asDouble).sum(), addends, Operation.ADD);
  }

  @Override
  double localDerivativeWithRespectTo(Value x) {
    if (children.contains(x)) {
      return 1;
    }
    return 0;
  }
}

class Product extends Value {

  Product(Collection<Value> multiplicands) {
    super(
        multiplicands.stream().mapToDouble(Value::asDouble).reduce(1, (a, b) -> a * b),
        multiplicands,
        Operation.MULTIPLY);
  }

  @Override
  double localDerivativeWithRespectTo(Value x) {
    if (!children.contains(x)) {
      return 0;
    }
    double p = 1;
    boolean found = false;
    for (Value v : children) {
      if (!found && v == x) {
        found = true;
        continue;
      }
      p *= v.asDouble();
    }
    return p;
  }
}

class Tanh extends Value {

  Tanh(Value v) {
    super(Math.tanh(v.asDouble()), ImmutableList.of(v), Operation.TANH);
  }

  @Override
  double localDerivativeWithRespectTo(Value x) {
    return 1 - data * data; // tanh'(x) = 1 - tanh(x) * tanh(x)
  }
}

class Power extends Value {

  private final double e;

  Power(Value v, double e) {
    super(Math.pow(v.asDouble(), e), ImmutableList.of(v), Operation.POW);
    this.e = e;
  }

  @Override
  double localDerivativeWithRespectTo(Value x) {
    if (!children.contains(x)) {
      return 0;
    }
    return e * Math.pow(x.data, e - 1);
  }

  @Override
  void appendTo(StringBuilder sb) {
    children.get(0).appendTo(sb);
    sb.append('^');
    sb.append(e);
  }
}

class Exponent extends Value {

  Exponent(Value v) {
    super(Math.exp(v.asDouble()), ImmutableList.of(v), Operation.EXP);
  }

  @Override
  double localDerivativeWithRespectTo(Value x) {
    if (!children.contains(x)) {
      return 0;
    }
    return data;
  }
}
