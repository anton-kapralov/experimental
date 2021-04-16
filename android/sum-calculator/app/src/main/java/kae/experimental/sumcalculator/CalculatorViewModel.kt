package kae.experimental.sumcalculator

import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import kae.experimental.calculator.CalculatorClient
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.cancelChildren
import kotlinx.coroutines.launch

class CalculatorViewModel(private val calculatorClient: CalculatorClient = CalculatorClient()) :
    ViewModel() {

    private val scope: CoroutineScope = CoroutineScope(Dispatchers.IO)

    private val sumLiveData = MutableLiveData<Int>()

    fun sum(): MutableLiveData<Int> {
        return sumLiveData
    }

    fun calculateSum(addends: Iterable<Int>) {
        scope.launch { calculatorClient.sum(addends, sumLiveData) }
    }

    override fun onCleared() {
        super.onCleared()
        scope.coroutineContext.cancelChildren()
    }

}