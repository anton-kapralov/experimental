package kae.experimental.kotlin.coroutines

import kotlinx.coroutines.coroutineScope
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking

//
// https://kotlinlang.org/docs/coroutines-basics.html#your-first-coroutine
//

fun main() = runBlocking { // this: CoroutineScope
    doWorld()
    println("Done")
}

private suspend fun doWorld() = coroutineScope {
    launch { // launch a new coroutine and continue
        delay(2000L) // non-blocking delay for 2 second (default time unit is ms)
        println("World 2") // print after delay
    }
    launch {
        delay(1000L)
        println("World 1")
    }
    println("Hello") // main coroutine continues while a previous one is delayed
}
