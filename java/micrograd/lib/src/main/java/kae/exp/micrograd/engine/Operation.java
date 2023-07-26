package kae.exp.micrograd.engine;

public enum Operation {
  NONE(""),
  ADD("+"),
  MULTIPLY("*"),
  TANH("tanh"),
  EXP("e^");

  private final String symbol;

  Operation(String symbol) {
    this.symbol = symbol;
  }

  @Override
  public String toString() {
    return symbol;
  }
}
