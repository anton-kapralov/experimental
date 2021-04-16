package kae.experimental.sumcalculator

import android.os.Bundle
import android.view.View
import android.widget.EditText
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.ViewModelProvider
import java.lang.Integer.parseInt

class MainActivity : AppCompatActivity() {

    private val calculatorViewModel by lazy {
        ViewModelProvider(this).get(CalculatorViewModel::class.java)
    }

    override fun onCreate(savedInstanceState: Bundle?) {

        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        calculatorViewModel.sum()
            .observe(this, { findViewById<EditText>(R.id.c).setText(it.toString()) })
    }

    fun onCalculate(view: View) {
        val a = parseInt(findViewById<EditText>(R.id.a).text.toString())
        val b = parseInt(findViewById<EditText>(R.id.b).text.toString())

        println("Calculating $a + $b")

        calculatorViewModel.calculateSum(listOf(a, b))
    }
}