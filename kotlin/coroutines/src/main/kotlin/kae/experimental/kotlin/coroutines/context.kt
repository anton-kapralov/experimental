package kae.experimental.kotlin.coroutines

import kotlinx.coroutines.*
import java.util.concurrent.Executors

//
// https://kotlinlang.org/docs/coroutine-context-and-dispatchers.html#dispatchers-and-threads
//

fun main() = runBlocking<Unit> {
  launch {
    println("main runBlocking   ${Thread.currentThread().name}")
  }
  launch(Dispatchers.Unconfined) {
    println("Unconfined         ${Thread.currentThread().name}")
  }
  launch(Dispatchers.Default) {
    println("Default            ${Thread.currentThread().name}, ${coroutineContext[Job]}")
  }
  val coroutineDispatcher = Executors.newSingleThreadExecutor().asCoroutineDispatcher()
  launch(coroutineDispatcher) {
    println("Executor           ${Thread.currentThread().name}")
  }
  coroutineDispatcher.close()
}
