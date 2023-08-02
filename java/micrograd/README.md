# micrograd

This is a Java implementation of the [micrograd](https://github.com/karpathy/micrograd), a tiny
Autograd engine.

Big thanks to [**@karpathy**](https://github.com/karpathy) for his amazing 
[video](https://youtu.be/VMj-3S1tku0).

### Example

```java
var mlp = new MultilayerPerceptron(new Random(), new int[] {3, 4, 4, 1});
var mse = LossFunction.meanSquareError();

var xs =
    ImmutableList.of(
        Value.listOf(2, 3, -1),
        Value.listOf(3, -1, 0.5),
        Value.listOf(0.5, 1, 1),
        Value.listOf(1, 1, -1));
var ys = Value.listOf(1, -1, -1, 1);

ImmutableList<Value> ypred = null;
for (int i = 0; i < 50; i++) {
  ypred = xs.stream().map(mlp::call).map(l -> l.get(0)).collect(toImmutableList());
  var loss = mse.calculate(ys, ypred);

  mlp.clear();
  loss.backward();
  mlp.descend(0.01);

  System.out.printf("%d. %f\n", i, loss.asDouble());
}
System.out.println(ypred.stream().map(Value::asDouble).collect(toImmutableList()));
```