package kae.exp.micrograd.engine;

public enum Operation {
  NONE(""),
  ADD("+"),
  SUBTRACT("-"),
  MULTIPLY("*"),
  TANH("tanh"),
  POW("^"),
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
